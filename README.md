### 环境准备

* [ev3dev][1] "严肃"开发者版本固件，直接上debian :)
* 闲置在家的小米随身usb wifi插入USB竟然能用，感谢最新内核

![](/ev3-with-xiaomi-wifi.png)

### 中文TTS
 ev3是有可以播放pcm的喇叭的，自带的Sound python库不支持中文。既然已经可以上网了，那我们就用目前世界第一的中文TTS——讯飞。
 
 * 申请讯飞云账号，开通tts webapi
 * 装python3-pip，然后装各种http依赖包
 
 中途死机2次...难道夏天玩插两个usb+sd卡全开太热了？然后断电重启，装pip花了20分钟，然后pip install也是1多分钟，这cpu是真卡啊。
 
 * [讯飞TTS Py模块](XFOnlineTTS.py)





[1]: https://www.ev3dev.org/
[2]: https://github.com/ev3dev/ev3dev-buildscripts
