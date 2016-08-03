package handlers

import (
  "sync"
  "text/template"
  "net/http"
  "path/filepath"
)

type TemplateHandler struct {
  Filename string
  UseBase bool
  once sync.Once
  template *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if t.UseBase {
    t.once.Do(func() {
      t.template = template.Must(template.ParseFiles(filepath.Join("templates", "base.html"),
        filepath.Join("templates", t.Filename)))
    })

    t.template.ExecuteTemplate(w, "base", r)
  } else {
    t.once.Do(func() {
      t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.Filename)))
    })

    t.template.Execute(w, r)
  }
}
