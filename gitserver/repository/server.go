package repository

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func CloneRepo(repoUrl string, deferCleanup bool, tempDirPrefix string) (string, error) {

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
