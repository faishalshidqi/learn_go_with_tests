package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleHeader       = "Title: "
	descriptionHeader = "Description: "
	TagsHeader        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func readPostBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore "---" line
	buffer := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}

	bodyLine := strings.TrimSuffix(buffer.String(), "\n")
	return bodyLine
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	titleLine := readMetaLine(titleHeader)
	descriptionLine := readMetaLine(descriptionHeader)
	tagsLine := strings.Split(readMetaLine(TagsHeader), ", ")
	bodyLine := readPostBody(scanner)

	post := Post{
		Title:       titleLine,
		Description: descriptionLine,
		Tags:        tagsLine,
		Body:        bodyLine,
	}
	return post, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
