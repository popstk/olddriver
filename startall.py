#!/usr/bin/env python
# -*- coding=utf-8 -*-
import os
import sys
import subprocess

services = {
    'db': [
        'docker mongodb',
        'docker-compose -d'
    ],
    'server': [
        'rpcserver',
        'nohup python rpcserver/server.py &'
    ],
    'spider': [
        'scrapyd',
        'scrapyd',
    ]
}


def start_process(names):
    s = []
    for name in names:
        if name.lower() == 'all':
            s.extend(services.keys())
        if name in services:
            s.append(name)

    for name in set(s):
        msg, cmd = services[name]
        print('start ' + msg)
        subprocess.call(cmd)

if __name__ == '__main__':
    if len(sys.argv) > 0:
        start_process(sys.argv)
