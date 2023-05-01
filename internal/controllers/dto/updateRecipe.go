package dto

type UpdateRecipeDTO struct {
	ID     int    `json:"id"`
	Column string `json:"column"`
	Value  string `json:"value"`
}
