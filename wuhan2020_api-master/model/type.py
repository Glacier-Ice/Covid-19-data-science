
class Archive:
    def __init__(self):
        self.province = ""
        self.city = '0',
        self.publish_time = '00:00:00',
        self.publish_date = "0000-00-00",
        self.title = ""
        self.content = "",
        self.link = "",
        self.links_to_pic = "",
        self.announce_type = 0
    
    def toDict(self) -> dict:
        return {
            "province": self.province,
            "publish_time": self.publish_time, 
            "publish_date": self.publish_date,
            "title": self.title,
            "content": self.content,
            "link": self.link,
            "links_to_pic": self.links_to_pic ,
            "announce_type": self.announce_type,
        }

class Response:
    def __init__(self, code = 0, msg = "", data=None):
        self.code = code
        self.message = msg
        self.data = data


def FailResp():
    return Response(code=-1, msg="发生错误")

def SuccessResp(data):
    return Response(code=0, msg="成功", data=data)

def createArchive(data:dict)->Archive:
    result = Archive()
    result.province = data.get("province", "")
    result.publish_time = data.get("publish_time",'00:00:00')
    result.publish_date = data.get("publish_date", "0000-00-00")
    result.title = data.get("title", "")
    result.content = data.get("content", "")
    result.link = data.get("link", "")
    result.links_to_pic = data.get("links_to_pic", "")
    result.announce_type = data.get("announce_type", "")
    return result    