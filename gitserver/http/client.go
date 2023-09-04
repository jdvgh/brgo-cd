package gitserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func SendCloneRepoRequest(repoUrl, brgoGitServerBaseUrl string) (string, error) {
	repoPath := ""
	targetUrl := fmt.Sprintf("%v%v", brgoGitServerBaseUrl, CloneRepoEndpoint)
	cloneBodyRepo := &CloneRepoBody{RepoUrl: repoUrl}
	body, err := json.Marshal(cloneBodyRepo)
	if err != nil {
		return repoPath, errors.New(fmt.Sprintf("Failed to Marshal Json of %v - err : %v", cloneBodyRepo, err))
	}
	requestBody := bytes.NewBuffer(body)
	log.Printf("Sending POST to %v with Body %v", targetUrl, requestBody.String())

	resp, err := http.Post(targetUrl, "application/json", requestBody)
	if err != nil {
		return repoPath, errors.New(fmt.Sprintf("Failed sending POST to %v with Body %v - err: %v", targetUrl, requestBody.String(), err))
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()

	response := &CloneRepoResponse{}
	err = decoder.Decode(response)
	if err != nil {
		return repoPath, errors.New(fmt.Sprintf("Could not parse response : \n %v \n err - %v", resp, err))
	}
	repoPath = response.GitDir
	if !response.Ok {
		return repoPath, errors.New(fmt.Sprintf("Cloning of repo failed in server - response not ok : %v \n", response))
	}
	return repoPath, nil
}
