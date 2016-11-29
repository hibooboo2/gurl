package restclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// APIClient ...
type APIClient struct {
	Verbose   bool
	Host      string
	AuthToken string
	Log       io.Writer
	User      map[string]interface{}
}

var (
	// ErrForbidden ...
	ErrForbidden = errors.New("Forbidden")
	// ErrUnauthorized ...
	ErrUnauthorized = errors.New("Unauthorized")
)

// Get ...
func (c *APIClient) Get(endpoint string, dest interface{}) error {
	endpoint = fmt.Sprintf(`%s/%s`, c.Host, endpoint)
	return c.apiRequest("GET", endpoint, nil, dest)
}

// Delete ...
func (c *APIClient) Delete(endpoint string, payload, dest interface{}) error {
	endpoint = fmt.Sprintf(`%s/%s`, c.Host, endpoint)
	return c.apiRequest("DELETE", endpoint, payload, dest)
}

// Post ...
func (c *APIClient) Post(endpoint string, payload, dest interface{}) error {
	endpoint = fmt.Sprintf(`%s/%s`, c.Host, endpoint)
	return c.apiRequest("POST", endpoint, payload, dest)
}

// Put ...
func (c *APIClient) Put(endpoint string, payload, dest interface{}) error {
	endpoint = fmt.Sprintf(`%s/%s`, c.Host, endpoint)
	return c.apiRequest("PUT", endpoint, payload, dest)
}

func (c *APIClient) apiRequest(method, endpoint string, payload, dest interface{}) error {
	start := time.Now()
	err := c.ExecuteRequest(method, endpoint, payload, dest)
	total := time.Since(start)
	if c.Verbose {
		msg := fmt.Sprintf("Request took: %s\n", total.String())
		if err != nil {
			msg = fmt.Sprintf("%s\n%s", msg, err.Error())
		}
		switch c.Log != nil {
		case true:
			c.Log.Write([]byte(msg))
		case false:
			log.Println(msg)
		}
	}
	return err
}

// ExecuteRequest ... Executes a request raw to the method provided and using the raw endpoint
// provided. As well as marshaling to the payload and dest given they are not nil.
func (c *APIClient) ExecuteRequest(method, endpoint string, payload, dest interface{}) error {
	if !strings.HasPrefix(endpoint, "http") {
		endpoint = fmt.Sprintf("http://%s", endpoint)
	}
	var r *http.Request
	var err error
	var pldData []byte
	switch method {
	case "POST", "PUT", "DELETE", "GET":
		if payload != nil {
			pldData, err = json.Marshal(payload)
			if err != nil {
				return err
			}
			pld := strings.NewReader(string(pldData))

			r, err = http.NewRequest(method, endpoint, pld)
		} else {
			r, err = http.NewRequest(method, endpoint, nil)
		}
	default:
		return fmt.Errorf("unsupported request method: %s", method)
	}

	if err != nil {
		return err
	}

	r.Header.Add("content-type", "application/json")
	if c.AuthToken != "" {
		r.Header.Add("token", c.AuthToken)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch {
	case resp.StatusCode == 403:
		log.Println("Url was ", endpoint)
		log.Printf("Body was: %s", string(data))
		return ErrForbidden
	case resp.StatusCode == 401:
		return ErrUnauthorized
	case resp.StatusCode >= 400:
		log.Println("Url was ", endpoint)
		log.Printf("Body was: %s", string(data))
		return fmt.Errorf("Got statusCode: %d expected status below 400", resp.StatusCode)
	}

	if dest != nil {
		err = json.Unmarshal(data, dest)
		if err != nil {
			return err
		}
	}
	return nil
}
