package main

import (
	"learnGoWithTests_ReadingFilesChapter/blogposts"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("../posts/"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
