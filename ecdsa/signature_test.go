package ecdsa

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

//验证通过私钥和椭圆曲线，hash等生成压缩签名，再验签后返回公钥的过程
func Test_SignCompact(t *testing.T) {
	// 	-----------------------------------------------------------------   数据初始化  -------------------------------------------------------------------------- //
	pubKey, err := ParsePubKey(pub.SerializeCompressed(), S256()) //验证公钥是否有效
	fmt.Println("公钥:", pubKey)
	if err != nil {
		t.Log(err)
		return
	}

	// 	-----------------------------------------------------------------   发送方对数据进行处理  -----------------------------------------------------------------  //
	ciphered, err := Encrypt(pubKey, []byte(testMessage)) //使用公钥对信息进行加密
	fmt.Println("加密后：", ciphered)
	if err != nil {
		t.Log(err)
		return
	}

	testHash := sha256.New()
	testHash.Write([]byte(ciphered))
	hash := testHash.Sum(nil)
	sign, err := SignCompact(S256(), prk, hash, true) //私钥加签
	if err != nil {
		t.Log(err)
		return
	}

	fmt.Printf("序列化后的签名: %x\n", sign)
	fmt.Println("签名长度位数为：", len(sign))

	if err != nil {
		t.Log(err)
		return
	}
	// 	-----------------------------------------------------------------  接收方对数据进行解析 ---------------------------------------------------------------------	//
	pubb, _, err := RecoverCompact(S256(), sign, hash) //公钥验签
	t.Log("还原后的公钥", pubb)
	equal := pubb.IsEqual(pubKey)
	t.Log("对比结果为", equal)

	plaintext, err := Decrypt(prk, ciphered) // 解密数据
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println("原文为：", string(plaintext))
}

var prkY *PrivateKey
var pubY *PublicKey

//验证通过私钥和椭圆曲线，hash等生成压缩签名，再验签后返回公钥的过程
func Test_SignCompact2(t *testing.T) {
	// 	----------------------------------------------------------------- 发送方  数据初始化  -------------------------------------------------------------------------- //
	pubKey, err := ParsePubKey(pub.SerializeCompressed(), S256()) //验证公钥是否有效
	fmt.Println("公钥:", pubKey)
	if err != nil {
		t.Log(err)
		return
	}

	// 	-----------------------------------------------------------------   发送方对数据进行处理  -----------------------------------------------------------------  //
	ciphered, err := Encrypt(pubKey, []byte(testMessage)) //使用公钥对信息进行加密
	fmt.Println("加密后：", ciphered)
	if err != nil {
		t.Log(err)
		return
	}

	testHash := sha256.New()
	testHash.Write([]byte(ciphered))
	hash := testHash.Sum(nil)
	sign, err := SignCompact(S256(), prkY, hash, true) //私钥加签
	if err != nil {
		t.Log(err)
		return
	}

	fmt.Printf("序列化后的签名: %x\n", sign)
	fmt.Println("签名长度位数为：", len(sign))

	if err != nil {
		t.Log(err)
		return
	}
	// 	-----------------------------------------------------------------  接收方对数据进行解析 ---------------------------------------------------------------------	//

	pubb, _, err := RecoverCompact(S256(), sign, hash)         //公钥验签
	pp, err := ParsePubKey(pubb.SerializeCompressed(), S256()) //验证公钥是否有效
	if err != nil {
		fmt.Println("err：", err)
	} else {
		fmt.Println("pp：", pp)
	}
	t.Log("还原后的公钥", pubb)
	equal := pubb.IsEqual(pubY)
	t.Log("对比结果为", equal)
	plaintext, err := Decrypt(prk, ciphered) // 解密数据
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println("原文为：", string(plaintext))
}
func getAddress() string {
	return "矿工地址"
}

//验证通过私钥和椭圆曲线，hash等生成压缩签名，再验签后返回公钥的过程
func Test_SignCompact3(t *testing.T) {
	// 	----------------------------------------------------------------- 管理员发送方  数据初始化  -------------------------------------------------------------------------- //
	address := getAddress() //矿工地址
	hash := sha256.Sum256([]byte(address))
	sign, err := SignCompact(S256(), prkY, hash[:], true) //私钥加签
	if err != nil {
		t.Log(err)
		return
	}

	fmt.Printf("序列化后的签名: %x\n", sign)
	fmt.Println("签名长度位数为：", len(sign))

	if err != nil {
		t.Log(err)
		return
	}
	// 	-----------------------------------------------------------------  接收方对数据进行解析 ---------------------------------------------------------------------	//

	pubb, _, err := RecoverCompact(S256(), sign, hash[:]) //公钥验签
	isEqual := pubb.IsEqual(pubY)
	if !isEqual {
		fmt.Println("错误")
	}
	pp, err := ParsePubKey(pubb.SerializeCompressed(), S256()) //验证公钥是否有效
	if err != nil {
		fmt.Println("err：", err)
	} else {
		fmt.Println("pp：", pp)
	}
	address2 := getAddress() //矿工地址
	hash2 := sha256.Sum256([]byte(address2))
	t.Log("还原后的Hash", hash)
	t.Log("我的Hash", hash2)

}
