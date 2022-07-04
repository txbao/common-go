
func (m *default{{.upperStartCamelObject}}Model) FindOne({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRow(&resp, {{.cacheKeyVariable}}, func(conn sqlx.SqlConn, v interface{}) error {
		query :=  fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		return conn.QueryRow(v, query, {{.lowerStartCamelPrimaryKey}})
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}query := fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
	var resp {{.upperStartCamelObject}}
	err := m.conn.QueryRow(&resp, query, {{.lowerStartCamelPrimaryKey}})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}

func (m *default{{.upperStartCamelObject}}Model) FindOneNoCache(id int64) ({{.upperStartCamelObject}}, error) {
	var model {{.upperStartCamelObject}}
	err := m.db.Where("id=?", id).
		First(&model).
        Limit(1).
		Error
	return model, err
}

func (m *default{{.upperStartCamelObject}}Model) FindAll(id int64) ([]{{.upperStartCamelObject}}, error) {
	var model []{{.upperStartCamelObject}}
	err := m.db.Where("id > ?", 0).
		Find(&model).
		Order("id ASC").
		Error
	return model, err
}

//分页
func (m *default{{.upperStartCamelObject}}Model) FindPaging(field string, limit int, page int) ([]{{.upperStartCamelObject}}, int64, int64, error) {
	var count int64 = 0
	if field == "" {
		field = "*"
	}
	var res []{{.upperStartCamelObject}}
	db := m.db.Model(&{{.upperStartCamelObject}}{}).
		Where("is_del=?", 0)
	err := db.Limit(limit).
		Order("id desc").
		Offset((page - 1) * limit).
		Scan(&res).
		Limit(-1).Offset(-1).
		Count(&count).
		Error
	//总页面数
    totalPage := int64(math.Ceil(float64(count) / float64(limit)))
    return res, count, totalPage, err
}