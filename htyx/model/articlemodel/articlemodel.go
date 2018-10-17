package articlemodel

import (
	. "htyx/model"
)

type GetSlideShowResponse struct {
	Id     uint   `json:"id"`
	Imgsrc string `json:"imgsrc"`
}

type GetListByCateIdResponse struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Imgsrc string `json:"imgsrc"`
	Detail string `json:"detail"`
}
type GetOneResponse struct {
	Id           uint   `json:"id"`
	Imgsrc       string `json:"imgsrc"`
	Detail       string `json:"detail"`
	Date         string `json:"date"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Authorimgsrc string `json:"authorimgsrc"`
	Content      string `json:"content"`
	Likenum      uint   `json:"likenum"`
	Likestatus   uint8  `json:"likestatus" db:"like_status"`
}

func GetSlideShow() ([]GetSlideShowResponse, error) {
	gssr := []GetSlideShowResponse{}
	query := "select id,imgsrc from article where isslideshow=1 and isdelete=1 "
	err := DB.Self.Select(&gssr, query)
	if err != nil {
		return nil, err
	}
	return gssr, nil

func GetListByCateId(categoryid int, start int) ([]GetListByCateIdResponse, error) {
	Gl := []GetListByCateIdResponse{}
	query := "select id,title,imgsrc,detail from article where category_id=? and isdelete=1 limit ?, 10 ;"
	err := DB.Self.Select(&Gl, query, categoryid, start)
	if err != nil {
		return nil, err

	}
	return Gl, nil
}

func GetOne(uid uint64, id int, cltype int) (*GetOneResponse, error) {
	var likestatus uint8
	data := GetOneResponse{}
	query := "select id,imgsrc,detail,date,title,author,authorimgsrc,content,likenum from article where id=? and isdelete=1"
	err := DB.Self.Get(&data, query, id)
	if err != nil {
		return nil, err
	}
	query = "SELECT like_status FROM USERLIKE WHERE uid=? AND cl_id=? AND cl_type=?"
	err = DB.Self.Get(&likestatus, query, uid, id, cltype)
	if likestatus == 1 {
		data.Likestatus = 1

	}

	return &data, nil

}
