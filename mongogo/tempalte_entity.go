package mongogo

func GetTemplate() string {

	return " package //Your package name " +

		"type {{ .Name }} struct{" +

		"{{ range .Attributes }}" +
		"{{ .Name }} {{ .DisplayType}} `json:\"{{.ToJsonFormat}}\" bson:\"{{.ToJsonFormat}}\"" +
		"{{ end }}" +

		"}"
}
