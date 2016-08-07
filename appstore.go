package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"golang.org/x/net/context"

	"github.com/docker/libcompose/config"
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/options"
)

type AppStore struct {
}

type Installer interface {
	Install(appName string)
	Uninstall(appName string)
	IsInstalled(appName string) bool
}

func NewAppStore() AppStore {
	return AppStore{}
}

func (store *AppStore) Install(appName string) {
	app := store.FindApp(appName)
	if app != nil {
		app.Project.Up(context.Background(), options.Up{})
	}
}

func (store *AppStore) IsInstalled(appName string) bool {
	app := store.FindApp(appName)
	if app == nil {
		return false
	}

	containers, _ := app.Project.Containers(context.Background(), project.Filter{})
	return len(containers) > 0
}

func (store *AppStore) Uninstall(appName string) {
	app := store.FindApp(appName)
	if app != nil {
		app.Project.Down(context.Background(), options.Down{})
	}
}

type App struct {
	Project project.APIProject
	Name    string
}

func (store *AppStore) FindApp(appName string) *App {
	for _, app := range store.Apps() {
		if app.Name == appName {
			return &app
		}
	}
	return nil
}

func (store *AppStore) Apps() []App {
	packageDir := "./packages"
	files, _ := ioutil.ReadDir(packageDir)
	apps := make([]App, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			project, err := docker.NewProject(&docker.Context{
				Context: project.Context{
					ComposeFiles: []string{path.Join(packageDir, f.Name(), "docker-compose.yml")},
					ProjectName:  f.Name(),
				},
			}, &config.ParseOptions{
				Interpolate: true,
				Validate:    true,
				Postprocess: func(configs map[string]*config.ServiceConfig) (map[string]*config.ServiceConfig, error) {
					for _, service := range configs {
						if service.Labels == nil {
							service.Labels = map[string]string{}
						}
						service.Labels["traefik.frontend.rule"] = fmt.Sprintf("Host:%s.{domain}", f.Name())
					}
					return configs, nil
				},
			})

			if err == nil {
				apps = append(apps, App{
					Project: project,
					Name:    f.Name(),
				})
			} else {
				log.Println(err)
			}
		}
	}

	return apps
}
