package ginupload

import (
	"github.com/gin-gonic/gin"
	"go-dev/common"
	"go-dev/component"
	uploadbusiness "go-dev/module/upload/biz"
	uploadstorage "go-dev/module/upload/storage"
	"strconv"
	"strings"

	"net/http"
)

func ListImageHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		arr := strings.Split(c.Param(common.ImageIds), ",")

		ids := make([]int, len(arr))

		for i, e := range arr {
			id, err := strconv.Atoi(e)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			ids[i] = id
		}

		store := uploadstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := uploadbusiness.NewListImageBiz(store)

		result, err := biz.List(c.Request.Context(), ids)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
