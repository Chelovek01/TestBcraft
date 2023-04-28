package dto

type UpdateRecipeDTO struct {
	Field        []string `json:"field"`
	ID           int      `json:"id"`
	Description  string   `json:"description"`
	Ingredients  string   `json:"ingredients"`
	PrepareSteps string   `json:"prepare_steps"`
}
