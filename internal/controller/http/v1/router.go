package v1

import (
	"net/http"

	"github.com/v1adhope/web-service-employees/internal/usecase"
)

func New(mux *http.ServeMux, u usecase.Employee) {
	NewRoutes(&Routes{mux, u})
}
