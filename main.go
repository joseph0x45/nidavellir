package main

import (
	"embed"
	"html/template"
)

var version = "debug"

//go:embed pico.min.css
var picoCSS template.CSS

//go:embed templates
var templatesFS embed.FS

var templates *template.Template

func init() {
}

func main() {
}
