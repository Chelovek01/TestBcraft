package postgresadapter

import (
	"TestBcraft/internal/domain/entity"
	"TestBcraft/pkg/logger"
	"context"
	"fmt"
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
		"insert into recipe (id, name, description, ingredients, prepare_steps) values($1,$2,$3,$4,$5)",
		recipe.ID,
		recipe.Name,
		recipe.Description,
		recipe.Ingredients,
		recipe.PrepareSteps,
	)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err

}

func (rs *recipeStorage) GetOne(id int) (*entity.Recipe, error) {

	var oneRecipe entity.Recipe

	rows, err := rs.db.Query(context.Background(), "select * from recipe where id=$1", id)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	if rows.Next() {
		err = rows.Scan(&oneRecipe.ID, &oneRecipe.Name, &oneRecipe.Description, &oneRecipe.Ingredients, &oneRecipe.PrepareSteps)
		if err != nil {
			return nil, err
		}
	}

	return &oneRecipe, err
}
func (rs *recipeStorage) Update(id int, column string, value string) error {
	query := fmt.Sprintf("update recipe set %s='%s' where id=%d", column, value, id)
	_, err := rs.db.Exec(context.Background(), query)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}
func (rs *recipeStorage) Delete(recipe *entity.Recipe) error {
	_, err := rs.db.Exec(context.Background(), "delete from recipe where id=$1", recipe.ID)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return err
}

func (rs *recipeStorage) GetALL() ([]*entity.Recipe, error) {

	var resArr []*entity.Recipe

	rows, err := rs.db.Query(context.Background(), "select * from recipe")
	for rows.Next() {
		var objectRecipe entity.Recipe

		err = rows.Scan(&objectRecipe.ID, &objectRecipe.Name, &objectRecipe.Description,
			&objectRecipe.Ingredients, &objectRecipe.PrepareSteps)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		resArr = append(resArr, &objectRecipe)
	}

	return resArr, err
}
