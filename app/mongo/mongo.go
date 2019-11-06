package mongo

import (
	"context"
	"time"

	logger "github.com/isayme/go-logger"
	"github.com/isayme/go-requestbin/app/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// Mongo mongo session
type Mongo struct {
	client   *mongo.Client
	database string
}

// NewMongo new mongo session
func NewMongo(config *conf.Mongo) *Mongo {
	logger.Debugw("mongodb", "uri", config.URI)

	timeout := time.Second * 5

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Ping(ctx, readpref.PrimaryPreferred())
	if err != nil {
		panic(err)
	}

	cs, _ := connstring.Parse(config.URI)

	return &Mongo{
		client:   client,
		database: cs.Database,
	}
}

func (m *Mongo) GetCollection(name string) *mongo.Collection {
	return m.client.Database(m.database).Collection(name)
}
