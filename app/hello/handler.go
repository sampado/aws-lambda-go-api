package hello

import (
	"fmt"
	"net/http"

	"github.com/apex/gateway/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	userID := ""
	requestContext, ok := gateway.RequestContext(r.Context())
	if ok && requestContext.Authorizer != nil {
		userID = requestContext.Authorizer.IAM.UserID
	}
	if userID == "" {
		userID = r.FormValue("name")
	}

	fmt.Fprintf(w, HelloMessage(userID))
}
