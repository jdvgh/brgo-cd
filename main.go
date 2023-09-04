package main

import (
	"log"

	"github.com/jdvgh/brgo-cd/gitserver"
)

func main() {
	repoUrl := "https://github.com/jdvgh/brgo-cd.git"
	err := gitserver.CloneRepo(repoUrl)
	if err != nil {
		log.Fatalf("gitserver.CloneRepo(%v) failed", repoUrl)
	}

}
