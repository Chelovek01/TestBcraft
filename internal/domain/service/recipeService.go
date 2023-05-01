package service

import (
	"TestBcraft/internal/domain/entity"
	"TestBcraft/pkg/logger"
)

type RecipeStorage interface {
	Create(recipe *entity.Recipe) error
	GetOne(id int) (*entity.Recipe, error)
	Update(id int, column string, value string) error
	Delete(recipe *entity.Recipe) error
	GetALL() ([]*entity.Recipe, error)
}

type recipeService struct {
	storage RecipeStorage
}

func NewRecipeService(storage RecipeStorage) *recipeService {
	return &recipeService{storage: storage}
}
func (r *recipeService) Create(recipe *entity.Recipe) error {
	err := r.storage.Create(recipe)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return err
}

func (r *recipeService) GetOne(id int) (*entity.Recipe, error) {
	recipe, err := r.storage.GetOne(id)

	return recipe, err
}

func (r *recipeService) Update(id int, column string, value string) error {
	err := r.storage.Update(id, column, value)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func (r *recipeService) Delete(recipe *entity.Recipe) error {
	err := r.storage.Delete(recipe)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return err
}
func (r *recipeService) GetAll() ([]*entity.Recipe, error) {
	recipe, err := r.storage.GetALL()
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return recipe, err
}
