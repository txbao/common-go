package utils

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/big"
)

//https://github.com/liamylian/x-rsa
/*
// 使用方式
//已生成
	//私钥
	private_key := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC8ewMrlZtwPTjCX8n2bkMcBWHAj+HBdjzEi8aQhuPyezoNLWAnT1ufxBBe7s2TK6/duq4qORUk2liEa8mXJt0dhN4f1HwzZJIRvf1DRL1Oy+8J8n7W7ESj8oU6eJdC+mNgXgcuJkdGIs0WGfTfMO/0oQX0e1F022TkE3Z1GnWxqrnvOmvjZsQcr67RrQxHkpHXWzzFsV6CvyVF+gMNRPj7Jb6vZ0Uz+LZgNNMkJJ2/XLPUZllrrn0eIpBgDcFwMfnuIH+tf3KF0ZLqjwAZG3/LGS3v6LeRDzSxTJVN/xdcVnMp/kKd2e413bEu8RfiOrXVfKWQ5bmHwaqgi2IgbJjrAgMBAAECggEBALbEGhjjY+z8kNN7C16ZIR5E4btWSjXNL7PvTkbLZrK1Z19rxreXrmNIPg+wRUwWqty5rzVmUOiBiiAuXTbjCrsbIkF/TEiRl0r5DmNyfpewDUV6DiDeztijzSkb0qGyJzdFNnU7zRVnnbN2zAq5a5WtpxahliVYsYawkMeRNttl5YItmWKFGTfuSP8oe0D2jGR252OrwMKpWNtRDV0GTDcttG4BFv7sc0SITrBJCpArrcL8TqHlBf9hp+Fk7hOqhTAN1+hqlsggoSMCwLoJzfDJJoDYfT8ZhngolB4Z7TbfuO2se6IylN7BKO3R6OaAR5PyWc7eo18SpOAXxFNInKECgYEA5y0mX+VBMrqz+pqezYZZrXTE/3WiOPGWK4aBIGrs0Dom3gn/9Fjqx0RcyMoLsyS4vhpNu0OE8qACScWqvuSZzH4FajujaJr8b3J+SyR6GgnLbJAZftwt3oayeeXZ51Fgzb+E4F3tRIbbJMBzQs8qD0Yc8hOy7MPtlbgC/wOUR+MCgYEA0LgwNvNI0Rk/8rZ1IHrppDfG1Gmw2ZlmzeKG7N+0hg9e+pLVvonxPEeJbULSoHY/Y5WwGbZ7/j+8njcDRmE/DSkhcMxNLGWaJpaf5ALjemsDO/bX7B0ZEO/ZQxM0LxCSPnhyXoZYxUvjjlAbRb67OcUtKarH+c2DMXEnohhm6VkCgYAYj7h7Cc2CiUmj/5eNuI+AmV5w0du5dxmAaFsBybp9aSBGCQPzvxq6ry24e5cAoo8qS3skwUi3yn5Tvjek7n7UtjL3FX0IGSdm4+A9NZPvfkjpsOm0i8on0WHXB6++HuxND+MbP398/2L7JaNZXs9WLhnXMnaZJzvv40wWMhgoBwKBgHfnKAOZVV6SA9OaD8Vd9vTY+gcyxF7tKkridKKFjP5qvCR00AfK1OPNMNysw7Vc2OOLrYPH+ok/8kptzhrTaNZVp5lWNItpvtmmNCpXonjEGzKKhtZkjO7W6tgGozbabTkoFoPI6BxHkW1DBbfWb1YruNsz+fe+oTx5Bcyv+gbZAoGASf/tmlt2p3+R54Xys7T3lKHEb6pr2n/dcX6+6DCSMK34xIwfBBPslBWGqdFHU7e7h5mJSfvYnhEwk1MLXJ1RXmo8bzf5ZgOYWAE6OiOR84e7dkXyZT1lVoneYzjdI3qQgT9+gD8UKKinfpIwC6yWlcsUngOZ+zeG041KjNNj1xA="
	//公钥
	public_key := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvHsDK5WbcD04wl/J9m5DHAVhwI/hwXY8xIvGkIbj8ns6DS1gJ09bn8QQXu7Nkyuv3bquKjkVJNpYhGvJlybdHYTeH9R8M2SSEb39Q0S9TsvvCfJ+1uxEo/KFOniXQvpjYF4HLiZHRiLNFhn03zDv9KEF9HtRdNtk5BN2dRp1saq57zpr42bEHK+u0a0MR5KR11s8xbFegr8lRfoDDUT4+yW+r2dFM/i2YDTTJCSdv1yz1GZZa659HiKQYA3BcDH57iB/rX9yhdGS6o8AGRt/yxkt7+i3kQ80sUyVTf8XXFZzKf5CndnuNd2xLvEX4jq11XylkOW5h8GqoItiIGyY6wIDAQAB"
	privateKey := bytes.NewBufferString(utils.FormatPrivateKey(private_key))
	publicKey := bytes.NewBufferString(utils.FormatPublicKey(public_key))

	//生成公钥 私钥
	//publicKey := bytes.NewBufferString("")
	//privateKey := bytes.NewBufferString("")
	//err := utils.XrsaCreateKeys(publicKey, privateKey, 2048)
	//if err != nil {
	//	return
	//}
	//fmt.Println("publicKey",publicKey,"privateKey",privateKey)
	//生成公钥 私钥 END

	xrsa, err := utils.XrsaNewXRsa(publicKey.Bytes(), privateKey.Bytes())
	if err != nil {
		fmt.Println("err",err)
		return
	}

	fmt.Println("EEEE")
	data := "Hello, World"
	encrypted, _ := xrsa.XrsaPublicEncrypt(data)
	fmt.Println("encrypted加密",encrypted)
	decrypted, _ := xrsa.XrsaPrivateDecrypt(encrypted)
	fmt.Println("decrypted解密",decrypted)
	sign, _ := xrsa.XrsaSign(data)

	fmt.Println("sign",sign)
	sign = sign + "11"
	err = xrsa.XrsaVerify(data, sign)
	fmt.Println("验证",err)


--------------------------PHP
//私钥
            $private_key = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC8ewMrlZtwPTjCX8n2bkMcBWHAj+HBdjzEi8aQhuPyezoNLWAnT1ufxBBe7s2TK6/duq4qORUk2liEa8mXJt0dhN4f1HwzZJIRvf1DRL1Oy+8J8n7W7ESj8oU6eJdC+mNgXgcuJkdGIs0WGfTfMO/0oQX0e1F022TkE3Z1GnWxqrnvOmvjZsQcr67RrQxHkpHXWzzFsV6CvyVF+gMNRPj7Jb6vZ0Uz+LZgNNMkJJ2/XLPUZllrrn0eIpBgDcFwMfnuIH+tf3KF0ZLqjwAZG3/LGS3v6LeRDzSxTJVN/xdcVnMp/kKd2e413bEu8RfiOrXVfKWQ5bmHwaqgi2IgbJjrAgMBAAECggEBALbEGhjjY+z8kNN7C16ZIR5E4btWSjXNL7PvTkbLZrK1Z19rxreXrmNIPg+wRUwWqty5rzVmUOiBiiAuXTbjCrsbIkF/TEiRl0r5DmNyfpewDUV6DiDeztijzSkb0qGyJzdFNnU7zRVnnbN2zAq5a5WtpxahliVYsYawkMeRNttl5YItmWKFGTfuSP8oe0D2jGR252OrwMKpWNtRDV0GTDcttG4BFv7sc0SITrBJCpArrcL8TqHlBf9hp+Fk7hOqhTAN1+hqlsggoSMCwLoJzfDJJoDYfT8ZhngolB4Z7TbfuO2se6IylN7BKO3R6OaAR5PyWc7eo18SpOAXxFNInKECgYEA5y0mX+VBMrqz+pqezYZZrXTE/3WiOPGWK4aBIGrs0Dom3gn/9Fjqx0RcyMoLsyS4vhpNu0OE8qACScWqvuSZzH4FajujaJr8b3J+SyR6GgnLbJAZftwt3oayeeXZ51Fgzb+E4F3tRIbbJMBzQs8qD0Yc8hOy7MPtlbgC/wOUR+MCgYEA0LgwNvNI0Rk/8rZ1IHrppDfG1Gmw2ZlmzeKG7N+0hg9e+pLVvonxPEeJbULSoHY/Y5WwGbZ7/j+8njcDRmE/DSkhcMxNLGWaJpaf5ALjemsDO/bX7B0ZEO/ZQxM0LxCSPnhyXoZYxUvjjlAbRb67OcUtKarH+c2DMXEnohhm6VkCgYAYj7h7Cc2CiUmj/5eNuI+AmV5w0du5dxmAaFsBybp9aSBGCQPzvxq6ry24e5cAoo8qS3skwUi3yn5Tvjek7n7UtjL3FX0IGSdm4+A9NZPvfkjpsOm0i8on0WHXB6++HuxND+MbP398/2L7JaNZXs9WLhnXMnaZJzvv40wWMhgoBwKBgHfnKAOZVV6SA9OaD8Vd9vTY+gcyxF7tKkridKKFjP5qvCR00AfK1OPNMNysw7Vc2OOLrYPH+ok/8kptzhrTaNZVp5lWNItpvtmmNCpXonjEGzKKhtZkjO7W6tgGozbabTkoFoPI6BxHkW1DBbfWb1YruNsz+fe+oTx5Bcyv+gbZAoGASf/tmlt2p3+R54Xys7T3lKHEb6pr2n/dcX6+6DCSMK34xIwfBBPslBWGqdFHU7e7h5mJSfvYnhEwk1MLXJ1RXmo8bzf5ZgOYWAE6OiOR84e7dkXyZT1lVoneYzjdI3qQgT9+gD8UKKinfpIwC6yWlcsUngOZ+zeG041KjNNj1xA=";
	//公钥
	        $public_key = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvHsDK5WbcD04wl/J9m5DHAVhwI/hwXY8xIvGkIbj8ns6DS1gJ09bn8QQXu7Nkyuv3bquKjkVJNpYhGvJlybdHYTeH9R8M2SSEb39Q0S9TsvvCfJ+1uxEo/KFOniXQvpjYF4HLiZHRiLNFhn03zDv9KEF9HtRdNtk5BN2dRp1saq57zpr42bEHK+u0a0MR5KR11s8xbFegr8lRfoDDUT4+yW+r2dFM/i2YDTTJCSdv1yz1GZZa659HiKQYA3BcDH57iB/rX9yhdGS6o8AGRt/yxkt7+i3kQ80sUyVTf8XXFZzKf5CndnuNd2xLvEX4jq11XylkOW5h8GqoItiIGyY6wIDAQAB";

            $rsa = new XRsa(Util::formatPublicKey($public_key), Util::formatPrivateKey($private_key));

            $data = "Hello, World";
            $encrypted = $rsa->publicEncrypt($data);
            echo "加密：".$encrypted.PHP_EOL;
            $decrypted = $rsa->privateDecrypt($encrypted);
            echo "解密：".$decrypted.PHP_EOL;
            $sign = $rsa->sign($data);
            echo "SIGN:".$sign.PHP_EOL;
            $is_valid = $rsa->verify($data, $sign);
            echo "结果：".$is_valid.PHP_EOL;

            exit();

*/

