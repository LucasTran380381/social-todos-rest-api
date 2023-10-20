package storage

import (
	"context"
	"social-todos-rest-api/module/item/model"
)

func (store *sqlStore) DeleteItem(ctx context.Context, id int) error {
	return store.db.
		Table(model.TodoItem{}.TableName()).
		Where("id=?", id).
		Update("status", "delete").Error
}
