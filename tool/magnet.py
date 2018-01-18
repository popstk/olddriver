#!/usr/bin/env python
# -*- coding=utf-8 -*-
import requests
import re
import json
import sys
import os
from search import db

cookie = ''
max_depth = 40
viewed_urls = []
found_magnets = []
ignore_url_param = True
ignore_html_label = True

session = requests.Session()
session.headers.update({'Cookie': cookie})


def scan_page(url, depth=0):
    if url in viewed_urls:
        return
    if (depth > max_depth):
        return

    print('Entering: ' + url)
    sys.stdout.flush()

    try:
        result = session.get(url, timeout=60)
        if not (result.status_code >= 400 and result.status_code < 500):
            result.raise_for_status()
        viewed_urls.append(url)
    except Exception:
        scan_page(url, depth)
        return
    result_text = result.content
    magnet_list = get_magnet_links(result_text)
    sub_urls = get_sub_urls(result_text, url)
    page_title = get_page_title(result_text)
    new_resource = {'title': page_title, 'magnets': magnet_list, 'url': url}
    
    if len(magnet_list) > 0 and not db.search(tinydb.Query().title == page_title):
        for magnet in magnet_list:
            print('Found magnet: ' + magnet)
            sys.stdout.flush()
        db.insert(new_resource)
    
    for sub_url in sub_urls:
        scan_page(sub_url, depth + 1)


def get_sub_urls(result_text, url):
    urls = set(re.findall(r'<a.*?href=[\'"](.*?)[\'"].*?>', result_text))
    sub_urls = []
    for sub_url in urls:
        sub_url = sub_url.strip()
        if sub_url == '':
            continue
        if 'javascript:' in sub_url or 'mailto:' in sub_url:
            continue
        if sub_url[0:4] == 'http':
            try:
                if (get_url_prefix(sub_url)[1] != get_url_prefix(url)[1]):
                    continue
            except Exception:
                continue
        elif sub_url[0:1] == '/':
            sub_url = get_url_prefix(url)[0] + '://' + get_url_prefix(url)[1] + sub_url
        else:
            sub_url = url + '/' + sub_url
        sub_url = re.sub(r'#.*$', '', sub_url)
        sub_url = re.sub(r'//$', '/', sub_url)
        if ignore_url_param:
            sub_url = re.sub(r'\?.*$', '', sub_url)
        if sub_url not in viewed_urls:
            sub_urls.append(sub_url)
    return sub_urls


def get_url_prefix(url):
    domain_match = re.search(r'(.*?)://(.*?)/', url)
    if (domain_match):
        return (domain_match.group(1), domain_match.group(2))
    else:
        domain_match = re.search(r'(.*?)://(.*)$', url)
        return (domain_match.group(1), domain_match.group(2))
    
