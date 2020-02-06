import pymysql.cursors

# Connect to the database

def connect()->pymysql.Connection:
   return pymysql.connect(
        host='localhost',
        user='user',
        password='passwd',
        db='db',
        charset='utf8mb4',
        cursorclass=pymysql.cursors.DictCursor
    )