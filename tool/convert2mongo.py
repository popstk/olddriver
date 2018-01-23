import pymongo
import json

client = pymongo.MongoClient()
db = client['spider']
v = json.loads(open('resource_list.json', encoding='utf8').read())
db.llss.insert_many(v)
