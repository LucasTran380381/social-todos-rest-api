package biz

import (
	"context"
	"social-todos-rest-api/module/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(storage GetItemStorage) *getItemBiz {
	return &getItemBiz{store: storage}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	item, err := biz.store.GetItem(ctx, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, err
	}

	if item.Status == "delete" {
		return nil, model.ErrItemIsDeleted
	}

	return item, nil
}
