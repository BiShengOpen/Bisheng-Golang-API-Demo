package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/hmac"
	"crypto/sha256"
)

// Sign 签名
// 参数：
//      priKey ：私钥
//      data   ：待签名的数据
// 返回值：
//      签名后的数据
func Sign(privateKey string, data []byte) ([]byte, error) {
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	sign, err := crypto.Sign(data, priKey)
	if err != nil {
		return nil, err
	}
    return sign, nil
}

// Verify 验证签名
// 参数：
//      sign   ：签名数据
//      data   ：待签名的数据
//      pubKey ：公钥
// 返回值：
//      验证结果
//验证签名
func Verify(sign[]byte, data []byte, publicKey string) (bool, error) {
	pub, err := hex.DecodeString(publicKey)
	if err != nil {
		return false, err
	}
	return crypto.VerifySignature(pub, data, sign[:64]), nil
}

// CreatePubKey 由私钥生成公钥
// 参数：
//      sk   ：16进制私钥
// 返回值：
//      公钥
func CreatePubKey(privateKey string) (string, error) {
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	publicKey := &ecdsa.PublicKey{
		Curve: priKey.Curve,
		X:     priKey.X,
		Y:     priKey.Y,
	}
    pubKey := crypto.CompressPubkey(publicKey)
	//pubKey := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)
	return hex.EncodeToString(pubKey),nil
}

func Hmacsha256(message []byte) ([]byte, error) {
	hmacsha := hmac.New(sha256.New, nil)
	_, err := hmacsha.Write(message)
	if err != nil {
		return nil, err
	}
	result := hmacsha.Sum(nil)
	return result, nil
}