package transport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todos-rest-api/module/item/biz"
	"social-todos-rest-api/module/item/storage"
	"social-todos-rest-api/shared"
	"strconv"
)

func DeleteItemById(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		store := storage.NewMySQLStore(db)
		business := biz.NewDeleteItemBiz(store)

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}
		if err := business.DeleteItemById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, shared.Response{Data: true})
	}
}
