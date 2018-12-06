package schema

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// M object map
type M map[string]interface{}

// RequestInfo request info
type RequestInfo struct {
	Method        string `json:"method" bson:"method"`
	Path          string `json:"path" bson:"path"`
	ContentType   string `json:"contentType,omitempty" bson:"contentType"`
	ContentLength int    `json:"contentLength,omitempty" bson:"contentLength"`
	IP            string `json:"ip" bson:"ip"`
	Headers       M      `json:"headers" bson:"headers"`
	Query         M      `json:"query,omitempty" bson:"query"`
	Data          string `json:"data,omitempty" bson:"data"`
	Form          M      `json:"form,omitempty" bson:"form"`
	Body          M      `json:"body,omitempty" bson:"body"`
}

// Request request record in db
type Request struct {
	ID      bson.ObjectId `json:"_id" bson:"_id"`
	Slug    string        `json:"slug" bson:"slug"`
	Request *RequestInfo  `json:"request" bson:"request"`
	Created time.Time     `json:"created" bson:"created"`
}
