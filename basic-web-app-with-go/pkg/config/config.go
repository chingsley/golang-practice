package config

import (
	"html/template"
	"log"
)

// AppConfig holds the aplication config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
