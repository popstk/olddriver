#!/usr/bin/env python
# -*- coding=utf-8 -*-
from gevent import monkey
from flask import Flask
from jsonrpc.backend.flask import api as flask_api
from search import db
from gevent import pywsgi
import json

monkey.patch_all()

app = Flask(__name__)
app.add_url_rule('/rpc', 'api', flask_api.as_view(), methods=['POST'])


@flask_api.dispatcher.add_method
def search(keyword):
    print("Search keyword: " + keyword.encode('utf8'))
    from bson.json_util import dumps
    cursor = db.find(
        {'title': {'$regex': '.*' + keyword + '.*', '$options': 'i'}})
    return json.loads(dumps(list(cursor)))


if __name__ == '__main__':
    server = pywsgi.WSGIServer(('0.0.0.0', 8923), app)
    server.serve_forever()
