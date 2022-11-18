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

package clients

import (
	"crypto/tls"
	"net/http"

	"github.com/kirubaCathrin/provider-git/internal/clients/git"
	"github.com/kirubaCathrin/provider-git/internal/clients/rest"
)

// Config provides configuration for the git client
type Config struct {
	Token     string
	BaseURL   string
	TLSConfig *tls.Config
}

// NewClient creates new git Client with provided base URL and credentials
func NewClient(c Config) *rest.Client {
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: c.TLSConfig,
		},
	}
	return &rest.Client{
		Token:      c.Token,
		BaseURL:    c.BaseURL,
		HTTPClient: &httpClient,
	}
}

// NewRepositoryClient creates a new client for the Repository api
func NewRepositoryClient(c Config) git.KeyClientAPI {
	return NewClient(c)
}

// NewAccessKeyClient creates a new client for the access key api
//func NewAccessKeyClient(c Config) git.KeyClientAPI {
	//return NewClient(c)
//}
