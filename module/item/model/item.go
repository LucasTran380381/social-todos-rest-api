package model

import (
	"errors"
	"social-todos-rest-api/shared"
	"strings"
)

var (
	ErrRequiredTitle = errors.New("title must not empty")
	ErrItemIsDeleted = errors.New("this items is deleted")
)

type TodoItem struct {
	shared.SQLBase
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func (TodoItem) TableName() string {
	return "todo_item"
}

type TodoItemCreation struct {
	Id      int
	Title   string
	Content string
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

func (item TodoItemCreation) Validate() error {
	item.Title = strings.TrimSpace(item.Title)
	if item.Title == "" {
		return ErrRequiredTitle
	}
	return nil
}

type TodoItemUpdate struct {
	Title, Content, Status string
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
