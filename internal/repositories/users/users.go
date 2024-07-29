package users

import (
	"context"
	"errors"

	"github.com/ADAGroupTcc/ms-users-api/exceptions"
	"github.com/ADAGroupTcc/ms-users-api/internal/domain"
	"github.com/ADAGroupTcc/ms-users-api/pkg/mongorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const USER_COLLECTION = "users"

type Repository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Get(ctx context.Context, id primitive.ObjectID) (*domain.User, error)
	List(ctx context.Context, userIds []primitive.ObjectID, limit int64, offset int64) ([]*domain.User, error)
	Update(ctx context.Context, id primitive.ObjectID, fields bson.M) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type userRepository struct {
	db *mongo.Database
}

func New(db *mongo.Database) Repository {
	return &userRepository{db}
}

func (h *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	filter := bson.M{"email": user.Email, "cpf": user.CPF}
	err := user.Read(ctx, h.db, USER_COLLECTION, filter, user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = user.Create(ctx, h.db, USER_COLLECTION, user)
			if err != nil {
				return nil, exceptions.New(exceptions.ErrDatabaseFailure, err)
			}
			return user, nil
		}
	}

	return nil, exceptions.New(exceptions.ErrUserAlreadyExists, err)
}

func (h *userRepository) Get(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	user := &domain.User{}
	err := user.Read(ctx, h.db, USER_COLLECTION, bson.M{"_id": id}, user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, exceptions.New(exceptions.ErrUserNotFound, err)
		}

		return nil, exceptions.New(exceptions.ErrDatabaseFailure, err)
	}
	return user, nil
}

func (h *userRepository) List(ctx context.Context, userIds []primitive.ObjectID, limit int64, offset int64) ([]*domain.User, error) {
	var users []*domain.User = make([]*domain.User, 0)
	var filter bson.M
	if len(userIds) > 0 {
		filter = bson.M{"_id": bson.M{"$in": userIds}}
	}
	err := mongorm.List(ctx, h.db, USER_COLLECTION, filter, &users, options.Find().SetLimit(limit).SetSkip(offset*limit))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return users, nil
		}
		return nil, exceptions.New(exceptions.ErrDatabaseFailure, err)
	}
	return users, nil
}

func (h *userRepository) Update(ctx context.Context, id primitive.ObjectID, fields bson.M) error {
	user := &domain.User{}
	err := user.Update(ctx, h.db, USER_COLLECTION, bson.M{"_id": id}, fields, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return exceptions.New(exceptions.ErrUserNotFound, err)
		}
		return exceptions.New(exceptions.ErrDatabaseFailure, err)
	}
	return nil
}

func (h *userRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	user := &domain.User{}
	err := user.Delete(ctx, h.db, USER_COLLECTION, bson.M{"_id": id})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return exceptions.New(exceptions.ErrUserNotFound, err)
		}
		return exceptions.New(exceptions.ErrDatabaseFailure, err)
	}
	return nil
}
