#-*- coding: utf-8 -*-
import requests
import re
import time
import hashlib
import base64
import struct

URL = "http://api.xfyun.cn/v1/service/v1/tts"
AUE = "raw"
APPID = "5b6e7538"
API_KEY = "c48b0054abf83a4ee3043195064c0236"

def _getHeader():
        curTime = str(int(time.time()))
        param = "{\"aue\":\""+AUE+"\",\"auf\":\"audio/L16;rate=16000\",\"voice_name\":\"xiaoyan\",\"engine_type\":\"intp65\"}"
        paramBase64 = base64.b64encode(param.encode('utf-8'))
        m2 = hashlib.md5()
        m2.update((API_KEY + curTime + paramBase64.decode('utf-8')).encode('utf8'))
        checkSum = m2.hexdigest()
        header ={
                'X-CurTime':curTime,
                'X-Param':paramBase64,
                'X-Appid':APPID,
                'X-CheckSum':checkSum,
                'X-Real-Ip':'127.0.0.1',
                'Content-Type':'application/x-www-form-urlencoded; charset=utf-8',
        }
        return header

def _getBody(text):
        data = {'text':text}
        return data
		
def _writeFile(file, content):
    with open(file, 'wb') as f:
        f.write(content)
    f.close()

def play(speaker, text):
	r = requests.post(URL,headers=_getHeader(),data=_getBody(text))
	contentType = r.headers['Content-Type']
	if contentType == "audio/mpeg":
		_writeFile('/tmp/tts.pcm', r.content)
		speaker.play('/tmp/tts.pcm')
	else:
		print(r.text)
