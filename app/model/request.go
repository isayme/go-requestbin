package model

import (
	"context"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/isayme/go-requestbin/app/constant"
	"github.com/isayme/go-requestbin/app/manager"
	"github.com/isayme/go-requestbin/app/mongo"
	"github.com/isayme/go-requestbin/app/schema"
)

// Request data model for requestinfo
type Request struct {
}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) getCollection(ctx context.Context) (*mongo.Session, *mgo.Collection) {
	m := manager.Get()
	s := m.Session.Copy()
	c := s.GetCollection(constant.CollectionRequest)
	return s, c
}

func (r *Request) List(ctx context.Context, slug string) ([]*schema.Request, error) {
	s, c := r.getCollection(ctx)
	defer s.Close()

	result := []*schema.Request{}
	err := c.Find(bson.M{"slug": slug}).Limit(100).Sort("-_id").All(&result)
	return result, err
}

func (r *Request) Create(ctx context.Context, slug string, info *schema.RequestInfo) (*schema.Request, error) {
	s, c := r.getCollection(ctx)
	defer s.Close()

	now := bson.Now()

	request := &schema.Request{
		Request: info,
	}

	request.ID = bson.NewObjectId()
	request.Slug = slug

	request.Created = now

	if err := c.Insert(request); err != nil {
		return nil, err
	}

	return request, nil
}
