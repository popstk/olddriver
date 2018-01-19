# -*- coding: utf-8 -*-
import scrapy
import re


def get_magnet_links(result_text):
    result_text = re.sub(r'<[\s\S]*?>', '', result_text)
    # 可能是截断的，先拼起来
    result_text = re.sub(r'([^0-9a-zA-Z])([0-9a-zA-Z]{5,30})[^0-9a-zA-Z]{5,30}([0-9a-zA-Z]{5,30})([^0-9a-zA-Z])', r'\1\2\3\4', result_text)

    # 40位和32位的磁力链接
    hashes = list(set(re.findall(r'[^0-9a-fA-F]([0-9a-fA-F]{40})[^0-9a-fA-F]', result_text)))
    hashes.extend(list(set(re.findall(r'[^0-9a-zA-Z]([0-9a-zA-Z]{32})[^0-9a-zA-Z]', result_text))))
    return list(set([hash_value.lower() for hash_value in hashes]))


def get_dupan_links(result_text):
    pairs = re.findall(r'\W(\w{8})\W*(\w{4})\W', result_text)
    return list(set(p[0] + '#' + p[1] for p in pairs))


class HacgSpider(scrapy.Spider):
    name = 'Hacg'
    allowed_domains = ['www.llss.tv']
    start_urls = ['http://www.llss.tv/wp/']

    def parse(self, response):
        for href in response.css('a::attr(href)'):
            full_url = response.urljoin(href.extract())
            yield scrapy.Request(full_url, callback=self.parse_page)


    def parse_page(self, response):
        yield {
            'title': response.css('title::text').extract_first(),
            'url': response.url,
            'magnets': get_magnet_links(response.body_as_unicode()),
            'baidu': get_dupan_links(response.body_as_unicode())
        }
