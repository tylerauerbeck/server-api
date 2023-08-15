// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package testclient

import (
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
	"go.infratographer.com/server-api/internal/ent/generated"
)

type TestClient interface {
}

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) TestClient {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	ServerAttribute     ServerAttribute     "json:\"serverAttribute\" graphql:\"serverAttribute\""
	ServerComponent     ServerComponent     "json:\"serverComponent\" graphql:\"serverComponent\""
	ServerComponentType ServerComponentType "json:\"serverComponentType\" graphql:\"serverComponentType\""
	ServerProvider      generated.Provider  "json:\"serverProvider\" graphql:\"serverProvider\""
	Server              Server              "json:\"server\" graphql:\"server\""
	ServerType          ServerType          "json:\"serverType\" graphql:\"serverType\""
	Entities            []Entity            "json:\"_entities\" graphql:\"_entities\""
	Service             Service             "json:\"_service\" graphql:\"_service\""
}
type Mutation struct {
	ServerAttributeCreate     ServerAttributeCreatePayload     "json:\"serverAttributeCreate\" graphql:\"serverAttributeCreate\""
	ServerAttributeUpdate     ServerAttributeUpdatePayload     "json:\"serverAttributeUpdate\" graphql:\"serverAttributeUpdate\""
	ServerAttributeDelete     ServerAttributeDeletePayload     "json:\"serverAttributeDelete\" graphql:\"serverAttributeDelete\""
	ServerComponentCreate     ServerComponentCreatePayload     "json:\"serverComponentCreate\" graphql:\"serverComponentCreate\""
	ServerComponentUpdate     ServerComponentUpdatePayload     "json:\"serverComponentUpdate\" graphql:\"serverComponentUpdate\""
	ServerComponentDelete     ServerComponentDeletePayload     "json:\"serverComponentDelete\" graphql:\"serverComponentDelete\""
	ServerComponentTypeCreate ServerComponentTypeCreatePayload "json:\"serverComponentTypeCreate\" graphql:\"serverComponentTypeCreate\""
	ServerComponentTypeUpdate ServerComponentTypeUpdatePayload "json:\"serverComponentTypeUpdate\" graphql:\"serverComponentTypeUpdate\""
	ServerComponentTypeDelete ServerComponentTypeDeletePayload "json:\"serverComponentTypeDelete\" graphql:\"serverComponentTypeDelete\""
	ServerProviderCreate      ServerProviderCreatePayload      "json:\"serverProviderCreate\" graphql:\"serverProviderCreate\""
	ServerProviderUpdate      ServerProviderUpdatePayload      "json:\"serverProviderUpdate\" graphql:\"serverProviderUpdate\""
	ServerProviderDelete      ServerProviderDeletePayload      "json:\"serverProviderDelete\" graphql:\"serverProviderDelete\""
	ServerCreate              ServerCreatePayload              "json:\"serverCreate\" graphql:\"serverCreate\""
	ServerUpdate              ServerUpdatePayload              "json:\"serverUpdate\" graphql:\"serverUpdate\""
	ServerDelete              ServerDeletePayload              "json:\"serverDelete\" graphql:\"serverDelete\""
	ServerTypeCreate          ServerTypeCreatePayload          "json:\"serverTypeCreate\" graphql:\"serverTypeCreate\""
	ServerTypeUpdate          ServerTypeUpdatePayload          "json:\"serverTypeUpdate\" graphql:\"serverTypeUpdate\""
	ServerTypeDelete          ServerTypeDeletePayload          "json:\"serverTypeDelete\" graphql:\"serverTypeDelete\""
}
