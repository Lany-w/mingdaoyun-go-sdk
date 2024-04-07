package request

import (
	"bytes"
	"encoding/json"
	"github.com/Lany-w/mingdaoyun-go-sdk/params"
	"io/ioutil"
	"net/http"
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

	RequestClient := &http.Client{}
	resp, err := RequestClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic("POST request failed:" + err.Error())
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)
	_body, _ := ioutil.ReadAll(resp.Body)
	return _body
}
