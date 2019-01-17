import scrapy
import re
import requests
from urllib.parse import urlparse


class TaoHuaSpider(scrapy.Spider):
    name = 'Taohua'
    start_urls = ('http://thzhd.us/')

    def parse(self, response):
        for href in response.css("#newurllink > a::attr(href)"):
            url = href.extract()
            self.logger.info("from url ", url)
            yield scrapy.Request(url, self.parse_main_index)

    def parse_main_index(self, response):
        domain = urlparse(response.url).netloc
        self.logger.info("domain", domain)

