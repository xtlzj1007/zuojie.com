# -*- coding: utf-8 -*-
import scrapy
from todayMoive.items import TodaymoiveItem

class WuhanmoivespiderSpider(scrapy.Spider):
    name = "wuHanMoiveSpider"
    allowed_domains = ["jycinema.com"]
    start_urls = ['http://jycinema.com/']

    def parse(self, response):
        subSelector = response.xpath('//*[@id="hotfilmlist"]')
        items = []
        for sub in subSelector:
            item = TodaymoiveItem()
            item["moiveName"] =sub.xpath("./ul/li[1]/div/div/a/span").extract()
            items.append(item)
        return items