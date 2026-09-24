import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"

function RecipeDetails() {
  const { id } = useParams()

  const [recipe, setRecipe] = useState<any>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const getRecipeDetails = async () => {
      try {
        const response = await fetch(
          `http://localhost:8080/recipe/${id}/details`
        )

        if (!response.ok) {
          throw new Error("Recipe details not found")
        }

        const data = await response.json()

        console.log("RECIPE DETAILS:", data)

        setRecipe(data)
      } catch (error) {
        console.error("Failed to load recipe details:", error)
      } finally {
        setLoading(false)
      }
    }

    getRecipeDetails()
  }, [id])

  if (loading) {
    return (
      <div className="recipe-details-page">
        <p>Loading recipe...</p>
      </div>
    )
  }

  if (!recipe) {
    return (
      <div className="recipe-details-page">
        <p>Recipe not found.</p>
      </div>
    )
  }

  return (
    <div className="recipe-details-page">
      <div className="recipe-details-card">

        <img
          className="recipe-details-image"
          src={recipe.image}
          alt={recipe.title}
        />

        <div className="recipe-details-content">

          <h1>{recipe.title}</h1>

          <div className="recipe-info">

            <div className="recipe-info-item">
              <span>⏱</span>
              <div>
                <strong>Ready in</strong>
                <p>{recipe.readyInMinutes} minutes</p>
              </div>
            </div>

            <div className="recipe-info-item">
              <span>🍽</span>
              <div>
                <strong>Servings</strong>
                <p>{recipe.servings}</p>
              </div>
            </div>

          </div>

          {recipe.summary && (
            <section className="recipe-section">
              <h2>Summary</h2>

              <div
                dangerouslySetInnerHTML={{
                  __html: recipe.summary,
                }}
              />
            </section>
          )}

          {recipe.ingredients &&
            recipe.ingredients.length > 0 && (
              <section className="recipe-section">
                <h2>Ingredients</h2>

                <ul className="ingredients-list">
                  {recipe.ingredients.map(
                    (ingredient: string, index: number) => (
                      <li key={index}>{ingredient}</li>
                    )
                  )}
                </ul>
              </section>
            )}

          {recipe.instructions && (
            <section className="recipe-section">
              <h2>Instructions</h2>

              <div
                dangerouslySetInnerHTML={{
                  __html: recipe.instructions,
                }}
              />
            </section>
          )}

        </div>
      </div>
    </div>
  )
}

export default RecipeDetails