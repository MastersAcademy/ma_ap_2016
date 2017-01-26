package main

import "fmt"

type Image interface {
	DisplayImage()
}

type RealImage struct {
	filename string
}
func NewRealImage(filename string) *RealImage {
	image := &RealImage{filename:filename}
	image.loadImageFromDisk()
	return image
}

func (image *RealImage) loadImageFromDisk() {
	fmt.Println("Load from disk: ", image.filename)
}
func (image *RealImage) DisplayImage() {
	fmt.Println("Displaying ", image.filename)
}

type ProxyImage struct {
	filename string
	image Image
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename:filename}
}

func (pi *ProxyImage) DisplayImage() {
	if pi.image == nil {
		pi.image = NewRealImage(pi.filename)
	}
	pi.image.DisplayImage()
}

func main() {
	var eagerImage, lazyImage Image
	eagerImage = NewRealImage("realImage")   // loadFileFromDisk() immediately
	lazyImage  = NewProxyImage("proxyImage") // loadFileFromDisk() deferred
	// file is already loaded and display gets called directly
	eagerImage.DisplayImage()
	// load file from disk
	// and then forward display call to the real image
	lazyImage.DisplayImage()
}
