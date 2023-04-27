package entity

type Recipe struct {
	ID           int    `json:"id"`
	Description  string `json:"description"`
	Ingredients  string `json:"ingredients"`
	PrepareSteps string `json:"prepare_steps"`
}
