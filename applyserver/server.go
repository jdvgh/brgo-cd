package applyserver

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func KustomizeFolder(folderToKustomize string) (string, error) {

	folderToKustomizeContent, err := os.ReadDir(folderToKustomize)
	if err != nil {
		return "", errors.New(fmt.Sprintln("Could not read folderToKustomize", folderToKustomize, "err: ", err))
	}

	log.Println("Content of cloned Repo:")
	for _, fileName := range folderToKustomizeContent {
		fmt.Println(fileName.Name())
	}
	fileName, err := os.CreateTemp(folderToKustomize, "out.yaml")
	if err != nil {
		return fileName.Name(), errors.New(fmt.Sprintln("Could not create TempFile for kustomization out at:", folderToKustomize, "/out.yaml err:", err))
	}
	log.Println("Created tempFile :", fileName.Name())
	_, lookErr := exec.LookPath("kubectl")
	if lookErr != nil {
		return fileName.Name(), errors.New(fmt.Sprintln("Could not find kubectl executable:", lookErr))
	}
	kubectlKustomizeCmd := exec.Command("kubectl", "kustomize", "--output", fileName.Name(), folderToKustomize)
	kubectlKustomizeCmdOut, err := kubectlKustomizeCmd.Output()
	if err != nil {

		return fileName.Name(), errors.New(fmt.Sprintln("Could not kustomize:", err))
	}
	log.Println(kubectlKustomizeCmdOut)

	return fileName.Name(), nil

}

func ApplyFile(fileName, kubeConfigPath string) error {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return errors.New(fmt.Sprintln("Could not read file: ", fileName, " err: ", err))
	}
	log.Println("File content: ", string(fileContent))
	_, lookErr := exec.LookPath("kubectl")
	if lookErr != nil {
		return errors.New(fmt.Sprintln("Could not find kubectl executable:", lookErr))
	}
	kubectlCmd := exec.Command("kubectl", "apply", "-f", fileName)
	kubectlCmd.Env = []string{} 
	kubectlCmd.Env = append(kubectlCmd.Env, fmt.Sprintf("KUBECONFIG=%v", kubeConfigPath))
	kubectlCmdOut, err := kubectlCmd.Output()
	if err != nil {
		return errors.New(fmt.Sprintln("Could not apply: ", fileName, " err: ", err))

	}
	log.Println(string(kubectlCmdOut))
	return nil
}
