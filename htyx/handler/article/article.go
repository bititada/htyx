package article

import (
	"htyx/handler"
	"htyx/lib/errno"
	"htyx/model/articlemodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSlideShow(c *gin.Context) {
	data, err := articlemodel.GetSlideShow()
	handler.SendResponse(c, err, data)
}
func GetListByCateId(c *gin.Context) {

	categoryid, err1 := strconv.Atoi(c.Param("categoryid"))
	start, err2 := strconv.Atoi(c.Param("start"))

	if (err1 != nil) || (err2 != nil) || (categoryid < 0) || (start < 0) {
		handler.SendResponse(c, errno.ErrValidation, nil)
	}
	data, err := articlemodel.GetListByCateId(categoryid, start)
	handler.SendResponse(c, err, data)
}
func GetOne(c *gin.Context) {
	uidin, _ := c.Get("uid")
	uid := uidin.(uint64)
	id, err1 := strconv.Atoi(c.Param("id"))
	cltype, err2 := strconv.Atoi(c.Param("cltype"))
	if (err1 != nil) || (err2 != nil) || (id < 0) || (cltype < 0) || (cltype > 3) {
		handler.SendResponse(c, errno.ErrValidation, nil)
	}
	data, err := articlemodel.GetOne(uid, id, cltype)
	handler.SendResponse(c, err, data)

}
