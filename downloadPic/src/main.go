package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
)

const (
	API_URL    = `http://api.lovebizhi.com/macos_v4.php?a=category&spdy=1&tid=3&order=hot&color_id=3&device=105&uuid=436e4ddc389027ba3aef863a27f6e6f9&mode=0&retina=0&client_id=1008&device_id=31547324&model_id=105&size_id=0&channel_id=70001&screen_width=1920&screen_height=1200&bizhi_width=1920&bizhi_height=1200&version_code=19&language=zh-Hans&jailbreak=0&mac=&p={pid}`
	FILE_PATH  = "./img"
	NUM_THREED = 10
)

var (
	reg *regexp.Regexp
)

func main() {
	reg, _ = regexp.Compile("([^/]+)$")

	if data, success := getResource(API_URL); success {
		if urls, success := resolveJSON(data); success {

			var pipe = make(chan string, 100)

			for _, url := range urls {

				if url != "" {
					pipe <- url
				} else {
					break
				}
			}

			ok := runWithGoroutine(pipe, NUM_THREED)

			for i := 0; i < NUM_THREED; i++ {
				<-ok
			}

		}
	}
}

// 多协程
func runWithGoroutine(pipeIn <-chan string, numOfGoroutine int) <-chan bool {
	isOK := make(chan bool)

	for i := 0; i < numOfGoroutine; i++ {
		go func(count int) {
			for {
				select {
				case url := <-pipeIn:
					runtime.Gosched()
					log.Printf("Threed %d : ", count)
					if imgData, success := getResource(url); success {
						saveImg(imgData, FILE_PATH+"/"+reg.FindString(url))
					}
				default:
					isOK <- true
					break
				}
			}
		}(i)
	}

	return isOK
}

// 从 json 中解析出图片 url
func resolveJSON(jsonData []byte) ([100]string, bool) {
	result := [100]string{}
	unmarshelData := make(map[string]interface{})

	if err := json.Unmarshal(jsonData, &unmarshelData); err != nil {
		log.Println("Unmarshal err : ", err.Error())
		return result, false
	}

	data := unmarshelData["data"].([]interface{})
	for index, dataItem := range data {
		item := dataItem.(map[string]interface{})
		image := item["image"].(map[string]interface{})
		original := image["original"].(string)
		result[index] = original
	}

	return result, true
}

// 通过 Get 请求获得数据
func getResource(url string) ([]byte, bool) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get file ", url, " failed : ", err.Error())
		return []byte{}, false
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read response err : ", err.Error())
		return []byte{}, false
	}

	return data, true
}

// 保存图片文件
func saveImg(data []byte, path string) bool {
	f, err := os.Create(path)
	if err != nil {
		log.Println("Open file err : ", err.Error())
		return false
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		log.Println("Write in file err : ", err.Error())
		return false
	}

	log.Println("Writed file in ", path)

	return true
}
