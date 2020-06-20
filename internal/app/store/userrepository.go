package store

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go_test_learning/internal/app/model"
	"time"
)

type UserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(s *Store) *UserRepository {
	return &UserRepository{
		coll: s.db.Collection("user"),
	}
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if !u.ID.IsZero() {
		return nil, errors.New("user already have ObjectID")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := r.coll.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		insertedUser := *u
		insertedUser.ID = oid
		return &insertedUser, nil
	} else {
		return nil, errors.New("user create has not returned ObjectID")
	}
}

func (r *UserRepository) SelectUser(id primitive.ObjectID) (*model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user *model.User
	err := r.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) SelectUsers(skip int64, limit int64) ([]model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userPipeline := mongo.Pipeline{
		bson.D{{"$skip", skip}},
		bson.D{{"$limit", limit}},
	}

	cur, err := r.coll.Aggregate(ctx, userPipeline)
	var users []model.User
	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {
		var user model.User
		err := cur.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, err
}

func (r *UserRepository) Update(u *model.User) (*model.User, error) {
	if u.ID.IsZero() {
		return nil, errors.New("can not update user without ObjectID")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := r.coll.ReplaceOne(ctx, bson.M{"_id": u.ID}, u)
	if err != nil {
		return nil, err
	}
	if oid, ok := res.UpsertedID.(primitive.ObjectID); ok {
		return r.SelectUser(oid)
	} else {
		return nil, errors.New("user update has not returned ObjectID")
	}
}
