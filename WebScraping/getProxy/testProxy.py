import re
import threading 
import csv 
from urllib import request


class TestProxy(object):
    def __init__(self):
        self.sFile = "proxy.csv"
        self.dFile = "alive.csv"
        self.URL = "https://www.baidu.com"
        self.threading = 40
        self.timeout = 3
        self.regex = re.compile("baidu")
        self.aliveList = []

        self.run()

    def run(self):
        with open(self.sFile,"r",encoding="utf-8") as fp:
            lines = fp.readlines()
            line = lines.pop()
            while lines:
                for i in range(self.threading):
                    t = threading.Thread(target=self.LinkWithProxy,args=(line,))
                    t.start()
                    if lines:
                        line = lines.pop()
                    else:
                        continue 
        with open(self.dFile,"w",encoding="utf-8") as fp:
            for i in range(len(self.aliveList)):
                fp.write(self.aliveList[i])
        
    def LinkWithProxy(self,line):
        lineList = line.split(",")
        # print(lineList)
        protocol = lineList[-1].strip("\n").lower()
        server = protocol + r"://" + lineList[0] + ":" + lineList[1]
        # print(server)
        opener = request.build_opener(request.ProxyHandler({protocol:server}))
        request.install_opener(opener)
        try:
            response = request.urlopen(self.URL,timeout=self.timeout)
        except:
            print("%s connect failed" %server)
        else:
            try:
                string = str(response.read())
                # print(string)
            except:
                print("%s connect failed" % server)
                return
            if self.regex.search(string):
                print("%s connect sucess ............" % server)
                self.aliveList.append(line)


if __name__ == '__main__':
    t = TestProxy()
    
    