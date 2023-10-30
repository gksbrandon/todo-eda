package rest

import (
	"context"

	"github.com/gksbrandon/todo-eda/users/userspb"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	const apiRoot = "/api/users"

	gateway := runtime.NewServeMux()
	err := userspb.RegisterUsersServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	// Register OAuth2 endpoints
	registerOAuthEndpoints(mux)

	// mount the GRPC gateway
	mux.Mount(apiRoot, gateway)

	return nil
}

func registerOAuthEndpoints(r chi.Router) {
	oauth := OAuth{
		ProviderConfigs: map[string]*oauth2.Config{},
	}

	dbxConfig := &oauth2.Config{
		ClientID:     "4iglelte8wxd775",
		ClientSecret: "soxa4wdky615xuv",
		Scopes:       []string{"account_info.read", "files.metadata.read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.dropbox.com/oauth2/authorize",
			TokenURL: "https://api.dropboxapi.com/oauth2/token",
		},
	}
	oauth.ProviderConfigs["dropbox"] = dbxConfig

	r.Get("/oauth/{provider}/connect", oauth.Connect)
	r.Get("/oauth/{provider}/callback", oauth.Callback)
}
