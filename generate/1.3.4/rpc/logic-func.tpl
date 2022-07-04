{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}in {{.request}}{{if .stream}},stream {{.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) ({{if .hasReply}}{{.response}},{{end}} error) {
	// todo: add your logic here and delete this line
	//数据验证
        /*
        formStruct := &valid.Index{
            ActivityId: in.ActivityId,
        }
        if err := validators.Valid(formStruct); err != nil {
            return nil, errorrpc.Parameter.WithErrMsg(err.Error()).AddDetail(in)
        }
        */
	
	return {{if .hasReply}}&{{.responseType}}{},{{end}} nil
}
