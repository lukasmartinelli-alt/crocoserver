package main

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/options"
	"golang.org/x/net/context"
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

func (store *AppStore) FindApp(appName string) *App {
	for _, app := range store.Apps() {
		if app.Name == appName {
			return &app
		}
	}
	return nil
}

func (store *AppStore) Apps() []App {
	packageDir := "./apps"
	files, _ := ioutil.ReadDir(packageDir)
	apps := make([]App, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			app, err := ParseApp(path.Join(packageDir, f.Name()))
			if err == nil {
				apps = append(apps, app)
			} else {
				log.Println(err)
			}
		}
	}

	return apps
}
