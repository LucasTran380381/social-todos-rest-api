package transport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todos-rest-api/module/item/biz"
	"social-todos-rest-api/module/item/model"
	"social-todos-rest-api/module/item/storage"
	"social-todos-rest-api/shared"
)

func CreateNewItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var item model.TodoItemCreation
		if err := ctx.ShouldBind(&item); err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}

		store := storage.NewMySQLStore(db)
		business := biz.NewCreateItemBiz(store)
		if err := business.CreateNewItem(ctx.Request.Context(), &item); err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, shared.Response{Data: item.Id})
	}
}
