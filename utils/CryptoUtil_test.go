package utils

import (
	"testing"
	"encoding/hex"

	"github.com/BishengOpen/Bisheng-Golang-API-Demo/config"
)

func Test_signedAndVerify(t *testing.T){
	msg := "test"
	digest, err := Hmacsha256([]byte(msg))

	signed, err := Sign(config.SECRET_KEY, digest)
	if nil != err {
		t.Error(err)
	} else {
		t.Log(hex.EncodeToString(signed), len(signed))
	}

	pubKey, err := CreatePubKey(config.SECRET_KEY)
	if nil != err {
		t.Error(err)
	}else{
		t.Log("pubKey:", pubKey)
	}

	havePassed, err := Verify(signed, []byte(digest), pubKey)
	if nil != err {
		t.Error(err)
	}else{
		if havePassed {
			t.Log("passed")
		}else{
			t.Log("failed")
		}
	}
}