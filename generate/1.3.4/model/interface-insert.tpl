Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error)
InsertGorm(data *{{.upperStartCamelObject}}) (*{{.upperStartCamelObject}}, error)
InsertTransaction(ctx context.Context,tx *sql.Tx, data *{{.upperStartCamelObject}}) (sql.Result,error)