# coding:utf-8

from scrapy.contrib.downloadermiddleware.useragent import UserAgentMiddleware

class CustomUserAgent(UserAgentMiddleware):
    def process_request(self,request,spider):
        ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
        request.headers.setdefault("User-Agent",ua)
        