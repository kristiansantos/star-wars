package httphelper

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func GetWithTimeout(url string) (*http.Response, error) {
	ctx, execTimeOut := context.WithTimeout(context.Background(), time.Second*60)
	defer execTimeOut()

	request, requestError := http.NewRequestWithContext(ctx, "GET", url, nil)

	if requestError != nil {
		return nil, requestError
	}

	response, responseError := http.DefaultClient.Do(request)

	if responseError != nil {
		return nil, requestError
	}

	return response, nil
}

func GetParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func GetBody(body io.Reader, v interface{}) (err error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, v); err != nil {
		return
	}

	return nil
}
