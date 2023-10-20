package storage

import (
	"context"
	"social-todos-rest-api/module/item/model"
	"social-todos-rest-api/shared"
)

func (store *sqlStore) GetItems(ctx context.Context, paging *shared.Paging, filter *model.Filter) ([]model.TodoItem, error) {
	db := store.db.Table(model.TodoItem{}.TableName()).Where("status <> ?", "delete")
	err := db.Error
	if err != nil {
		return nil, err
	}

	if filter != nil {
		if filter.Status != nil {
			db = db.Where("status=?", filter.Status)

		}

		err = db.Error

		if err != nil {
			return nil, err
		}
	}

	if db.Select("id").Count(&paging.Total).Error != nil {
		return nil, store.db.Select("id").Count(&paging.Total).Error
	}

	var result []model.TodoItem
	err = db.Select("*").Order("updated_at desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
