# crocoserver

A server for people without technical knowledge to host your Open Source webapps.
The **crocoserver** project wants to make it possible for anyone to setup their
own server on a supported cloud providers and install the server apps they like from the
public app store.

Things you can do with **crocoserver**:

- Run your own Dropbox alternative with [Owncloud](crocoserver/packages/owncloud/)
- Run your own wiki with [MediaWiki](https://www.mediawiki.org/wiki/MediaWiki)
- Run your own Slack like chat with [Mattermost](http://www.mattermost.org/)
- Run your own online Latex editor with [ShareLatex](http://sharelatex.com/)
- Run your own web diagram tool with [draw.io](http://draw.io)
- Run your own spreadsheet software with [EtherCalc](https://ethercalc.net/)
- Run your own git server with [Gogs](http://draw.io)
- Run your own web analytics platform with [Piwik](https://piwik.org/)

## For Users

You can choose a local provider to run **corocserver** and use
beautiful and state of the art Open Source apps without any
knowledge about managing a server. It gives control over
your data back to you.

You can install **crocoserver** by visiting http://crocoserver.com and choosing
a local provider.

## For Developers

**crocoserver** is a Python library and server using the Docker Compose API to run
containerized services.  Any `docker-compose.yml` file is a valid **app package**
that can be deployed to the server.

### Install

**crocoserver** is easy to install. All you need is to [install Docker on your Linux machine](https://docs.docker.com/linux/step_one/).


### Create a new App

Packaging a new application for **crocoserver** and making it available for users
to install with a single click is straightforward.

1. Ensure your application is containerized and provides public Docker images
2. Create a `docker-compose.yml` file and add a new metadata section
3. Create a PR to the **crocoserver** repo

### Architecture

The **crocoserver** is using the [libcompose library](https://godoc.org/github.com/docker/libcompose)
to provision the actual containers from a [Compose file](https://docs.docker.com/compose/compose-file/).

The available apps are all in the `apps` directory at the time of build and are shipped with the application.
The apps are divided into two namespaces:
- **user**: All apps that can be installed by a user
- **system**: System level apps that are needed to run crocoserver (like the reverse proxy)

Note that **crocoserver** itself is also an app. It just a system level app that manages other apps
and configures the reverse proxy app.

### How are Apps shipped

As of now apps need to be present in the `apps` directory at build time since they are included as binary
assets in the executable. In the future apps should be pulled directly from GitHub and stored in a directory
so advanced users can make changes.

### App Package Format

*This is experimental. The package format is not stable yet*

The app package format is defined as a folder `apps/<app-id>`containing a
[Docker Compose File ](https://docs.docker.com/compose/compose-file/) called `docker-compose.yml` and an additional `metadata.yml` file containing useful information
about the app package.

You can validate an app package with **crocoserver** CLI.

```
crocoserver validate <package-location>
```

### Build

Package the `apps` and the `gui` as binary assets.

```
go-bindata apps/... gui/...
```

Install the go package.

```
go install
```

