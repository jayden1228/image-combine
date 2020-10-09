package image

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

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
func CombineImage(imgs []ImageLayer, bgProperty BgProperty) (*image.RGBA, error) {
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
func SaveImage(path string, m image.Image) error {
	var opt jpeg.Options
	opt.Quality = 100
	out, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return err
	}

	return jpeg.Encode(out, m, &opt)
}

// 存储png
func SavePngImage(path string, m image.Image) error {
	out, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return err
	}
	enc := png.Encoder{
		CompressionLevel: png.BestSpeed,
		BufferPool:       nil,
	}
	return enc.Encode(out, m)
}