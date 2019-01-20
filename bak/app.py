#!/usr/bin/env python
# -*- coding=utf-8 -*-
from gevent import monkey
from flask import Flask
from flask_restful import Api
from gevent import pywsgi
from resources.search import Search
from resources.spider import Spider

monkey.patch_all()

app = Flask(__name__)
api = Api(app)

api.add_resource(Search, '/search/<keyword>')
api.add_resource(Spider, '/spider/', '/spider/<name>',
                 resource_class_kwargs={'url': 'http://localhost:6800'})


if __name__ == '__main__':
    server = pywsgi.WSGIServer(('0.0.0.0', 8765), app)
    server.serve_forever()
