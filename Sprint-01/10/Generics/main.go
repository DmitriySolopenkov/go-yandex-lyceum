package main

import "context"

// func sum(nums []int) int {
// 	total := 0
// 	for _, n := range nums {
// 		total += n
// 	}
// 	return total
// }

func sum[T any](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// Repository — обобщённый тип для работы с базой данных
type Repository[T any] struct {
	db *sqlx.DB
}

// NewRepository — создаёт новый экземпляр репозитория
func NewRepository[T any](db *sqlx.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

// Add — добавляет запись в базу данных
func (r *Repository[T]) Add(entity T, ctx context.Context) error {
	_, err := r.db.NamedExecContext(ctx, "INSERT INTO your_table_name_here (field1, field2) VALUES (:field1, :field2)", entity)
	return err
}

// GetById — получает запись из базы данных по идентификатору
func (r *Repository[T]) GetById(id int, ctx context.Context) (T, error) {
	var entity T
	err := r.db.GetContext(ctx, &entity, "SELECT * FROM your_table_name_here WHERE id = ? AND is_active = ?", id, true)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

// Get — выполняет поиск записи в базе данных по параметрам
func (r *Repository[T]) Get(params T, ctx context.Context) T {
	var entity T
	r.db.GetContext(ctx, &entity, "SELECT * FROM your_table_name_here WHERE field1 = :field1 AND field2 = :field2", params)
	return entity
}

// GetAll — получает все записи из базы данных
func (r *Repository[T]) GetAll(ctx context.Context) ([]T, error) {
	var entities []T
	err := r.db.SelectContext(ctx, &entities, "SELECT * FROM your_table_name_here")
	if err != nil {
		return entities, err
	}
	return entities, nil
}

// Update — обновляет запись в базе данных
func (r *Repository[T]) Update(entity T, ctx context.Context) error {
	_, err := r.db.NamedExecContext(ctx, "UPDATE your_table_name_here SET field1 = :field1, field2 = :field2 WHERE id = :id", entity)
	return err
}
