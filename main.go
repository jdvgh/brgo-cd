package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jdvgh/brgo-cd/applyserver"
	"github.com/jdvgh/brgo-cd/gitserver"
)

func main() {
	kubeConfigPath, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		log.Fatalf("Environ KUBECONFIG not set - please set it")
	}
	_, err := os.ReadFile(kubeConfigPath)
	if err != nil {
		log.Fatalf("Could not Read kubeconfig at : %v - err: %v ", kubeConfigPath, err)
	}
	repoUrl := "https://github.com/jdvgh/brgo-cd.git"
	gitDir, err := gitserver.CloneRepo(repoUrl, false)
	if err != nil {
		log.Fatalf("gitserver.CloneRepo(%v) failed - err : %v", repoUrl, err)
	}
	defer os.RemoveAll(gitDir)
	kustomizePath := filepath.Join(gitDir, "k8s", "overlays", "k3s")
	fileName, err := applyserver.KustomizeFolder(kustomizePath)
	if err != nil {
		log.Fatalf("applyserver.Kustomiyefolder(%v) failed - err: %v", kustomizePath, err)
	}
	err = applyserver.ApplyFile(fileName, kubeConfigPath)
	if err != nil {

		log.Fatalf("applyserver.ApplyFile(%v,%v) failed - err: %v", fileName, kubeConfigPath, err)
	}
}
