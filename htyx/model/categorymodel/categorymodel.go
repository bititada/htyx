package categorymodel

import (
	. "htyx/model"
)

type GetCategoryResponse struct {
	Id         uint   `json:"id"`
	Title      string `json:"title"`
	Cateimgsrc string `json:"cateimgsrc"`
}

func GetCategory(cltype int) ([]GetCategoryResponse, error) {

	Gcr := []GetCategoryResponse{}
	query := "select id,title,cateimgsrc from category where cl_type=? ;"
	err := DB.Self.Select(&Gcr, query, cltype)
	if err != nil {
		return nil, err

	}
	return Gcr, nil
}
