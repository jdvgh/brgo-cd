package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jdvgh/brgo-cd/applyserver"
	gitserver "github.com/jdvgh/brgo-cd/gitserver/http"
)

func main() {
	kubeConfigPath, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		log.Fatalln("Environ KUBECONFIG not set - please set it")
	}
	repoUrl, ok := os.LookupEnv("REPO_URL")
	if !ok {
		log.Fatalln("Environ REPO_URL not set - please set it")
	}
	brgoGitServerBaseUrl, ok := os.LookupEnv("BRGO_GIT_SERVER_BASE_URL")
	if !ok {
		log.Println("Environ BRGO_GIT_SERVER_BASE_URL not set - using default http://0.0.0.0:8080")
		brgoGitServerBaseUrl = "http://0.0.0.0:8080"
	}
	_, err := os.ReadFile(kubeConfigPath)
	if err != nil {
		log.Fatalf("Could not Read kubeconfig at : %v - err: %v ", kubeConfigPath, err)
	}

	gitDir, err := gitserver.SendCloneRepoRequest(repoUrl, brgoGitServerBaseUrl)
	// gitDir, err := repository.CloneRepo(repoUrl, false)
	if err != nil {
		log.Fatalf("repository.CloneRepo(%v) failed - err : %v", repoUrl, err)
	}
	defer os.RemoveAll(gitDir)
	kustomizePath := filepath.Join(gitDir, "k8s", "overlays", "k3s")
	fileName, err := applyserver.KustomizeFolder(kustomizePath)
	if err != nil {
		log.Fatalf("applyserver.Kustomiyefolder(%v) failed - err: %v", kustomizePath, err)
	}
	defer os.Remove(fileName)
	err = applyserver.ApplyFile(fileName, kubeConfigPath)
	if err != nil {

		log.Fatalf("applyserver.ApplyFile(%v,%v) failed - err: %v", fileName, kubeConfigPath, err)
	}
}
