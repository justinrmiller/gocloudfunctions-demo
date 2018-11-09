package function

import (
	"fmt"
	"log"
	"net/http"

	"context"

	vision "cloud.google.com/go/vision/apiv1"
)

// F prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func F(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {

		log.Fatalf("Failed to create client: %v", err)
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer file.Close()

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	for _, label := range labels {
		fmt.Fprint(w, label.Description)
	}
}
