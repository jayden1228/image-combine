package main

import (
	"image-combine/image"
	"image-combine/model"
	"image-combine/walk"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 遍历所有目录
	result, _ := walk.GetAllFileIncludeSubFolder("/Users/chuangkegongchang/Downloads/Reveal")
	// 过滤空目录
	result = model.FilterEmptyImages(result)
	// 拼接图片
	CombineSprites(result, 1280, 720)
}

func CombineSprites(resources []model.ImageResource, width int, height int) {
	for _, v := range resources {
		var deletePath []string
		_, combineName := filepath.Split(v.Dir)
		combineFileName := combineName + ".jpg"

		var images []image.ImageLayer
		for i, v := range v.Images {
			deletePath = append(deletePath, v)
			tempImg, _ := image.LoadImage(v)
			images = append(images, image.ImageLayer{
				Image: tempImg,
				XPos:  i * width,
				YPos:  0,
			})
		}

		bg := image.BgProperty{
			Width:  width * len(images),
			Length: height,
		}
		// 合成图片
		res, err := image.CombineImage(images, bg)
		if err != nil {
			log.Printf("Error generating banner: %+v\n", err)
		}

		// 保存图片
		err = image.SaveImage(combineFileName, v.Dir, res)
		if err != nil {
			log.Printf("Error creating image file: %+v\n", err)
			return
		}

		// 删除源文件
		for _, v := range deletePath {
			_ = os.Remove(v)
		}

		log.Println("Image Generated")
	}
}
