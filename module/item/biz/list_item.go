package biz

import (
	"context"
	"social-todos-rest-api/module/item/model"
	"social-todos-rest-api/shared"
)

type ListItemStorage interface {
	GetItems(ctx context.Context, paging *shared.Paging, filter *model.Filter) ([]model.TodoItem, error)
}
type listItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(storage ListItemStorage) *listItemBiz {
	return &listItemBiz{store: storage}
}

func (biz *listItemBiz) GetItems(ctx context.Context, paging *shared.Paging, filter *model.Filter) ([]model.TodoItem, error) {
	return biz.store.GetItems(ctx, paging, filter)
}