const (
	CHAR_SET               = "UTF-8"
	BASE_64_FORMAT         = "UrlSafeNoPadding"
	RSA_ALGORITHM_KEY_TYPE = "PKCS8"
	RSA_ALGORITHM_SIGN     = crypto.SHA256
)

type XRsa struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

//
//通用格式化
func formatKey(key string) string {
	keyLen := len(key)
	keyStr := "\n"
	for i := 0; i <= keyLen; i = i + 64 {
		s := 0 + i
		e := 64 + i
		if e > keyLen {
			e = keyLen
		}
		keyStr += key[s:e] + "\n"
	}
	return keyStr
}

//格式化私钥
func FormatPrivateKey(privateKey string) string {
	return "-----BEGIN RSA PRIVATE KEY-----" + formatKey(privateKey) + "-----END RSA PRIVATE KEY-----"
}

// 格式化公钥
func FormatPublicKey(publicKey string) string {
	return "-----BEGIN PUBLIC KEY-----" + formatKey(publicKey) + "-----END PUBLIC KEY-----"
}

// see https://github.com/buf1024/golib/blob/master/crypt/rsa.go

// copy from crypt/rsa/pkcs1v15.go
var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

// copy from crypt/rsa/pkcs1v15.go
func rsaEncrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// copy from crypt/rsa/pkcs1v15.go
func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	// Special case: crypto.Hash(0) is used to indicate that the data is
	// signed directly.
	if hash == 0 {
		return inLen, nil, nil
	}

	hashLen = hash.Size()
	if inLen != hashLen {
		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
	}
	prefix, ok := hashPrefixes[hash]
	if !ok {
		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
	}
	return
}

