package httpHelper

import (
	"api/common/helpers"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type HttpHeader struct {
	Label string
	Value string
}

var httpClient *http.Client

func defaultHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}
	return httpClient
}

// HttpGet Fetches url and return response
func HttpGet(url string, response any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, response)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf(
			"HTTP Error %d. Failed to GET %s. Response: %s",
			resp.StatusCode,
			url,
			helpers.PrintAny(response),
		)
	}
	return nil
}

// HttpPost Post data to url and return response
func HttpPost(url string, headers []HttpHeader, body any, response any) error {
	jsonBody, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return err
	}
	for _, header := range headers {
		req.Header.Set(header.Label, header.Value)
	}

	client := defaultHttpClient()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, response)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf(
			"HTTP Error %d. Failed to POST %s. Response: %s",
			resp.StatusCode,
			url,
			helpers.PrintAny(response),
		)
	}
	return nil
}

// Download file from url and save to destination
func HttpDownloadFile(url, destination string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf(
			"HTTP Error %d. Failed to GET %s. Response: %s",
			resp.StatusCode,
			url,
			resp.Status,
		)
	}

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
