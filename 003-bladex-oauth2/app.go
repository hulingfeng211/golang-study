package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	tokenUrl    = "http://169.24.2.82:8107/blade-auth/oauth/token"
	userInfoUrl = "http://169.24.2.82:8107/blade-auth/oauth/user-info"
	noticeUrl   = "http://169.24.2.82:8107/blade-desk/notice/my-notices"
	tenantId    = "000000"
)

func encode(appid, secret string) string {
	encoding := base64.StdEncoding
	//dest := make([]byte)
	return encoding.EncodeToString([]byte(appid + ":" + secret))
	// return string(dest)
}
func passwordEncode(orgin string) string {
	h := md5.New()
	h.Write([]byte(orgin))
	return hex.EncodeToString(h.Sum(nil))
	//return fmt.Sprintf("%x", string(md5.New().Sum([]byte(orgin))))
}

type BladexHeader struct {
	BladeAuth     string
	Authorization string
	TenantId      string
}

type BladexToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    uint   `json:"expires_in"`
	Scope        string `json:"scope"`
	TenantId     string `json:"tenant_id"`
	UserName     string `json:"user_name"`
	RealName     string `json:"real_name"`
	Avatar       string `json:"avatar"`
	ClientId     string `json:"client_id"`
	RoleName     string `json:"role_name"`
	License      string `json:"license"`
	PostId       string `json:"post_id"`
	UserId       string `json:"user_id"`
	RoleId       string `json:"role_id"`
	NickName     string `json:"nick_name"`
	OauthId      string `json:"oauth_id"`
	DeptId       string `json:"dept_id"`
	Account      string `json:"account"`
	Jti          string `json:"jti"`
}

/*
*
 */
func getAccessToken(appId, appSecret, tenantId, username, password string) (BladexToken, error) {
	authorization := "Basic " + encode(appId, appSecret)
	data := make(url.Values)
	data["grant_type"] = []string{"password"}
	data["scope"] = []string{"all"}
	data["username"] = []string{username}
	data["password"] = []string{passwordEncode(password)}
	client := http.Client{}
	var req *http.Request
	req, _ = http.NewRequest("POST", tokenUrl, strings.NewReader(data.Encode()))
	//cookies := &http.Cookie{}
	//cookies.
	fmt.Println(authorization)
	fmt.Println("=================")
	req.Header.Add("Blade-Auth", "")
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Tenant-Id", tenantId)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, _, ret := req.BasicAuth()
	if !ret {
		log.Fatal(ret)
	}
	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	checkError(err)
	restObj := BladexToken{}
	json.Unmarshal(result, &restObj)
	return restObj, err
}

func getNoticeInfo(token, appId, appSecret string) {
	var req *http.Request
	req, _ = http.NewRequest("GET", noticeUrl, nil)
	//cookies := &http.Cookie{}
	//cookies.
	//fmt.Println(authorization)
	//fmt.Println("=================")
	req.Header.Add("Blade-Auth", "bearer "+token)
	//token:=getAccessToken()
	req.Header.Add("Authorization", "Basic "+encode(appId, appSecret))
	req.Header.Add("Tenant-Id", tenantId)
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	checkError(err)
	//restObj := BladexToken{}
	//json.Unmarshal(result, &restObj)
	fmt.Println(string(result))

	//return restObj, err
}

func getUserInfo(token string) {
	var req *http.Request
	req, _ = http.NewRequest("GET", userInfoUrl, nil)
	//cookies := &http.Cookie{}
	//cookies.
	//fmt.Println(authorization)
	//fmt.Println("=================")
	// req.Header.Add("Blade-Auth", "")
	//token:=getAccessToken()
	req.Header.Add("Authorization", "bearer "+token)
	// req.Header.Add("Tenant-Id", tenantId)
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	checkError(err)
	//restObj := BladexToken{}
	//json.Unmarshal(result, &restObj)
	fmt.Println(string(result))
	//return restObj, err
}

func main() {
	// header := BladexHeader{
	// 	BladeAuth:     "Blade-Auth",
	// 	Authorization: "Authorization",
	// 	TenantId:      "Tenant-Id",
	// }
	appId := "sword"
	appSecret := "sword_secret"
	//tenantId := "000000"

	//tokenUrl := "http://169.24.2.82:8107/blade-auth/oauth/token"
	//fmt.Print("Password(md5):")
	//fmt.Println(passwordEncode("admin"))

	//grant_type=password&scope=all&username=admin&password=21232f297a57a5a743894a0e4a801fc3
	//data["tenantId"] = []string{tenantId}
	resultObj, err := getAccessToken(appId, appSecret, tenantId, "admin", "admin")
	checkError(err)
	fmt.Println(resultObj.AccessToken)
	getUserInfo(resultObj.AccessToken)
	getNoticeInfo(resultObj.AccessToken, appId, appSecret)
}

/*
*判空操作,一旦遭遇错误程序就退出
 */
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
