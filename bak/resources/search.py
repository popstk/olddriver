#!/usr/bin/env python
# -*- coding=utf-8 -*-
from flask_restful import Resource
import pymongo
import json


class Search(Resource):
    def __init__(self):
        client = pymongo.MongoClient()
        self.db = client['spider'].llss

    def get(self, keyword):
        print("Search keyword: " + keyword.encode('utf8'))
        from bson.json_util import dumps
        cursor = self.db.find(
            {'title': {'$regex': '.*' + keyword + '.*', '$options': 'i'}})
        return json.loads(dumps(list(cursor)))
