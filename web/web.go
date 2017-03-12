package web

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"myiot_hub/sensor"
	"net/http"
	_ "strconv"
	_ "strings"
	"time"
)

/*
	This api is responsible for the creation, and data point adding of the dryer cycle
*/

func PostPoint(s sensor.SensorData, tag string) string {
	url := "http://ec2-52-23-227-218.compute-1.amazonaws.com:8000/data/?access_data_key=%242b%2412%24l24KcoFv0wVZvmuAoraEcOFJ4EWDRDv6hq.6f8qSGypj2wuyuIMR."
	//url = "http://httpbin.org/post"
	payload := []byte("{\"tag\": \"" + tag + "\", \"temperature\" : " + fmt.Sprint(s.Temp()) + ", \"humidity\" : " + fmt.Sprint(s.Humidity()) + " , \"orientation\" : \"0 0 0\"}")
	fmt.Println(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(payload)
	return string(body)

}

func CreateDryingCycle(tag string) string {
	url := "http://ec2-52-23-227-218.compute-1.amazonaws.com:8000/cycles/?access_data_key=%242b%2412%24l24KcoFv0wVZvmuAoraEcOFJ4EWDRDv6hq.6f8qSGypj2wuyuIMR."
	//url = "http://httpbin.org/post"
	payload := []byte("{\"time\": \"" + fmt.Sprint(time.Now().Unix()) + "\" , \"tag\" : \"" + tag + "\"}")
	fmt.Println(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(payload)
	return string(body)
}

func FinishDryingCycle(tag string, recommendations string) string {
	url := "http://ec2-52-23-227-218.compute-1.amazonaws.com:8000/cycles/finish/?access_data_key=%242b%2412%24l24KcoFv0wVZvmuAoraEcOFJ4EWDRDv6hq.6f8qSGypj2wuyuIMR.&tag=" + tag
	//url = "http://httpbin.org/post"
	payload := []byte("{\"time\": \"" + fmt.Sprint(time.Now().Unix()) + "\" , \"tag\" : \"" + tag + "\" , \"recommendations\" : \"" + recommendations + "\"}")
	fmt.Println(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(payload)
	return string(body)
}

func Web() {
}
