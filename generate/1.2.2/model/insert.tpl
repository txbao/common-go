/*
func (m *default{{.upperStartCamelObject}}Model) Insert(data {{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{if .containsIndexCache}}{{.keys}}
    ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		return conn.Exec(query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.ExecNoCache(query, {{.expressionValues}})
	{{end}}{{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.conn.Exec(query, {{.expressionValues}}){{end}}
	return ret,err
}
*/
//添加
func (m *default{{.upperStartCamelObject}}Model) InsertGorm(data *{{.upperStartCamelObject}}) (*{{.upperStartCamelObject}}, error) {
	data.CreatedAt = time.Now().Local().Unix()
	data.UpdatedAt = time.Now().Local().Unix()
	if err := m.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}