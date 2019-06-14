#!/usr/bin/env python3
"""Test endpoints for gogame
"""

import json
import requests
import provoj
import time

import os
os.system("mongo gogame --eval 'db.dropDatabase()'")


URL = "http://127.0.0.1:5000"

t = provoj.NewTest("gogame")

registerData = {
    "name": "user00",
    "password": "user00password",
    "email": "user00@email.com",
}
r = requests.post(URL + "/register", json=registerData)
t.rStatus("post /register", r)
jsonR = r.json()
print(jsonR)


loginData = {
    "email": "user00@email.com",
    "password": "user00password",
}
r = requests.post(URL + "/login", json=loginData)
t.rStatus("post /login", r)
jsonR = r.json()
print(jsonR)

userid = jsonR["user"]["id"]
r = requests.get(URL + "/resources/"+ userid)
t.rStatus("get /resources", r)
jsonR = r.json()
print(jsonR)

time.sleep(4)
r = requests.get(URL + "/resources/"+ userid)
t.rStatus("get /resources", r)
jsonR = r.json()
print(jsonR)

t.printScores()

