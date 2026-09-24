package main

import (
	recipeApplication "recipe-plateform/internal/recipe/application"
	"recipe-plateform/internal/recipe/infrastructure/spoonacular"
	"recipe-plateform/internal/recipe/presentation/httprecipe"

	"recipe-plateform/configs"
	"recipe-plateform/internal/user/application"
	"recipe-plateform/internal/user/infrastructure/jwt"
	"recipe-plateform/internal/user/infrastructure/mySql"
	"recipe-plateform/internal/user/presentation/middleware"
	"recipe-plateform/internal/user/presentation/userhttp"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	configs.ConnectDB()

	userRepo := mySql.UserImpl{
		DB: configs.DB,
	}

	recipeRepo2 := spoonacular.RecipeImpl{
		DB:     configs.DB,
		APIKey: configs.SpoonacularApiKey,
	}

	favoriteRepo := spoonacular.FavoriteImpl{
		DB: configs.DB,
	}

	createUser := application.AddUser{
		UserRepo: userRepo,
	}

	getUser := application.GetUser{
		UserRepo: userRepo,
	}

	deleteUser := application.DeleteUser{
		UserRepo: userRepo,
	}

	updateUser := application.UpdateUser{
		UserRepo: userRepo,
	}
	getAlluser := application.GetAllUser{
		UserRepo: userRepo,
	}

	getRandomRecipe := recipeApplication.GetRandomRecipe{
		RecipeRepo: recipeRepo2,
	}
	getRecipeByID := recipeApplication.GetRecipeByID{
		RecipeRepo: recipeRepo2,
	}
	getRecipeDetails := recipeApplication.GetRecipeDetails{
		RecipeRepo: recipeRepo2,
	}

	searchRecipe := recipeApplication.SearchRecipe{
		RecipeRepo: recipeRepo2,
	}

	searchByIngredient := recipeApplication.SearchByIngredient{
		RecipeRepo: recipeRepo2,
	}

	getHandler := httprecipe.RecipeHandler{
		GetRandomRecipe: getRandomRecipe,
	}

	getRecipeByIDHandler := httprecipe.GetRecipeByIDHandler{
		GetRecipeByID: getRecipeByID,
	}

	getRecipeDetailsHandler := httprecipe.GetRecipeDetailsHandler{
		GetRecipeDetails: getRecipeDetails,
	}
	jwtService := jwt.JWTService{
		SecretKey: configs.JWTSecret,
	}

	loginUser := application.LoginUser{
		UserRepo:     userRepo,
		TokenService: jwtService,
	}

	registerUser := application.RegisterUser{
		UserRepo: userRepo,
	}

	favoriteHandler := httprecipe.FavoriteHandler{
		AddFavorite: recipeApplication.AddFavorite{
			FavoriteRepo: favoriteRepo,
			RecipeRepo:   recipeRepo2,
		},
		RemoveFavorite: recipeApplication.RemoveFavorite{
			FavoriteRepo: favoriteRepo,
		},
		GetFavorites: recipeApplication.GetFavorites{
			FavoriteRepo: favoriteRepo,
		},
	}

	userHandler := userhttp.UserHandler{
		CreateUser: createUser,
		GetUser:    getUser,
		DeleteUser: deleteUser,
		UpdateUser: updateUser,
		GetAllUser: getAlluser,
	}

	loginHandler := userhttp.LoginHandler{
		LoginUser: loginUser,
	}

	registerHandler := userhttp.RegisterHandler{
		RegisterUser: registerUser,
	}
	searchRecipeHandler := httprecipe.SearchRecipeHandler{
		SearchRecipe:       searchRecipe,
		SearchByIngredient: searchByIngredient,
	}

	authMiddleware := middleware.AuthMiddleware(jwtService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	protected := router.Group("/users")
	protected.Use(authMiddleware)

	// user DB handling part
	router.GET("/users/:id", userHandler.GetByID)
	router.PATCH("/users/:id", userHandler.Update)
	router.DELETE("/users/:id", userHandler.Delete)
	router.GET("/users", userHandler.GetAll)

	// recipe part third party implementation

	router.GET("/recipe/:id/details", getRecipeDetailsHandler.GetDetails)
	router.GET("/recipe/search", searchRecipeHandler.Search)
	router.GET("/recipe/search/ingredient", searchRecipeHandler.SearchIngredient)
	router.GET("/recipe/random", getHandler.GetRandom)
	router.GET("/recipe/:id", getRecipeByIDHandler.GetByID)

	// user login in authentication part
	router.POST("/login", loginHandler.Login)
	router.POST("/users", registerHandler.Register)

	// favorite part for the recipe
	favorites := router.Group("/favorites")
	favorites.Use(authMiddleware)

	favorites.POST("/:id", favoriteHandler.Add)
	favorites.DELETE("/:id", favoriteHandler.Remove)
	favorites.GET("", favoriteHandler.Get)

	router.Run(":8080")
}
