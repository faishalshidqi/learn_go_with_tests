package templating

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
	"strings"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostViewModel(post Post, renderer *PostRenderer) postViewModel {
	viewModel := postViewModel{Post: post}
	viewModel.HTMLBody = template.HTML(markdown.ToHTML([]byte(post.Body), renderer.markdownParser, nil))
	return viewModel
}

type PostRenderer struct {
	templates      *template.Template
	markdownParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	prser := parser.NewWithExtensions(extensions)
	return &PostRenderer{templates: templ, markdownParser: prser}, nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	if err := r.templates.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	if err := r.templates.ExecuteTemplate(w, "blog.gohtml", newPostViewModel(post, r)); err != nil {
		return err
	}

	return nil
}

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
