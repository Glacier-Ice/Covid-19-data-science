#coding=utf-8
from flask import Flask,request
from control import readCtl, writeCtl

app: Flask = Flask(__name__)
app.register_blueprint(readCtl)
app.register_blueprint(writeCtl)

if __name__ == '__main__':
    app.run()

