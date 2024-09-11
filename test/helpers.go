package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func getURL(path string) string {
	port := os.Getenv("TODO_PORT")
	return fmt.Sprintf("http://localhost:%s/%s", port, path)
}
func request(path string, method string, values map[string]any, cookies ...map[string]string) ([]byte, error) {
	var (
		data []byte
		err  error
	)

	if len(values) > 0 {
		data, err = json.Marshal(values)
		if err != nil {
			return nil, err
		}
	}

	url := getURL(path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	if len(cookies) > 0 {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
		for key, val := range cookies[0] {
			jar.SetCookies(req.URL, []*http.Cookie{
				{
					Name:  key,
					Value: val,
				},
			})
		}

		client.Jar = jar
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	return io.ReadAll(resp.Body)
}
