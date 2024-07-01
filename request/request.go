package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Lany-w/mingdaoyun-go-sdk/params"
	"github.com/kpango/fastime"
	"github.com/kpango/glg"
)

//var RequestClient *http.Client

func Do(url string, params params.MingDaoRequest) []byte {
	body, _ := json.Marshal(params)
	//fmt.Printf("url:%v, %+v \n", url, string(body))

	/* client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    20,               // 最大空闲连接数
			IdleConnTimeout: 30 * time.Second, // 空闲连接的超时时间
		},
	} */
	location, _ := time.LoadLocation("Asia/Shanghai")

	fastime.SetLocation(location)
	errlog := glg.FileWriter("./error_mingdaoyun.log", 0666)
	defer errlog.Close()
	glg.Get().SetMode(glg.BOTH).AddLevelWriter(glg.ERR, errlog)
	RequestClient := &http.Client{}
	resp, err := RequestClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		glg.Error("MingDaoYun Request Error:" + err.Error())
		glg.Error(string(body))
		panic("POST request failed:" + err.Error())
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)
	_body, _ := ioutil.ReadAll(resp.Body)
	return _body
}
