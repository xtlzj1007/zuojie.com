# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html


class GetproxyPipeline(object):
    def process_item(self, item, spider):
        fileName = "proxy.csv"
        with open(fileName,"a",encoding="utf-8") as fp:
        #     # for i in range(len(item["ip"])):
        #     #     fp.write(item["ip"][i]+",")
        #     #     fp.write(item["port"][i]+",")
        #     #     fp.write(item["loction"][i]+",")
        #     #     fp.write(item["protocol"][i]+"\n")
        #     # for i in range(len(item["ip"])):
            fp.write(item["ip"]+",")
            fp.write(item["port"]+",")
            fp.write(item["loction"]+",")
            fp.write(item["protocol"]+"\n")
        return item
