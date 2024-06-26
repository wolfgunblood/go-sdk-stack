package stack

import (
	"github.com/wolfgunblood/go-sdk-stack/client"
	"github.com/wolfgunblood/go-sdk-stack/logger"
	models "github.com/wolfgunblood/go-sdk-stack/model"
)

type Client struct {
    apiKey    string
    apiSecret string
    baseURL   string
    httpClient *client.HttpClient
}

func NewClient(apiKey, apiSecret, baseURL string) *Client {
    return &Client{
        apiKey:    apiKey,
        apiSecret: apiSecret,
        baseURL:   baseURL,
        httpClient: client.NewHttpClient(),
    }
}

func (c *Client) GetUser(userID string) (*models.User, error) {
    logger.Info.Println("Get User")
    //api call to get user
    return nil, ErrUserNotFound
}
