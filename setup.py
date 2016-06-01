from setuptools import setup, find_packages
import sys

import crocoserver

open_kwds = {}
if sys.version_info > (3,):
    open_kwds['encoding'] = 'utf-8'

with open('README.md', **open_kwds) as f:
    readme = f.read()

setup(
    name='crocoserver',
    version='0.1',
    description="A Docker based app server for non developers",
    long_description=readme,
    classifiers=[],
    keywords='Docker',
    author='Lukas Martinelli',
    author_email='me@lukasmartinelli.ch',
    url='https://github.com/lukasmartinelli/foyapp',
    license='GPL-v3',
    packages=find_packages(exclude=[]),
    include_package_data=True,
    install_requires=['docker-compose==1.7.1'],
    scripts=['bin/crocoserver']
)
