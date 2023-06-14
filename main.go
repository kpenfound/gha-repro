package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	baf := client.Host().Directory(".").File("baf.tar.gz")

	c := client.Container().From("alpine").
		WithExec([]string{"apk", "add", "tar"}).
		WithMountedFile("/baf.tar.gz", baf).
		WithExec([]string{"cp", "/baf.tar.gz", "/newtar.tar.gz"})

	_, err = c.Directory("/").File("newtar.tar.gz").Export(ctx, "./exported.tar.gz")
	if err != nil {
		panic(err)
	}
}
