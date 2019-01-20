#!/usr/bin/env python
# -*- coding=utf-8 -*-
from flask_restful import Resource
from scrapyd_api import ScrapydAPI


def listjobs(project, name):
    for status in ('running', 'finished', 'pending'):
        for job in project[status]:
            job['status'] = status
            job['spider'] = name + '.' + job['spider']
            yield job


def reduce_spiders(jobs):
    spiders = {}
    for job in jobs:
        if job['spider'] not in spiders:
            spiders[job['spider']] = [job]
        else:
            spiders[job['spider']].append(job)
    return [reduce(lambda x, y: x if x['start_time'] > y['start_time'] else y, v) for v in spiders.itervalues()]


class Spider(Resource):
    def __init__(self, url):
        self.scrapyd = ScrapydAPI(url)

    def get(self):
        data = []
        for p in self.scrapyd.list_projects():
            spiders = reduce_spiders(listjobs(self.scrapyd.list_jobs(p), p))
            data.extend(spiders)
        return data

    def post(self, name):
        p, s = name.split('.')
        jobs = self.scrapyd.list_jobs(p)
        for job in (jobs['running'] + jobs['pending']):
            if job['spider'] == s:
                return 'Already Running'
        return self.scrapyd.schedule(p, s)


if __name__ == '__main__':
    s = Spider('http://45.76.111.18:6800')
    print(s.get())
