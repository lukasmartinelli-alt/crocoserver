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

### Deploy

**crocoserver** is surprisingly easy to install.
All you need is Python and Docker on your Linux machine.

1. Install Python
2. [Install Docker](https://docs.docker.com/linux/step_one/)
3. Install **crocoserver** with `pip install crocoserver`

### Create a new App

Packaging a new application for **crocoserver** and making it available for users
to install with a single click is straightforward.

1. Ensure your application is containerized and provides public Docker images
2. Create a `docker-compose.yml` file and add a new metadata section
3. Create a PR to the **crocoserver** repo

### Package Format

The package format is defined as the standard
[Docker Compose File Reference](https://docs.docker.com/compose/compose-file/).
That's all that you need to provide.
