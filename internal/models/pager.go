package models

type Pager struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}
