package examples

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/almubarok/canopusgo"
)

func GetToken() {
	privKey, err := ioutil.ReadFile("/your/path/to/M-00001.key")
	if err != nil {
		fmt.Println(err)
	}
	privPem, err := ioutil.ReadFile("/your/path/to/M-00001.pem")
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{Timeout: time.Second * time.Duration(60)}

	cano := canopusgo.CreateService("snap", privKey, privPem, "M-0001", "yoursecret", client)

	res, err := cano.GetToken()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", res)
}
