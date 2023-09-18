// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package graphclient

import (
	"context"
	"net/http"
	"time"

	"github.com/Yamashou/gqlgenc/client"
	"go.infratographer.com/x/gidx"
)

type GraphClient interface {
	GetOwnerServers(ctx context.Context, id gidx.PrefixedID, orderBy *ServerOrder, httpRequestOptions ...client.HTTPRequestOption) (*GetOwnerServers, error)
	GetServer(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServer, error)
	GetServerCPU(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServerCPU, error)
	GetServerCPUType(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServerCPUType, error)
	GetServerChassisType(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServerChassisType, error)
	GetServerType(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServerType, error)
	ServerCPUCreate(ctx context.Context, input CreateServerCPUInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUCreate, error)
	ServerCPUDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUDelete, error)
	ServerCPUTypeCreate(ctx context.Context, input CreateServerCPUTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUTypeCreate, error)
	ServerCPUTypeDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUTypeDelete, error)
	ServerCPUTypeUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerCPUTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUTypeUpdate, error)
	ServerCPUUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerCPUInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUUpdate, error)
	ServerChassisTypeCreate(ctx context.Context, input CreateServerChassisTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerChassisTypeCreate, error)
	ServerChassisTypeDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerChassisTypeDelete, error)
	ServerChassisTypeUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerChassisTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerChassisTypeUpdate, error)
	ServerCreate(ctx context.Context, input CreateServerInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCreate, error)
	ServerDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerDelete, error)
	ServerTypeCreate(ctx context.Context, input CreateServerTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerTypeCreate, error)
	ServerTypeDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerTypeDelete, error)
	ServerTypeUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerTypeUpdate, error)
	ServerUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerUpdate, error)
}

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) GraphClient {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	ServerChassis         ServerChassis         "json:\"serverChassis\" graphql:\"serverChassis\""
	ServerChassisType     ServerChassisType     "json:\"serverChassisType\" graphql:\"serverChassisType\""
	ServerComponent       ServerComponent       "json:\"serverComponent\" graphql:\"serverComponent\""
	ServerComponentType   ServerComponentType   "json:\"serverComponentType\" graphql:\"serverComponentType\""
	ServerCPU             ServerCPU             "json:\"serverCPU\" graphql:\"serverCPU\""
	ServerCPUType         ServerCPUType         "json:\"serverCPUType\" graphql:\"serverCPUType\""
	ServerHardDrive       ServerHardDrive       "json:\"serverHardDrive\" graphql:\"serverHardDrive\""
	ServerHardDriveType   ServerHardDriveType   "json:\"serverHardDriveType\" graphql:\"serverHardDriveType\""
	ServerMemory          ServerMemory          "json:\"serverMemory\" graphql:\"serverMemory\""
	ServerMemoryType      ServerMemoryType      "json:\"serverMemoryType\" graphql:\"serverMemoryType\""
	ServerMotherboard     ServerMotherboard     "json:\"serverMotherboard\" graphql:\"serverMotherboard\""
	ServerMotherboardType ServerMotherboardType "json:\"serverMotherboardType\" graphql:\"serverMotherboardType\""
	ServerNetworkCard     ServerNetworkCard     "json:\"serverNetworkCard\" graphql:\"serverNetworkCard\""
	ServerNetworkCardType ServerNetworkCardType "json:\"serverNetworkCardType\" graphql:\"serverNetworkCardType\""
	ServerNetworkPort     ServerNetworkPort     "json:\"serverNetworkPort\" graphql:\"serverNetworkPort\""
	ServerPowerSupply     ServerPowerSupply     "json:\"serverPowerSupply\" graphql:\"serverPowerSupply\""
	ServerPowerSupplyType ServerPowerSupplyType "json:\"serverPowerSupplyType\" graphql:\"serverPowerSupplyType\""
	ServerProvider        ServerProvider        "json:\"serverProvider\" graphql:\"serverProvider\""
	Server                Server                "json:\"server\" graphql:\"server\""
	ServerType            ServerType            "json:\"serverType\" graphql:\"serverType\""
	Entities              []Entity              "json:\"_entities\" graphql:\"_entities\""
	Service               Service               "json:\"_service\" graphql:\"_service\""
}
type Mutation struct {
	ServerChassis               ServerChassisCreatePayload         "json:\"serverChassis\" graphql:\"serverChassis\""
	ServerChassisUpdate         ServerChassisUpdatePayload         "json:\"serverChassisUpdate\" graphql:\"serverChassisUpdate\""
	ServerChassisDelete         ServerChassisDeletePayload         "json:\"serverChassisDelete\" graphql:\"serverChassisDelete\""
	ServerChassisTypeCreate     ServerChassisTypeCreatePayload     "json:\"serverChassisTypeCreate\" graphql:\"serverChassisTypeCreate\""
	ServerChassisTypeUpdate     ServerChassisTypeUpdatePayload     "json:\"serverChassisTypeUpdate\" graphql:\"serverChassisTypeUpdate\""
	ServerChassisTypeDelete     ServerChassisTypeDeletePayload     "json:\"serverChassisTypeDelete\" graphql:\"serverChassisTypeDelete\""
	ServerComponentCreate       ServerComponentCreatePayload       "json:\"serverComponentCreate\" graphql:\"serverComponentCreate\""
	ServerComponentUpdate       ServerComponentUpdatePayload       "json:\"serverComponentUpdate\" graphql:\"serverComponentUpdate\""
	ServerComponentDelete       ServerComponentDeletePayload       "json:\"serverComponentDelete\" graphql:\"serverComponentDelete\""
	ServerComponentTypeCreate   ServerComponentTypeCreatePayload   "json:\"serverComponentTypeCreate\" graphql:\"serverComponentTypeCreate\""
	ServerComponentTypeUpdate   ServerComponentTypeUpdatePayload   "json:\"serverComponentTypeUpdate\" graphql:\"serverComponentTypeUpdate\""
	ServerComponentTypeDelete   ServerComponentTypeDeletePayload   "json:\"serverComponentTypeDelete\" graphql:\"serverComponentTypeDelete\""
	ServerCPUCreate             ServerCPUCreatePayload             "json:\"serverCPUCreate\" graphql:\"serverCPUCreate\""
	ServerCPUUpdate             ServerCPUUpdatePayload             "json:\"serverCPUUpdate\" graphql:\"serverCPUUpdate\""
	ServerCPUDelete             ServerCPUDeletePayload             "json:\"serverCPUDelete\" graphql:\"serverCPUDelete\""
	ServerCPUTypeCreate         ServerCPUTypeCreatePayload         "json:\"serverCPUTypeCreate\" graphql:\"serverCPUTypeCreate\""
	ServerCPUTypeUpdate         ServerCPUTypeUpdatePayload         "json:\"serverCPUTypeUpdate\" graphql:\"serverCPUTypeUpdate\""
	ServerCPUTypeDelete         ServerCPUTypeDeletePayload         "json:\"serverCPUTypeDelete\" graphql:\"serverCPUTypeDelete\""
	ServerHardDrive             ServerHardDriveCreatePayload       "json:\"serverHardDrive\" graphql:\"serverHardDrive\""
	ServerHardDriveUpdate       ServerHardDriveUpdatePayload       "json:\"serverHardDriveUpdate\" graphql:\"serverHardDriveUpdate\""
	ServerHardDriveDelete       ServerHardDriveDeletePayload       "json:\"serverHardDriveDelete\" graphql:\"serverHardDriveDelete\""
	ServerHardDriveType         ServerHardDriveTypeCreatePayload   "json:\"serverHardDriveType\" graphql:\"serverHardDriveType\""
	ServerHardDriveTypeUpdate   ServerHardDriveTypeUpdatePayload   "json:\"serverHardDriveTypeUpdate\" graphql:\"serverHardDriveTypeUpdate\""
	ServerHardDriveTypeDelete   ServerHardDriveTypeDeletePayload   "json:\"serverHardDriveTypeDelete\" graphql:\"serverHardDriveTypeDelete\""
	ServerMemory                ServerMemoryCreatePayload          "json:\"serverMemory\" graphql:\"serverMemory\""
	ServerMemoryUpdate          ServerMemoryUpdatePayload          "json:\"serverMemoryUpdate\" graphql:\"serverMemoryUpdate\""
	ServerMemoryDelete          ServerMemoryDeletePayload          "json:\"serverMemoryDelete\" graphql:\"serverMemoryDelete\""
	ServerMemoryType            ServerMemoryTypeCreatePayload      "json:\"serverMemoryType\" graphql:\"serverMemoryType\""
	ServerMemoryTypeUpdate      ServerMemoryTypeUpdatePayload      "json:\"serverMemoryTypeUpdate\" graphql:\"serverMemoryTypeUpdate\""
	ServerMemoryTypeDelete      ServerMemoryTypeDeletePayload      "json:\"serverMemoryTypeDelete\" graphql:\"serverMemoryTypeDelete\""
	ServerMotherboard           ServerMotherboardCreatePayload     "json:\"serverMotherboard\" graphql:\"serverMotherboard\""
	ServerMotherboardUpdate     ServerMotherboardUpdatePayload     "json:\"serverMotherboardUpdate\" graphql:\"serverMotherboardUpdate\""
	ServerMotherboardDelete     ServerMotherboardDeletePayload     "json:\"serverMotherboardDelete\" graphql:\"serverMotherboardDelete\""
	ServerMotherboardType       ServerMotherboardTypeCreatePayload "json:\"serverMotherboardType\" graphql:\"serverMotherboardType\""
	ServerMotherboardTypeUpdate ServerMotherboardTypeUpdatePayload "json:\"serverMotherboardTypeUpdate\" graphql:\"serverMotherboardTypeUpdate\""
	ServerMotherboardTypeDelete ServerMotherboardTypeDeletePayload "json:\"serverMotherboardTypeDelete\" graphql:\"serverMotherboardTypeDelete\""
	ServerNetworkCard           ServerNetworkCardCreatePayload     "json:\"serverNetworkCard\" graphql:\"serverNetworkCard\""
	ServerNetworkCardUpdate     ServerNetworkCardUpdatePayload     "json:\"serverNetworkCardUpdate\" graphql:\"serverNetworkCardUpdate\""
	ServerNetworkCardDelete     ServerNetworkCardDeletePayload     "json:\"serverNetworkCardDelete\" graphql:\"serverNetworkCardDelete\""
	ServerNetworkCardType       ServerNetworkCardTypeCreatePayload "json:\"serverNetworkCardType\" graphql:\"serverNetworkCardType\""
	ServerNetworkCardTypeUpdate ServerNetworkCardTypeUpdatePayload "json:\"serverNetworkCardTypeUpdate\" graphql:\"serverNetworkCardTypeUpdate\""
	ServerNetworkCardTypeDelete ServerNetworkCardTypeDeletePayload "json:\"serverNetworkCardTypeDelete\" graphql:\"serverNetworkCardTypeDelete\""
	ServerNetworkPort           ServerNetworkPortCreatePayload     "json:\"serverNetworkPort\" graphql:\"serverNetworkPort\""
	ServerNetworkPortUpdate     ServerNetworkPortUpdatePayload     "json:\"serverNetworkPortUpdate\" graphql:\"serverNetworkPortUpdate\""
	ServerNetworkPortDelete     ServerNetworkPortDeletePayload     "json:\"serverNetworkPortDelete\" graphql:\"serverNetworkPortDelete\""
	ServerPowerSupply           ServerPowerSupplyCreatePayload     "json:\"serverPowerSupply\" graphql:\"serverPowerSupply\""
	ServerPowerSupplyUpdate     ServerPowerSupplyUpdatePayload     "json:\"serverPowerSupplyUpdate\" graphql:\"serverPowerSupplyUpdate\""
	ServerPowerSupplyDelete     ServerPowerSupplyDeletePayload     "json:\"serverPowerSupplyDelete\" graphql:\"serverPowerSupplyDelete\""
	ServerPowerSupplyType       ServerPowerSupplyTypeCreatePayload "json:\"serverPowerSupplyType\" graphql:\"serverPowerSupplyType\""
	ServerPowerSupplyTypeUpdate ServerPowerSupplyTypeUpdatePayload "json:\"serverPowerSupplyTypeUpdate\" graphql:\"serverPowerSupplyTypeUpdate\""
	ServerPowerSupplyTypeDelete ServerPowerSupplyTypeDeletePayload "json:\"serverPowerSupplyTypeDelete\" graphql:\"serverPowerSupplyTypeDelete\""
	ServerProviderCreate        ServerProviderCreatePayload        "json:\"serverProviderCreate\" graphql:\"serverProviderCreate\""
	ServerProviderUpdate        ServerProviderUpdatePayload        "json:\"serverProviderUpdate\" graphql:\"serverProviderUpdate\""
	ServerProviderDelete        ServerProviderDeletePayload        "json:\"serverProviderDelete\" graphql:\"serverProviderDelete\""
	ServerCreate                ServerCreatePayload                "json:\"serverCreate\" graphql:\"serverCreate\""
	ServerUpdate                ServerUpdatePayload                "json:\"serverUpdate\" graphql:\"serverUpdate\""
	ServerDelete                ServerDeletePayload                "json:\"serverDelete\" graphql:\"serverDelete\""
	ServerTypeCreate            ServerTypeCreatePayload            "json:\"serverTypeCreate\" graphql:\"serverTypeCreate\""
	ServerTypeUpdate            ServerTypeUpdatePayload            "json:\"serverTypeUpdate\" graphql:\"serverTypeUpdate\""
	ServerTypeDelete            ServerTypeDeletePayload            "json:\"serverTypeDelete\" graphql:\"serverTypeDelete\""
}
type GetOwnerServers struct {
	Entities []*struct {
		Servers struct {
			Edges []*struct {
				Node *struct {
					ID   gidx.PrefixedID "json:\"id\" graphql:\"id\""
					Name string          "json:\"name\" graphql:\"name\""
				} "json:\"node\" graphql:\"node\""
			} "json:\"edges\" graphql:\"edges\""
		} "json:\"servers\" graphql:\"servers\""
	} "json:\"_entities\" graphql:\"_entities\""
}
type GetServer struct {
	Server struct {
		ID          gidx.PrefixedID "json:\"id\" graphql:\"id\""
		Name        string          "json:\"name\" graphql:\"name\""
		Description *string         "json:\"description\" graphql:\"description\""
		Owner       struct {
			ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
		} "json:\"owner\" graphql:\"owner\""
		Location struct {
			ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
		} "json:\"location\" graphql:\"location\""
		ServerProvider struct {
			ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
		} "json:\"serverProvider\" graphql:\"serverProvider\""
		CreatedAt time.Time "json:\"createdAt\" graphql:\"createdAt\""
		UpdatedAt time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
	} "json:\"server\" graphql:\"server\""
}
type GetServerCPU struct {
	ServerCPU struct {
		ID     gidx.PrefixedID "json:\"id\" graphql:\"id\""
		Serial string          "json:\"serial\" graphql:\"serial\""
		Server struct {
			ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
		} "json:\"server\" graphql:\"server\""
		CreatedAt     time.Time "json:\"createdAt\" graphql:\"createdAt\""
		UpdatedAt     time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
		ServerCPUType struct {
			ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
		} "json:\"serverCPUType\" graphql:\"serverCPUType\""
	} "json:\"serverCPU\" graphql:\"serverCPU\""
}
type GetServerCPUType struct {
	ServerCPUType struct {
		ID         gidx.PrefixedID "json:\"id\" graphql:\"id\""
		Vendor     string          "json:\"vendor\" graphql:\"vendor\""
		Model      string          "json:\"model\" graphql:\"model\""
		CoreCount  int64           "json:\"coreCount\" graphql:\"coreCount\""
		ClockSpeed string          "json:\"clockSpeed\" graphql:\"clockSpeed\""
		CreatedAt  time.Time       "json:\"createdAt\" graphql:\"createdAt\""
		UpdatedAt  time.Time       "json:\"updatedAt\" graphql:\"updatedAt\""
	} "json:\"serverCPUType\" graphql:\"serverCPUType\""
}
type GetServerChassisType struct {
	ServerChassisType struct {
		ID          gidx.PrefixedID "json:\"id\" graphql:\"id\""
		Model       string          "json:\"model\" graphql:\"model\""
		Vendor      string          "json:\"vendor\" graphql:\"vendor\""
		Height      string          "json:\"height\" graphql:\"height\""
		CreatedAt   time.Time       "json:\"createdAt\" graphql:\"createdAt\""
		UpdatedAt   time.Time       "json:\"updatedAt\" graphql:\"updatedAt\""
		IsFullDepth bool            "json:\"isFullDepth\" graphql:\"isFullDepth\""
	} "json:\"serverChassisType\" graphql:\"serverChassisType\""
}
type GetServerType struct {
	ServerType struct {
		ID    gidx.PrefixedID "json:\"id\" graphql:\"id\""
		Name  string          "json:\"name\" graphql:\"name\""
		Owner struct {
			ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
		} "json:\"owner\" graphql:\"owner\""
		CreatedAt time.Time "json:\"createdAt\" graphql:\"createdAt\""
		UpdatedAt time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
	} "json:\"serverType\" graphql:\"serverType\""
}
type ServerCPUCreate struct {
	ServerCPUCreate struct {
		ServerCPU struct {
			ID     gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Serial string          "json:\"serial\" graphql:\"serial\""
			Server struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"server\" graphql:\"server\""
			CreatedAt     time.Time "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt     time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
			ServerCPUType struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"serverCPUType\" graphql:\"serverCPUType\""
		} "json:\"serverCPU\" graphql:\"serverCPU\""
	} "json:\"serverCPUCreate\" graphql:\"serverCPUCreate\""
}
type ServerCPUDelete struct {
	ServerCPUDelete struct {
		DeletedID gidx.PrefixedID "json:\"deletedID\" graphql:\"deletedID\""
	} "json:\"serverCPUDelete\" graphql:\"serverCPUDelete\""
}
type ServerCPUTypeCreate struct {
	ServerCPUTypeCreate struct {
		ServerCPUType struct {
			ID         gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Vendor     string          "json:\"vendor\" graphql:\"vendor\""
			Model      string          "json:\"model\" graphql:\"model\""
			CoreCount  int64           "json:\"coreCount\" graphql:\"coreCount\""
			ClockSpeed string          "json:\"clockSpeed\" graphql:\"clockSpeed\""
			CreatedAt  time.Time       "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt  time.Time       "json:\"updatedAt\" graphql:\"updatedAt\""
		} "json:\"serverCPUType\" graphql:\"serverCPUType\""
	} "json:\"serverCPUTypeCreate\" graphql:\"serverCPUTypeCreate\""
}
type ServerCPUTypeDelete struct {
	ServerCPUTypeDelete struct {
		DeletedID gidx.PrefixedID "json:\"deletedID\" graphql:\"deletedID\""
	} "json:\"serverCPUTypeDelete\" graphql:\"serverCPUTypeDelete\""
}
type ServerCPUTypeUpdate struct {
	ServerCPUTypeUpdate struct {
		ServerCPUType struct {
			ID         gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Vendor     string          "json:\"vendor\" graphql:\"vendor\""
			Model      string          "json:\"model\" graphql:\"model\""
			CoreCount  int64           "json:\"coreCount\" graphql:\"coreCount\""
			ClockSpeed string          "json:\"clockSpeed\" graphql:\"clockSpeed\""
			CreatedAt  time.Time       "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt  time.Time       "json:\"updatedAt\" graphql:\"updatedAt\""
		} "json:\"serverCPUType\" graphql:\"serverCPUType\""
	} "json:\"serverCPUTypeUpdate\" graphql:\"serverCPUTypeUpdate\""
}
type ServerCPUUpdate struct {
	ServerCPUUpdate struct {
		ServerCPU struct {
			ID     gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Serial string          "json:\"serial\" graphql:\"serial\""
			Server struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"server\" graphql:\"server\""
			CreatedAt     time.Time "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt     time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
			ServerCPUType struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"serverCPUType\" graphql:\"serverCPUType\""
		} "json:\"serverCPU\" graphql:\"serverCPU\""
	} "json:\"serverCPUUpdate\" graphql:\"serverCPUUpdate\""
}
type ServerChassisTypeCreate struct {
	ServerChassisTypeCreate struct {
		ServerChassisType struct {
			ID          gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Model       string          "json:\"model\" graphql:\"model\""
			Vendor      string          "json:\"vendor\" graphql:\"vendor\""
			Height      string          "json:\"height\" graphql:\"height\""
			CreatedAt   time.Time       "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt   time.Time       "json:\"updatedAt\" graphql:\"updatedAt\""
			IsFullDepth bool            "json:\"isFullDepth\" graphql:\"isFullDepth\""
		} "json:\"serverChassisType\" graphql:\"serverChassisType\""
	} "json:\"serverChassisTypeCreate\" graphql:\"serverChassisTypeCreate\""
}
type ServerChassisTypeDelete struct {
	ServerChassisTypeDelete struct {
		DeletedID gidx.PrefixedID "json:\"deletedID\" graphql:\"deletedID\""
	} "json:\"serverChassisTypeDelete\" graphql:\"serverChassisTypeDelete\""
}
type ServerChassisTypeUpdate struct {
	ServerChassisTypeUpdate struct {
		ServerChassisType struct {
			ID          gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Model       string          "json:\"model\" graphql:\"model\""
			Vendor      string          "json:\"vendor\" graphql:\"vendor\""
			Height      string          "json:\"height\" graphql:\"height\""
			CreatedAt   time.Time       "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt   time.Time       "json:\"updatedAt\" graphql:\"updatedAt\""
			IsFullDepth bool            "json:\"isFullDepth\" graphql:\"isFullDepth\""
		} "json:\"serverChassisType\" graphql:\"serverChassisType\""
	} "json:\"serverChassisTypeUpdate\" graphql:\"serverChassisTypeUpdate\""
}
type ServerCreate struct {
	ServerCreate struct {
		Server struct {
			ID             gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Name           string          "json:\"name\" graphql:\"name\""
			ServerProvider struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"serverProvider\" graphql:\"serverProvider\""
			Owner struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"owner\" graphql:\"owner\""
			Location struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"location\" graphql:\"location\""
			CreatedAt time.Time "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
		} "json:\"server\" graphql:\"server\""
	} "json:\"serverCreate\" graphql:\"serverCreate\""
}
type ServerDelete struct {
	ServerDelete struct {
		DeletedID gidx.PrefixedID "json:\"deletedID\" graphql:\"deletedID\""
	} "json:\"serverDelete\" graphql:\"serverDelete\""
}
type ServerTypeCreate struct {
	ServerTypeCreate struct {
		ServerType struct {
			ID    gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Name  string          "json:\"name\" graphql:\"name\""
			Owner struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"owner\" graphql:\"owner\""
			CreatedAt time.Time "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
		} "json:\"serverType\" graphql:\"serverType\""
	} "json:\"serverTypeCreate\" graphql:\"serverTypeCreate\""
}
type ServerTypeDelete struct {
	ServerTypeDelete struct {
		DeletedID gidx.PrefixedID "json:\"deletedID\" graphql:\"deletedID\""
	} "json:\"serverTypeDelete\" graphql:\"serverTypeDelete\""
}
type ServerTypeUpdate struct {
	ServerTypeUpdate struct {
		ServerType struct {
			ID    gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Name  string          "json:\"name\" graphql:\"name\""
			Owner struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"owner\" graphql:\"owner\""
			CreatedAt time.Time "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
		} "json:\"serverType\" graphql:\"serverType\""
	} "json:\"serverTypeUpdate\" graphql:\"serverTypeUpdate\""
}
type ServerUpdate struct {
	ServerUpdate struct {
		Server struct {
			ID        gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Name      string          "json:\"name\" graphql:\"name\""
			CreatedAt time.Time       "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt time.Time       "json:\"updatedAt\" graphql:\"updatedAt\""
		} "json:\"server\" graphql:\"server\""
	} "json:\"serverUpdate\" graphql:\"serverUpdate\""
}

