{{if .hasComment}}{{.comment}}
{{end}}{{.method}}(ctx context.Context{{if .hasReq}}, in *{{.pbRequest}}{{end}}, opts ...grpc.CallOption) ({{if .notStream}}*{{.pbResponse}}, {{else}}{{.streamBody}},{{end}} error)