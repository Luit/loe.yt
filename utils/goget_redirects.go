package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/google/go-github/github"
)

const redirectBody = `<meta name="go-import" content="loe.yt/%[1]s git https://github.com/loeyt/%[1]s">
<meta http-equiv="refresh" content="0; url=%[2]s">
`

func main() {
	redirects := make(map[string]string)

	client := github.NewClient(nil)
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "loeyt", nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, repo := range repos {
		if repo.Name == nil {
			log.Fatal("got nil repo.Name")
		}
		if repo.Homepage == nil {
			redirects[*repo.Name] = ""

		} else {
			redirects[*repo.Name] = *repo.Homepage
		}
	}
	f, err := os.OpenFile("./public/_redirects", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for pkg, redirect := range redirects {
		if _, err = fmt.Fprintf(f, "/%[1]s/* /%[1]s 200\n", pkg); err != nil {
			log.Fatal(err)
		}

		otherPath := "https://github.com/loeyt/" + pkg
		if redirect != "" {
			otherPath = redirect
		}
		err = ioutil.WriteFile(path.Join("public", pkg+".html"), []byte(fmt.Sprintf(redirectBody, pkg, otherPath)), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = f.Close(); err != nil {
		log.Fatal(err)
	}
}
