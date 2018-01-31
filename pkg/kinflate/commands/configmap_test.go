/*
Copyright 2017 The Kubernetes Authors.

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

package commands

import (
	"testing"
)

func TestNewAddConfigMapIsNotNil(t *testing.T) {
	if NewCmdAddConfigMap(nil) == nil {
		t.Fatal("NewCmdAddConfigMap shouldn't be nil")
	}
}

func TestAddConfigValidation_NoName(t *testing.T) {
	config := addConfigMap{}

	if config.Validate([]string{}) == nil {
		t.Fatal("Validation should fail if no name is specified")
	}
}

func TestAddConfigValidation_MoreThanOneName(t *testing.T) {
	config := addConfigMap{}

	if config.Validate([]string{"name", "othername"}) == nil {
		t.Fatal("Validation should fail if more than one name is specified")
	}
}

func TestAddConfigValidation_Flags(t *testing.T) {
	tests := []struct {
		name       string
		config     addConfigMap
		shouldFail bool
	}{
		{
			name: "env-file-source and literal are both set",
			config: addConfigMap{
				LiteralSources: []string{"one", "two"},
				EnvFileSource:  "three",
			},
			shouldFail: true,
		},
		{
			name: "env-file-source and from-file are both set",
			config: addConfigMap{
				FileSources:   []string{"one", "two"},
				EnvFileSource: "three",
			},
			shouldFail: true,
		},
		{
			name:       "we don't have any option set",
			config:     addConfigMap{},
			shouldFail: true,
		},
		{
			name: "we have from-file and literal ",
			config: addConfigMap{
				LiteralSources: []string{"one", "two"},
				FileSources:    []string{"three", "four"},
			},
			shouldFail: false,
		},
	}

	for _, test := range tests {
		if test.config.Validate([]string{"name"}) == nil && test.shouldFail {
			t.Fatalf("Validation should fail if %s", test.name)
		} else if test.config.Validate([]string{"name"}) != nil && !test.shouldFail {
			t.Fatalf("Validation should succeed if %s", test.name)
		}
	}
}
