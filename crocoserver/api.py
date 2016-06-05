"""API app store"""
from flask import Flask, jsonify
from . import package

API_V1 = '/api/v1/'

app = Flask(__name__, static_url_path='')
store = package.AppStore()


@app.route(API_V1 + "apps", methods=['GET'])
def list_apps():
    return jsonify(apps=[app.serialize() for app in store.apps.values()])


@app.route(API_V1 + "apps/<app_id>", methods=['PUT'])
def install_app(app_id):
    app = store.apps[app_id]
    app.install()
    return jsonify(app.serialize())
