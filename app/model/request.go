package model

import (
	"context"
	"time"

	"github.com/isayme/go-requestbin/app/constant"
	"github.com/isayme/go-requestbin/app/manager"
	"github.com/isayme/go-requestbin/app/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Request data model for requestinfo
type Request struct {
}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) getCollection(ctx context.Context) *mongo.Collection {
	m := manager.Get()
	c := m.Mongo.GetCollection(constant.CollectionRequest)
	return c
}

func (r *Request) List(ctx context.Context, slug string) ([]*schema.Request, error) {
	c := r.getCollection(ctx)

	option := &options.FindOptions{}
	option.SetLimit(100)
	option.SetSkip(0)
	option.SetSort(bson.D{
		bson.E{
			Key:   "_id",
			Value: -1,
		},
	})
	cursor, err := c.Find(ctx, bson.M{"slug": slug}, option)
	if err != nil {
		return nil, err
	}

	result := []*schema.Request{}
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *Request) Create(ctx context.Context, slug string, info *schema.RequestInfo) (*schema.Request, error) {
	c := r.getCollection(ctx)

	now := time.Now()

	request := &schema.Request{
		Request: info,
	}

	request.ID = primitive.NewObjectID()
	request.Slug = slug

	request.Created = now

	_, err := c.InsertOne(ctx, request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
