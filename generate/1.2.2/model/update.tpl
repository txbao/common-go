
func (m *default{{.upperStartCamelObject}}Model) Update(data {{.upperStartCamelObject}}) error {
    return nil
    /*
	{{if .withCache}}{{.keys}}
    _, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		return conn.Exec(query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.Exec(query, {{.expressionValues}}){{end}}
	return err
	*/
}

//gorm修改
func (m *default{{.upperStartCamelObject}}Model) UpdateGorm(data *{{.upperStartCamelObject}}) error {
    data.UpdatedAt = time.Now().Local().Unix()
	err := m.db.Model(&{{.upperStartCamelObject}}{}).
		Where("id = ?", data.Id).
		Updates(data).
		Error
	if err != nil {
		return err
	}
	return nil
}
