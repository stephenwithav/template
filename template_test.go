package template

import (
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTemplate(t *testing.T) {
	tt := []struct {
		desc  string
		given struct {
			name  string
			funcs []Option
			tmpl  string
		}
		want string
	}{
		{
			desc: "",
			given: struct {
				name  string
				funcs []Option
				tmpl  string
			}{
				name: "test",
				funcs: []Option{
					Funcs(FuncMap{"str": strconv.Itoa}),
				},
				tmpl: `Hi {{str 42}}`,
			},
			want: "Hi 42",
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			sb := &strings.Builder{}
			tmpl, err := New(
				tc.given.name,
				append(tc.given.funcs, Parse(tc.given.tmpl))...,
			)
			if err != nil {
				t.Errorf("Template construction error: %s\n", err)
				return
			}
			err = tmpl.Execute(sb, tc.given.tmpl)
			if err != nil {
				t.Errorf("Template execution error: %s\n", err)
				return
			}
			if d := cmp.Diff(tc.want, sb.String()); d != "" {
				t.Errorf("TestTemplate (-want +got):\n%s", d)
			}
		})
	}
}

func TestTemplateMust(t *testing.T) {
	tt := []struct {
		desc  string
		given struct {
			name  string
			funcs []Option
			tmpl  string
		}
		want        string
		expectPanic bool
	}{
		{
			desc: "panic due to splithostport's invalid return type",
			given: struct {
				name  string
				funcs []Option
				tmpl  string
			}{
				name: "test",
				funcs: []Option{
					Funcs(FuncMap{
						"str":           strconv.Itoa,
						"splithostport": net.SplitHostPort,
					}),
				},
				tmpl: `Hi {{str 42}} {{splithostport "yahoo.com"}}`,
			},
			want:        "Hi 42",
			expectPanic: true,
		},
		{
			desc: "panic due to splithostport's invalid return type even when not called",
			given: struct {
				name  string
				funcs []Option
				tmpl  string
			}{
				name: "test",
				funcs: []Option{
					Funcs(FuncMap{
						"str":           strconv.Itoa,
						"splithostport": net.SplitHostPort,
					}),
				},
				tmpl: `Hi {{str 42}}`,
			},
			want:        "Hi 42",
			expectPanic: true,
		},
		{
			desc: "no panic expected",
			given: struct {
				name  string
				funcs []Option
				tmpl  string
			}{
				name: "test",
				funcs: []Option{
					Funcs(FuncMap{
						"str": strconv.Itoa,
					}),
				},
				tmpl: `Hi {{str 42}}`,
			},
			want:        "Hi 42",
			expectPanic: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			defer func(expectPanic bool) {
				r := recover()
				if r != nil && !expectPanic {
					t.Error("A panic occurred unexpectedly")
				}
			}(tc.expectPanic)
			Must(New(
				tc.given.name,
				append(tc.given.funcs, Parse(tc.given.tmpl))...,
			))
		})
	}
}
