package middle

import "net/http"

type HttpMiddleWare func(next http.HandlerFunc) http.HandlerFunc
