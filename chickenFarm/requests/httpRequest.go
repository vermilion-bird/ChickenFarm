package requests

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
