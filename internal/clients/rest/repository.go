/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kirubaCathrin/provider-git/internal/clients/git"
)

func (c *Client) CreateRepository(ctx context.Context, repo git.Repo, key git.AccessKey) (git.AccessKey, error) {
	payload := UploadKeyPayload{
		Key: PublicSSHKey{
			Text:  key.Key,
			Label: key.Label,
		},
		Permission: key.Permission,
	}

	marshalledPayload, err := json.Marshal(payload)
	if err != nil {
		return git.AccessKey{}, err
	}

	url := c.BaseURL + fmt.Sprintf("/rest/keys/1.0/projects/%s/repos/%s/ssh",
		url.PathEscape(repo.ProjectKey), url.PathEscape(repo.Repo))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshalledPayload))
	if err != nil {
		return git.AccessKey{}, err
	}

	var response KeyDescription
	if err := c.sendRequest(req, &response); err != nil {
		return git.AccessKey{}, err
	}
	return git.AccessKey{
		ID:         response.Key.ID,
		Key:        response.Key.Text,
		Label:      response.Key.Label,
		Permission: response.Permission,
	}, nil
}

// PublicSSHKey represents the public ssh key
type PublicSSHKey struct {
	// Text contains the public key
	Text string `json:"text"`
	// Labels describes the public key
	Label string `json:"label"`
}

// UploadKeyPayload defines api object for key upload
type UploadKeyPayload struct {
	// Key defines the type of public ssh key
	Key PublicSSHKey `json:"key"`
	// Permissions defines the access level for the access key in git server
	Permission string `json:"permission"`
}

// KeyDescription describes a specific accesskey in git server
type KeyDescription struct {
	// Key contains info about the access key
	Key KeyInfo `json:"key"`
	// Repository contains information about the repository where the access key is added
	Repository RepositoryInfo `json:"repository"`
	// Permission is the level of permission the access key has been granted
	Permission string `json:"permission"`
}

// KeyInfo contains the information about the access key
type KeyInfo struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Label string `json:"label"`
}

// RepositoryInfo contains information about the repository
type RepositoryInfo struct {
	Name    string `json:"name"`
	ID      int    `json:"id"`
	Project ProjectInfo
}

// ProjectInfo contains information on the project
type ProjectInfo struct {
	Key string `json:"key"`
}
