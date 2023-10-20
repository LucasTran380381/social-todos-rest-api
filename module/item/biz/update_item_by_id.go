package biz

import (
	"context"
	"social-todos-rest-api/module/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, id int, update *model.TodoItemUpdate) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(storage UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store: storage}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, data *model.TodoItemUpdate) error {
	item, err := biz.store.GetItem(ctx, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}

	if item.Status == "delete" {
		return model.ErrItemIsDeleted
	}

	err = biz.store.UpdateItem(ctx, id, data)
	if err != nil {
		return err
	}

	return nil
}
