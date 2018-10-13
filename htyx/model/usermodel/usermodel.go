package usermodel

import (
	"htyx/lib/token"
	. "htyx/model"

	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type CreateTokenResponse struct {
	Token string `json:"token"`
}

func ChangeLikeStatus(uid uint64, clid int, cltype int) {
	var likestatus int
	var tablename string
	switch cltype {
	case 1:
		tablename = "article"
	case 2:
		tablename = "audio"
	case 3:
		tablename = "video"
	}
	query := "SELECT like_status FROM USERLIKE WHERE uid=? AND cl_id=? AND cl_type=?"
	_ = DB.Self.Get(&likestatus, query, uid, clid, cltype)
	switch likestatus {
	case 0:
		insertlikestatus := "INSERT INTO userlike ( uid, cl_id,cl_type ,like_status) VALUES ( ?,?,?,1 )"
		_, _ = DB.Self.Exec(insertlikestatus, uid, clid, cltype)
		changelikenum := "update " + tablename + " set likenum=likenum+1 where id=?"
		_, _ = DB.Self.Exec(changelikenum, clid)
	case 1:
		updatelikestatus := "update userlike set like_status=2 where uid=? and cl_id=? and cl_type=?"
		_, _ = DB.Self.Exec(updatelikestatus, uid, clid, cltype)
		changelikenum := "update " + tablename + " set likenum=likenum-1 where id=?"
		_, _ = DB.Self.Exec(changelikenum, clid)
	case 2:
		updatelikestatus := "update userlike set like_status=1 where uid=? and cl_id=? and cl_type=?"
		_, _ = DB.Self.Exec(updatelikestatus, uid, clid, cltype)
		changelikenum := "update " + tablename + " set likenum=likenum+1 where id=?"
		_, _ = DB.Self.Exec(changelikenum, clid)
	}
}

func ChangeCommentLikeStatus(uid uint64, commentid int) {
	var likestatus int
	query := "SELECT like_status FROM userlikecomment WHERE uid=? AND comment_id=? "
	_ = DB.Self.Get(&likestatus, query, uid, commentid)
	switch likestatus {
	case 0:
		insertlikestatus := "INSERT INTO userlikecomment ( uid, comment_id ,like_status) VALUES ( ?,?,1 )"
		_, _ = DB.Self.Exec(insertlikestatus, uid, commentid)
		changelikenum := "update comments set likenum=likenum+1 where id=?"
		_, _ = DB.Self.Exec(changelikenum, commentid)
	case 1:
		updatelikestatus := "update userlikecomment set like_status=2 where uid=? and comment_id=? "
		_, _ = DB.Self.Exec(updatelikestatus, uid, commentid)
		changelikenum := "update comments set likenum=likenum-1 where id=?"
		_, _ = DB.Self.Exec(changelikenum, commentid)
	case 2:
		updatelikestatus := "update userlikecomment set like_status=1 where uid=? and comment_id=?"
		_, _ = DB.Self.Exec(updatelikestatus, uid, commentid)
		changelikenum := "update comments set likenum=likenum+1 where id=?"
		_, _ = DB.Self.Exec(changelikenum, commentid)
	}

}

func CreateToken(openid string) (*CreateTokenResponse, error) {
	var uid uint64
	query := "select uid from user where openid=?"
	_ = DB.Self.Get(&uid, query, openid)
	if uid == 0 {
		insertop := "insert into user (openid) values (?)"
		res, err := DB.Self.Exec(insertop, openid)
		if err != nil {
			log.Error("usermodel insert openid mistake", err)
			return nil, err
		}
		uidx, _ := res.LastInsertId()
		uid = uint64(uidx)
		return updatetoken(uid)

	} else {
		return updatetoken(uid)
	}
}

func updatetoken(uid uint64) (*CreateTokenResponse, error) {
	var py token.PayLoad
	var ctr CreateTokenResponse
	var errtk error
	py.Uid = uid
	ctr.Token, errtk = token.Sign(py, viper.GetString("jwt_secret"))
	if errtk != nil {
		log.Error("usermodel create token fail ", errtk)
		return nil, errtk
	}
	updateto := "update user set token=? where uid=?"
	_, err := DB.Self.Exec(updateto, ctr.Token, uid)
	if err != nil {
		log.Error("usermodel update token fail", err)
		return nil, err
	}
	return &ctr, nil
}
