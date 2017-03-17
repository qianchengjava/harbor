// Fetch prints the content found at a URL.
package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	time.Sleep(60 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	var client = &http.Client{
		Timeout:   time.Second * 30,
		Transport: tr,
	}

	for _, url := range os.Args[1:] {

		resp, err := client.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//		fmt.Printf("%s", b)

		if strings.Contains(string(b), "Clarity Seed App") {
			fmt.Printf("sucess!\n")
		} else {
			fmt.Println("the response does not contain \"Harbor\"!")

			fmt.Println(string(b))
			os.Exit(1)
		}

	}
}
