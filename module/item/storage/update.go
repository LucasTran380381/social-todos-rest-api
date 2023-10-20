package storage

import (
	"context"
	"social-todos-rest-api/module/item/model"
)

func (store *sqlStore) UpdateItem(ctx context.Context, id int, update *model.TodoItemUpdate) error {
	return store.db.Where("id=?", id).Updates(update).Error
}
