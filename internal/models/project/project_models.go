package project

type DataModel struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	OrgId  int    `json:"orgId"`
	Org    string `json:"org"`
}
