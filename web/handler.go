package web

import (
	"net/http"
	"text/template"

	"github.com/gabrielcerdam-olxautos/goreddit"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewHandler(store goreddit.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}
	h.Use(middleware.Logger)
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
		r.Get("/new", h.ThreadsCreate())
	})
	return h
}

type Handler struct {
	*chi.Mux

	store goreddit.Store
}

const threadsListHTML = `
<h1>Threads</h1>
{{range .Threads}}
	<dt><strong>{{.Title}}</strong></dt>
	<dd><strong>{{.Description}}</strong></dd>
{{end}}
`

func (h *Handler) ThreadsList() http.HandlerFunc {
	type data struct {
		Threads []goreddit.Thread
	}

	tmpl := template.Must(template.New("").Parse(threadsListHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data{Threads: tt})
	}
}

const threadCreate = `
<h1>New Thread</h1>
<form action="/threads" method="POST">
	<table>
		<tr>
		<td>Title</td>
		<td><input type="text" name="title"\></td>
		</tr>
		<tr>
		<td>Description</td>
		<td><input type="text" name="description"\></td>
		</tr>
	</table>
	<button type"submit">Create threads</button>
</form>
`

func (h *Handler) ThreadsCreate() http.HandlerFunc {
	tmpl := template.Must(template.New("").Parse(threadCreate))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}
