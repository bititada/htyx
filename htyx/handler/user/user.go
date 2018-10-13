package user

import (
	"encoding/json"
	"fmt"
	"htyx/handler"
	"htyx/lib/errno"
	"htyx/model/usermodel"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GiveLikeRequest struct {
	Cltype int `json:"cltype"`
	Clid   int `json:"clid"`
}
type CommentLikeRequest struct {
	Commentid int `json:"commentid"`
}
type CreateTokenRequest struct {
	Code string `json:"code"`
}
type WxResponse struct {
	Openid     string `json:"openid"`
	Sessionkey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	ErrMsg     string `json:"errMsg"`
}

//此函数可以判断参数具体错误实现复杂判断
func (self *GiveLikeRequest) checkparam() error {

	if self.Clid < 0 {

		//可以写具体错误
		return errno.ErrValidation
	}
	if self.Cltype < 0 {

		return errno.ErrValidation
	}
	return nil
}
func (self *CommentLikeRequest) checkparam() error {

	if self.Commentid < 0 {

		return errno.ErrValidation
	}
	return nil
}

func CreateToken(c *gin.Context) {
	var ctr CreateTokenRequest
	var wxr WxResponse
	if err := c.BindJSON(&ctr); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	resp, err := http.Get(fmt.Sprintf(viper.GetString("wxurl"), ctr.Code))
	respbody, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		handler.SendResponse(c, errno.ErrWx, nil)
		return
	}
	err = json.Unmarshal(respbody, &wxr)
	if err != nil {
		handler.SendResponse(c, errno.ErrWx, nil)
		return
	}
	if wxr.Errcode != 0 {
		handler.SendResponse(c, errno.ErrWx, nil)
		return
	}
	data, err := usermodel.CreateToken(wxr.Openid)

	if err != nil {
		handler.SendResponse(c, errno.ErrWx, nil)
		return
	}

	handler.SendResponse(c, err, data)
}
func ChangeLikeStatus(c *gin.Context) {
	uidin, _ := c.Get("uid")
	uid := uidin.(uint64)
	var glr GiveLikeRequest

	if err := c.BindJSON(&glr); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := glr.checkparam(); err != nil {
		handler.SendResponse(c, err, nil)
		return

	} else {
		usermodel.ChangeLikeStatus(uid, glr.Clid, glr.Cltype)
	}
}
func ChangeCommentLikeStatus(c *gin.Context) {
	uidin, _ := c.Get("uid")
	uid := uidin.(uint64)
	var clr CommentLikeRequest

	if err := c.BindJSON(&clr); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := clr.checkparam(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	} else {
		usermodel.ChangeCommentLikeStatus(uid, clr.Commentid)
	}
}
