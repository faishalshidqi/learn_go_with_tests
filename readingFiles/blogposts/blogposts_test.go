package blogposts

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("should return BlogPosts struct in correct amount", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Post 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Post 2
Tags: rust, borrow-checker
---
B
ASDASDASDASDAS
HA`
		)
		fileSystem := fstest.MapFS{
			"hello-world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}
		posts, err := NewPostsFromFS(fileSystem)
		if err != nil {
			t.Fatalf("NewPostsFromFS() failed: %v", err)
		}
		assertEqual(t, len(posts), len(fileSystem))
		assertEqual(t, posts[0], Post{Title: "Post 1", Description: "Post 1", Tags: []string{"tdd", "go"}, Body: "Hello\nWorld"})
		assertEqual(t, posts[1], Post{Title: "Post 2", Description: "Post 2", Tags: []string{"rust", "borrow-checker"}, Body: "B\nASDASDASDASDAS\nHA"})
	})
	t.Run("should return error when failing to read filesystem", func(t *testing.T) {
		_, err := NewPostsFromFS(StubFailingFS{})
		assertError(t, err)
	})
}

func assertEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected: %+v, got: %+v", got, expected)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected error but got none")
	}
}
