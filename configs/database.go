package configs

import (
	"log"
	"os"

	recipeEntity "recipe-plateform/internal/recipe/domain/entity"
	"recipe-plateform/internal/user/domain/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var SpoonacularApiKey string

func ConnectDB() {

	JWTSecret = os.Getenv("JWT_SECRET")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Some error occur couldn't connect to db")
	}

	dp_path := os.Getenv("DB_PATH")

	db, err := gorm.Open(sqlite.Open(dp_path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	SpoonacularApiKey = os.Getenv("SPOONACULAR_API_KEY")

	if SpoonacularApiKey == "" {
		log.Fatal("The API key is missing in .env")
	}

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB = db

	err = db.AutoMigrate(&entity.User{}, recipeEntity.Recipe{}, &recipeEntity.Ingredient{}, recipeEntity.Favorite{})
	if err != nil {
		log.Fatal("The error occurs")
	}
	log.Println("database connected successfully")
}
