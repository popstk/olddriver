#!/usr/bin/env python
# -*- coding=utf-8 -*-
from flask import Flask
from jsonrpc.backend.flask import api as flask_api
from search import db, Query

app = Flask(__name__)
app.add_url_rule('/rpc', 'api', flask_api.as_view(), methods=['POST'])


@flask_api.dispatcher.add_method
def search(keyword):
    print("Search keyword: " + keyword)
    return db.search(Query().title.test(lambda v, t: t in v.lower(), keyword.lower()))


@app.route('/index')
@app.route('/')
def index():
    return app.send_static_file('index.html')

if __name__ == '__main__':
    app.run(host='127.0.0.1', port=5000, debug=True)
