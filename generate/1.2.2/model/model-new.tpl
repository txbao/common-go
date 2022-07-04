
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn, db *gorm.DB{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: sqlc.NewConn(conn, c, func(o *cache.Options) {
              o.Expiry = 5 * time.Minute
        }){{else}}conn:conn{{end}},
		table:      {{.table}},
		db:    db,
	}
}
