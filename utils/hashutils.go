package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func Sha512Hex(data []byte) (string, error) {
	v, err := Sha512WithSalt(data, nil, nil)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(v), nil
}
func Sha512WithSaltHex(data []byte, saltPrepend []byte, saltAppend []byte) (string, error) {
	v, err := Sha512WithSalt(data, saltPrepend, saltAppend)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(v), nil
}

func Sha512(data []byte) ([]byte, error) {
	return Sha512WithSalt(data, nil, nil)
}

func Sha512WithSalt(data []byte, saltPrepend []byte, saltAppend []byte) ([]byte, error) {
	h := sha512.New()
	_, err := h.Write(saltPrepend)
	if err != nil {
		return nil, err
	}
	_, err = h.Write(data)
	if err != nil {
		return nil, err
	}
	_, err = h.Write(saltAppend)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func HmacSha512HexEncoded(secret string, message []byte) (string, error) {
	h := hmac.New(sha512.New, []byte(secret))
	_, err := h.Write(message)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
