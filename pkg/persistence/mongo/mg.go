package mongo

import (
	"dreamt/pkg/models"
	"dreamt/pkg/persistence"
	"fmt"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionDreams = "dreams"
)

type MGController struct {
	uri          string
	database     string
	queryTimeout time.Duration
	persistence.DatabaseController
}

func NewMGController(uri string, queryTimeout time.Duration) persistence.DatabaseController {
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	return MGController{uri: uri, queryTimeout: queryTimeout, database: "test"}
}

func (m MGController) close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func (m MGController) connect() (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.queryTimeout)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	return client, ctx, cancel, err
}

func (m MGController) query(client *mongo.Client, ctx context.Context,
	coll string, query, field interface{}) (*mongo.Cursor, error) {
	collection := client.Database(m.database).Collection(coll)
	return collection.Find(ctx, query, options.Find().SetProjection(field))
}

func (m MGController) GetDreams() ([]models.DreamHeader, error) {
	fmt.Println("I'm in get  dreams")
	dreamHeaders := []models.DreamHeader{}
	// get document from mongo
	client, ctx, cancel, err := m.connect()
	if err != nil {
		panic(err)
	}
	defer m.close(client, ctx, cancel)

	var filter, option interface{}

	filter = bson.D{
		// {"maths", bson.D{{"$gt", 70}}},
	}
	option = bson.D{
		// {"_id", 0},
	}

	curr, err := m.query(client, ctx, collectionDreams, filter, option)
	if err != nil {
		return nil, err
	}

	var results []bson.D
	var resultsD []persistence.DreamMG

	if err := curr.All(ctx, &resultsD); err != nil {
		return nil, err
	}

	fmt.Println("Result of Fetch operation")
	fmt.Println(results)

	for _, res := range resultsD {
		dreamHeaders = append(dreamHeaders, models.DreamHeader{
			ID:    res.ID,
			Title: res.Title,
		})
	}

	return dreamHeaders, nil
}

func (m MGController) WriteDreams(dream models.Dream) (int64, error) {
	client, ctx, cancel, err := m.connect()
	if err != nil {
		panic(err)
	}
	defer m.close(client, ctx, cancel)

	mgDoc := persistence.DreamMG{
		Title:   dream.Title,
		Content: dream.Content,
		Tags:    dream.Tags,
	}

	collection := client.Database(m.database).Collection(collectionDreams)
	res, err := collection.InsertOne(ctx, mgDoc)
	if err != nil {
		return -1, err
	}

	fmt.Println("Result of InsertOne")
	fmt.Println(res)

	return -1, nil
}

func (m MGController) GetDream(title string) (models.Dream, error) {
	fmt.Println("I'm in get single dream")
	// get document from mongo
	client, ctx, cancel, err := m.connect()
	if err != nil {
		panic(err)
	}
	defer m.close(client, ctx, cancel)

	// var results []bson.D
	var resultsD persistence.DreamMG

	collection := client.Database(m.database).Collection(collectionDreams)

	if err := collection.FindOne(ctx, bson.M{"title": title}).Decode(&resultsD); err != nil {
		return models.Dream{}, err
	}

	fmt.Println("Result of Fetch operation")
	fmt.Println(resultsD)

	return models.Dream{
		ID:      resultsD.ID,
		Title:   resultsD.Title,
		Content: resultsD.Content,
		Tags:    resultsD.Tags,
	}, nil
}

func (m MGController) DeleteDream(title string) error {
	fmt.Println("I'm in get single dream")
	// get document from mongo
	client, ctx, cancel, err := m.connect()
	if err != nil {
		panic(err)
	}
	defer m.close(client, ctx, cancel)

	collection := client.Database(m.database).Collection(collectionDreams)

	res, err := collection.DeleteOne(ctx, bson.M{"title": title})
	if err != nil {
		return err
	}

	fmt.Println("delete count")
	fmt.Println(res.DeletedCount)

	return nil
}
