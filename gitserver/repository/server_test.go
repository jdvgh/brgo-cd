package repository_test

import (
	"testing"

	"github.com/jdvgh/brgo-cd/gitserver/repository"
)

func TestCloneRepo(t *testing.T) {
	repoUrl := "https://github.com/jdvgh/sample-manifests.git"
	_, err := repository.CloneRepo(repoUrl, true, "")
	if err != nil {
		t.Errorf("gitserver.CloneRepo(%v) failed", repoUrl)
	}
}
