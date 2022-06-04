package httphelper

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

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
