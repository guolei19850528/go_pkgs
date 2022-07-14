package coder

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
	"strings"
)

type Coder interface {
	Decode(r io.Reader) (image.Image, error)
	DecodeByName(name string) (image.Image, error)
	Encode(w io.Writer, m image.Image, o interface{}) error
}

type PngCoder struct {
}

func (ptr *PngCoder) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func (ptr *PngCoder) DecodeByName(name string) (image.Image, error) {
	r, err := os.Open(name)
	if err == nil {
		defer r.Close()
		return ptr.Decode(r)
	}
	return nil, err
}

func (ptr *PngCoder) Encode(w io.Writer, m image.Image, o interface{}) error {
	return png.Encode(w, m)
}

type JpegCoder struct {
}

func (ptr *JpegCoder) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

func (ptr *JpegCoder) DecodeByName(name string) (image.Image, error) {
	r, err := os.Open(name)
	if err == nil {
		defer r.Close()
		return ptr.Decode(r)
	}
	return nil, err
}

func (ptr *JpegCoder) Encode(w io.Writer, m image.Image, o interface{}) error {
	return jpeg.Encode(w, m, o.(*jpeg.Options))
}

type GifCoder struct {
}

func (ptr *GifCoder) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

func (ptr *GifCoder) DecodeByName(name string) (image.Image, error) {
	r, err := os.Open(name)
	if err == nil {
		defer r.Close()
		return ptr.Decode(r)
	}
	return nil, err
}

func (ptr *GifCoder) Encode(w io.Writer, m image.Image, o interface{}) error {
	return gif.Encode(w, m, o.(*gif.Options))
}

func NewCoder(name string) (Coder, error) {
	if strings.ToLower(path.Ext(name)) == ".png" {
		return new(PngCoder), nil
	}
	if strings.ToLower(path.Ext(name)) == ".jpeg" || strings.ToLower(path.Ext(name)) == ".jpg" {
		return new(JpegCoder), nil
	}
	if strings.ToLower(path.Ext(name)) == ".gif" {
		return new(GifCoder), nil
	}
	return nil, fmt.Errorf("%s extension not supported", name)
}

func GetImgByName(name string) (image.Image, error) {
	imgCoder, err := NewCoder(name)
	if err == nil {
		return imgCoder.DecodeByName(name)
	}
	return nil, err
}
