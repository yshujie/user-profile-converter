package mongodb

import (
    "context"
    "time"

    "my-go-tool/pkg/logging"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// LoggingClient is a wrapper around a mongo.Client that logs queries.
type LoggingClient struct {
    *mongo.Client
}

// LoggingDatabase is a wrapper around a mongo.Database that logs queries.
type LoggingDatabase struct {
    *mongo.Database
}

// LoggingCollection is a wrapper around a mongo.Collection that logs queries.
type LoggingCollection struct {
    *mongo.Collection
}

func Connect() (*LoggingClient, error) {
    uri := os.Getenv("MONGODB_URI")
    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        return nil, err
    }

    return &LoggingClient{client}, nil
}

func (c *LoggingClient) Database(name string, opts ...*options.DatabaseOptions) *LoggingDatabase {
    db := c.Client.Database(name, opts...)
    return &LoggingDatabase{db}
}

func (db *LoggingDatabase) Collection(name string, opts ...*options.CollectionOptions) *LoggingCollection {
    coll := db.Database.Collection(name, opts...)
    return &LoggingCollection{coll}
}

func (c *LoggingCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
    start := time.Now()
    result, err := c.Collection.InsertOne(ctx, document, opts...)
    elapsed := time.Since(start)
    logging.LogQuery("InsertOne", document, elapsed, err)
    return result, err
}

func (c *LoggingCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
    start := time.Now()
    cursor, err := c.Collection.Find(ctx, filter, opts...)
    elapsed := time.Since(start)
    logging.LogQuery("Find", filter, elapsed, err)
    return cursor, err
}

func (c *LoggingCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
    start := time.Now()
    result := c.Collection.FindOne(ctx, filter, opts...)
    elapsed := time.Since(start)
    logging.LogQuery("FindOne", filter, elapsed, result.Err())
    return result
}