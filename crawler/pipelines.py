# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html
import json
import pymongo


class CrawlerPipeline(object):
    def __init__(self):
        self.client = pymongo.MongoClient()
        self.db = self.lient['spider']

    def process_item(self, item, spider):
        if item['magnets'] or item['baidu']:
            self.db.llss.insert_one(item)
        return item
