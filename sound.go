package ev3play

import (
	"time"
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"net/http"
	"net/url"
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"
	"os/exec"
	"github.com/gorilla/mux"
)

const(
	appkey="c48b0054abf83a4ee3043195064c0236"
	appid="5b6e7538"
)

func RegisterSoundHandlers(router *mux.Router){
	router.HandleFunc("/text/{value}", playText).Methods("GET")
	router.HandleFunc("/body", playData).Methods("POST")
}


func playData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data,err := ioutil.ReadAll(r.Body)
	if err != nil {
		ioutil.WriteFile("/tmp/audio.wav", data, os.ModePerm)
		exec.Command("/usr/bin/aplay", "-q", "/tmp/audio.wav").Run()
	}
}

func playText(w http.ResponseWriter, r *http.Request) {
	text := mux.Vars(r)["value"]
	curTime := time.Now().Unix()
	param := "{\"aue\":\"raw\",\"auf\":\"audio/L16;rate=16000\",\"voice_name\":\"xiaoyan\",\"engine_type\":\"intp65\"}"
	param = base64.URLEncoding.EncodeToString([]byte(param));
	sum := md5.Sum([]byte(fmt.Sprintf("%s%d%s", appkey, curTime, param)))
	sumStr := hex.EncodeToString(sum[:])
	bodyValues:= url.Values{}
	bodyValues.Add("text", text);
	bodyStr := bodyValues.Encode()
	req,err := http.NewRequest("POST", "http://api.xfyun.cn/v1/service/v1/tts", bytes.NewBuffer([]byte(bodyStr)))
	if err!=nil{
		http.Error(w, "upstream error:" + err.Error(), 500)
		return
	}

	req.Header.Add("X-CurTime", fmt.Sprintf("%d",curTime))
	req.Header.Add("X-Param", param)
	req.Header.Add("X-Appid", appid)
	req.Header.Add("X-CheckSum", sumStr)
	req.Header.Add("X-Real-Ip", "127.0.0.1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "upstream error:" + err.Error(), 500)
		return
	}

	if 200<=rsp.StatusCode && rsp.StatusCode < 300{
		defer rsp.Body.Close()
		audio,err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			http.Error(w, "upstream error:" + err.Error(), 500)
			return
		}
		ioutil.WriteFile("/tmp/audio.wav", audio, os.ModePerm)
		exec.Command("/usr/bin/aplay", "-q", "/tmp/audio.wav").Run()
		return
	}
	http.Error(w, "upstream error:" + fmt.Sprintf("status code %d", rsp.StatusCode), 500)
}
