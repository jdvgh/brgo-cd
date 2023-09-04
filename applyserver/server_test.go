package applyserver_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jdvgh/brgo-cd/applyserver"
	"github.com/stretchr/testify/assert"
)

func TestKustomizeFolder(t *testing.T) {

	const fileContentKustomizationYaml = ` 
configMapGenerator:
- name: example-configmap-1
  files:
  - application.properties
`
	const fileContentApplicationProperties = `BAR=Foo
FOO=Bar
`
	const expectedFileContent = `apiVersion: v1
data:
  application.properties: |
    BAR=Foo
    FOO=Bar
kind: ConfigMap
metadata:
  name: example-configmap-1-6c79t224df
`
	dirName, err := os.MkdirTemp("", "kustomizeFolder")
	if err != nil {
		t.Errorf("Could not create tmpDir for testing: %v", err)
	}
	defer os.RemoveAll(dirName)
	filePathKustomizationYaml := filepath.Join(dirName, "kustomization.yaml")
	filePathApplicationProperties := filepath.Join(dirName, "application.properties")
	err = os.WriteFile(filePathKustomizationYaml, []byte(fileContentKustomizationYaml), 0644)
	if err != nil {
		t.Errorf("Could not create file %v - err %v", filePathKustomizationYaml, err)
	}
	defer os.Remove(filePathKustomizationYaml)
	err = os.WriteFile(filePathApplicationProperties, []byte(fileContentApplicationProperties), 0644)
	if err != nil {
		t.Errorf("Could not create file %v - err %v", filePathApplicationProperties, err)
	}
	defer os.Remove(filePathApplicationProperties)
	fileName, err := applyserver.KustomizeFolder(dirName)
	if err != nil {
		t.Errorf("Could not kustomize folder %v with files %v and %v - err : %v", dirName, filePathKustomizationYaml, filePathApplicationProperties, err)
	}
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		t.Errorf("Could not read fileContent of %v - err %v", fileName, err)
	}
	assert.Equal(t, expectedFileContent, string(fileContent))
}
