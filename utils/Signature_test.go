package utils

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/BishengOpen/Bisheng-Golang-API-Demo/config"

)

func generateParameters() map[string]string {
	params := make(map[string]string)
	params["APIKey"] = config.API_KEY
	params["SignatureMethod"] = "secp256k1"
	params["SignatureVersion"] = "1"
	//params["Timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)

	return params
}

func generateSignature() (string, error) {

	params := generateParameters()

	strRequestPath := fmt.Sprintf("/test1/%s/market", config.API_KEY)

	return ProcessSign(params, "GET", config.HOST, strRequestPath, config.SECRET_KEY)
}


func Test_digest_hash(t *testing.T) {
	expectSignature := "TC+tQ8CTqbaO28MI0QJddmB+F3fQZ80NUrIRpiyKN90Y9dZE62mqKPRo/bYIalgijVil9n5z7if60ZovW7v4ogA="
	signature, err := generateSignature()
	if nil != err {
		t.Log(err)
	}

	assert.Equal(t, signature, expectSignature, "The signatures should be the same")
}


func Benchmark_private_signature_performance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		params := generateParameters()

		hostName := config.HOST
		strRequestPath := fmt.Sprintf("/test1/%s/market", config.API_KEY)
		secretKey := config.SECRET_KEY

		ProcessSign(params, "GET", hostName, strRequestPath, secretKey)
	}
}
