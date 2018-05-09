package twitch

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

func parseURLQuery(i interface{}) url.Values {
	q, _ := query.Values(i)
	v := url.Values{}
	for key, value := range q {
		for i := 0; i < len(value); i++ {
			if !isZero(value[i]) {
				v.Add(key, value[i])
			}
		}
	}
	return v
}

func (c Client) doRequest(method, uri string, body io.Reader) (*http.Response, error) {
	r, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Client-ID", c.ClientID)

	res, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 401 {
		return nil, errors.New("server returned 401. likely caused due to an invalid client id")
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("server returned non 200 status code. status code: %v", res.StatusCode)
	}

	return res, nil
}

func isZero(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
