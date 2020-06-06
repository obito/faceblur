package main

import (
	"image"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/Kagami/go-face"
	"gocv.io/x/gocv"
)

const dataDir = "images"

func main() {
	extList := []string{
		"jpg",
		"png",
		"jpeg",
	}

	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		log.Fatal(err)
	}
	defer rec.Close()

	log.Print("FaceBlur started..")

	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		extSplited := strings.Split(file.Name(), ".")
		ext := extSplited[len(extSplited)-1]

		if !stringInSlice(ext, extList) {
			return
		}

		img := filepath.Join(dataDir, file.Name())

		faces, err := rec.RecognizeFile(img)
		if err != nil {
			log.Fatalf("Can't recognize: %v", err)
		}

		log.Print(file.Name()+" Faces ", len(faces))
		photo := gocv.IMRead(img, gocv.IMReadColor)
		for _, face := range faces {
			imgFace := photo.Region(face.Rectangle)

			gocv.GaussianBlur(imgFace, &imgFace, image.Pt(75, 75), 0, 0, gocv.BorderDefault)
			imgFace.Close()
		}

		gocv.IMWrite("dist/"+extSplited[0]+"-blurred."+ext, photo)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
