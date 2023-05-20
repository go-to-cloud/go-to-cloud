package repositories

type OrgLite struct {
	OrgId   uint   `gorm:"column:orgId"`
	OrgName string `gorm:"column:orgName"`
}
