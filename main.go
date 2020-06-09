package main

import (
	"flag"
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
		"mp4",
	}

	var cnnMode = flag.Bool("cnn", false, "run with ccn mode")
	flag.Parse()

	rec, err := face.NewRecognizer("data")
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

		if ext == "mp4" {
			video, err := gocv.OpenVideoCapture(img)
			if err != nil {
				log.Printf("error opening video: %v\n", err)
				return
			}

			defer video.Close()

			videoHeight := int(video.Get(gocv.VideoCaptureFrameHeight))
			videoWidth := int(video.Get(gocv.VideoCaptureFrameWidth))
			videoWrite, err := gocv.VideoWriterFile("dist/"+extSplited[0]+"-blurred.mp4", "MP4V", video.Get(gocv.VideoCaptureFPS), videoWidth, videoHeight, true)

			if err != nil {
				log.Fatalf("error init write video: %v", err)
			}

			for video.IsOpened() {
				imgMat := gocv.NewMat()

				frame := video.Read(&imgMat)

				if !frame {
					break
				}

				jpgFile, err := gocv.IMEncode(gocv.JPEGFileExt, imgMat)

				if err != nil {
					log.Fatal(err)
				}

				var faces []face.Face
				if *cnnMode {
					faces, err = rec.RecognizeCNN(jpgFile)
				} else {
					faces, err = rec.Recognize(jpgFile)
				}

				if err != nil {
					log.Fatalf("Can't recognize: %v", err)
				}

				log.Print(file.Name()+" Faces ", len(faces))

				for _, face := range faces {
					imgFace := imgMat.Region(face.Rectangle)

					gocv.GaussianBlur(imgFace, &imgFace, image.Pt(75, 75), 0, 0, gocv.BorderDefault)
					imgFace.Close()
				}

				videoWrite.Write(imgMat)
			}

			videoWrite.Close()
			video.Close()
		} else {

			var faces []face.Face
			if *cnnMode {
				faces, err = rec.RecognizeFileCNN(img)
			} else {
				faces, err = rec.RecognizeFile(img)
			}

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
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
