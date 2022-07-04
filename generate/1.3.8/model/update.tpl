//修改
func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, session sqlx.Session, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		if session != nil {
		    return session.ExecCtx(ctx, query, {{.expressionValues}})
		}
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
//gorm修改
func (m *default{{.upperStartCamelObject}}Model) UpdateGormMap(ctx context.Context, mp map[string]interface{}, id int64) error {
	mp["updated_at"] = time.Now().Local().Unix()
	err := m.db.Model(&{{.upperStartCamelObject}}{}).
		Where("id = ?", id).
		Updates(mp).
		Error
	if err != nil {
		return err
	}
	_ = m.CacheDel(ctx, id)
	return nil
}
//修改-分布式事务专用
func (m *default{{.upperStartCamelObject}}Model) UpdateTransaction(ctx context.Context,tx *sql.Tx, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _, err := tx.Exec(query, {{.expressionValues}})
    _ = m.CacheDel(ctx, {{if .containsIndexCache}}newData{{else}}data{{end}}.Id)
    return err
}

//缓存删除
func (m *default{{.upperStartCamelObject}}Model) CacheDel(ctx context.Context, {{.upperStartCamelPrimaryKey}} int64) error {
	{{if .withCache}}

data, err:=m.FindOne(ctx, {{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}
	{{if .containsIndexCache}}

{{end}}	{{.keys}}
	if err := m.CachedConn.DelCacheCtx(ctx,  {{.keyValues}}); err != nil {
    		logx.Error("DelCacheCtxErr:删除Model缓存失败：", err.Error())
    		return err
    	}
	{{else}}
    if err := m.CachedConn.DelCacheCtx(ctx,  {{.keyValues}}); err != nil {
        		logx.Error("DelCacheCtxErr:删除Model缓存失败：", err.Error())
        		return err
        	}
    {{end}}
	return err
}

//logic开启事务
func (m *default{{.upperStartCamelObject}}Model) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}