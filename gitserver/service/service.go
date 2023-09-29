package main

import (
	"errors"
	"fmt"
	gitserver "github.com/jdvgh/brgo-cd/gitserver"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/exec"
)

func (s *Server) CloneRepo(ctx context.Context, req *gitserver.CloneRepoRequest) (*gitserver.CloneRepoResponse, error) {
	log.Printf("Received message from client: %s", req)

	folder, err := cloneRepo(req.RepoUrl, true, "")
	res := fmt.Sprint(err)
	if err == nil {
		res = "OK"
	}
    return &gitserver.CloneRepoResponse{RepoUrl: req.RepoUrl, BrgoGitServerBaseUrl: req.BrgoGitServerBaseUrl, Result: res, Folder: folder}, nil
}

func cloneRepo(repoUrl string, deferCleanup bool, tempDirPrefix string) (string, error) {

	dirName, err := os.MkdirTemp(tempDirPrefix, "repoTarget")
	if err != nil {
		return dirName, errors.New(fmt.Sprintln("Could not create tmpDir for cloning:", err))
	}
	if deferCleanup {
		defer os.RemoveAll(dirName)
	}
	log.Println("Created tempDir:", dirName)
	_, lookErr := exec.LookPath("git")
	if lookErr != nil {
		return dirName, errors.New(fmt.Sprintln("Could not find git executable:", lookErr))
	}
	gitCmd := exec.Command("git", "clone", repoUrl, dirName)
	gitCmdOut, err := gitCmd.Output()
	if err != nil {

		return dirName, errors.New(fmt.Sprintln("Could not clone:", err))
	}
	log.Println(gitCmdOut)
	log.Println("Successfully cloned repo:", repoUrl)

	cloneContent, err := os.ReadDir(dirName)
	if err != nil {
		return dirName, errors.New(fmt.Sprintln("Could not read content of cloned repo:", repoUrl, " at :", dirName))
	}
	log.Println("Content of cloned Repo:")
	for _, fileName := range cloneContent {
		log.Println(fileName)
	}
	return dirName, nil
}
