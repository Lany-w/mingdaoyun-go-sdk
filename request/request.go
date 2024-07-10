package request

import (
	"bytes"
	"encoding/json"
	"github.com/Lany-w/mingdaoyun-go-sdk/params"
	"github.com/kpango/fastime"
	"github.com/kpango/glg"
	"github.com/natefinch/lumberjack"
	"io/ioutil"
	"net/http"
	"time"
)

var GlgLogger *glg.Glg

//var RequestClient *http.Client
func init() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	fastime.SetLocation(location)
	GlgLogger = glg.Get().SetMode(glg.BOTH).AddLevelWriter(glg.DEBG, &lumberjack.Logger{
		Filename:   "mingdaoyun.log",
		MaxSize:    64, // megabytes
		MaxBackups: 5,
		MaxAge:     7,     //days
		Compress:   false, // disabled by default
	}).AddLevelWriter(glg.ERR, &lumberjack.Logger{
		Filename:   "error_mingdaoyun.log",
		MaxSize:    64, // megabytes
		MaxBackups: 5,
		MaxAge:     7,     //days
		Compress:   false, // disabled by default
	})
}

func Do(url string, params params.MingDaoRequest) []byte {
	body, _ := json.Marshal(params)
	//GlgLogger.Debug(string(body))
	//fmt.Printf("url:%v, %+v \n", url, string(body))

	/* client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    20,               // 最大空闲连接数
			IdleConnTimeout: 30 * time.Second, // 空闲连接的超时时间
		},
	} */
	RequestClient := &http.Client{}
	resp, err := RequestClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		GlgLogger.Error(string(body))
		GlgLogger.Error(err.Error())
		panic("POST request failed:" + err.Error())
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)
	_body, _ := ioutil.ReadAll(resp.Body)
	//GlgLogger.Debug(string(_body))
	return _body
}
