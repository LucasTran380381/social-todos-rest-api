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

func GetItems(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var query struct {
			shared.Paging
			model.Filter
		}
		if err := ctx.ShouldBindQuery(&query); err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}
		query.Format()

		store := storage.NewMySQLStore(db)
		business := biz.NewListItemBiz(store)

		items, err := business.GetItems(ctx, &query.Paging, &query.Filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, shared.Response{Data: items})
	}
}
