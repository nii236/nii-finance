# -*- coding: utf-8 -*-
"""
Spyder Editor

This is a temporary script file.
"""

# pip install websocket-client
# pip install protobuf
# pandas
# numpy



#import pandas.io.data as web
#import websocket
#import numpy as np
#import requests


import test_pb2



from werkzeug.wrappers import Request, Response
from werkzeug.serving import run_simple
#from jsonrpc import JSONRPCResponseManager, dispatcher

import uuid
import requests
import json

#// units ns|us|ms|s|m|h
#// http://127.0.0.1:8081/registry?ttl=10s

payload={"Name": "foo.bar",
          "Nodes": [{
              "Port": 9091,
              "Address": "127.0.0.1",
              "Id": "foo.bar-017da09a-734f-11e5-8136-68a86d0d36b6"
              }]}

def register():
    register_uri = "http://192.168.99.100:8081/registry"
    service = "go.micro.srv.yahoo"
    headers = {'content-type': 'application/json'}
    payload = {
        "name": service,
        "nodes": [{
            "id": service + "-" + str(uuid.uuid4()),
            "address": "192.168.0.7",
            "port": 4000,
        }],
    }
    requests.post(register_uri, data=json.dumps(payload), headers=headers)

@Request.application
def application(request):
    # Dispatcher is dictionary {<method_name>: callable}

    # dispatcher["Say.Hello"] = lambda s: "hello " + s["name"]

    msg = test_pb2.MessageData()
    msg.ParseFromString(request.data)

    #response = JSONRPCResponseManager.handle(
    #    request.data, dispatcher)
    #return Response(response.json, mimetype='application/json')

    return Response(msg.SerializeToString(), mimetype='application/octet-stream')


if __name__ == '__main__':
    print "registering service"
    register()
    print "running service"
    run_simple('localhost', 4000, application)




#r = requests.post('http://192.168.99.100:8081', json=payload)


#r = requests.get('http://192.168.99.100:8081/registry')
#print r.status_code
#print r.headers['content-type']

#r = requests.get('http://192.168.99.100:8081/registry')
#print r.text


#ws = websocket.WebSocket()
#ws.connect("ws://192.168.99.100:8081/broker?topic=go.micro.srv.BitstampRecorder")
#ws.connect("ws://192.168.99.100:8081/broker?topic=go.micro.srv.TickRecorder")
#print ws.recv()

#ws://192.168.0.12:8081/broker?topic=go.micro.srv.TickRecorder


#foo = np.random.randint(0, 10, (4,5,6))

#print foo

#print foo.mean(axis=0)


#DAX = web.DataReader(name='GOOG', data_source='yahoo', start='2000-1-1')
#DAX.info()
#DAX['Close'].plot()

#print DAX.tail()
