package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"exam/client/request"
	"exam/client/response"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	// 用户相关接口
	registerURL = "/office/user/register"
	loginURL    = "/office/user/login"

	// 短链接相关接口
	createShortLinkURL        = "/office/short_link/create"
	deleteShortLinkURL        = "/office/short_link/delete"
	updateShortLinkURL        = "/office/short_link/update"
	shareShortLinkURL         = "/office/short_link/share"
	listShortLinksURL         = "/office/short_link/list"
	searchShortLinksURL       = "/office/short_link/search"
	redirectToOriginalLinkURL = "/office/short_link"
	summarizeURL              = "/office/short_link/summarize"
)

type TestClient struct {
	addr   string
	port   int
	client *http.Client
}

func (c *TestClient) Register(req *request.RegisterUserReq) (*response.RegisterResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s:%d%s", c.addr, c.port, registerURL)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	var registerResp response.RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&registerResp); err != nil {
		return nil, err
	}

	return &registerResp, nil
}

func (c *TestClient) Login(req *request.LoginUserReq) (*response.LoginResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s:%d%s", c.addr, c.port, loginURL)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	var loginResp response.LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, err
	}

	return &loginResp, nil
}

func (c *TestClient) CreateShortLink(req *request.CreateShortLinkReq, cookies ...*http.Cookie) (*response.CreateResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s:%d%s", c.addr, c.port, createShortLinkURL)
	fmt.Println(url)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// 添加 cookies 到请求
	for _, cookie := range cookies {
		httpReq.AddCookie(cookie)
	}

	// 查找名为 "wps_id" 的 cookie
	var wpsIDCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "wps_id" {
			wpsIDCookie = cookie
			break
		}
	}

	if wpsIDCookie == nil {
		fmt.Println("wps_id cookie not found")
		return nil, errors.New("wps_id cookie not found")
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return nil, errors.New("status not 200")
	}

	var createResp response.CreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		return nil, err
	}

	return &createResp, nil
}

func (c *TestClient) ShareShortLink(req *request.ShareShortLinkReq, cookies ...*http.Cookie) (*response.ShareResponse, error) {
	url := fmt.Sprintf("http://%s:%d%s?link_id=%s&user_ids=%s", c.addr, c.port, shareShortLinkURL, req.LinkId, strings.Join(req.UserIds, ","))
	httpReq, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	// 添加 cookies 到请求
	for _, cookie := range cookies {
		httpReq.AddCookie(cookie)
	}

	// 查找名为 "wps_id" 的 cookie
	var wpsIDCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "wps_id" {
			wpsIDCookie = cookie
			break
		}
	}

	if wpsIDCookie == nil {
		fmt.Println("wps_id cookie not found")
		return nil, errors.New("wps_id cookie not found")
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	var shareResp response.ShareResponse
	if err := json.NewDecoder(resp.Body).Decode(&shareResp); err != nil {
		return nil, err
	}

	return &shareResp, nil
}

// UpdateShortLink 处理
func (c *TestClient) UpdateShortLink(req *request.UpdateShortLinkReq, cookies ...*http.Cookie) (*response.UpdateResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s:%d%s", c.addr, c.port, updateShortLinkURL)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// 添加 cookies 到请求
	for _, cookie := range cookies {
		httpReq.AddCookie(cookie)
	}

	// 查找名为 "wps_id" 的 cookie
	var wpsIDCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "wps_id" {
			wpsIDCookie = cookie
			break
		}
	}

	if wpsIDCookie == nil {
		fmt.Println("wps_id cookie not found")
		return nil, errors.New("wps_id cookie not found")
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	var updateResp response.UpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateResp); err != nil {
		return nil, err
	}

	return &updateResp, nil
}

// DeleteShortLink 处理
func (c *TestClient) DeleteShortLink(req *request.DeleteShortLinkReq, cookies ...*http.Cookie) (*response.DeleteResp, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s:%d%s", c.addr, c.port, deleteShortLinkURL)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// 添加 cookies 到请求
	for _, cookie := range cookies {
		httpReq.AddCookie(cookie)
	}

	// 查找名为 "wps_id" 的 cookie
	var wpsIDCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "wps_id" {
			wpsIDCookie = cookie
			break
		}
	}

	if wpsIDCookie == nil {
		fmt.Println("wps_id cookie not found")
		return nil, errors.New("wps_id cookie not found")
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	var deleteResp response.DeleteResp
	if err := json.NewDecoder(resp.Body).Decode(&deleteResp); err != nil {
		return nil, err
	}

	return &deleteResp, nil
}

