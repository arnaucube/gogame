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

token = jsonR["token"]
headers = {"Authorization": "Bearer " + token}

r = requests.get(URL + "/", headers=headers)
t.rStatus("get /", r)
jsonR = r.json()
print(jsonR)

userid = jsonR["user"]["id"]
r = requests.get(URL + "/resources", headers=headers)
t.rStatus("get /resources", r)
jsonR = r.json()
print(jsonR)

time.sleep(1)
r = requests.get(URL + "/resources", headers=headers)
t.rStatus("get /resources", r)
jsonR = r.json()
print(jsonR)

r = requests.get(URL + "/planets", headers=headers)
t.rStatus("post /planets/:userid", r)
jsonR = r.json()
print(jsonR)
print(jsonR["planets"][0])
planetid = jsonR["planets"][0]["id"]

d = {
        "planetid": planetid,
        "building": "metalplant",
}
r = requests.post(URL + "/buildings", json=d, headers=headers)
t.rStatus("post /building/:userid", r)
jsonR = r.json()
print(jsonR)

d = {
        "planetid": planetid,
        "building": "ressearchlab",
}
r = requests.post(URL + "/buildings", json=d, headers=headers)
t.rStatus("post /building/:userid", r)
jsonR = r.json()
print(jsonR)

t.printScores()

