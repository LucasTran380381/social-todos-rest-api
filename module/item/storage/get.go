package storage

import (
	"context"
	"social-todos-rest-api/module/item/model"
)

func (store *sqlStore) GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error) {
	var item model.TodoItem
	err := store.db.Where(condition).First(&item).Error
	return &item, err
}
