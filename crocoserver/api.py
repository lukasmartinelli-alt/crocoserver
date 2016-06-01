"""API app store"""
from flask import Flask, jsonify, request

import package

API_V1 = '/api/v1/'

store = AppStore()


@app.route(API_V1 + "apps", methods=['GET'])
def list_apps():
    return jsonify(apps=store.apps)
