package function

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"testing"
)

func TestF(t *testing.T) {
	ctx := context.Background()

	a := Article{
		Author: "Isaac Asimov",
		Title:  "The Last Question",
		URL:    "http://www.multivax.com/last_question.html",
	}

	data, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}

	m := Message{Data: base64.StdEncoding.EncodeToString(data)}

	if err := F(ctx, m); err != nil {
		t.Fatal(err)
	}
}
