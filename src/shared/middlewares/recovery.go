package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/tools/communication"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				comm := communication.New()

				// span := tracer.StartSpan("http.request", tracer.ResourceName("/api/v1/placements"))

				// span.SetTag("http.url", r.URL.Path)
				// span.SetTag("http.query", r.URL.RawQuery)

				// span.SetTag("api.error.get_placements", err)
				// span.Finish(tracer.WithError(errors.New(fmt.Sprintf("%s", err))))

				fmt.Println(err)

				response := comm.ResponseError(500, "error", errors.New("internal server error"))
				response.JSON(w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
