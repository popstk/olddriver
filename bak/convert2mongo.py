import pymongo
import json

# wget
# https://raw.githubusercontent.com/Chion82/hello-old-driver/master/resource_list.json

client = pymongo.MongoClient()
db = client['spider']
v = json.loads(open('resource_list.json').read().decode('utf8'))
db.llss.insert_many(v)
