// middleware/error_handler.go
package middleware

import (
	"log"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered from panic:", r)

				// Xử lý lỗi ở đây, có thể gửi một phản hồi lỗi hoặc thực hiện các bước xử lý khác.

				// Gửi một phản hồi lỗi 500 Internal Server Error
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		// Chuyển hành động đến handler tiếp theo trong chuỗi middleware
		next.ServeHTTP(w, r)
	})
}