// copy from crypt/rsa/pkcs1v15.go
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}
func unLeftPad(input []byte) (out []byte) {
	n := len(input)
	t := 2
	for i := 2; i < n; i++ {
		if input[i] == 0xff {
			t = t + 1
		} else {
			if input[i] == input[0] {
				t = t + int(input[1])
			}
			break
		}
	}
	out = make([]byte, n-t)
	copy(out, input[t:])
	return
}

func RsaPrivateEncrypt(privt *rsa.PrivateKey, data []byte) ([]byte, error) {
	signData, err := rsa.SignPKCS1v15(nil, privt, crypto.Hash(0), data)
	if err != nil {
		return nil, err
	}
	return signData, nil
}

func RsaPublicDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	decData, err := rsaPublicDecrypt(pub, crypto.Hash(0), nil, data)

	if err != nil {
		return nil, err
	}
	return decData, nil
}

func rsaPublicDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (out []byte, err error) {
	hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := (pub.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, fmt.Errorf("length illegal")
	}

	c := new(big.Int).SetBytes(sig)
	m := rsaEncrypt(new(big.Int), pub, c)
	em := leftPad(m.Bytes(), k)
	out = unLeftPad(em)

	err = nil
	return
}

//

func XrsaCreateKeys(publicKeyWriter, privateKeyWriter io.Writer, keyLength int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return err
	}
	derStream, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	err = pem.Encode(privateKeyWriter, block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = pem.Encode(publicKeyWriter, block)
	if err != nil {
		return err
	}

	return nil
}

