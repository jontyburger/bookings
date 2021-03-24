package config

import (
	"text/template"

	"github.com/gorilla/sessions"
)

// AppConfig struct to hold the template cache
type AppConfig struct {
	CacheOn       bool
	TemplateCache map[string]*template.Template
	Session       *sessions.CookieStore
	InProduction  bool
}
