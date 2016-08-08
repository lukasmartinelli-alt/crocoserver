package main

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"

	"github.com/docker/libcompose/config"
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
)

// Defines app package format

type App struct {
	Project  project.APIProject
	Name     string
	Metadata AppMetadata
}

type AppMetadata struct {
	Name        string `json:"name"`
	Level       string `json:"level"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl" yaml:"icon"`
	ProjectUrl  string `json:"projectUrl" yaml:"url"`
}

func parseProject(composePath string) (project.APIProject, error) {

	appName := path.Base(path.Dir(composePath))
	assignTraefikRule := func(configs map[string]*config.ServiceConfig) (map[string]*config.ServiceConfig, error) {
		for _, service := range configs {
			if service.Labels == nil {
				service.Labels = map[string]string{}
			}
			service.Labels["traefik.frontend.rule"] = fmt.Sprintf("Host:%s.localhost", appName)
		}
		return configs, nil
	}

	return docker.NewProject(&docker.Context{
		Context: project.Context{
			ComposeFiles: []string{composePath},
			ProjectName:  appName,
		},
	}, &config.ParseOptions{
		Interpolate: true,
		Validate:    true,
		Postprocess: assignTraefikRule,
	})
}

func parseMetadata(metadataPath string) (AppMetadata, error) {
	var metadata AppMetadata
	data, err := ioutil.ReadFile(metadataPath)
	if err != nil {
		return metadata, err
	}

	yaml.Unmarshal(data, &metadata)
	return metadata, err
}

func ParseApp(appDir string) (App, error) {
	name := path.Base(appDir)
	composePath := path.Join(appDir, "docker-compose.yml")
	metadataPath := path.Join(appDir, "metadata.yml")

	project, err := parseProject(composePath)
	if err != nil {
		return App{}, err
	}
	metadata, err := parseMetadata(metadataPath)
	if err != nil {
		return App{}, nil
	}
	return App{
		Name:     name,
		Metadata: metadata,
		Project:  project,
	}, err
}
