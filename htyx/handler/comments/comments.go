package comments

import (
	"htyx/handler"
	"htyx/lib/errno"
	"htyx/model/commentsmodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	Content string `json:"content"`
	Author  string `json:"author"`
	Clid    int    `json:"clid"`
	Cltype  int    `json:"cltype"`
	Sonid   int    `json:"sonid"`
}

func (self *CreateCommentRequest) checkparam() error {

	if len(self.Content) > 30 {

		return errno.ErrComment
	}
	if self.Sonid < 0 {

		return errno.ErrValidation
	}
	return nil
}
func GetNewComments(c *gin.Context) {
	uidin, _ := c.Get("uid")
	uid := uidin.(uint64)
	clid, err1 := strconv.Atoi(c.Param("clid"))
	cltype, err2 := strconv.Atoi(c.Param("cltype"))
	start, _ := strconv.Atoi(c.Param("start"))
	if (err1 != nil) || (err2 != nil) || (clid < 0) || (cltype < 0) || (cltype > 3) {
		handler.SendResponse(c, errno.ErrValidation, nil)
	}
	data, err := commentsmodel.GetNewComments(uid, clid, cltype, start)
	handler.SendResponse(c, err, data)
}
func GetHotComments(c *gin.Context) {
	uidin, _ := c.Get("uid")
	uid := uidin.(uint64)
	clid, err1 := strconv.Atoi(c.Param("clid"))
	cltype, err2 := strconv.Atoi(c.Param("cltype"))
	start, _ := strconv.Atoi(c.Param("start"))
	if (err1 != nil) || (err2 != nil) || (clid < 0) || (cltype < 0) || (cltype > 3) {
		handler.SendResponse(c, errno.ErrValidation, nil)
	}
	data, err := commentsmodel.GetHotComments(uid, clid, cltype, start)
	handler.SendResponse(c, err, data)
}
func CreateComment(c *gin.Context) {
	var ccr CreateCommentRequest
	if err := c.BindJSON(&ccr); err != nil {

		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := ccr.checkparam(); err != nil {

		handler.SendResponse(c, err, nil)
		return

	} else {

		commentsmodel.CreateComment(ccr.Clid, ccr.Cltype, ccr.Author, ccr.Content, ccr.Sonid)
		handler.SendResponse(c, errno.OK, nil)
	}
}
