#!flask/bin/python
from flask import Flask, jsonify, request
from flask_cors import CORS
import ev3dev2
import ev3dev2.motor
import ev3dev2.sensor.lego as sensor
import ev3dev2.led
import ev3dev2.sound
import ev3dev2.display
import XFOnlineTTS as tts

app = Flask(__name__)
CORS(app)

default_global = {'dev': ev3dev2, 'motor': ev3dev2.motor, 'sensor': sensor, 'leds': ev3dev2.led.Leds(), 'sound': ev3dev2.sound.Sound(), 'display': ev3dev2.display, 'tts': tts}
my_global = default_global.copy()

@app.route('/clear')
def clear():
    my_global = default_global.copy()

@app.route('/run', methods=['POST'])
def run():
	script = request.data
	exec(script, my_global)
	return 'ok', 200

if __name__ == '__main__':
	app.run(debug=True,host='0.0.0.0')
