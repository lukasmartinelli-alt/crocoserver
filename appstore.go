package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path"
	"time"

	"golang.org/x/net/context"

	"github.com/boltdb/bolt"
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/options"
)

type BoltInstaller struct {
	db            *bolt.DB
	installBucket []byte
}

type AppStore struct {
	installer Installer
}

func NewBoltInstaller(dbPath string) (*BoltInstaller, error) {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	store := BoltInstaller{db: db, installBucket: []byte("Installations")}
	return &store, err
}

type Installer interface {
	Install(appName string)
	Uninstall(appName string)
	IsInstalled(appName string) bool
}

func (store *BoltInstaller) Install(appName string) {
	_ = store.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(store.installBucket)
		bucket.Put([]byte(appName), []byte{1})
		return err
	})
}

func (store *BoltInstaller) Uninstall(appName string) {
	_ = store.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(store.installBucket)
		bucket.Put([]byte(appName), []byte{0})
		return err
	})
}

func (store *BoltInstaller) IsInstalled(appName string) bool {
	installed := false
	store.db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(store.installBucket)
		if err != nil {
			return err
		}
		state := bucket.Get([]byte(appName))
		if bytes.Equal(state, []byte{1}) {
			installed = true
		}
		return err
	})
	return installed
}

func NewAppStore(dbPath string) AppStore {
	installer, _ := NewBoltInstaller(dbPath)
	return AppStore{
		installer: installer,
	}
}

func (store *AppStore) Install(appName string) {
	app := store.FindApp(appName)
	if app != nil {
		app.Project.Up(context.Background(), options.Up{})
		store.installer.Install(appName)
	}
}

func (store *AppStore) IsInstalled(appName string) bool {
	return store.installer.IsInstalled(appName)
}

func (store *AppStore) Uninstall(appName string) {
	app := store.FindApp(appName)
	if app != nil {
		app.Project.Down(context.Background(), options.Down{})
		store.installer.Uninstall(appName)
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
			}, nil)

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
