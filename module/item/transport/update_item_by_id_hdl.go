package transport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todos-rest-api/module/item/biz"
	"social-todos-rest-api/module/item/model"
	"social-todos-rest-api/module/item/storage"
	"social-todos-rest-api/shared"
	"strconv"
)

func UpdateItemById(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		store := storage.NewMySQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}

		var updateItem model.TodoItemUpdate
		if err := ctx.ShouldBind(&updateItem); err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}

		if err := business.UpdateItemById(ctx.Request.Context(), id, &updateItem); err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, shared.Response{Data: true})
	}
}
