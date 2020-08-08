package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func main() {

	tempImg1, err := LoadImage("./Antelope canyon Close.jpg")
	if err != nil {
		log.Println(err)
		return
	}

	tempImg2, err := LoadImage("./Antelope canyon Far.jpg")
	if err != nil {
		log.Println(err)
		return
	}
	imgs := []ImageLayer{
		ImageLayer{
			Image: tempImg1,
			XPos:  0,
			YPos:  0,
		},
		ImageLayer{
			Image: tempImg2,
			XPos:  1280,
			YPos:  0,
		},
	}

	bg := BgProperty{
		Width:  1280 * 2,
		Length: 720,
	}

	res, err := ImageCombine(imgs, bg)
	if err != nil {
		log.Printf("Error generating banner: %+v\n", err)
	}

	err = SaveImage("output.png", res)
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return
	}

	log.Println("Image Generated")
}

//素材
type ImageLayer struct {
	Image image.Image
	XPos  int
	YPos  int
}

//背景
type BgProperty struct {
	Width  int
	Length int
}

// 图片合并
func ImageCombine(imgs []ImageLayer, bgProperty BgProperty) (*image.RGBA, error) {
	//创建背景
	bgImg := image.NewRGBA(image.Rect(0, 0, bgProperty.Width, bgProperty.Length))

	//创建需要合并图片
	for _, img := range imgs {
		//set image offset
		offset := image.Pt(img.XPos, img.YPos)

		//combine the image
		draw.Draw(bgImg, img.Image.Bounds().Add(offset), img.Image, image.ZP, draw.Over)
	}

	return bgImg, nil
}

// 加载图片
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// 存储
func SaveImage(f string, m image.Image) error {
	var opt jpeg.Options
	opt.Quality = 80
	out, err := os.Create("./" + f)
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return err
	}

	return jpeg.Encode(out, m, &opt)
}