// ListShortLinks 处理
func (c *TestClient) ListShortLinks(req *request.ListShortLinksReq, cookies ...*http.Cookie) (*response.ListResp, error) {
	url := fmt.Sprintf("http://%s:%d%s?page=%d&page_size=%d", c.addr, c.port, listShortLinksURL, req.Page, req.PageSize)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// 添加 cookies 到请求
	for _, cookie := range cookies {
		httpReq.AddCookie(cookie)
	}

	// 查找名为 "wps_id" 的 cookie
	var wpsIDCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "wps_id" {
			wpsIDCookie = cookie
			break
		}
	}

	if wpsIDCookie == nil {
		fmt.Println("wps_id cookie not found")
		return nil, errors.New("wps_id cookie not found")
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	var listResp response.ListResp
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}

	return &listResp, nil
}

// SearchShortLinksByLinkID 处理
func (c *TestClient) SearchShortLinksByLinkID(req *request.ListShortLinksByLinkIDReq, cookies ...*http.Cookie) (*response.SearchResp, error) {
	url := fmt.Sprintf("http://%s:%d%s?keyword=%s&page=%d&page_size=%d", c.addr, c.port, searchShortLinksURL, req.Keyword, req.Page, req.PageSize)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// 添加 cookies 到请求
	for _, cookie := range cookies {
		httpReq.AddCookie(cookie)
	}

	// 查找名为 "wps_id" 的 cookie
	var wpsIDCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "wps_id" {
			wpsIDCookie = cookie
			break
		}
	}

	if wpsIDCookie == nil {
		fmt.Println("wps_id cookie not found")
		return nil, errors.New("wps_id cookie not found")
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	var searchResp response.SearchResp
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, err
	}

	return &searchResp, nil
}

// RedirectToOriginalLink 处理
func (c *TestClient) RedirectToOriginalLink(linkId string, cookies ...*http.Cookie) error {
	url := fmt.Sprintf("http://%s:%d%s/%s", c.addr, c.port, redirectToOriginalLinkURL, linkId)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// 添加 cookies 到请求
	for _, cookie := range cookies {
		httpReq.AddCookie(cookie)
	}

	// 查找名为 "wps_id" 的 cookie
	var wpsIDCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "wps_id" {
			wpsIDCookie = cookie
			break
		}
	}

	if wpsIDCookie == nil {
		fmt.Println("wps_id cookie not found")
		return errors.New("wps_id cookie not found")
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMovedPermanently {
		return errors.New("status not 301")
	}

	return nil
}

// Summarize 处理
func (c *TestClient) Summarize(req *request.SummarizeReq, cookies ...*http.Cookie) (*response.SummarizeResp, error) {
	url := fmt.Sprintf("http://%s:%d%s?link_id=%s", c.addr, c.port, summarizeURL, req.LinkId)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// 添加 cookies 到请求
	for _, cookie := range cookies {
		httpReq.AddCookie(cookie)
	}

	// 查找名为 "wps_id" 的 cookie
	var wpsIDCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "wps_id" {
			wpsIDCookie = cookie
			break
		}
	}

	if wpsIDCookie == nil {
		fmt.Println("wps_id cookie not found")
		return nil, errors.New("wps_id cookie not found")
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	var summarizeResp response.SummarizeResp
	if err := json.NewDecoder(resp.Body).Decode(&summarizeResp); err != nil {
		return nil, err
	}

	return &summarizeResp, nil
}

func NewClient(addr string, port int) *TestClient {
	client := http.DefaultClient
	client.Transport = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConnsPerHost:   1024,
		WriteBufferSize:       32 << 10,
		ReadBufferSize:        32 << 10,
		IdleConnTimeout:       15 * time.Second,
		ResponseHeaderTimeout: 3 * time.Minute,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	}
	return &TestClient{
		addr:   addr,
		port:   port,
		client: client,
	}
}