const GetOwnerServersDocument = `query GetOwnerServers ($id: ID!, $orderBy: ServerOrder) {
	_entities(representations: {__typename:"ResourceOwner",id:$id}) {
		... on ResourceOwner {
			servers(orderBy: $orderBy) {
				edges {
					node {
						id
						name
					}
				}
			}
		}
	}
}
`

func (c *Client) GetOwnerServers(ctx context.Context, id gidx.PrefixedID, orderBy *ServerOrder, httpRequestOptions ...client.HTTPRequestOption) (*GetOwnerServers, error) {
	vars := map[string]interface{}{
		"id":      id,
		"orderBy": orderBy,
	}

	var res GetOwnerServers
	if err := c.Client.Post(ctx, "GetOwnerServers", GetOwnerServersDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetServerDocument = `query GetServer ($id: ID!) {
	server(id: $id) {
		id
		name
		description
		owner {
			id
		}
		location {
			id
		}
		serverProvider {
			id
		}
		createdAt
		updatedAt
	}
}
`

func (c *Client) GetServer(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServer, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetServer
	if err := c.Client.Post(ctx, "GetServer", GetServerDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetServerCPUDocument = `query GetServerCPU ($id: ID!) {
	serverCPU(id: $id) {
		id
		serial
		server {
			id
		}
		createdAt
		updatedAt
		serverCPUType {
			id
		}
	}
}
`

func (c *Client) GetServerCPU(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServerCPU, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetServerCPU
	if err := c.Client.Post(ctx, "GetServerCPU", GetServerCPUDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetServerCPUTypeDocument = `query GetServerCPUType ($id: ID!) {
	serverCPUType(id: $id) {
		id
		vendor
		model
		coreCount
		clockSpeed
		createdAt
		updatedAt
	}
}
`

func (c *Client) GetServerCPUType(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServerCPUType, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetServerCPUType
	if err := c.Client.Post(ctx, "GetServerCPUType", GetServerCPUTypeDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetServerChassisTypeDocument = `query GetServerChassisType ($id: ID!) {
	serverChassisType(id: $id) {
		id
		model
		vendor
		height
		createdAt
		updatedAt
		isFullDepth
	}
}
`

func (c *Client) GetServerChassisType(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServerChassisType, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetServerChassisType
	if err := c.Client.Post(ctx, "GetServerChassisType", GetServerChassisTypeDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetServerTypeDocument = `query GetServerType ($id: ID!) {
	serverType(id: $id) {
		id
		name
		owner {
			id
		}
		createdAt
		updatedAt
	}
}
`

func (c *Client) GetServerType(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetServerType, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetServerType
	if err := c.Client.Post(ctx, "GetServerType", GetServerTypeDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerCPUCreateDocument = `mutation ServerCPUCreate ($input: CreateServerCPUInput!) {
	serverCPUCreate(input: $input) {
		serverCPU {
			id
			serial
			server {
				id
			}
			createdAt
			updatedAt
			serverCPUType {
				id
			}
		}
	}
}
`

func (c *Client) ServerCPUCreate(ctx context.Context, input CreateServerCPUInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUCreate, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res ServerCPUCreate
	if err := c.Client.Post(ctx, "ServerCPUCreate", ServerCPUCreateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerCPUDeleteDocument = `mutation ServerCPUDelete ($id: ID!) {
	serverCPUDelete(id: $id) {
		deletedID
	}
}
`

func (c *Client) ServerCPUDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUDelete, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res ServerCPUDelete
	if err := c.Client.Post(ctx, "ServerCPUDelete", ServerCPUDeleteDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerCPUTypeCreateDocument = `mutation ServerCPUTypeCreate ($input: CreateServerCPUTypeInput!) {
	serverCPUTypeCreate(input: $input) {
		serverCPUType {
			id
			vendor
			model
			coreCount
			clockSpeed
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) ServerCPUTypeCreate(ctx context.Context, input CreateServerCPUTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUTypeCreate, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res ServerCPUTypeCreate
	if err := c.Client.Post(ctx, "ServerCPUTypeCreate", ServerCPUTypeCreateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerCPUTypeDeleteDocument = `mutation ServerCPUTypeDelete ($id: ID!) {
	serverCPUTypeDelete(id: $id) {
		deletedID
	}
}
`

func (c *Client) ServerCPUTypeDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUTypeDelete, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res ServerCPUTypeDelete
	if err := c.Client.Post(ctx, "ServerCPUTypeDelete", ServerCPUTypeDeleteDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerCPUTypeUpdateDocument = `mutation ServerCPUTypeUpdate ($id: ID!, $input: UpdateServerCPUTypeInput!) {
	serverCPUTypeUpdate(id: $id, input: $input) {
		serverCPUType {
			id
			vendor
			model
			coreCount
			clockSpeed
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) ServerCPUTypeUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerCPUTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUTypeUpdate, error) {
	vars := map[string]interface{}{
		"id":    id,
		"input": input,
	}

	var res ServerCPUTypeUpdate
	if err := c.Client.Post(ctx, "ServerCPUTypeUpdate", ServerCPUTypeUpdateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerCPUUpdateDocument = `mutation ServerCPUUpdate ($id: ID!, $input: UpdateServerCPUInput!) {
	serverCPUUpdate(id: $id, input: $input) {
		serverCPU {
			id
			serial
			server {
				id
			}
			createdAt
			updatedAt
			serverCPUType {
				id
			}
		}
	}
}
`

func (c *Client) ServerCPUUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerCPUInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCPUUpdate, error) {
	vars := map[string]interface{}{
		"id":    id,
		"input": input,
	}

	var res ServerCPUUpdate
	if err := c.Client.Post(ctx, "ServerCPUUpdate", ServerCPUUpdateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerChassisTypeCreateDocument = `mutation ServerChassisTypeCreate ($input: CreateServerChassisTypeInput!) {
	serverChassisTypeCreate(input: $input) {
		serverChassisType {
			id
			model
			vendor
			height
			createdAt
			updatedAt
			isFullDepth
		}
	}
}
`

func (c *Client) ServerChassisTypeCreate(ctx context.Context, input CreateServerChassisTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerChassisTypeCreate, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res ServerChassisTypeCreate
	if err := c.Client.Post(ctx, "ServerChassisTypeCreate", ServerChassisTypeCreateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerChassisTypeDeleteDocument = `mutation ServerChassisTypeDelete ($id: ID!) {
	serverChassisTypeDelete(id: $id) {
		deletedID
	}
}
`

func (c *Client) ServerChassisTypeDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerChassisTypeDelete, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res ServerChassisTypeDelete
	if err := c.Client.Post(ctx, "ServerChassisTypeDelete", ServerChassisTypeDeleteDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerChassisTypeUpdateDocument = `mutation ServerChassisTypeUpdate ($id: ID!, $input: UpdateServerChassisTypeInput!) {
	serverChassisTypeUpdate(id: $id, input: $input) {
		serverChassisType {
			id
			model
			vendor
			height
			createdAt
			updatedAt
			isFullDepth
		}
	}
}
`

func (c *Client) ServerChassisTypeUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerChassisTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerChassisTypeUpdate, error) {
	vars := map[string]interface{}{
		"id":    id,
		"input": input,
	}

	var res ServerChassisTypeUpdate
	if err := c.Client.Post(ctx, "ServerChassisTypeUpdate", ServerChassisTypeUpdateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerCreateDocument = `mutation ServerCreate ($input: CreateServerInput!) {
	serverCreate(input: $input) {
		server {
			id
			name
			serverProvider {
				id
			}
			owner {
				id
			}
			location {
				id
			}
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) ServerCreate(ctx context.Context, input CreateServerInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerCreate, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res ServerCreate
	if err := c.Client.Post(ctx, "ServerCreate", ServerCreateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerDeleteDocument = `mutation ServerDelete ($id: ID!) {
	serverDelete(id: $id) {
		deletedID
	}
}
`

func (c *Client) ServerDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerDelete, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res ServerDelete
	if err := c.Client.Post(ctx, "ServerDelete", ServerDeleteDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerTypeCreateDocument = `mutation ServerTypeCreate ($input: CreateServerTypeInput!) {
	serverTypeCreate(input: $input) {
		serverType {
			id
			name
			owner {
				id
			}
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) ServerTypeCreate(ctx context.Context, input CreateServerTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerTypeCreate, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res ServerTypeCreate
	if err := c.Client.Post(ctx, "ServerTypeCreate", ServerTypeCreateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerTypeDeleteDocument = `mutation ServerTypeDelete ($id: ID!) {
	serverTypeDelete(id: $id) {
		deletedID
	}
}
`

func (c *Client) ServerTypeDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*ServerTypeDelete, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res ServerTypeDelete
	if err := c.Client.Post(ctx, "ServerTypeDelete", ServerTypeDeleteDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerTypeUpdateDocument = `mutation ServerTypeUpdate ($id: ID!, $input: UpdateServerTypeInput!) {
	serverTypeUpdate(id: $id, input: $input) {
		serverType {
			id
			name
			owner {
				id
			}
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) ServerTypeUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerTypeInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerTypeUpdate, error) {
	vars := map[string]interface{}{
		"id":    id,
		"input": input,
	}

	var res ServerTypeUpdate
	if err := c.Client.Post(ctx, "ServerTypeUpdate", ServerTypeUpdateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ServerUpdateDocument = `mutation ServerUpdate ($id: ID!, $input: UpdateServerInput!) {
	serverUpdate(id: $id, input: $input) {
		server {
			id
			name
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) ServerUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateServerInput, httpRequestOptions ...client.HTTPRequestOption) (*ServerUpdate, error) {
	vars := map[string]interface{}{
		"id":    id,
		"input": input,
	}

	var res ServerUpdate
	if err := c.Client.Post(ctx, "ServerUpdate", ServerUpdateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}
