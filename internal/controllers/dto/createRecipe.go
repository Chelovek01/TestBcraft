package dto

type Recipe struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Ingredients  string `json:"ingredients"`
	PrepareSteps string `json:"prepare_steps"`
}

type CreateRecipeDTO struct {
	Action string `json:"action"`
	Recipe Recipe `json:"recipe"`
}
