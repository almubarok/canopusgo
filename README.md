![GitHub](https://img.shields.io/github/license/almubarok/canopusgo)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/almubarok/canopusgo)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/almubarok/canopusgo)
![GitHub all releases](https://img.shields.io/github/downloads/almubarok/canopusgo/total)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/almubarok/canopusgo?sort=semver)

## Overview

Canopusgo is implementation of canopus service using Golang.

## Usage

Create service first:
```go
key, err := ioutil.ReadFile("/your/path/to/M-00001.key")
if err != nil {
  fmt.Println(err)
}
pem, err := ioutil.ReadFile("/your/path/to/M-0001.pem")
if err != nil {
  fmt.Println(err)
}
client := &http.Client{Timeout: time.Second * time.Duration(60)}

cano := canopusgo.CreateService("snap", key, pem, "M-0001", "yoursecret", client)
```

Then you can user this service to create payment

```go
func (cano *Canopus) GenerateSignature(payload []byte) (string, error)
func (cano *Canopus) GetToken() (string, error)
func (cano *Canopus) GetAvailableMethod(amount float64) ([]PaymentMethod, error)
func (cano *Canopus) GenerateCart(payload CartPayload, paymentMethod PaymentMethod) (CartResponse, error)
```

Check some examples [canopusgo](https://github.com/almubarok/canopusgo/tree/main/examples)