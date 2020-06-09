# faceblur


Detect and blur faces in less than a second with the help of OpenCV and dlib.

# Install

You have to install OpenCV. You can see instruction [here](https://github.com/hybridgroup/gocv#how-to-install).

You also have to install dlib, you can get instruction [here](https://github.com/Kagami/go-face#requirements).

# How to use


Clone this repo, put the images that you want to blur in the images folder.

Run the program, and you will get the blurred faces in dist folder.

```
git clone https://github.com/obito/faceblur
cd faceblur
go mod download
go run .
```

* `-cnn` flag is used to tell if it should use CNN for recognition. Default is false.

Unfortunately I can't provide binaries because of OpenCV limitations, I can't cross compile CGO.

# Example

Before             |  After
:-------------------------:|:-------------------------:
![](https://raw.githubusercontent.com/obito/faceblur/master/images/img.jpg)  |  ![](https://raw.githubusercontent.com/obito/faceblur/master/dist/img-blurred.jpg)


Before (without CNN)           |  After (without CNN)
:-------------------------:|:-------------------------:
![](https://raw.githubusercontent.com/obito/faceblur/master/assets/head-pose-face-detection-male.gif)  |  ![](https://raw.githubusercontent.com/obito/faceblur/master/assets/head-pose-face-detection-male-blurred.gif)

# Improve recognition accuracy

* CNN recognizing (takes more time).
* Try different tolerance values of `ClassifyThreshold`
* Try different size/padding/jittering values of `NewRecognizerWithConfig`
* [Train](https://blog.dlib.net/2017/02/high-quality-face-recognition-with-deep.html) network (`dlib_face_recognition_resnet_model_v1.dat`) on your own data.