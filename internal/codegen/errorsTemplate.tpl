{{ range .Errors }}
{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Is{{.CamelValue}}(err error) bool {
	return errgen.IsError(err, {{ .Name }}_{{ .Value }}, {{ .HTTPCode }})
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Error{{ .CamelValue }}(format string, args ...interface{}) *errors.Error {
	return errgen.NewError({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}, format, args...)
}

{{- end }}