package user

type Org struct {
	Key         uint   `json:"key"`
	Id          uint   `json:"id"`
	CreatedAt   string `json:"created_at"`
	Name        string `json:"name"` // 组织名称
	MemberCount uint   `json:"member_count"`
	Remark      string `json:"remark"`
}
