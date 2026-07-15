package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	mu         sync.Mutex
	cache      map[string]cacheEntry
}

type cacheEntry struct {
	price     float64
	expiresAt time.Time
}

type MarketResponse struct {
	Name      string  `json:"name"`
	Set       string  `json:"set"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	Source    string  `json:"source"`
	Timestamp string  `json:"timestamp"`
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		cache: make(map[string]cacheEntry),
	}
}

func (c *Client) Query(name, setName string) (MarketResponse, error) {
	cacheKey := strings.ToLower(strings.TrimSpace(name + "|" + setName))
	c.mu.Lock()
	if entry, ok := c.cache[cacheKey]; ok && time.Now().Before(entry.expiresAt) {
		c.mu.Unlock()
		return MarketResponse{
			Name:  name,
			Set:   setName,
			Price: entry.price,
			Source: "cache",
		}, nil
	}
	c.mu.Unlock()

	url := fmt.Sprintf("%s/card?name=%s&set=%s", c.baseURL, name, setName)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return MarketResponse{}, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return MarketResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return MarketResponse{}, fmt.Errorf("market API error %d: %s", resp.StatusCode, string(body))
	}

	var mr MarketResponse
	if err := json.NewDecoder(resp.Body).Decode(&mr); err != nil {
		return MarketResponse{}, err
	}

	c.mu.Lock()
	c.cache[cacheKey] = cacheEntry{
		price:     mr.Price,
		expiresAt: time.Now().Add(15 * time.Minute),
	}
	c.mu.Unlock()

	return mr, nil
}
