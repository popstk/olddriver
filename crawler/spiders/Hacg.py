# -*- coding: utf-8 -*-
import scrapy
import re


def get_magnet_links(result_text):
    result_text = re.sub(r'<[\s\S]*?>', '', result_text)

    result_text = re.sub(
        r'([^0-9a-zA-Z])([0-9a-zA-Z]{5,30})[^0-9a-zA-Z]{5,30}([0-9a-zA-Z]{5,30})([^0-9a-zA-Z])', r'\1\2\3\4', result_text)

    hashes = list(
        set(re.findall(r'[^0-9a-fA-F]([0-9a-fA-F]{40})[^0-9a-fA-F]', result_text)))
    hashes.extend(
        list(set(re.findall(r'[^0-9a-zA-Z]([0-9a-zA-Z]{32})[^0-9a-zA-Z]', result_text))))
    magnets = list(set([hash_value.lower()
                        for hash_value in hashes if not hash_value.lower() in found_magnets]))

    found_magnets.extend(magnets)
    return magnets


def get_dupan_links(result_text):
    return []


class HacgSpider(scrapy.Spider):
    name = 'Hacg'
    allowed_domains = ['www.llss.tv']
    start_urls = ['http://www.llss.tv/']

    def parse(self, response):
        for href in response.css('a::attr(href)'):
            full_url = response.urljoin(href.extract())
            yield scrapy.Request(full_url, callback=self.parse_page)

    def parse_page(self, response):
        yield {
            'title': response.css('title').extract()[0],
            'url': response.url,
            'magnets': get_magnet_links(response.body_as_unicode()),
            'baidu': get_dupan_links(response.body_as_unicode())
        }
