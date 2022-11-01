import requests
from bs4 import BeautifulSoup
import re

def get_source(url):
    response = requests.get(url)
    response.encoding = 'utf-8'
    return response.text

def get_info(source):
    soup = BeautifulSoup(source, 'html5lib')
    ul = soup.find('ul', class_='t clearfix')
    _weather = ul.find_all('p')
    _day= ul.find_all('h1')
    _temperature_i = ul.find_all('i')
    _temperature_h = ul.find_all('span')
    temperature_re = re.compile(r'[0-9]{1,}℃')
    weather_re = re.compile(r'([\u4e00-\u9fa5]{2,})')
    day_re =re.compile(r'([0-9\u4e00-\u9fa5]{2,})')
    weather = weather_re.findall(str(_weather))
    day = day_re.findall(str(_day))
    temperature = temperature_re.findall(str(_temperature_i))
    temperature_h = temperature_re.findall(str(_temperature_h))
    return day,weather,temperature,temperature_h
if __name__ == '__main__':
    url = "http://www.weather.com.cn/weather/101040100.shtml"
    source = get_source(url)
    _day,_weather,_temperature,_temperature_h = get_info(source)
    for i in range(int(len(_day)/2-1)):
        if i == 0:
            print(_day[2*i]+" "+_day[2*i+1]+" "+_weather[4*i+1]+"  "+_temperature[i])
        else:
            print(_day[2*i]+" "+_day[2*i+1]+" "+_weather[4*i+1]+"  "+_temperature[i]+" ~ "+_temperature_h[i])