package category

import (
	"htyx/handler"
	"htyx/lib/errno"
	"htyx/model/categorymodel"
	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/lexkong/log"
)

func GetCategory(c *gin.Context) {
	cltype, err := strconv.Atoi(c.Param("cltype"))
	if (err != nil) || (cltype < 0) {
		handler.SendResponse(c, errno.ErrValidation, nil)
	}
	data, err := categorymodel.GetCategory(cltype)
	handler.SendResponse(c, err, data)
}
