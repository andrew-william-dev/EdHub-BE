package middleware

import "net/http"

func CORS_MIDDLEWARE ( next http.HandlerFunc ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if(r.Method == http.MethodOptions) {
			w.WriteHeader((http.StatusOK))
			return
		}

		next.ServeHTTP(w, r)
	}
}
