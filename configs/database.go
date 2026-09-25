package configs

import (
	"log"
	"os"

	recipeEntity "recipe-plateform/internal/recipe/domain/entity"
	"recipe-plateform/internal/user/domain/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var SpoonacularApiKey string

func ConnectDB() {

	// Local development ke liye .env load karo.
	// Render par environment variables already available honge.
	_ = godotenv.Load()

	JWTSecret = os.Getenv("JWT_SECRET")

	// PostgreSQL connection string
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is missing")
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	SpoonacularApiKey = os.Getenv("SPOONACULAR_API_KEY")

	if SpoonacularApiKey == "" {
		log.Fatal("The API key is missing")
	}

	DB = db

	err = db.AutoMigrate(
		&entity.User{},
		&recipeEntity.Recipe{},
		&recipeEntity.Ingredient{},
		&recipeEntity.Favorite{},
	)

	if err != nil {
		log.Fatal("AutoMigrate failed: ", err)
	}

	log.Println("PostgreSQL database connected successfully")
}
