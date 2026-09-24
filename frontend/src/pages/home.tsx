import { useEffect, useState } from "react"
import { Link } from "react-router-dom"

function Home() {
  const token = localStorage.getItem("token")

  const [recipes, setRecipes] = useState<any[]>([])
  const [favoriteIds, setFavoriteIds] = useState<number[]>([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    if (!token) {
      return
    }

    const getData = async () => {
      setLoading(true)

      try {
        // Get random recipes
        const recipeResponse = await fetch(
          "http://localhost:8080/recipe/random?number=6"
        )

        const recipeData = await recipeResponse.json()

        setRecipes(recipeData)

        // Get user's favorites
        const favoriteResponse = await fetch(
          "http://localhost:8080/favorites",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        )

        const favoriteData = await favoriteResponse.json()

        const ids = favoriteData.map((recipe: any) => recipe.id)

        setFavoriteIds(ids)
      } catch (error) {
        console.error("Failed to load data:", error)
      } finally {
        setLoading(false)
      }
    }

    getData()
  }, [token])

  const handleFavorite = async (recipeId: number) => {
    const token = localStorage.getItem("token")

    if (!token) {
      return
    }

    const isFavorite = favoriteIds.includes(recipeId)

    try {
      if (isFavorite) {
        // Remove from favorites
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
        // Add to favorites
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

  // Guest Home
  if (!token) {
    return (
      <div className="home-page">
        <section className="hero-section">
          <h1>Welcome to Recipe Platform</h1>

          <p>
            Discover delicious recipes, search for meals by name or
            ingredient, and save your favorite recipes in one place.
          </p>
        </section>

        <section className="guide-section">
          <h2>What can you do here?</h2>

          <div className="guide-grid">
            <div className="guide-card">
              <div className="guide-icon">🔍</div>

              <h3>Search Recipes</h3>

              <p>
                Search for recipes by their name or by an ingredient
                you already have at home.
              </p>
            </div>

            <div className="guide-card">
              <div className="guide-icon">❤️</div>

              <h3>Save Favorites</h3>

              <p>
                Found a recipe you like? Save it to your favorites
                so you can easily find it later.
              </p>
            </div>

            <div className="guide-card">
              <div className="guide-icon">🍳</div>

              <h3>Discover Recipes</h3>

              <p>
                Explore different recipes and discover new ideas
                for your next meal.
              </p>
            </div>
          </div>
        </section>

        <section className="how-section">
          <h2>How to use Recipe Platform?</h2>

          <div className="steps">
            <div className="step">
              <span>1</span>

              <div>
                <h3>Create an account</h3>

                <p>
                  Register your account to access all recipe features.
                </p>
              </div>
            </div>

            <div className="step">
              <span>2</span>

              <div>
                <h3>Search for recipes</h3>

                <p>
                  Search by recipe name or ingredient.
                </p>
              </div>
            </div>

            <div className="step">
              <span>3</span>

              <div>
                <h3>Save your favorites</h3>

                <p>
                  Save recipes you want to try later.
                </p>
              </div>
            </div>
          </div>
        </section>
      </div>
    )
  }

  // Logged-in Home
  return (
    <div className="home-page">
      <section className="welcome-section">
        <h1>Welcome Back 👋</h1>

        <p>
          Here are some recipes you might enjoy today.
        </p>
      </section>

      <section className="random-recipes-section">
        <h2>Discover Recipes</h2>

        {loading ? (
          <p className="home-message">
            Loading recipes...
          </p>
        ) : recipes.length === 0 ? (
          <p className="home-message">
            No recipes found.
          </p>
        ) : (
          <div className="recipe-grid">
            {recipes.map((recipe) => {
              const isFavorite = favoriteIds.includes(recipe.id)

              return (
                <div
                  className="recipe-card"
                  key={recipe.id}
                >
                  <div className="recipe-image-container">
                    {/* Recipe image clickable */}
                    <Link to={`/recipe/${recipe.id}`}>
                      <img
                        className="recipe-image"
                        src={recipe.image}
                        alt={recipe.title}
                      />
                    </Link>

                    {/* Favorite button */}
                    <button
                      className="favorite-icon"
                      onClick={(e) => {
                        e.stopPropagation()
                        handleFavorite(recipe.id)
                      }}
                      title={
                        isFavorite
                          ? "Remove from favorites"
                          : "Add to favorites"
                      }
                    >
                      {isFavorite ? "♥" : "♡"}
                    </button>
                  </div>

                  {/* Recipe information clickable */}
                  <Link
                    to={`/recipe/${recipe.id}`}
                    className="recipe-card-link"
                  >
                    <div className="recipe-content">
                      <h2>{recipe.title}</h2>

                      <p>
                        ⏱ Ready in: {recipe.readyInMinutes} minutes
                      </p>

                      <p>
                        🍽 Servings: {recipe.servings}
                      </p>
                    </div>
                  </Link>
                </div>
              )
            })}
          </div>
        )}
      </section>
    </div>
  )
}

export default Home