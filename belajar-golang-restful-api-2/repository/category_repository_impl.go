package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/RyouRio/belajar-golang-restful-api-2/helper"
	"github.com/RyouRio/belajar-golang-restful-api-2/model/domain"
)

type CategoryRepositoryImpl struct {
}
func NewCategoryRepository() CategoryRepository{
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category *domain.Category)*domain.Category {
	sql := "INSERT INTO category(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	
	category.Id = int(id)
	return category

}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category *domain.Category) *domain.Category {
	sql := "UPDATE category SET name = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	rowsAffected, err := result.RowsAffected() // check when the database not found (best practice)
	helper.PanicIfError(err)

	if rowsAffected == 0 {
		panic(errors.New("category not found"))
	}

	return category
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category *domain.Category) {
	sql := "DELETE FROM category WHERE id = ?"
	result, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err) 

	rowsAffected, err := result.RowsAffected()
	helper.PanicIfError(err)

	if rowsAffected == 0 {
		panic(errors.New("category id not found"))
	}
}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (*domain.Category, error) {
	sql := "SELECT id, name FROM category WHERE id = ?"
	row := tx.QueryRowContext(ctx, sql, categoryId)

	category := domain.Category{} // dibutuhkan untuk scan datanya
	err := row.Scan(&category.Id, &category.Name)

	if err != nil {
		return &category, errors.New("category not found")
	}
	return &category, nil
}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []*domain.Category {
	sql := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []*domain.Category // slice kosong
	for rows.Next() {
		category := &domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}
	return categories

	/*var categories []domain.Category  // slice kosong
for rows.Next() {                 // loop sampai tidak ada baris lagi
    category := domain.Category{} // buat struct baru tiap iterasi
    err := rows.Scan(&category.Id, &category.Name)
    helper.PanicIfError(err)

    categories = append(categories, category) // tambah ke slice
}
return categories
```

---

## Ilustrasi Prosesnya

Misalnya database punya 3 data:
```
| id | name        |
|----|-------------|
| 1  | Electronics |
| 2  | Fashion     |
| 3  | Food        |
```
```
rows.Next() → iterasi 1 → scan → categories = [{1, Electronics}]
rows.Next() → iterasi 2 → scan → categories = [{1, Electronics}, {2, Fashion}]
rows.Next() → iterasi 3 → scan → categories = [{1, Electronics}, {2, Fashion}, {3, Food}]
rows.Next() → false → loop berhenti
return categories*/
}

