package handlers

//func updateRecipe(c *gin.Context, p *pgxpool.Pool) {
//
//	var updateObject dto.UpdateRecipeDTO
//
//	err := c.ShouldBindJSON(&updateObject)
//	if err != nil {
//		logger.ErrorLogger.Println(err)
//	}
//
//	recipeStorage := postgresadapter.NewRecipeStorage(p)
//
//	recipeService := service.RecipeStorage(recipeStorage)
//
//	err = recipeService.Update()
//
//}
