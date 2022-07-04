FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)
FindOneNoCache({{.lowerStartCamelPrimaryKey}} {{.dataType}}) ({{.upperStartCamelObject}}, error)
FindAll(id int64) ([]{{.upperStartCamelObject}}, error)
FindPaging(field string, Limit int, page int) ([]{{.upperStartCamelObject}}, int64, int64, error)