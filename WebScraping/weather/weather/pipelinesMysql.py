# coding:utf-8

from .models import session,WeatherMysql

class WeatherPipeline(object):
    def process_item(self, item, spider):
        data = {
        "cityName":item["cityName"],
        "cityDate":item["cityDate"][0]+item["cityDate"][1],
        "temperature":item["temperature"],
        "weather":item["weather"],
        "temperatureRange":item["temperatureRange"],
        "wind":item["wind"][0]+item["wind"][1]+item["wind"][2],
        "air":item["air"],
        }
        # print(data)
        weather = WeatherMysql(**data)
        session.add(weather)
        # 提交即保存到数据库:
        session.commit()
        return item 