from tinydb import TinyDB, Query
import json

db = TinyDB('data/hacg.json')
v = json.loads(open('resource_list.json', encoding='utf8').read())
for item in v:
    db.insert(item)
