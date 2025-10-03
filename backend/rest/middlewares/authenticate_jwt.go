package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticationJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")

		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//fmt.Println("-----------", headerArr)
		accessToken := headerArr[1]

		tokenparts := strings.Split(accessToken, ".")

		if len(tokenparts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtheader := tokenparts[0]
		jwtpayload := tokenparts[1]
		signature := tokenparts[2]

		msg := jwtheader + "." + jwtpayload

		//cnf := config.GetConfig()

		byteArrSecret := []byte(m.cnf.JwtSecretKey)
		byteArrmsg := []byte(msg)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrmsg)

		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)

		if newSignature != signature {
			http.Error(w, "Unauthorized. Tui mc hacker vaag", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
