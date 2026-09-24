import { useState } from "react"
import { useNavigate } from "react-router-dom"

function Search() {
  const [query, setQuery] = useState("")
  const [recipes, setRecipes] = useState<any[]>([])

  const navigate = useNavigate()

  const searchRecipes = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!query.trim()) {
      return
    }

    const response = await fetch(
      `http://localhost:8080/recipe/search?query=${encodeURIComponent(query)}`
    )

    const data = await response.json()

    console.log("SEARCH RESPONSE:", data)
    console.log("NUMBER OF RECIPES:", data.length)

    setRecipes(data)
  }

  return (
    <div className="search-page">
      <h1>Search Recipes</h1>

      <form className="search-form" onSubmit={searchRecipes}>
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
        {recipes.map((recipe) => (
          <div
            className="recipe-card"
            key={recipe.id}
            onClick={() => navigate(`/recipe/${recipe.id}/details`)}
          >
            <img
              className="recipe-image"
              src={recipe.image}
              alt={recipe.title}
            />

            <div className="recipe-content">
              <h2>{recipe.title}</h2>

              <p>
                ⏱ Ready in: {recipe.readyInMinutes} minutes
              </p>

              <p>
                🍽 Servings: {recipe.servings}
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default Search
