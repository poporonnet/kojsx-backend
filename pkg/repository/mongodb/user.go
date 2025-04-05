package mongodb

import (
	"context"
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/utils"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/repository/mongodb/entity"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	client Client
}

func NewUserRepository(c Client) *UserRepository {
	return &UserRepository{client: c}
}

func (u UserRepository) CreateUser(d domain.User) error {
	role := 2
	if d.IsVerified() {
		role = 1
	}
	if d.IsAdmin() {
		role = 0
	}

	e := entity.User{
		ID:       d.GetID(),
		Name:     d.GetName(),
		Email:    d.GetEmail(),
		Password: d.GetPassword(),
		Role:     role,
	}

	_, err := u.client.Cli.Database("kojs").Collection("user").InsertOne(context.Background(), e)
	if err != nil {
		utils.SugarLogger.Errorf("failed to create user: %v", err)
		return err
	}

	return nil
}

func (u UserRepository) FindAllUsers() ([]domain.User, error) {
	filter := &bson.D{}
	cursor, err := u.client.Cli.Database("kojs").Collection("user").Find(context.Background(), filter)
	if err != nil {
		utils.SugarLogger.Errorf("failed to find users: %v", err)
		return []domain.User{}, fmt.Errorf("failed to find users: %w", err)
	}

	var user []entity.User
	if err := cursor.All(context.Background(), &user); err != nil {
		utils.SugarLogger.Errorf("failed to decode users: %v", err)
		return []domain.User{}, fmt.Errorf("failed to decode users: %w", err)
	}

	res := make([]domain.User, len(user))
	for i, v := range user {
		res[i] = v.ToDomain()
	}
	return res, nil
}

func (u UserRepository) FindUserByID(id id.SnowFlakeID) (*domain.User, error) {
	filter := &bson.M{"_id": id}

	result := u.client.Cli.Database("kojs").Collection("user").FindOne(context.Background(), filter)

	var user entity.User
	if err := result.Decode(&user); err != nil {
		utils.SugarLogger.Errorf("failed to decode user: %v", err)
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}
	res := user.ToDomain()
	return &res, nil
}

func (u UserRepository) FindUserByName(name string) (*domain.User, error) {
	filter := &bson.M{"name": name}

	result := u.client.Cli.Database("kojs").Collection("user").FindOne(context.Background(), filter)

	var user entity.User
	if err := result.Decode(&user); err != nil {
		utils.SugarLogger.Errorf("failed to decode user: %v", err)
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}
	res := user.ToDomain()
	return &res, nil
}

func (u UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	filter := &bson.M{"email": email}

	result := u.client.Cli.Database("kojs").Collection("user").FindOne(context.Background(), filter)

	var user entity.User
	if err := result.Decode(&user); err != nil {
		utils.SugarLogger.Errorf("failed to decode user: %v", err)
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}
	res := user.ToDomain()
	return &res, nil
}

func (u UserRepository) UpdateUser(d domain.User) error {
	//TODO implement me
	panic("implement me")
}
