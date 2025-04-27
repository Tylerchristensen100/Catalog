package helpers

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func LoadTemplates(dirs ...string) (*template.Template, error) {
	tmpl := template.New("")

	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".html" {
				return nil
			}

			relPath, err := filepath.Rel(dirs[0], path) // Use the first dir as the base
			if err != nil {
				return err
			}

			templateName := filepath.ToSlash(relPath)
			templateName = strings.TrimSuffix(templateName, ".html")

			t, err := tmpl.New(templateName).ParseFiles(path)
			if err != nil {
				return err
			}
			tmpl = t

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return tmpl, nil
}
