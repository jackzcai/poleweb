package render

import (
	"html/template"
	"net/http"

	"github.com/TruenoCB/poleweb/internal/utils"
)

type HTML struct {
	Data       interface{}
	Name       string
	Template   *template.Template
	IsTemplate bool
}
type HTMLRender struct {
	Template *template.Template
}

func (h *HTML) Render(w http.ResponseWriter, code int) error {
	h.WriteContentType(w)
	w.WriteHeader(code)
	if h.IsTemplate {
		err := h.Template.ExecuteTemplate(w, h.Name, h.Data)
		return err
	}
	_, err := w.Write(utils.StringToBytes(h.Data.(string)))
	return err
}

func (h *HTML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, "text/html; charset=utf-8")
}
