import { useEffect, useState } from "react"
import { Link, useLocation, useNavigate } from "react-router-dom"

function Navbar() {
  const navigate = useNavigate()
  const location = useLocation()

  const [isLoggedIn, setIsLoggedIn] = useState(
    !!localStorage.getItem("token")
  )

  useEffect(() => {
    setIsLoggedIn(!!localStorage.getItem("token"))
  }, [location])

  const handleLogout = () => {
    localStorage.removeItem("token")
    setIsLoggedIn(false)
    navigate("/")
  }

  return (
    <nav className="navbar">
      <div className="navbar-container">

        <Link to="/" className="navbar-logo">
          Recipe Platform
        </Link>

        <div className="navbar-links">

          {!isLoggedIn ? (
            <>
              <Link to="/login">Login</Link>
              <Link to="/register">Register</Link>
            </>
          ) : (
            <>
              <Link to="/search">🔍 Search</Link>

              <Link to="/favorites">
                ❤️ Favorites
              </Link>

              <button
                className="logout-button"
                onClick={handleLogout}
              >
                Logout
              </button>
            </>
          )}

        </div>

      </div>
    </nav>
  )
}

export default Navbar