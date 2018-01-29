#!/usr/bin/env python
# -*- coding=utf-8 -*-
from gevent import monkey
from flask import Flask
from jsonrpc.backend.flask import api as flask_api
from search import db
from gevent import pywsgi
from scrapyd_api import ScrapydAPI

monkey.patch_all()

app = Flask(__name__)
app.add_url_rule('/rpc', 'api', flask_api.as_view(), methods=['POST'])
scrapyd = ScrapydAPI('http://45.76.111.18:6800')


@flask_api.dispatcher.add_method
def search(keyword):
    print("Search keyword: " + keyword.encode('utf8'))
    from bson.json_util import dumps
    cursor = db.find(
        {'title': {'$regex': '.*' + keyword + '.*', '$options': 'i'}})
    return json.loads(dumps(list(cursor)))


@flask_api.dispatcher.add_method
def listjobs():
    data = {}
    for p in scrapyd.list_projects(): 
        jobs = scrapyd.list_jobs(p)
        for status in ('running', 'finished', 'pending'):
            for job in jobs[status]:
                spider = "%s.%s" % (p, job.pop('spider'))
                job['status'] = status
                if spider not in data:
                    data[spider] = []
                data[spider].append(job)
    return [{'spider': k, 'jobs': v} for k, v in data.iteritems()]


@flask_api.dispatcher.add_method
def startspider(name):
    p = name.split('.')
    return scrapyd.schedule(p[0], '.'.join(p[1:]))


if __name__ == '__main__':
    server = pywsgi.WSGIServer(('0.0.0.0', 8765), app)
    server.serve_forever()
