Update(ctx context.Context, session sqlx.Session, newData *{{.upperStartCamelObject}}) error
UpdateGorm(ctx context.Context, data *{{.upperStartCamelObject}}) error
UpdateGormMap(ctx context.Context, mp map[string]interface{}, id int64) error
UpdateTransaction(ctx context.Context,tx *sql.Tx, newData *{{.upperStartCamelObject}}) error
CacheDel(ctx context.Context, id int64) error
Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error