package spoonacular

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"recipe-plateform/internal/recipe/domain"
	"recipe-plateform/internal/recipe/domain/entity"

	"gorm.io/gorm"
)

type RecipeImpl struct {
	APIKey string
	DB     *gorm.DB
}

type FavoriteImpl struct {
	DB *gorm.DB
}

type randomRecipeResponse struct {
	Recipes []struct {
		Id             int    `json:"id"`
		Title          string `json:"title"`
		Image          string `json:"image"`
		ReadyInMinutes int    `json:"readyInMinutes"`
		Servings       int    `json:"servings"`
	} `json:"recipes"`
}

func (r RecipeImpl) GetRandomRecipe(number int) ([]domain.Recipe, error) {

	url := fmt.Sprintf(
		"https://api.spoonacular.com/recipes/random?number=%d&apiKey=%s",
		number,
		r.APIKey,
	)

	fmt.Println("Request URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Error Response:", string(body))
		return nil, fmt.Errorf("spoonacular returned status %d", resp.StatusCode)
	}

	var apiResponse randomRecipeResponse

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	fmt.Println("Recipes Received:", len(apiResponse.Recipes))

	recipes := make([]domain.Recipe, 0, len(apiResponse.Recipes))

	for _, recipe := range apiResponse.Recipes {

		fmt.Println("Recipe:", recipe.Title)

		recipes = append(recipes, domain.Recipe{
			Id:             recipe.Id,
			Title:          recipe.Title,
			Image:          recipe.Image,
			ReadyInMinutes: recipe.ReadyInMinutes,
			Servings:       recipe.Servings,
		})
		recipeEntity := entity.Recipe{
			Id:             recipe.Id,
			Title:          recipe.Title,
			Image:          recipe.Image,
			ReadyInMinutes: recipe.ReadyInMinutes,
			Servings:       recipe.Servings,
		}
		if err := r.DB.FirstOrCreate(&recipeEntity).Error; err != nil {
			return nil, err
		}
	}

	return recipes, nil
}

