/*
Copyright 2016 The Kubernetes Authors.

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

// Package main is the entrypoint for tabled
package main

import (
	"log"

	"github.com/nyrahul/tabled/pkg/json2table"
	"github.com/nyrahul/tabled/pkg/version"
)

func main() {
	log.Printf("version: %s\n", version.Version)
	json2table.Json2Table()
}
