package main

import (
	"context"
	"github.com/google/go-cmp/cmp"
	gitserver "github.com/jdvgh/brgo-cd/gitserver"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func Test_SendCloneRepo(t *testing.T) {

	tests := []struct {
		name             string
		givenRequest     *gitserver.CloneRepoRequest
		expectedResponse *gitserver.CloneRepoResponse
		expectedError    bool
	}{
		{
			name: "Service returns sent message",
			givenRequest: &gitserver.CloneRepoRequest{
				RepoUrl:              "https://github.com/jdvgh/sample-manifests.git",
				BrgoGitServerBaseUrl: "tcp://0.0.0.0:50051",
			},
			expectedResponse: &gitserver.CloneRepoResponse{
				RepoUrl:              "https://github.com/jdvgh/sample-manifests.git",
				BrgoGitServerBaseUrl: "tcp://0.0.0.0:50051",
				Result:               "OK",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ctx := context.Background()

			s := Server{}
			// when
			resp, err := s.CloneRepo(ctx, tc.givenRequest)

			// then
			if err == nil && tc.expectedError {
				t.Errorf("Error was nil but expected error.")
				return
			}

			if err != nil {
				t.Errorf("Got Error: %v", err)
				return
			}

			if diff := cmp.Diff(tc.expectedResponse, resp, protocmp.IgnoreFields(resp, "folder"), protocmp.Transform()); diff != "" {
				t.Errorf("got %v, want %v, diff (-want +got) %s", resp, tc.expectedResponse, diff)
			}
		})
	}
}
