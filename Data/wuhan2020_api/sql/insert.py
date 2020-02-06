from .conn import connect
from model.type import Archive

connection = connect()


def insert(data: Archive):
    with connection.cursor() as cursor:
        # Create a new record
        sql = "INSERT INTO `archive` \
            (`province`,`city`,`publish_time`,`publish_date`,`title`,`content`,`link`,`links_to_pic`,`announce_type`) \
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)"
        
        cursor.execute(sql, (data.province, data.city, data.publish_time, data.publish_date, data.title, data.content, data.link,data.links_to_pic,data.announce_type))
        connection.commit()