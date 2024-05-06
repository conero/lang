# http://www.mca.gov.cn/article/sj/xzqh/2019/201908/201908271607.html 数据获取
# 2019年9月9日 星期一

import requests
import bs4
import json

u = 'http://www.mca.gov.cn/article/sj/xzqh/2019/201908/201908271607.html'
r = requests.get(u)

bs = bs4.BeautifulSoup(r.content, "lxml")
trs = bs.find_all('tr', attrs={'height': 19})

# 否是为汉字
def isHz(c):
    return '\u4e00' <= c <= '\u9fa5'

# 城市
data = []
# 盛世分别代表一二级行政单位
province = ''
city = ''
for tr in trs:
    td = tr.find_all('td')
    code = td[1].text
    name = td[2].text

    # 一级单位无空格
    if isHz(name[0:1]):
        province = name
        city = ""
    else:
        # 二级单位有一个空格
        if isHz(name[1:2]):
            city = name.strip()
        name = name.strip()
    
    # 数据生成
    data.append({
        'name': name,
        'city': city,
        'province': province,
        'code': code
    })
    # print(code+','+name)


# 文件写入
# 改写法，中文被改写为 [\u] 格式
# fw = open('./zxqy.json', 'w', encoding='utf-8')
fw = open('./zxqy.json', 'w')
fw.write(json.dumps(data))
fw.close()
print('文件已写入！')
