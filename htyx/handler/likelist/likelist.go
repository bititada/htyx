package likelist

import (
	"htyx/handler"
	"htyx/lib/errno"
	"htyx/model/likelistmodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetListByCltype(c *gin.Context) {
	uidin, _ := c.Get("uid")
	uid := uidin.(uint64)
	start, err1 := strconv.Atoi(c.Param("start"))
	ty, err := strconv.Atoi(c.Param("cltype"))
	if (err != nil) || (ty < 0) || (err1 != nil) || (start < 0) {
		handler.SendResponse(c, errno.ErrValidation, nil)
	}
	data, err := likelistmodel.GetListByCltype(uid, ty, start)
	handler.SendResponse(c, err, data)
}
