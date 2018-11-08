package function

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"cloud.google.com/go/datastore"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

type Message struct {
	Data string `json:"data"`
}

type Article struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author string `json:"author"`
}

func init() {
	config = &configuration{}
}

var (
	config *configuration
)

// configFunc sets the global configuration; it's overridden in tests.
var configFunc = defaultConfigFunc

type configuration struct {
	datastoreClient *datastore.Client
	err             error
	once            sync.Once
}

type envError struct {
	name string
}

func (e *envError) Error() string {
	return fmt.Sprintf("%s environment variable unset or missing", e.name)
}

func F(ctx context.Context, m Message) error {
	config.once.Do(func() { configFunc() })

	data, err := base64.StdEncoding.DecodeString(m.Data)
	if err != nil {
		return err
	}

	var article Article
	err = json.Unmarshal(data, &article)
	if err != nil {
		return err
	}

	key := datastore.NameKey("Article", article.Title, nil)

	_, err = config.datastoreClient.Put(ctx, key, &article)
	if err != nil {
		return err
	}

	return nil
}

func defaultConfigFunc() {
	projectId := os.Getenv("GCP_PROJECT")
	if projectId == "" {
		config.err = &envError{"GCP_PROJECT"}
		return
	}

	stackdriverExporter, err := stackdriver.NewExporter(stackdriver.Options{ProjectID: projectId})
	if err != nil {
		config.err = err
		return
	}

	trace.RegisterExporter(stackdriverExporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	client, err := datastore.NewClient(context.Background(), projectId)
	if err != nil {
		config.err = err
		return
	}

	config.datastoreClient = client
}
