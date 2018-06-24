# -*- coding: utf-8 -*-
import scrapy
import re
import requests
from urlparse import urlparse

splits = [u'本站不提供下载']


def get_magnet_links(result):
    result = re.sub(r'<[\s\S]*?>', '', result)
    # 可能是截断的，先拼起来
    result = re.sub(r'([^0-9a-zA-Z])([0-9a-zA-Z]{5,30})[^0-9a-zA-Z]{5,30}([0-9a-zA-Z]{5,30})([^0-9a-zA-Z])', r'\1\2\3\4', result)
    # 40位和32位的磁力链接
    hashes = re.findall(r'\b([0-9a-fA-F]{40})\b', result)
    hashes.extend(re.findall(r'\b([0-9a-fA-Z]{32})\b', result))
    return [hash_value.lower() for hash_value in hashes]


def get_dupan_links(result):
    pairs = re.findall(ur'\W(1\w{7})(提取|密码|：|\s)+(\w{4}|8酱)\W', result)
    return[p[0] + '#' + p[2] for p in pairs]


class HacgSpider(scrapy.Spider):
    name = 'Hacg'
    start_urls = ['http://acg.gy/']

    def parse(self, response):
        for href in response.css('a::attr(href)'):
            url = href.extract() + "/wp"
            domain = urlparse(url).netloc
            setattr(HacgSpider, 'allowed_domains', [domain])
            self.logger.info("starting from %s", domain)
            yield scrapy.Request(url, self.parse_index)

    def parse_index(self, response):
        for href in response.css('a::attr(href)'):
            full_url = response.urljoin(href.extract())
            if re.match(r'.*/\d+\.html', full_url):
                yield scrapy.Request(full_url, self.parse_page)

        last_page = response.css('#wp_page_numbers ul li a')[-1]
        if last_page.css('::text').extract_first() == '>':
            url = response.urljoin(last_page.css('::attr(href)').extract_first())
            self.logger.info('Next page is %s', url)
            # dont filter this url
            yield scrapy.Request(url, self.parse_index)

    def parse_page(self, response):
        contents = response.css('div[class="entry-content"]').xpath('string(.)').extract()
        contents.extend(response.css('div[class="comment-content"] p').xpath('string(.)').extract())

        magnets = []
        baidu = []
        for content in contents:
            for split in splits:
                content = content.replace(split, '')
            # self.logger.info(content)
            magnets.extend(get_magnet_links(content))
            baidu.extend(get_dupan_links(content))

        yield {
            'title': response.css('title::text').extract_first(),
            'url': response.url,
            'magnets': list(set(magnets)),
            'baidu': list(set(baidu))
        }