func XrsaNewXRsa(publicKey []byte, privateKey []byte) (*XRsa, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	block, _ = pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pri, ok := priv.(*rsa.PrivateKey)
	if ok {
		return &XRsa{
			publicKey:  pub,
			privateKey: pri,
		}, nil
	} else {
		return nil, errors.New("private key not supported")
	}
}

func (r *XRsa) XrsaPublicEncrypt(data string) (string, error) {
	partLen := r.publicKey.N.BitLen()/8 - 11
	chunks := XrsaSplit([]byte(data), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bts, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bts)
	}

	return base64.RawURLEncoding.EncodeToString(buffer.Bytes()), nil
}

func (r *XRsa) XrsaPrivateDecrypt(encrypted string) (string, error) {
	partLen := r.publicKey.N.BitLen() / 8
	raw, err := base64.RawURLEncoding.DecodeString(encrypted)
	chunks := XrsaSplit([]byte(raw), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}

	return buffer.String(), err
}

func (r *XRsa) XrsaPrivateEncrypt(data string) (string, error) {
	partLen := r.publicKey.N.BitLen()/8 - 11
	chunks := XrsaSplit([]byte(data), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bts, err := RsaPrivateEncrypt(r.privateKey, chunk)
		if err != nil {
			return "", err
		}

		buffer.Write(bts)
	}

	return base64.RawURLEncoding.EncodeToString(buffer.Bytes()), nil
}

func (r *XRsa) XrsaPublicDecrypt(encrypted string) (string, error) {
	partLen := r.publicKey.N.BitLen() / 8
	raw, err := base64.RawURLEncoding.DecodeString(encrypted)
	chunks := XrsaSplit([]byte(raw), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := RsaPublicDecrypt(r.publicKey, chunk)

		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}

	return buffer.String(), err
}

func (r *XRsa) XrsaSign(data string) (string, error) {
	h := RSA_ALGORITHM_SIGN.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, r.privateKey, RSA_ALGORITHM_SIGN, hashed)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(sign), err
}

func (r *XRsa) XrsaVerify(data string, sign string) error {
	h := RSA_ALGORITHM_SIGN.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	decodedSign, err := base64.RawURLEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(r.publicKey, RSA_ALGORITHM_SIGN, hashed, decodedSign)
}

func XrsaSplit(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
