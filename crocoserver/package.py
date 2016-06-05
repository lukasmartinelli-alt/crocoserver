"""
Packaging format for foyapp
"""
import os
import fnmatch
import yaml

from compose.config.environment import Environment
from compose.cli.command import get_project, get_config_path_from_options


PACKAGE_PATH = os.getenv(
    'PACKAGE_PATH',
    os.path.join(os.path.dirname(__file__), 'packages')
)


class AppPackage:
    """
    An app package consists of a docker compose file
    on the local filesystem
    """
    def __init__(self, docker_compose_path):
        self.docker_compose_path = docker_compose_path
        self.docker_compose_content = yaml.load(open(self.docker_compose_path))

    def install(self):
        self.project.up()

    @property
    def id(self):
        return os.path.basename(os.path.dirname(self.docker_compose_path))

    @property
    def description(self):
        return self.metadata['description']

    @property
    def project(self):
        project = get_project(os.path.dirname(self.docker_compose_path))
        return project

    @property
    def is_installed(self):
        containers = self.project.containers(stopped=True)
        return len(containers) > 0

    def serialize(self):
        return {
            "id": self.id,
            "is_installed": self.is_installed
        }


class AppStore():
    """
    The app store relies on a directory tree containing
    docker-compose.yml files and a  README.md file describing the package
    """
    def __init__(self, package_path=PACKAGE_PATH):
        self.apps = {app.id: app for app in _find_packages(package_path)}


def _find_packages(package_path):
    for path, _, filenames in os.walk(package_path):
        for filename in fnmatch.filter(filenames, 'docker-compose.yml'):
            yield AppPackage(os.path.join(path, filename))