func (r RecipeImpl) GetRecipebyID(id int) (*domain.Recipe, error) {

	var recipeEntity entity.Recipe

	// 1. Pehle database mein search karo
	err := r.DB.First(&recipeEntity, id).Error

	// 2. Recipe DB mein mil gayi
	if err == nil {
		recipe := domain.Recipe{
			Id:             recipeEntity.Id,
			Title:          recipeEntity.Title,
			Image:          recipeEntity.Image,
			ReadyInMinutes: recipeEntity.ReadyInMinutes,
			Servings:       recipeEntity.Servings,
		}

		return &recipe, nil
	}

	// 3. Database mein koi actual error hai
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 4. DB mein recipe nahi mili
	// Ab Spoonacular API ko hit karo
	url := fmt.Sprintf(
		"https://api.spoonacular.com/recipes/%d/information?apiKey=%s",
		id,
		r.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 5. Spoonacular ka response read karo
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("SPOONACULAR RESPONSE:", string(body))
	// 6. JSON ko entity.Recipe mein convert karo
	err = json.Unmarshal(body, &recipeEntity)
	if err != nil {
		return nil, err
	}

	fmt.Printf("RECIPE ENTITY: %+v\n", recipeEntity)
	// 7. Spoonacular se mili recipe ko DB mein save karo
	err = r.DB.Create(&recipeEntity).Error
	if err != nil {
		return nil, err
	}

	// 8. Entity ko domain mein convert karo
	recipe := domain.Recipe{
		Id:             recipeEntity.Id,
		Title:          recipeEntity.Title,
		Image:          recipeEntity.Image,
		ReadyInMinutes: recipeEntity.ReadyInMinutes,
		Servings:       recipeEntity.Servings,
	}

	return &recipe, nil
}

func (r RecipeImpl) SearchByname(query string) ([]domain.Recipe, error) {

	var recipeEntities []entity.Recipe

	// 1. Pehle database mein existing recipes search karo
	err := r.DB.Where("title LIKE ?", "%"+query+"%").
		Find(&recipeEntities).Error

	if err != nil {
		return nil, err
	}

	// 2. Existing DB recipes ko result mein add karo
	recipes := make([]domain.Recipe, 0)
	existingIDs := make(map[int]bool)

	for _, recipeEntity := range recipeEntities {
		recipes = append(recipes, domain.Recipe{
			Id:             recipeEntity.Id,
			Title:          recipeEntity.Title,
			Image:          recipeEntity.Image,
			ReadyInMinutes: recipeEntity.ReadyInMinutes,
			Servings:       recipeEntity.Servings,
		})

		existingIDs[recipeEntity.Id] = true
	}

	// 3. Spoonacular se multiple recipes search karo
	url := fmt.Sprintf(
		"https://api.spoonacular.com/recipes/complexSearch?query=%s&number=10&apiKey=%s",
		query,
		r.APIKey,
	)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// 4. Spoonacular API error check
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"spoonacular API error: status %d, response: %s",
			resp.StatusCode,
			string(body),
		)
	}

	// 5. Spoonacular response
	var result struct {
		Results []entity.Recipe `json:"results"`
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	// 6. Spoonacular recipes ko process karo
	for _, recipeEntity := range result.Results {

		// Agar recipe DB mein already hai
		// to dobara save nahi karni
		if existingIDs[recipeEntity.Id] {
			continue
		}

		// New recipe DB mein save karo
		err = r.DB.Create(&recipeEntity).Error

		if err != nil {
			return nil, err
		}

		// ID ko existing mark karo
		existingIDs[recipeEntity.Id] = true

		// Domain result mein add karo
		recipes = append(recipes, domain.Recipe{
			Id:             recipeEntity.Id,
			Title:          recipeEntity.Title,
			Image:          recipeEntity.Image,
			ReadyInMinutes: recipeEntity.ReadyInMinutes,
			Servings:       recipeEntity.Servings,
		})
	}

	return recipes, nil
}
func (r RecipeImpl) SreachByingre(query string) ([]domain.Recipe, error) {

	fmt.Println("========== SearchByIngredient CALLED ==========")
	fmt.Println("QUERY:", query)

	// 1. Pehle database mein ingredient search karo
	var ingredient entity.Ingredient

	err := r.DB.Where("name = ?", query).First(&ingredient).Error

	if err == nil {

		var recipes []entity.Recipe

		err = r.DB.
			Joins("JOIN recipe_ingredients ON recipe_ingredients.recipe_id = recipes.id").
			Where("recipe_ingredients.ingredient_id = ?", ingredient.ID).
			Find(&recipes).Error

		if err != nil {
			return nil, err
		}

		if len(recipes) > 0 {
			result := make([]domain.Recipe, 0, len(recipes))

			for _, recipeEntity := range recipes {
				result = append(result, domain.Recipe{
					Id:             recipeEntity.Id,
					Title:          recipeEntity.Title,
					Image:          recipeEntity.Image,
					ReadyInMinutes: recipeEntity.ReadyInMinutes,
					Servings:       recipeEntity.Servings,
				})
			}

			return result, nil
		}
	}

	// 2. Database mein ingredient nahi mila
	// Spoonacular ko hit karo

	url := fmt.Sprintf(
		"https://api.spoonacular.com/recipes/findByIngredients?ingredients=%s&apiKey=%s",
		query,
		r.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 3. API error check
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"spoonacular API error: status %d, response: %s",
			resp.StatusCode,
			string(body),
		)
	}

	// 4. Spoonacular response structure
	var result []struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
		Image string `json:"image"`

		UsedIngredients []struct {
			Name string `json:"name"`
		} `json:"usedIngredients"`

		MissedIngredients []struct {
			Name string `json:"name"`
		} `json:"missedIngredients"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	for _, item := range result {
		fmt.Println("RECIPE:", item.Id, item.Title)
		fmt.Println("USED:", item.UsedIngredients)
		fmt.Println("MISSED:", item.MissedIngredients)
	}

	fmt.Println("TOTAL RECIPES:", len(result))

	recipes := make([]domain.Recipe, 0, len(result))

	// 5. Recipes + ingredients DB mein save karo
	for _, item := range result {

		// Recipe save karo
		recipeEntity := entity.Recipe{
			Id:    item.Id,
			Title: item.Title,
			Image: item.Image,
		}

		err = r.DB.FirstOrCreate(
			&recipeEntity,
			entity.Recipe{Id: item.Id},
		).Error

		if err != nil {
			return nil, err
		}

		// Used + missed ingredients collect karo
		ingredientNames := make(map[string]bool)

		for _, ingredient := range item.UsedIngredients {
			ingredientNames[ingredient.Name] = true
		}

		for _, ingredient := range item.MissedIngredients {
			ingredientNames[ingredient.Name] = true
		}

		// Har ingredient database mein save karo
		for name := range ingredientNames {

			var ingredientEntity entity.Ingredient

			err = r.DB.Where("name = ?", name).
				First(&ingredientEntity).Error

			if err == gorm.ErrRecordNotFound {

				ingredientEntity = entity.Ingredient{
					Name: name,
				}

				err = r.DB.Create(&ingredientEntity).Error

				if err != nil {
					return nil, err
				}

			} else if err != nil {
				return nil, err
			}

			// Recipe aur ingredient ko connect karo
			err = r.DB.Model(&recipeEntity).
				Association("Ingredients").
				Append(&ingredientEntity)

			if err != nil {
				return nil, err
			}
		}

		recipes = append(recipes, domain.Recipe{
			Id:    recipeEntity.Id,
			Title: recipeEntity.Title,
			Image: recipeEntity.Image,
		})
	}

	return recipes, nil
}

func (f FavoriteImpl) AddFavorite(userID uint, recipeID int) error {

	favorite := entity.Favorite{
		UserID:   userID,
		RecipeID: recipeID,
	}

	return f.DB.Create(&favorite).Error
}

func (f FavoriteImpl) RemoveFavorite(userID uint, recipeID int) error {

	return f.DB.
		Where("user_id = ? AND recipe_id = ?", userID, recipeID).
		Delete(&entity.Favorite{}).Error
}

func (f FavoriteImpl) GetFavorites(userID uint) ([]domain.Recipe, error) {

	var recipes []entity.Recipe

	err := f.DB.
		Joins("JOIN favorites ON favorites.recipe_id = recipes.id").
		Where("favorites.user_id = ?", userID).
		Find(&recipes).Error

	if err != nil {
		return nil, err
	}

	result := make([]domain.Recipe, 0, len(recipes))

	for _, recipe := range recipes {
		result = append(result, domain.Recipe{
			Id:             recipe.Id,
			Title:          recipe.Title,
			Image:          recipe.Image,
			ReadyInMinutes: recipe.ReadyInMinutes,
			Servings:       recipe.Servings,
		})
	}

	return result, nil
}
func (r RecipeImpl) GetRecipeDetails(id int) (*domain.RecipeDetails, error) {

	url := fmt.Sprintf(
		"https://api.spoonacular.com/recipes/%d/information?apiKey=%s",
		id,
		r.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"spoonacular API error: status %d, response: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var result struct {
		Id             int    `json:"id"`
		Title          string `json:"title"`
		Image          string `json:"image"`
		ReadyInMinutes int    `json:"readyInMinutes"`
		Servings       int    `json:"servings"`
		Summary        string `json:"summary"`
		Instructions   string `json:"instructions"`

		ExtendedIngredients []struct {
			Original string `json:"original"`
		} `json:"extendedIngredients"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	ingredients := make([]string, 0)

	for _, ingredient := range result.ExtendedIngredients {
		ingredients = append(
			ingredients,
			ingredient.Original,
		)
	}

	return &domain.RecipeDetails{
		Id:             result.Id,
		Title:          result.Title,
		Image:          result.Image,
		ReadyInMinutes: result.ReadyInMinutes,
		Servings:       result.Servings,
		Summary:        result.Summary,
		Instructions:   result.Instructions,
		Ingredients:    ingredients,
	}, nil
}
