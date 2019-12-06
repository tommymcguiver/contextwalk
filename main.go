package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/builder/dockerignore"
	"github.com/tonistiigi/fsutil"
)

// FIXME: What is err?
func walkFunc(path string, info os.FileInfo, err error) error {
	_, err2 := fmt.Println(path)
	return err2
}

func dockerIgnoreError(err error) {
	if err != nil {
		fmt.Println("Can't find .dockerignore file in $PWD")
		os.Exit(1)
	}
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	f, err := os.Open(path + "/.dockerignore")
	dockerIgnoreError(err)
	ignore, err := dockerignore.ReadAll(f)
	dockerIgnoreError(err)

	fsutil.Walk(
		context.Background(),
		path,
		&fsutil.WalkOpt{
			ExcludePatterns: ignore,
		},
		walkFunc,
	)
}
