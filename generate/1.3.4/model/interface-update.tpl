Update(ctx context.Context, data *{{.upperStartCamelObject}}) error
UpdateGorm(ctx context.Context, data *{{.upperStartCamelObject}}) error
UpdateTransaction(ctx context.Context,tx *sql.Tx, data *{{.upperStartCamelObject}}) error