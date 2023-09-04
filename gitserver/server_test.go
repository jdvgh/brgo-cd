package gitserver_test

import (
	"testing"

	"github.com/jdvgh/brgo-cd/gitserver"
)

func TestCloneRepo(t *testing.T) {
	repoUrl := "https://github.com/jdvgh/brgo-cd.git"
	_, err := gitserver.CloneRepo(repoUrl, true)
	if err != nil {
		t.Errorf("gitserver.CloneRepo(%v) failed", repoUrl)
	}
}
