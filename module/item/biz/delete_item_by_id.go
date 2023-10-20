package biz

import (
	"context"
	"social-todos-rest-api/module/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, id int) error
}
type deleteItemBiz struct {
	store DeleteItemStorage
}

func NewDeleteItemBiz(storage DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{store: storage}
}

func (biz *deleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	if _, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	if err := biz.store.DeleteItem(ctx, id); err != nil {
		return err
	}
	return nil
}
