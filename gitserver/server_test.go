package gitserver_test

import (
	"testing"

	"github.com/jdvgh/brgo-cd/gitserver"
)

func TestCloneRepo(t *testing.T) {
	repoUrl := "https://github.com/jdvgh/brgo-cd.git"
	err := gitserver.CloneRepo(repoUrl)
	if err != nil {
		t.Errorf("gitserver.CloneRepo(%v) failed", repoUrl)
	}
}
