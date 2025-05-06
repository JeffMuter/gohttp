// pages/pages.go
package pages

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Store all templates in a simple map
var templates = make(map[string]*template.Template)

// InitTemplates loads all templates from the templates directory
func InitTemplates() error {

	// parse all templates
	files, err := filepath.Glob("templates/*.html")
	if err != nil {
		return err
	}

	// parse each template file and store it by name
	for _, file := range files {
		name := filepath.Base(file[:len(file)-len(".html")])
		templates[name], err = template.ParseFiles(file)
		if err != nil {
			return err
		}
	}

	return nil
}

// RenderTemplate renders a template with provided data
func RenderTemplate(w http.ResponseWriter, templateName string, data any) {
	tmpl, ok := templates[templateName]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

func ServeHomepage(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "GoHTTP Homepage",
	}
	RenderTemplate(w, "homepage", data)
}

func ServeForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "GoHTTP Homepage",
	}
	RenderTemplate(w, "form", data)

}
