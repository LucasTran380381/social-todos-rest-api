package biz

import (
	"context"
	"social-todos-rest-api/module/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, item *model.TodoItemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(storage CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: storage}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, item *model.TodoItemCreation) error {
	if err := item.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateItem(ctx, item); err != nil {
		return err
	}

	return nil
}
