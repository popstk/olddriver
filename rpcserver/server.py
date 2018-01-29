#!/usr/bin/env python
# -*- coding=utf-8 -*-
from gevent import monkey
from flask import Flask
from jsonrpc.backend.flask import api as flask_api
from search import db
from gevent import pywsgi
from scrapyd_api import ScrapydAPI
import json

monkey.patch_all()

app = Flask(__name__)
app.add_url_rule('/rpc', 'api', flask_api.as_view(), methods=['POST'])
scrapyd = ScrapydAPI('http://localhost:6800')


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
                    data[spider] = {'jobs': [],
                                    'running': False, 'spider': spider}
                data[spider]['jobs'].append(job)
                if status == 'running' or status == 'pending':
                    data[spider]['running'] = True
    return data.values()


@flask_api.dispatcher.add_method
def startspider(name):
    p, s = name.split('.')
    jobs = scrapyd.list_jobs(p)
    for job in (jobs['running'] + jobs['pending']):
        if job['spider'] == s:
            return 'Already Running'
    return scrapyd.schedule(p, s)


if __name__ == '__main__':
    server = pywsgi.WSGIServer(('0.0.0.0', 8765), app)
    server.serve_forever()
