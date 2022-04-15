// Copyright 2021 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build nosigv4
// +build nosigv4

package sigv4

import (
	"net/http"
)

// NewSigV4RoundTripper returns an error when a new RoundTripper is created.
func NewSigV4RoundTripper(cfg *SigV4Config, next http.RoundTripper) (http.RoundTripper, error) {
	return nil, ErrSigV4Disabled
}

// SigV4Config is the configuration for signing remote write requests with
// AWS's SigV4 verification process.
type SigV4Config struct {
}

func (c *SigV4Config) Validate() error {
	return ErrSigV4Disabled
}

func (c *SigV4Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return c.Validate()
}
