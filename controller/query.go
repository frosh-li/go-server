package controller

type PageQuery struct {
	PageNum  int `form:"current" json:"current"`
	PageSize int `form:"pageSize" json:"pageSize"`
}
