package grpcHandlers

import (
	"MiddleApp/api/api/ServiceApiPb"
	"MiddleApp/internal/repositories/interfaces"
	"context"
)

type Handlers struct {
	Data interfaces.Repository

	ServiceApiPb.UnimplementedUsersServiceServer
}

func New(data interfaces.Repository) *Handlers {
	return &Handlers{Data: data}
}

func (h *Handlers) Read(ctx context.Context, in *ServiceApiPb.ReadRequest) (*ServiceApiPb.ReadResponse, error) {
	info := h.Data.Read()
	users := make([]*ServiceApiPb.ReadResponse_User, len(info))
	for i, user := range info {
		users[i] = &ServiceApiPb.ReadResponse_User{
			Id:      user.Id,
			Name:    user.Name,
			Surname: user.Surname,
			Age:     user.Age,
		}
	}

	response := ServiceApiPb.ReadResponse{Users: users}

	return &response, nil
}
