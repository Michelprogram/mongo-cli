package mongogo

func GetTemplate() string {

	return "package//Your package name\ntype {{ .Name }} struct{\n{{ range .Attributes }}\n\t{{ .Name }} {{ .DisplayType}} `json:\"{{.ToJsonFormat}}\" bson:\"{{.ToJsonFormat}}\"`\n{{ end }}\n}"
}
