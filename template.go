package template

import (
	"io/fs"
	texttemplate "text/template"
)

type Option func(*texttemplate.Template) error
type FuncMap texttemplate.FuncMap

func New(name string, options ...Option) (*texttemplate.Template, error) {
	tmpl := texttemplate.New(name)

	for _, option := range options {
		if err := option(tmpl); err != nil {
			return nil, err
		}
	}

	return tmpl, nil
}

func Funcs(funcMap FuncMap) Option {
	return func(t *texttemplate.Template) error {
		t.Funcs(texttemplate.FuncMap(funcMap))
		return nil
	}
}

func ParseFS(fsys fs.FS, patterns ...string) Option {
	return func(t *texttemplate.Template) error {
		_, err := t.ParseFS(fsys, patterns...)
		return err
	}
}

func ParseFiles(filenames ...string) Option {
	return func(t *texttemplate.Template) error {
		_, err := t.ParseFiles(filenames...)
		return err
	}
}

func ParseGlob(pattern string) Option {
	return func(t *texttemplate.Template) error {
		_, err := t.ParseGlob(pattern)
		return err
	}
}

func Parse(text string) Option {
	return func(t *texttemplate.Template) error {
		_, err := t.Parse(text)
		return err
	}
}
