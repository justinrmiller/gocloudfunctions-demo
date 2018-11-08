package function

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"

	"github.com/nfnt/resize"
)

// F prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func F(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer file.Close()
	originalImage, _, err := image.Decode(file)

	newImage := resize.Resize(500, 0, originalImage, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImage, nil)

	w.Header().Set("Content-Length", fmt.Sprint(buf.Len())) /* value: 7007 */
	w.Header().Set("Content-Type", "image/jpg")             /* value: image/png */
	w.Write(buf.Bytes())
}
