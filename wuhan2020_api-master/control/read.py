#coding=utf-8
from flask import Blueprint

read = Blueprint('read', __name__)

@read.route("/api/read", methods=['GET'])
def index():
    return "read"


