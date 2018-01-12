# -*- coding: utf-8 -*-
import scrapy
from getProxy.items import GetproxyItem

class A31fSpider(scrapy.Spider):
    name = "31f"
    allowed_domains = ["31f.cn"]
    start_urls = ['http://31f.cn/']

    def parse(self, response):
        data = response.xpath('/*')
        items = []
        for d in data:
            item = GetproxyItem()
            item["ip"] = d.xpath("//td[2]/text()").extract() # IP
            item["port"] = d.xpath("//td[3]/text()").extract() # port
            item["loction"] = d.xpath("//td[4]/a/text()").extract() # locaion
            item["protocol"] = d.xpath("//td[5]/a/text()").extract()
            items.append(item)
        return  items