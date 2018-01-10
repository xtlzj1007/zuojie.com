import logging
import getpass
import sys

# 定义MyLog类
class MyLog(object):
    '''
    这个类用于创建一个自用的log
    '''
    def __init__(self):
        user = getpass.getuser()
        self.logger = logging.getLogger(user)
        self.logger.setLevel(logging.DEBUG)
        logFile = './'+sys.argv[0][0:-3] +'.log'
        formatter = logging.Formatter("%(asctime)-12s %(levelname)-8s %(name)-10s %(message)-12s")
        # 日志显现在屏幕上并存储到日志文件中
        logHand = logging.FileHandler(logFile,mode='a',encoding="utf-8")
        logHand.setFormatter(formatter)
        logHand.setLevel(logging.ERROR) # 只有错误才会被记录到日志文件中

        logHandSt = logging.StreamHandler()
        logHandSt.setFormatter(formatter)
        self.logger.addHandler(logHand)
        self.logger.addHandler(logHandSt)

    # 日志的5个级别对应的5个方法
    def debug(self,msg):
        self.logger.debug(msg)

    def info(self,msg):
        self.logger.info(msg)
    
    def warn(self,msg):
        self.logger.warn(msg)

    def error(self,msg):
        self.logger.error(msg)

    def critical(self,msg):
        self.logger.critical(msg)

    
if __name__ == '__main__':
    mylog = MyLog()
    mylog.debug("这是一个bug")
    mylog.info("这是一个info")
    mylog.critical("这是一个critical")
    mylog.error("这是一个错误")