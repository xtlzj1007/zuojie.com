# coding:utf-8
from sqlalchemy import create_engine,Column,String,Integer,SmallInteger
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base 

conn_str = "mysql+pymysql://root:123456@localhost:3306/scrapydb?charset=utf8"
engine = create_engine(conn_str)
Base = declarative_base()
Session = sessionmaker(bind=engine)
session = Session()

class WeatherMysql(Base):
    __tablename__ = "weather"
    id = Column(Integer,primary_key=True, autoincrement=True)
    cityName = Column(String(64))
    cityDate = Column(String(32))
    temperature = Column(SmallInteger)
    weather = Column(String(32))
    temperatureRange = Column(String(32))
    wind = Column(String(128))
    air = Column(String(64))

if __name__ == '__main__':
    Base.metadata.create_all(engine)

