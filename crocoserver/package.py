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
        self.docker_compose_content = yaml.load(self.docker_compose_path)

    def install(self):
        self.project.up()

    @property
    def metadata(self):
        return self.docker_compose_content['metadata']

    @property
    def id(self):
        return self.metadata['id']

    @property
    def name(self):
        return self.metadata['name']

    @property
    def description(self):
        return self.metadata['description']

    @property
    def project(self):
        environment = Environment.from_env_file(self.docker_compose_path)
        config_path = get_config_path_from_options(self.docker_compose_path,
                                                   dict(), environment)
        project = get_project(self.docker_compose_path, config_path)
        return project


class AppStore():
    """
    The app store relies on a directory with a docker-compose.yml and a
    README.md file describing the package
    """
    def __init__(self, package_path=PACKAGE_PATH):
        self.apps = _find_packages(package_path)


def _find_packages(package_path):
    for _, _, filenames in os.walk(package_path):
        for filename in fnmatch.filter(filenames, 'docker-compose.yml'):
            yield AppPackage(filename)
