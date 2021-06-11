package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

const (
	DB="odoo14_0513"
	HOST="http://erp.gaowei.com/jsonrpc"
	PORT=80
	USERNAME="admin"
	PASSWORD="admin"

)

type ClientRequest struct {
	JsonRpcVersion string `json:"jsonrpc"`
	Method string `json:"method"`
	Params interface{} `json:"params"`
	Id int `json:"id"`

}
type Response struct {

	JsonRpc string `json:"jsonrpc"`
	Id int `json:"id"`
	Result interface{} `json:"result"`
	Error  string `json:"json"`
}
func randomID() int  {
	max:=1000000000
	//min:=0
	return rand.Intn(max)
}

func call(url string, service string , method string , args... interface{}) Response{
	body:= make(map[string]interface{})
	body["service"]=service
	body["method"]=method
	body["args"]=args
	return json_rpc(url, "call", body)
}
func json_rpc(url string,method string,params interface{}) Response  {
	 request:= ClientRequest{
		JsonRpcVersion: "2.0",
		Method:         method,
		Params:         params,
		Id:             randomID(),
	}
	log.Print(url)
	data,err:=json.Marshal(request)
	if err!=nil {
		log.Fatal(err)
	}
	log.Print(string(data))
	req, _ := http.NewRequest("GET", url, strings.NewReader(string(data)))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res,err:=client.Do(req)
	if err!=nil {

		panic("调用服务失败")
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	print(string(result))
	checkError(err)
	restObj := Response{}
	json.Unmarshal(result, &restObj)
	if restObj.Error !="" {
		panic(restObj.Error)
		//panic("调用服务失败")
	}
	return restObj
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	res := call(HOST, "common", "login", DB, USERNAME, PASSWORD)

	//# create a new note
	args:=make(map[string]interface{})
	args["color"]=8
	args["memo"]="aaaa"
	args["create_uid"]=int(res.Result.(float64))
	//args = {
	//	'color': 8,
	//		'memo': 'This is another note',
	//		'create_uid': int(res.Result.(float64)),
	//}
	res = call(HOST, "object", "execute", DB, int(res.Result.(float64)), PASSWORD, "note.note", "create", args)
	print(res.Result)
	//print("aaaaaaaaaa")
	
}
