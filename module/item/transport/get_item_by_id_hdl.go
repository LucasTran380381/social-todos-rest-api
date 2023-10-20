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

func GetItemById(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		store := storage.NewMySQLStore(db)
		business := biz.NewGetItemBiz(store)

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}

		item, err := business.GetItemById(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, shared.Response{Data: item})
	}
}
