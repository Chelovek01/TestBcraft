package postgresadapter

import (
	"TestBcraft/internal/controllers/dto"
	"TestBcraft/internal/domain/entity"
	"TestBcraft/pkg/logger"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type recipeStorage struct {
	db *pgxpool.Pool
}

func NewRecipeStorage(db *pgxpool.Pool) *recipeStorage {
	return &recipeStorage{db: db}
}

func (rs *recipeStorage) Create(recipe *entity.Recipe) error {

	_, err := rs.db.Exec(
		context.Background(),
		"insert into recipe (id, description, ingredients, prepare_steps) values($1,$2,$3,$4)",
		recipe.ID,
		recipe.Description,
		recipe.Ingredients,
		recipe.PrepareSteps,
	)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err

}

func (rs *recipeStorage) GetOne(id string) *entity.Recipe {
	return nil
}
func (rs *recipeStorage) Update(recipeDTO *dto.UpdateRecipeDTO) error {

	_, err := rs.db.Exec(context.Background(), "update ")

	return err
}
func (rs *recipeStorage) Delete(recipe *entity.Recipe) error {
	_, err := rs.db.Exec(context.Background(), "delete from recipe where id=$1", recipe.ID)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return err

}
func (rs *recipeStorage) GetALL(manyID []string) []*entity.Recipe {
	return nil
}
