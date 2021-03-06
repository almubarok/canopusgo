package examples

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/almubarok/canopusgo"
)

func GetAvailableMethod() {
	privKey, err := ioutil.ReadFile("/your/path/to/M-00001.key")
	if err != nil {
		fmt.Println(err)
	}
	privPem, err := ioutil.ReadFile("/your/path/to/M-00001.pem")
	if err != nil {
		fmt.Println(err)
	}
	timeout := time.Second * time.Duration(60)

	cano, err := canopusgo.CreateService(canopusgo.InitService{
		MerchantKey: privKey,
		MerchantPem: privPem,
		TimeOut:     timeout,
		MerchantID:  "M-0001",
		Secret:      "yoursecret",
	})
	if err != nil {
		fmt.Println(err)
	}

	res, err := cano.GetAvailableMethod(10000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", res)
}
