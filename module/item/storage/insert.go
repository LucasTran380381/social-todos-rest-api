package storage

import (
	"context"
	"social-todos-rest-api/module/item/model"
)

func (store *sqlStore) CreateItem(ctx context.Context, item *model.TodoItemCreation) error {
	return store.db.Create(&item).Error
}
