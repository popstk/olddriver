#!/usr/bin/env python
# -*- coding=utf-8 -*-
from flask_restful import Resource
from scrapyd_api import ScrapydAPI


class Spider(Resource):
    def __init__(self, url):
        self.scrapyd = ScrapydAPI(url)

    def get(self):
        data = {}
        for p in self.scrapyd.list_projects():
            jobs = self.scrapyd.list_jobs(p)
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

    def post(self, name):
        p, s = name.split('.')
        jobs = self.scrapyd.list_jobs(p)
        for job in (jobs['running'] + jobs['pending']):
            if job['spider'] == s:
                return 'Already Running'
        return self.scrapyd.schedule(p, s)
