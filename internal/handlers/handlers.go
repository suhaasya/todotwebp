package handlers

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// validate token somewhere here

	file, header, err := r.FormFile("image")

	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	isJpeg := header.Header.Get("Content-Type") == "image/jpeg"
	isPng := header.Header.Get("Content-Type") == "image/png"

	isJpegOrPng := !isJpeg && !isPng

	if isJpegOrPng {
		http.Error(w, "File must be a JPEG image", http.StatusBadRequest)
		return
	}

	var img image.Image

	if isJpeg {
		img, err = jpeg.Decode(file)
		if err != nil {
			http.Error(w, "Error while decoding jpeg image: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if isPng {
		img, err = png.Decode(file)
		if err != nil {
			http.Error(w, "Error while decoding png image: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)

	if err != nil {
		http.Error(w, "Error while encoding to webp: "+err.Error(), http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)

	if err := webp.Encode(buf, img, options); err != nil {
		http.Error(w, "Error while encoding to webp: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "image/webp")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.webp\"", strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename))))

	// Send the WebP image
	if _, err := buf.WriteTo(w); err != nil {
		http.Error(w, "Error while sending to response: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
