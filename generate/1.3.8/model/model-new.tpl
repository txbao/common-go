
func new{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, db *gorm.DB, c cache.CacheConf{{end}}) *default{{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: sqlc.NewConn(conn, c, func(o *cache.Options) { o.Expiry = 5 * time.Minute }){{else}}conn:conn{{end}},
		table:      {{.table}},
		db:         db,
	}
}
