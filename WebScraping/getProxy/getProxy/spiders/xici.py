# -*- coding: utf-8 -*-
import scrapy
from getProxy.items import GetproxyItem

class XiciSpider(scrapy.Spider):
    name = "xici"
    allowed_domains = ["xicidaili.com"]
    start_urls = []
    wds = ["nn","nt","wn","wt"]
    page = 40 
    for type in wds:
        for i in range(1,page+1):
            start_urls.append("http://www.xicidaili.com"+"/"+type+"/"+str(i))
    

    def parse(self, response):
        subSelect = response.xpath('//tr[@class=""]|//tr[@class="odd"]')
        items = []
        for sub in subSelect:
            item = GetproxyItem()
            item["ip"] = sub.xpath(".//td[2]/text()").extract()[0]
            item["port"] = sub.xpath("//td[3]/text()").extract()[0] # port
            if sub.xpath("//td[4]/a/text()"):
                item["loction"] = sub.xpath("//td[4]/a/text()").extract()[0]
            else:
                item["loction"] = sub.xpath("//td[4]/text()").extract()[0] # locaion
            item["protocol"] = sub.xpath("//td[6]/text()").extract()[0]
            items.append(item)
        return items 


