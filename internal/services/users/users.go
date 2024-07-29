package users

import (
	"context"

	"github.com/ADAGroupTcc/ms-users-api/exceptions"
	"github.com/ADAGroupTcc/ms-users-api/internal/domain"
	"github.com/ADAGroupTcc/ms-users-api/internal/helpers"
	"github.com/ADAGroupTcc/ms-users-api/internal/repositories/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Create(ctx context.Context, request domain.UserRequest) (*domain.User, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	List(ctx context.Context, queryParams helpers.QueryParams) (*domain.UserResponse, error)
	Update(ctx context.Context, id string, request domain.UserPatchRequest) error
	Delete(ctx context.Context, id string) error
}

type userService struct {
	userRepository users.Repository
}

func New(userRepository users.Repository) Service {
	return &userService{
		userRepository,
	}
}

func (h *userService) Create(ctx context.Context, request domain.UserRequest) (*domain.User, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	user := request.ToUser()

	return h.userRepository.Create(ctx, user)
}

func (h *userService) Get(ctx context.Context, id string) (*domain.User, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, exceptions.New(exceptions.ErrInvalidID, err)
	}
	return h.userRepository.Get(ctx, parsedId)
}

func (h *userService) List(ctx context.Context, queryParams helpers.QueryParams) (*domain.UserResponse, error) {
	var parsedUserIds []primitive.ObjectID = make([]primitive.ObjectID, 0)
	for _, id := range queryParams.UserIDs {
		parsedUserId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		parsedUserIds = append(parsedUserIds, parsedUserId)
	}

	users, err := h.userRepository.List(ctx, parsedUserIds, queryParams.Limit, queryParams.Offset)
	if err != nil {
		return nil, err
	}

	response := &domain.UserResponse{
		Users: users,
	}

	if len(users) > 0 {
		response.NextPage = queryParams.Offset + 1
	}

	return response, nil
}

func (h *userService) Update(ctx context.Context, id string, request domain.UserPatchRequest) error {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return exceptions.New(exceptions.ErrInvalidID, err)
	}

	err = request.Validate()
	if err != nil {
		return err
	}

	fieldsToUpdate := request.ToBsonM()

	return h.userRepository.Update(ctx, parsedId, fieldsToUpdate)
}

func (h *userService) Delete(ctx context.Context, id string) error {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return exceptions.New(exceptions.ErrInvalidID, err)
	}

	return h.userRepository.Delete(ctx, parsedId)
}
