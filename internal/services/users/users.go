package users

import (
	"context"

	"github.com/ADAGroupTcc/ms-users-api/exceptions"
	"github.com/ADAGroupTcc/ms-users-api/internal/domain"
	"github.com/ADAGroupTcc/ms-users-api/internal/repositories/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Create(ctx context.Context, request domain.UserRequest) (*domain.UserWithCategories, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	List(ctx context.Context, userIds []string, limit int64, offset int64) (*domain.UserResponse, error)
	ListWithCategories(ctx context.Context, userIds []string) ([]*domain.UserWithCategories, error)
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

func (h *userService) Create(ctx context.Context, request domain.UserRequest) (*domain.UserWithCategories, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	user := request.ToUser()

	user, err = h.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	userWithCategory, err := h.userRepository.Aggregate(ctx, []primitive.ObjectID{user.ID})
	if err != nil {
		return nil, err
	}

	if len(userWithCategory) < 1 {
		return nil, exceptions.New(exceptions.ErrDatabaseFailure, err)
	}

	return userWithCategory[0], err
}

func (h *userService) Get(ctx context.Context, id string) (*domain.User, error) {
	parsedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, exceptions.New(exceptions.ErrInvalidID, err)
	}
	return h.userRepository.Get(ctx, parsedId)
}

func (h *userService) List(ctx context.Context, userIds []string, limit int64, offset int64) (*domain.UserResponse, error) {
	parsedIds := ParseStringIdsToObjectId(userIds)

	users, err := h.userRepository.List(ctx, parsedIds, limit, offset)
	if err != nil {
		return nil, err
	}

	response := &domain.UserResponse{
		Users: users,
	}

	if len(users) > 0 {
		response.NextPage = offset + 1
	}

	return response, nil
}

func (h *userService) ListWithCategories(ctx context.Context, userIds []string) ([]*domain.UserWithCategories, error) {
	parsedIds := ParseStringIdsToObjectId(userIds)
	return h.userRepository.Aggregate(ctx, parsedIds)
}

func ParseStringIdsToObjectId(ids []string) []primitive.ObjectID {
	var parsedIds []primitive.ObjectID = make([]primitive.ObjectID, 0)
	for _, id := range ids {
		parsedId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		parsedIds = append(parsedIds, parsedId)
	}
	return parsedIds
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
