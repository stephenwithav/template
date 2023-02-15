# Purpose 

`text/template` chains methods together, which prevents some useful code patterns from being employed.  This package wraps `text/template` and provides an Options-style interface.

Example from [sqlc](https://github.com/kyleconroy/sqlc)'s internal code:

```go
template.New("table").
    Funcs(funcMap).
	ParseFS(
		template,
		"templates/*.tmpl",
		"templates/*/*.tmpl*",
	)
```

Rewritten in Options style:

```go
template.New("table",
	template.Funcs(funcMap),
	template.ParseFS(
		template,
		"templates/*.tmpl",
		"templates/*/*.tmpl*",
	))
```

Why does this matter?

It improves code reuse, enabling custom functionality like that in [kobra](https://github.com/stephenwithav/kobra).
