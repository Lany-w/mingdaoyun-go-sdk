package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Lany-w/mingdaoyun-go-sdk/params"
)

func Do(url string, params params.MingDaoRequest) []byte {
	body, _ := json.Marshal(params)
	//fmt.Printf("url:%v, %+v \n", url, string(body))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic("POST request failed:" + err.Error())
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)
	_body, _ := ioutil.ReadAll(resp.Body)
	return _body
}
