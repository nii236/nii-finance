# -*- coding: utf-8 -*-
"""
Created on Thu Jun 02 22:36:21 2016

@author: Jacob
"""

import test_pb2
import requests

payload = {"jsonrpc": "2.0",
           "method": "postMessage",
           "params": [{"name": "Jacob"}],
           "id": 99}

msg = test_pb2.MessageData()
msg.id = 1234
msg.name = "John Doe"

url = "http://127.0.0.1:4000"

# POST with protobuf 
r = requests.post(url, data=msg.SerializeToString())

print r.text
print r.headers['content-type']
print r.status_code

