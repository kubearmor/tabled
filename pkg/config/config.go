/*
Copyright 2023 Rahul Jadhav

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

// Package config is the entrypoint for tabled
package config

import (
	"io/ioutil"
	"log"

	"sigs.k8s.io/yaml"
)

type TableConfig struct {
	Title   string `yaml:"title"`
	Caption string `yaml:"caption"`
}

type ColConfig struct {
	Name  string   `yaml:"name"`
	Spec  string   `yaml:"spec"`
	Color []string `yaml:"color"`
}

type YamlConfig struct {
	Table   TableConfig `yaml:"table"`
	Columns []ColConfig `yaml:"columns"`
}

type Config struct {
	InFile  string
	YamlCfg YamlConfig
}

func LoadYAMLConfig(file string) YamlConfig {
	if file == "" {
		return YamlConfig{}
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed reading from yaml config <%s> error: %v", file, err)
	}
	log.Printf("reading from yaml config <%s>", file)
	y := YamlConfig{}
	err = yaml.Unmarshal([]byte(data), &y)
	if err != nil {
		log.Fatalf("failed to read yaml config. error: %v", err)
	}
	//	fmt.Printf("%+v\n", y)
	return y
}
