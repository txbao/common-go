
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{.keys}}
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return ret,err
}

//Grom添加
func (m *default{{.upperStartCamelObject}}Model) InsertGorm(data *{{.upperStartCamelObject}}) (*{{.upperStartCamelObject}}, error) {
	data.CreatedAt = time.Now().Local().Unix()
	data.UpdatedAt = time.Now().Local().Unix()
	if err := m.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
//添加-分布式事务专用
func (m *default{{.upperStartCamelObject}}Model) InsertTransaction(ctx context.Context,tx *sql.Tx, data *{{.upperStartCamelObject}}) (sql.Result,error) {
    query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	return tx.Exec(query)
}