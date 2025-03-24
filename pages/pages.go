package pages

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type TemplateConstruct struct {
	layouts   map[string]*template.Template
	fractions map[string]*template.Template
}

// TemplateData represents the data used to render any dynamic page.
type TemplateData struct {
	Data map[string]any
}

var tmplConstruct *TemplateConstruct

func RenderLayoutTemplate(w http.ResponseWriter, r *http.Request, templateName string, data TemplateData) {
	// Retrieve the template
	tmpl, ok := tmplConstruct.layouts[templateName]
	if !ok {
		log.Printf("Template not found: %s", templateName)
		http.Error(w, fmt.Sprintf("The template %s does not exist.", templateName), http.StatusInternalServerError)
		return
	}

	// Set content type and execute template
	w.Header().Set("Content-Type", "text/html")
	err := tmpl.ExecuteTemplate(w, "defaultLayout", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

func ServeHomepage(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Data: map[string]any{
			"Title": "GoHTTP",
		}}
	RenderLayoutTemplate(w, r, "homepage", data)
}
