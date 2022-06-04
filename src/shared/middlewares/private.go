package middlewares

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"gitlab.internal.b2w.io/team/a-tech/maps/portal/src/server/config"
// 	"gitlab.internal.b2w.io/team/a-tech/maps/portal/src/shared/providers/jwt"
// 	"gitlab.internal.b2w.io/team/a-tech/maps/portal/src/shared/providers/logger"
// 	"gitlab.internal.b2w.io/team/a-tech/maps/portal/src/shared/tools/communication"
// )

// func PrivateRoute(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log := logger.Instance
// 		cfg := config.Instance

// 		comm := communication.New()

// 		jwtConfig := jwt.JwtConfig{
// 			SecretKey:       cfg.Application.Jwt.SecretKey,
// 			Issuer:          cfg.Application.Jwt.Issuer,
// 			ExpirationHours: cfg.Application.Jwt.ExpirationHours,
// 		}
// 		jwt := jwt.New(jwtConfig)

// 		token, err := extractToken(r)
// 		if err != nil {
// 			log.Error("shared.middlewares.private.extractToken", err.Error())

// 			response := comm.ResponseError(400, "error", err)
// 			response.JSON(w)

// 			return
// 		}

// 		tokenPayload, err := jwt.CheckToken(token)
// 		if err != nil {
// 			log.Error("shared.middlewares.private.CheckToken", err.Error())

// 			response := comm.ResponseError(401, "error", err)
// 			response.JSON(w)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), UserEmailContext{}, tokenPayload.Email)

// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func extractToken(r *http.Request) (string, error) {
// 	log := logger.Instance

// 	token := r.Header.Get("Authorization")
// 	if token == "" {
// 		err := fmt.Errorf("%s", "No Authorization header provided")
// 		log.Error("shared.middlewares.private.Authorization", err.Error())

// 		return token, err
// 	}

// 	extractedToken := strings.Split(token, "Bearer ")

// 	if len(extractedToken) == 2 {
// 		token = strings.TrimSpace(extractedToken[1])
// 	} else {
// 		err := fmt.Errorf("%s", "Incorrect Format of Authorization Token")
// 		log.Error("shared.middlewares.private.SplitToken", err.Error())

// 		return token, err
// 	}

// 	return token, nil
// }
