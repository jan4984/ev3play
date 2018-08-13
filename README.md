### 环境准备

* [ev3dev][1] "严肃"开发者版本固件，直接上debian :)
* 闲置在家的小米随身usb wifi插入USB竟然能用，感谢最新内核
* py环境是自带的，本来想用js，但node新版本放弃了对armel架构的支持，官方建议用py就用吧

![](/ev3-with-xiaomi-wifi.png =320x180)

### 中文TTS
 ev3是有可以播放pcm的喇叭的，自带的Sound python库不支持中文。既然已经可以上网了，那我们就用目前世界第一的中文TTS——讯飞。
 
 * 申请讯飞云账号，开通tts webapi
 * 装python3-pip，然后装各种http依赖包
 
 中途死机2次...难道夏天玩插两个usb+sd卡全开太热了？然后断电重启，装pip花了20分钟，然后pip install也是1多分钟，这cpu是真卡啊。
 
 * [讯飞TTS Py模块](XFOnlineTTS.py) 。免费应用，每天就500调用量，就不隐appkey了。

### 远程控制web接口

 * 直接利用py动态解析，在客户端发程序上去。结果发现编译爆慢，可能得要5秒钟，很难接受。
 * [web服务器模块](webServer.py)
 * 然后用go重写了一下[马达控制http服务](motor.go)，加上前端[html按键检测控制](controller.html)

 ![](/camera-runner-browser-view.gif)
 ![](/camera-runner-God-view.gif)

设置一个LED并且用讯飞在线TTS播放一句话

```bash
echo 'leds.set_color("LEFT", "AMBER")
tts.play(sound, "我是乐高EV3机器人")' | http http://10.0.0.100:5000/run
```

*ev3 的 cpu实在太卡太卡，python环境下简直考验耐性。 有人尝试过树莓派封装的主控 https://www.dexterindustries.com/shop/brickpi-advanced-for-raspberry-pi/ 吗？ 这个要是功耗、稳定性、兼容性都 ok 的话，应该还挺爽*

[1]: https://www.ev3dev.org/
[2]: https://github.com/ev3dev/ev3dev-buildscripts
