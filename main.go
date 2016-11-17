package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hibooboo2/gurl/restclient"
)

var client = &restclient.APIClient{}
var method string
var endpoint string

func main() {
	log.SetFlags(0)

	var val interface{}
	if endpoint == "" {
		endpoint = os.Args[len(os.Args)-1]
	}
	if client.Verbose {
		log.Println("Endpoint: ", endpoint)
	}
	if endsInt(endpoint) {
		val = make(map[string]interface{})
	} else {
		val = []map[string]interface{}{}
	}

	var err error
	pld := makePayload()
	switch strings.ToLower(method) {
	case "get":
		err = client.Get(endpoint, &val)
	case "put":
		err = client.Put(endpoint, pld, &val)
	case "delete":
		err = client.Delete(endpoint, pld, &val)
	case "post":
		err = client.Post(endpoint, pld, &val)
	default:
		panic("Unsupported method :" + method)
	}
	if err == nil {
		jsonText, _ := json.MarshalIndent(val, "", "    ")
		fmt.Fprintln(os.Stdout, string(Format(jsonText)))
	} else {
		fmt.Fprintln(os.Stdout, err.Error())
	}
	time.Sleep(time.Millisecond * 100)
}

func endsInt(endpoint string) bool {
	parts := strings.Split(endpoint, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if id != 0 {
		log.Println("Resource id:", id)
	}
	return err != nil
}

var payloadString string

func makePayload() map[string]interface{} {
	pld := make(map[string]interface{})
	err := json.Unmarshal([]byte(payloadString), &pld)
	if err != nil {
		return make(map[string]interface{})
	}
	return pld
}

func humanSize(sizeORIG int) string {
	// Take int representing size of len([]byte)
	// Convert it to human readable size. ex: 3KB or 3GB
	size := float64(sizeORIG)
	switch {
	case size > 1024*1024*1024*1024: // TiB
		return fmt.Sprintf(`%.4v TiB`, size/1024*1024*1024*1024)
	case size > 1024*1024*1024: // GiB
		return fmt.Sprintf(`%.4v GiB`, size/1024*1024*1024)
	case size > 1024*1024: // MiB
		return fmt.Sprintf(`%.4v MiB`, size/1024*1024)
	case size > 1024: // KiB
		return fmt.Sprintf(`%.4v KiB`, size/1024)
	default:
		return fmt.Sprintf(`%.4v bytes`, size)
	}
}
