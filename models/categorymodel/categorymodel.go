package categorymodel

import (
	"quest-rolling-system-RnD-backup/config"
	"quest-rolling-system-RnD-backup/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories

}

func Add(category entities.Category) bool {
	result, err := config.DB.Exec(`INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?)`, category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	/*
		Jika berhasil menginsert data maka LastInsertId > 0
		Sebaliknya, jika gagal menginsert data maka LastInsertId = 0
	*/
	return LastInsertId > 0
}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)

	var category entities.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		panic(err)
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, category.Name, category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}

func Search(query string) []entities.Category {
	rows, err := config.DB.Query(`
		SELECT id, name, created_at, updated_at FROM categories
		WHERE name LIKE ?
	`, "%"+query+"%")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(
			&category.Id, 
			&category.Name, 
			&category.CreatedAt, 
			&category.UpdatedAt,
		); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}
