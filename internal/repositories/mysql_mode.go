package repositories

func disableOnlyFullGroupBy() *string {
	sql := "SET sql_mode=(SELECT REPLACE(@@sql_mode, 'ONLY_FULL_GROUP_BY', ''));"
	return &sql
}
