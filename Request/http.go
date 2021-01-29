package Request

import (
	"fmt"
	_ "io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var counter int = 0

func Send(method string, uri string, body string, wg *sync.WaitGroup) {

	client := &http.Client{}

	var methodName string
	methodName = strings.ToUpper(method)

	// Create request
	req, err := http.NewRequest(methodName, uri, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	/*respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	fmt.Println("response Status : ", resp.Status)
	/*fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))*/
	counter++
	wg.Done()
}

func GetCounter() int {
	return counter
}
