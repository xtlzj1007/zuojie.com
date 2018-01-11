# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html
import time

class WeatherPipeline(object):
    def process_item(self, item, spider):
        today = time.strftime("%F",time.localtime())
        fileName = today + ".txt"
        with open(fileName,"w",encoding="utf-8") as fp:
            fp.write(item["cityName"]+'\t')
            fp.write(item["cityDate"][0]+'\t'+ item["cityDate"][1]+'\t')
            fp.write(item["temperature"]+'\t')
            fp.write(item["weather"]+"\t")
            fp.write(item["temperatureRange"]+'\t')
            fp.write(item["wind"][0] + item["wind"][1] +item["wind"][2]+"\t")
            fp.write(item["air"]+"\n")
        return item
