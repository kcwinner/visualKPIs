package handlers

import (
  "sync"
  "text/template"
  "net/http"
  "path/filepath"
  "log"
)

type DataTemplateHandler struct {
  Filename string
  once sync.Once
  template *template.Template
}

func (t *DataTemplateHandler) RenderTemplate(w http.ResponseWriter, base string, viewModel interface{}) {
  t.once.Do(func() {
    t.template = template.Must(template.ParseFiles(filepath.Join("templates", base + ".html"),
      filepath.Join("templates", t.Filename)))
    })

  err := t.template.ExecuteTemplate(w, base, viewModel)
  if err != nil {
    log.Println(err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
