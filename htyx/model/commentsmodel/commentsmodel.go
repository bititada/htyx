package commentsmodel

import (
	"database/sql"

	. "htyx/model"

	"github.com/lexkong/log"
)

type CommentsResponse struct {
	Id         int            `json:"id" db:"aid"`
	Content    string         `json:"content" db:"acontent"`
	Author     string         `json:"author" db:"aauthor"`
	SonContent sql.NullString `json:"soncontent" db:"bcontent"`
	SonAuthor  sql.NullString `json:"sonauthor" db:"bauthor"`
	Likestatus sql.NullInt64  `json:"likestatus" db:"clikestatus"`
	Likenum    int            `json:"likenum" db:"alikenum"`
	Createtime string         `json:"createtime" db:"acreatetime"`
}

func GetNewComments(uid uint64, clid int, cltype int, start int) ([]CommentsResponse, error) {
	var csr []CommentsResponse
	query := "select a.id as aid,a.createtime as acreatetime ,a.content as acontent,a.author as aauthor,a.likenum as alikenum,b.content as bcontent,b.author as bauthor,c.like_status as clikestatus from comments a left join comments b on a.sonid=b.id  left join userlikecomment c on a.id=c.comment_id and c.uid=?  where a.cl_id=? and a.cl_type=? and a.isagree=1  order by a.id desc limit ?,20;"
	rows, err := DB.Self.Queryx(query, uid, clid, cltype, start)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var cr CommentsResponse
		err = rows.StructScan(&cr)
		if err != nil {
			return nil, err
		}
		csr = append(csr, cr)
	}

	return csr, nil
}
func GetHotComments(uid uint64, clid int, cltype int, start int) ([]CommentsResponse, error) {
	var csr []CommentsResponse
	query := "select a.id as aid,a.createtime as acreatetime,a.content as acontent,a.author as aauthor,a.likenum as alikenum,b.content as bcontent,b.author as bauthor,c.like_status as clikestatus from comments a left join comments b on a.sonid=b.id  left join userlikecomment c on a.id=c.comment_id and c.uid=?  where a.cl_id=? and a.cl_type=? and a.isagree=1   order by a.likenum desc limit ?,5;"
	rows, err := DB.Self.Queryx(query, uid, clid, cltype, start)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var cr CommentsResponse
		err = rows.StructScan(&cr)
		if err != nil {
			return nil, err
		}
		csr = append(csr, cr)
	}

	return csr, nil
}
func CreateComment(clid int, cltype int, author string, content string, sonid int) {
	insertcomment := "INSERT INTO comments ( cl_id,cl_type ,author,content,sonid) VALUES ( ?,?,?,?,?)"
	_, err := DB.Self.Exec(insertcomment, clid, cltype, author, content, sonid)
	if err != nil {
		log.Error("comments insert mistake", err)
	}

}
