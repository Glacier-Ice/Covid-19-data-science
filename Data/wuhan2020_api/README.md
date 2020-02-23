# wuhan2020_api
病毒爬虫提交平台



## Table of contents

+ 提交数据
  + [提交结构](#提交结构)
  + [通过python3提交](#通过python3提交)

+ 获取数据
  + [1、通过城市获取](#1、通过城市获取)
  + [2、通过省级获取](#2、通过省级获取)
  + [3、获取提交日志](#3、获取提交日志)



### 提交结构

> 接口POST http://url/api/add

文章

| 字段          | 说明     |
| ------------- | -------- |
| announce_type | 默认 0 |
| city          | 市级名   |
| content       | 主要内容 |
| link          | 原文链接 |
| links_to_pic  | 图片链接 |
| province      | 省名     |
| publish_date  | 发布日期 |
| publish_time  | 发布时间 |
| title         | 标题     |
| uploader      | 提交者   |



### 通过python3提交

> title重复会被判断为已存在

```python
import json
from urllib import request

headers = {'Content-Type': 'application/json'}

postdata = {
        'uploader': '风行',  # 提交用户名， 用于记录
    	'province': "test",
    	'city': 'test',
    	'publish_time': '00:00:00',
    	'publish_date': '0',
    	'title': 'teeeeeeeeeee',
    	'content': "测试测试测试",
    	'link': "http://wsjkw.hebei.gov.cn/content/content_45/395747.whtml",
    	'links_to_pic': '0',
    	'announce_type': 0
}

req = request.Request(
    url="http://url/api/add", 
    headers=headers,
    data=json.dumps(data, ensure_ascii=False).encode("UTF-8")
)

res = request.urlopen(req)
print(res.read().decode("utf-8"))
```





### 获取数据

#### 1、通过城市获取

>GET http://url/api/read?city=<城市名>

例如

GET http://url/api/read?city=test

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "announce_type": "0",
      "city": "test",
      "content": "ttttttttttttttttttttt",
      "link": "http://test.test",
      "links_to_pic": "teeeeeeessst",
      "province": "test",
      "publish_date": "0000-00-00",
      "publish_time": "00:00:00",
      "title": "test"
    }
  ]
}
```

#### 2、通过省级获取

> GET http://url/api/read?province=<省级名>
>
> 例如
>
> GET http://url/api/read?province=test

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "announce_type": "0",
      "city": "test",
      "content": "ttttttttttttttttttttt",
      "link": "http://test.test",
      "links_to_pic": "teeeeeeessst",
      "province": "test",
      "publish_date": "0000-00-00",
      "publish_time": "00:00:00",
      "title": "test"
    }
  ]
}
```



#### 3、获取提交日志

> GET http://url/api/logs?limit=1

```json
{
  "code": 0,
  "message": "成功",
  "data": [
    { 
        "city": "上海", 
        "ip": "120.227.30.249", 
        "province": "上海", 
        "time": "2020.02.10-19:10:44", 
        "uploader": "BeiTown"
    }
  ]
}
```

