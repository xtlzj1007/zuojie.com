# coding:utf-8

from  myLog  import MyLog

mylog = MyLog()


def testMylog():
    try:
        a = 1/0 
    except ZeroDivisionError as e :
        mylog.error("def testMylog "+str(e))
    

if __name__ == '__main__':
    testMylog()

    

