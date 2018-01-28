# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html
import pymongo
import urlparse


class CrawlerPipeline(object):
    def __init__(self):
        self.client = pymongo.MongoClient()
        self.db = self.client['spider']

    def process_item(self, item, spider):
        item['_id'] = urlparse.urlsplit(item['url']).path
        self.db.llss.update({'_id': item['_id']}, {'$set': item}, {'upsert': True})
        return item
