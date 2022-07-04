//修改
func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, data *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{.keys}}
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}

//gorm修改
func (m *default{{.upperStartCamelObject}}Model) UpdateGorm(ctx context.Context, data *{{.upperStartCamelObject}}) error {
    data.UpdatedAt = time.Now().Local().Unix()
	err := m.db.Model(&{{.upperStartCamelObject}}{}).
		Where("id = ?", data.Id).
		Updates(data).
		Error
	if err != nil {
		return err
	}
	_ = m.CacheDel(ctx, data.Id)
	return nil
}
//修改-分布式事务专用
func (m *default{{.upperStartCamelObject}}Model) UpdateTransaction(ctx context.Context,tx *sql.Tx, data *{{.upperStartCamelObject}}) error {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _, err := tx.Exec(query, {{.expressionValues}})
    _ = m.CacheDel(ctx, data.Id)
    return err
}