package bindings

import (
	"context"
	"errors"
	"net/http"

	"github.com/Southclaws/storyden/backend/pkg/services/authentication"
	"github.com/Southclaws/storyden/backend/pkg/transports/http/openapi"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/kr/pretty"
)

type Authentication struct{ s authentication.Service }

func NewAuthentication(s authentication.Service) Authentication { return Authentication{s} }

func (i *Authentication) Signin(ctx context.Context, request openapi.SigninRequestObject) any {
	return nil
}

func (i *Authentication) Signup(ctx context.Context, request openapi.SignupRequestObject) any {
	return nil
}

func (i *Authentication) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session, ok := i.s.DecodeSession(r); ok {
			ctx := authentication.AddUserToContext(r.Context(), session)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (i *Authentication) validator(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
	pretty.Println(ai.SecurityScheme)
	return errors.New("not allowed")
}
