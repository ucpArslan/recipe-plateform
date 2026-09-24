import { useEffect, useState } from "react"

function Favorites() {
  const [favorites, setFavorites] = useState<any[]>([])

  useEffect(() => {
    const token = localStorage.getItem("token")

    fetch("http://localhost:8080/favorites", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        setFavorites(data)
      })
  }, [])

  const removeFavorite = async (recipeId: number) => {
    const token = localStorage.getItem("token")

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
      setFavorites((prevFavorites) =>
        prevFavorites.filter((recipe) => recipe.id !== recipeId)
      )
    }
  }

  return (
    <div className="favorites-page">
      <h1>My Favorites</h1>

      {favorites.length === 0 ? (
        <p className="no-favorites">No favorite recipes yet.</p>
      ) : (
        <div className="recipe-grid">
          {favorites.map((recipe) => (
            <div className="recipe-card" key={recipe.id}>
              <img
                className="recipe-image"
                src={recipe.image}
                alt={recipe.title}
              />

              <div className="recipe-content">
                <h2>{recipe.title}</h2>

                <p>⏱ Ready in: {recipe.readyInMinutes} minutes</p>

                <p>🍽 Servings: {recipe.servings}</p>

                <button
                  className="remove-button"
                  onClick={() => removeFavorite(recipe.id)}
                >
                  Remove Favorite
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default Favorites