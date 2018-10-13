package likelistmodel

import (
	. "htyx/model"
)

type LikeListResponse struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Imgsrc string `json:"imgsrc"`
	Detail string `json:"detail"`
}

func GetListByCltype(uid uint64, ty int, start int) ([]LikeListResponse, error) {
	var query string
	LikeList := []LikeListResponse{}
	switch ty {
	case 1:
		query = "select article.id,article.title,article.imgsrc,article.detail from article,userlike where userlike.cl_type=1 and userlike.like_status=1 and userlike.uid=? and article.id=userlike.cl_id and article.isdelete=1 limit ?, 10 ;"
	case 2:
		query = "select audio.id,audio.title,audio.imgsrc from audio,userlike where userlike.cl_type=2 and userlike.like_status=1 and userlike.uid=? and audio.id=userlike.cl_id and audio.isdelete=1 limit ?, 10 ;"

	case 3:
		query = "select video.id,video.title,video.imgsrc from video,userlike where userlike.cl_type=3 and userlike.like_status=1 and userlike.uid=? and video.id=userlike.cl_id and video.isdelete=1 limit ?, 10 ;"

	}
	err := DB.Self.Select(&LikeList, query, uid, start)
	if err != nil {
		return nil, err

	}
	return LikeList, nil

}
