package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	DefaultApiBase = "http://127.0.0.1:40000"
)

var (
	apiBase string
)

func SetApiBase(url string) {
	apiBase = url
}

func GetApiBase() string {
	if apiBase != "" {
		return apiBase
	}
	if envApiBase := os.Getenv("MORELOGIN_API_BASE"); envApiBase != "" {
		return envApiBase
	}
	return DefaultApiBase
}

type MoreLoginClient struct {
	Url       string
	Method    string
	Payload   interface{}
	Headers   map[string]string
	Response  *http.Response
	parsedUrl *url.URL
	Query     map[string]string
}

type Option func(client *MoreLoginClient)

type ErrMsgV5 struct {
	Message string `json:"message"`
}

func NewMoreLoginClient(method, urlString string, opts ...Option) *MoreLoginClient {
	urlString = GetApiBase() + urlString
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	client := &MoreLoginClient{
		Method:    method,
		Url:       parsedUrl.String(),
		parsedUrl: parsedUrl,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func WithQuery(query map[string]interface{}) Option {
	return func(client *MoreLoginClient) {
		parsedQuery := make(map[string]string)
		if query != nil {
			queryParams := client.parsedUrl.Query()
			for k, v := range query {
				parsedValue := ""
				switch v.(type) {
				case string:
					parsedValue = v.(string)
				case int:
					parsedValue = strconv.Itoa(v.(int))
				case bool:
					parsedValue = strconv.FormatBool(v.(bool))
				}
				if parsedValue != "" {
					queryParams.Set(k, parsedValue)
					parsedQuery[k] = parsedValue
				}
			}
			client.parsedUrl.RawQuery = queryParams.Encode()
		}
		client.Url = client.parsedUrl.String()
		client.Query = parsedQuery
	}
}

func WithPayload(payload interface{}) Option {
	return func(client *MoreLoginClient) {
		client.Payload = payload
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(client *MoreLoginClient) {
		client.Headers = headers
	}
}

func (g *MoreLoginClient) SetHeaders(headers map[string]string) *MoreLoginClient {
	g.Headers = headers
	return g
}

func (g *MoreLoginClient) Do() (*MoreLoginClient, error) {
	g.Response = nil
	_payload, _ := json.Marshal(g.Payload)
	println("request body", g.Url, g.Method, string(_payload))
	req, err := http.NewRequest(g.Method, g.Url, bytes.NewReader(_payload))
	if err != nil {
		return nil, NewInternalError(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "mcp-morelogin "+Version+" Go/"+runtime.GOOS+"/"+runtime.GOARCH+"/"+runtime.Version())

	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return g, NewNetworkError(err)
	}

	g.Response = resp

	resStr, err := json.Marshal(resp.Body)
	if err != nil {
		return g, NewInternalError(err)
	}
	jsonString := string(resStr)
	fmt.Println(jsonString)

	// 检查响应状态码
	if !g.IsSuccess() {
		body, _ := ioutil.ReadAll(resp.Body)
		return g, NewAPIError(resp.StatusCode, body)
	}

	return g, nil
}

func (g *MoreLoginClient) IsSuccess() bool {
	if g.Response == nil {
		return false
	}

	successMap := map[int]struct{}{
		http.StatusOK:          struct{}{},
		http.StatusCreated:     struct{}{},
		http.StatusNoContent:   struct{}{},
		http.StatusFound:       struct{}{},
		http.StatusNotModified: struct{}{},
	}

	if _, ok := successMap[g.Response.StatusCode]; ok {
		return true
	}
	return false
}

func (g *MoreLoginClient) IsFail() bool {
	return !g.IsSuccess()
}

func (g *MoreLoginClient) GetRespBody() ([]byte, error) {
	return ioutil.ReadAll(g.Response.Body)
}

func (g *MoreLoginClient) HandleMCPResult(object any) (*mcp.CallToolResult, error) {
	_, err := g.Do()
	if err != nil {
		switch {
		case IsNetworkError(err):
			return mcp.NewToolResultError("Network error: Unable to connect to Morelogin API"), err
		case IsAPIError(err):
			moreLoginErr := err.(*MoreLoginError)
			return mcp.NewToolResultError(fmt.Sprintf("API error (%d): %s", moreLoginErr.Code, moreLoginErr.Details)), err
		default:
			return mcp.NewToolResultError(err.Error()), err
		}
	}

	// Handle no content case when object is nil
	if object == nil {
		return mcp.NewToolResultText("Operation completed successfully"), nil
	}

	body, err := g.GetRespBody()
	println("response body", g.Url, string(body))
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to read response body: %s", err.Error())),
			NewInternalError(err)
	}

	if err = json.Unmarshal(body, object); err != nil {
		errorMessage := fmt.Sprintf("Failed to parse response: %v", err)
		return mcp.NewToolResultError(errorMessage), NewInternalError(errors.New(errorMessage))
	}

	result, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to format response: %s", err.Error())),
			NewInternalError(err)
	}

	return mcp.NewToolResultText(string(result)), nil
}
