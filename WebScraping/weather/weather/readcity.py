# coding:utf-8

def readCity():
    with  open("weather/city.txt",encoding="utf8") as f:
        data =  f.readlines()
    return data
    
if __name__ == '__main__':
    print(readCity())
    