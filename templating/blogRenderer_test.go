package templating

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"io"
	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "Hello World",
			Body:        "This is a Post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := NewPostRenderer()
	if err != nil {
		return
	}
	t.Run("it converts a single Post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatalf("rendering failed: %s", err)
		}
		got := buf.String()
		/*
					want := `<h1>Hello World</h1>
			<p>This is a description</p>
			Tags: <ul><li>go</li><li>tdd</li></ul>`
		*/
		//assertEqual(t, want, got)
		approvals.VerifyString(t, got)
	})
	t.Run("it renderes an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "Hello World"}, {Title: "This is another Post"}}
		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatalf("rendering failed: %s", err)
		}
		got := buf.String()
		/*
			want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/this-is-another-post">This is another Post</a></li></ol>`

			assertEqual(t, want, got)
		*/
		approvals.VerifyString(t, got)
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = Post{
			Title:       "Hello World",
			Body:        "This is a Post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}

func assertEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("want: %v, got: %v\n", expected, got)
	}
}
