#coding=utf-8


#例子


import requests,json
data = {
            'province': "河北",
            'city': '0',
            'publish_time': '00:00:00',
            'publish_date': '0',
            'title': '0',
            'content': "河北省报告新型冠状病毒感染的肺炎新增确诊病例5例，其中，邯郸市2例，保定市2例，石家庄市1例。邯郸市为报告首例确诊病例。新增疑似病例5例，其中，承德市2例，秦皇岛市1例，廊坊市2例。　　截至1月25日24时，河北省累计报告确诊病例13例，其中死亡1例。确诊病例中，石家庄市5例、保定市3例、沧州市2例、邯郸市2例、承德市1例；死亡病例中，沧州市1例。累计报告疑似病例5例。　　目前追踪到密切接触者222人，均在接受隔离医学观察。　　邯郸市首例确诊病例信息                                                                                                        河北省卫生健康委员会                                                                                                            2020年1月26日\n",
            'link': "http://wsjkw.hebei.gov.cn/content/content_45/395747.whtml",
            'links_to_pic': '0',
            'announce_type': 0,
            'uploader':"1234", #写自己的ID
            'key':"12345" #密钥作为验证权限
        }
headers = {'Content-Type': 'application/json'}
rep = requests.post(headers=headers,data=json.dumps(data),url="http://wuhan.muxxs.com/api/write")
print(rep.text)
