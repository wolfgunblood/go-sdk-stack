package client

import (
	"net/http"
)

type Retrier struct {
    client Requester
    retries int
}

func NewRetrier(client Requester, retries int) *Retrier {
    return &Retrier{
        client: client,
        retries: retries,
    }
}

func (r *Retrier) Do(req *http.Request) (*http.Response, error) {
    var res *http.Response
    var err error

    for i := 0; i <= r.retries; i++ {
        res, err = r.client.Do(req)
        if err == nil {
            return res, nil
        }
    }

    return nil, err
}
