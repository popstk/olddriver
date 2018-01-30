import schedule
import time
from scrapyd_api import ScrapydAPI

scrapyd = ScrapydAPI('http://localhost:6800')


def job():
    for p in scrapyd.list_projects():
        for spider in scrapyd.list_spiders(p):
            res = scrapyd.schedule(p, spider)
            print('Start spider[%s.%s]' % (p, spider))

schedule.every().day.at("01:00").do(job)

while True:
    schedule.run_pending()
    time.sleep(1)
