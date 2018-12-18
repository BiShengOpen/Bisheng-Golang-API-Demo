package utils

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
	"strconv"
	"errors"
	"fmt"


	"gopkg.in/yaml.v2"
	"github.com/BishengOpen/Bisheng-Golang-API-Demo/config"

)

// Http GET请求
// strUrl: 请求的URL，strParams: string类型的请求参数, foo=xxx&bar=xxx
func HttpGetRequest(strUrl string, mapParams map[string]string) (string, error) {
	httpClient := &http.Client{}

	var strRequestUrl string
	if nil == mapParams {
		strRequestUrl = strUrl
	} else {
		strParams := Map2UrlQuery(mapParams)
		strRequestUrl = strUrl + "?" + strParams
	}

	// 构建Request
	request, err := http.NewRequest("GET", strRequestUrl, nil)
	if nil != err {
		return err.Error(), err
	}

	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error(), err
	}
	defer response.Body.Close()

	// 解析响应内容
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error(), err
	}

	// 解析StatusCode
	if 200 != response.StatusCode {
		fmt.Println(response.StatusCode)
		return "", errors.New(string(body))
	}

	return string(body), nil
}

// Http POST请求
// strUrl: 请求的URL，mapParams: map类型的请求参数
func HttpPostRequest(strUrl string, mapParams map[string]interface{}) (string, error) {
	httpClient := &http.Client{}

	jsonParams := ""
	if nil != mapParams {
		bytesParams, _ := json.Marshal(mapParams)
		jsonParams = string(bytesParams)
	}

	request, err := http.NewRequest("POST", strUrl, strings.NewReader(jsonParams))
	if nil != err {
		return err.Error(), err
	}
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")

	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if nil != err {
		return err.Error(), err
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error(), err
	}

	// 解析StatusCode
	if 200 != response.StatusCode {
		return "", errors.New(string(body))
	}

	return string(body), nil
}

// 进行签名后的HTTP GET请求
// mapParams: map类型的请求参数, key:value ，strRequest: API路由路径
func ApiKeyGet(mapParams map[string]string, strRequestPath string) (string, error) {
	strMethod := "GET"
	timestamp := time.Now().Unix()

	mapParams["APIKey"] = config.API_KEY
	mapParams["SignatureMethod"] = "secp256k1"
	mapParams["SignatureVersion"] = "1"
	mapParams["Timestamp"] = strconv.FormatInt(timestamp, 10)

	hostName := config.HOST
	strRequestPath = "/" + config.GATEWAY + "/" + config.API_KEY + strRequestPath
	sign, err := ProcessSign(mapParams, strMethod, hostName, strRequestPath, config.SECRET_KEY)
	if err != nil {
		return "", err
	}
	mapParams["Signature"] = sign

	strUrl := config.BISHENG_URL + strRequestPath

	return HttpGetRequest(strUrl, MapValueEncodeURI(mapParams))
}

// 进行签名后的HTTP POST请求
// mapParams: map类型的请求参数, key:value
// strRequest: API路由路径
// return: 请求结果
func ApiKeyPost(mapParams map[string]interface{}, strRequestPath string) (string, error) {
	strMethod := "POST"
	timestamp := time.Now().Unix()

	mapParams2Sign := make(map[string]string)
	mapParams2Sign["APIKey"] = config.API_KEY
	mapParams2Sign["SignatureMethod"] = "secp256k1"
	mapParams2Sign["SignatureVersion"] = "1"
	mapParams2Sign["Timestamp"] =  strconv.FormatInt(timestamp, 10)

	hostName := config.HOST
	strRequestPath = "/" + config.GATEWAY + "/" + config.API_KEY + strRequestPath
	sign, err := ProcessSign(mapParams2Sign, strMethod, hostName, strRequestPath, config.SECRET_KEY)
	if err != nil {
		return "", err
	}
	mapParams2Sign["Signature"] = sign

	strUrl := config.BISHENG_URL + strRequestPath + "?" + Map2UrlQuery(MapValueEncodeURI(mapParams2Sign))
	return HttpPostRequest(strUrl, mapParams)
}

// 构造签名
// mapParams: 送进来参与签名的参数, Map类型
// strMethod: 请求的方法 GET, POST......
// strHostUrl: 请求的主机
// strRequestPath: 请求的路由路径
// strSecretKey: 进行签名的密钥
// 1. 先hmacsha256 hash之后，在ecc签名，返回base64编码的签名
func ProcessSign(mapParams map[string]string, strMethod, strHostUrl, strRequestPath, strSecretKey string) (string, error) {
	// 参数处理, 参数名应按ASCII码进行排序(使用UTF-8编码, 其进行URI编码, 16进制字符必须大写)
	mapCloned := make(map[string]string)
	for key, value := range mapParams {
		mapCloned[key] = url.QueryEscape(value)
	}

	strParams := Map2UrlQueryBySort(mapCloned)
	strPayload := strMethod + "\n" + strHostUrl + "\n" + strRequestPath + "\n" + strParams

	hashedMsg, err := Hmacsha256([]byte(strPayload))
	if err!=nil{
		return "", err
	}
	return  CreatePrivateSign(hashedMsg, strSecretKey)
}

// 构造签名
// mapParams: 送进来参与签名的参数, Map类型
// strMethod: 请求的方法 GET, POST......
// strHostUrl: 请求的主机
// strRequestPath: 请求的路由路径
// strSecretKey: 进行签名的密钥
func CreatePrivateSign(hasedMsg []byte, strSecretKey string) (string, error) {
	signedMsg, err := Sign(strSecretKey, hasedMsg)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signedMsg), nil
}

// 对Map按着ASCII码进行排序
// mapValue: 需要进行排序的map
// return: 排序后的map
func MapSortByKey(mapValue map[string]string) map[string]string {
	var keys []string
	for key := range mapValue {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	mapReturn := make(map[string]string)
	for _, key := range keys {
		mapReturn[key] = mapValue[key]
	}

	return mapReturn
}

// 对Map的值进行URI编码
// mapParams: 需要进行URI编码的map
// return: 编码后的map
func MapValueEncodeURI(mapValue map[string]string) map[string]string {
	for key, value := range mapValue {
		valueEncodeURI := url.QueryEscape(value)
		mapValue[key] = valueEncodeURI
	}

	return mapValue
}

// 将map格式的请求参数转换为字符串格式的
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQuery(mapParams map[string]string) string {
	var strParams string
	for key, value := range mapParams {
		strParams += (key + "=" + value + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// 将map格式的请求参数转换为字符串格式的,并按照Map的key升序排列
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQueryBySort(mapParams map[string]string) string {
	var keys []string
	for key := range mapParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var strParams string
	for _, key := range keys {
		strParams += key + "=" + mapParams[key] + "&"
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

//加载配置文件，config应该是一个pointer
func LoadYamlConfig(config interface{}, fileURL string) error {
	yamlFile, err := ioutil.ReadFile(fileURL)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return err
	}

	return nil
}