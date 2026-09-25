import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"

function Search() {
  const [query, setQuery] = useState("")
  const [recipes, setRecipes] = useState<any[]>([])
  const [favoriteIds, setFavoriteIds] = useState<number[]>([])

  const navigate = useNavigate()

  // Load user's favorites
  useEffect(() => {
    const token = localStorage.getItem("token")

    if (!token) {
      return
    }

    const getFavorites = async () => {
      try {
        const response = await fetch(
          "http://localhost:8080/favorites",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        )

        if (!response.ok) {
          return
        }

        const data = await response.json()

        const ids = data.map((recipe: any) => recipe.id)

        setFavoriteIds(ids)
      } catch (error) {
        console.error("Failed to load favorites:", error)
      }
    }

    getFavorites()
  }, [])

  const searchRecipes = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!query.trim()) {
      return
    }

    try {
      const response = await fetch(
        `http://localhost:8080/recipe/search?query=${encodeURIComponent(query)}`
      )

      const data = await response.json()

      console.log("SEARCH RESPONSE:", data)
      console.log("NUMBER OF RECIPES:", data.length)

      setRecipes(data)
    } catch (error) {
      console.error("Search failed:", error)
    }
  }

  const handleFavorite = async (
    e: React.MouseEvent,
    recipeId: number
  ) => {
    // Card ka click trigger nahi hone dena
    e.stopPropagation()

    const token = localStorage.getItem("token")

    if (!token) {
      return
    }

    const isFavorite = favoriteIds.includes(recipeId)

    try {
      if (isFavorite) {
        // Remove favorite
        const response = await fetch(
          `http://localhost:8080/favorites/${recipeId}`,
          {
            method: "DELETE",
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        )

        if (response.ok) {
          setFavoriteIds((prev) =>
            prev.filter((id) => id !== recipeId)
          )
        }
      } else {
        // Add favorite
        const response = await fetch(
          `http://localhost:8080/favorites/${recipeId}`,
          {
            method: "POST",
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        )

        if (response.ok) {
          setFavoriteIds((prev) => [...prev, recipeId])
        }
      }
    } catch (error) {
      console.error("Favorite action failed:", error)
    }
  }

  return (
    <div className="search-page">

      <h1>Search Recipes</h1>

      <form
        className="search-form"
        onSubmit={searchRecipes}
      >
        <input
          type="text"
          placeholder="Search recipe..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />

        <button type="submit">
          Search
        </button>
      </form>

      <div className="recipe-grid">

        {recipes.map((recipe) => {

          const isFavorite = favoriteIds.includes(recipe.id)

          return (
            <div
              className="recipe-card"
              key={recipe.id}
              onClick={() =>
                navigate(`/recipe/${recipe.id}/details`)
              }
            >

              <div className="recipe-image-container">

                <img
                  className="recipe-image"
                  src={recipe.image}
                  alt={recipe.title}
                />

                <button
                  className="favorite-icon"
                  onClick={(e) =>
                    handleFavorite(e, recipe.id)
                  }
                  title={
                    isFavorite
                      ? "Remove from favorites"
                      : "Add to favorites"
                  }
                >
                  {isFavorite ? "♥" : "♡"}
                </button>

              </div>

              <div className="recipe-content">

                <h2>{recipe.title}</h2>

                <p>
                  ⏱ Ready in:{" "}
                  {recipe.readyInMinutes} minutes
                </p>

                <p>
                  🍽 Servings: {recipe.servings}
                </p>

              </div>

            </div>
          )
        })}

      </div>

    </div>
  )
}

export default Search