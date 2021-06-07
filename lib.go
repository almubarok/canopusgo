package canopusgo

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
)

func ValidateResponse(resp []byte) (CommonMessage, error) {
	var response CommonMessage
	var err error

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return CommonMessage{}, err
	}

	if response.Response.Result.Code != "00000000" {
		err = errors.New(response.Response.Result.Code + " " + response.Response.Result.Message)
		return CommonMessage{}, err
	}

	return response, nil
}

func VerifySignature(merchantPem []byte, response []byte, signature string) (bool, error) {
	block, _ := pem.Decode(merchantPem)
	if block == nil {
		return false, errors.New("failed to decode merchantPem")
	}

	pubkey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	sDec, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256(response)

	err = rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hashed[:], sDec)
	if err != nil {
		return false, err
	}
	return true, nil
}
