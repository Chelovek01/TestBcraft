package service

import (
	"TestBcraft/internal/domain/entity"
	"TestBcraft/pkg/logger"
	"context"
)

type RecipeStorage interface {
	Create(recipe *entity.Recipe) error
	GetOne(id string) *entity.Recipe
	Update(id string) *entity.Recipe
	Delete(recipe *entity.Recipe) error
	GetALL(manyID []string) []*entity.Recipe
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
func (r *recipeService) GetOne(ctx context.Context) entity.Recipe {
	return entity.Recipe{}
}
func (r *recipeService) Update(ctx context.Context) entity.Recipe {
	return entity.Recipe{}
}
func (r *recipeService) Delete(recipe *entity.Recipe) error {
	err := r.storage.Delete(recipe)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return err
}
func (r *recipeService) GetAll(ctx context.Context, manyId []string) []*entity.Recipe {
	return nil
}
