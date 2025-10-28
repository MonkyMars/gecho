package errors

import (
	"fmt"
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) HandleMethod(w http.ResponseWriter, r *http.Request, intendedMethod string) *utils.ResponseBuilder {
	method := r.Method
	if method != intendedMethod {
		return MethodNotAllowed(w).WithMessage(fmt.Sprintf("Method %s not allowed", method))
	}

	return nil
}
