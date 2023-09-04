package gitserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jdvgh/brgo-cd/gitserver/repository"
)

const DefaultHttpPort = "8080"
const CloneRepoEndpoint = "/cloneRepo"
const DefaultTmpDirPrefix = ""

type CloneRepoBody struct {
	RepoUrl string `json:"repo_url"`
}
type CloneRepoResponse struct {
	Ok     bool   `json:"ok"`
	GitDir string `json:"git_dir"`
}
var tmpDirPrefix string
func CreateServer() {
	httpPort, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		httpPort = DefaultHttpPort
	}
	tmpDirPrefix, ok = os.LookupEnv("TMP_DIR_PREFIX")
	if !ok {
		tmpDirPrefix = DefaultTmpDirPrefix
	}
	listeningTarget := fmt.Sprintf(":%v", httpPort)
	http.HandleFunc(CloneRepoEndpoint, cloneRepo)
	log.Println("Listening on :", listeningTarget)
	http.ListenAndServe(listeningTarget, nil)
}

func cloneRepo(w http.ResponseWriter, req *http.Request) {
	log.Printf("Got request\n%v\n", req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	response := &CloneRepoResponse{Ok: false, GitDir: ""}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	cloneRepoBody := &CloneRepoBody{}
	decoder.Decode(&cloneRepoBody)

	repoPath, err := repository.CloneRepo(cloneRepoBody.RepoUrl, false, tmpDirPrefix)
	if err != nil {
		log.Printf("Failed cloning repo %v - err %v", cloneRepoBody.RepoUrl, err)
	} else {
		log.Printf("Successfully cloned repo %v to %v", cloneRepoBody.RepoUrl, repoPath)
		response.Ok = true
	}
	response.GitDir = repoPath
	json.NewEncoder(w).Encode(response)
}
