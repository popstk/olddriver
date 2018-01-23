import pymongo
import json

client = pymongo.MongoClient()
db = client['spider']
v = json.loads(open('hacg_backup.json').read().decode('utf8'))
db.llss.insert_many(v)
