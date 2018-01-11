# -*- coding: utf-8 -*-
import os 
import scrapy
from weather.items import WeatherItem 
from weather.readcity import readCity


class BeijingspiderSpider(scrapy.Spider):
    name = "beiJingSpider"
    allowed_domains = ["www.tianqi.com/"]
    start_urls = []
    citys = readCity()
    for city in citys:
        start_urls.append('http://www.tianqi.com/'+city.strip("\n"))
        

    def parse(self, response):
        sub = response.xpath('//dl[@class="weather_info"]')
        items = []
        item = WeatherItem()
        item["cityName"] = sub.xpath('//dd[@class="name"]/h2/text()').extract()[0]  # CityName
        item["cityDate"] = sub.xpath('//dd[@class="week"]/text()').extract()[0].split("\u3000")   # Date
        item["temperature"] = sub.xpath('//dd[@class="weather"]/p/b/text()').extract()[0] # 气温 
        item["weather"] = sub.xpath('//dd[@class="weather"]/span/b/text()').extract()[0] #天气 
        item["temperatureRange"] = sub.xpath('//dd[@class="weather"]/span/text()').extract()[0] # 温度范围
        item["img"] = sub.xpath('//dd[@class="weather"]/i/img/@src').extract()[0] # img
        item["wind"] = sub.xpath('//dd[@class="shidu"]/b/text()').extract() # 风向
        item["air"] = sub.xpath('//dd[@class="kongqi"]/h5/text()').extract()[0] # 空气质量
        items.append(item) 
        return items 
