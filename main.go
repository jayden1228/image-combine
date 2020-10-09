package main

import (
	"image-combine/image"
	"image-combine/model"
	"image-combine/walk"
	"log"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
)

func main() {
	// 遍历所有目录
	result, _ := walk.GetAllFileIncludeSubFolder("/Users/chuangkegongchang/Downloads/game_resource/addition")
	// 过滤空目录
	result = model.FilterEmptyImages(result)
	// 排序
	for k, _ := range result {
		images := result[k].Images
		sort.Slice(images, func(i, j int) bool {
			re := regexp.MustCompile("[0-9]+")
			num1 := gconv.Int(re.Find([]byte(images[i])))
			num2 := gconv.Int(re.Find([]byte(images[j])))
			if num1 < num2 {
				return true
			}
			return false
		})
		result[k].Images = images
	}
	// 拼接图片
	// CombineSprites(result, 1280, 720)
	CombineLevelImage(result, "/Users/chuangkegongchang/Downloads/game_resource/output")
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
		outPath := gfile.Join(v.Dir, combineFileName)
		err = image.SaveImage(outPath, res)
		if err != nil {
			log.Printf("Error creating image file: %+v\n", err)
			return
		}

		// 删除源文件
		//for _, v := range deletePath {
		//	_ = os.Remove(v)
		//}

	log.Println("Image Generated")
}
}


func CombineLevelImage(resources []model.ImageResource, outDir string) {
	for _, v := range resources {
		var deletePath []string
		_, combineName := filepath.Split(v.Dir)
		combineFileName := combineName + ".png"

		if len(v.Images) ==  0 {
			continue
		}

		// 获取长宽
		tempImg, _ := image.LoadImage(v.Images[0])
		width := tempImg.Bounds().Size().X
		height := tempImg.Bounds().Size().Y

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
		outPath := gfile.Join(outDir, combineFileName)
		err = image.SavePngImage(outPath, res)
		if err != nil {
			log.Printf("Error creating image file: %+v\n", err)
			return
		}

		// 删除源文件
		//for _, v := range deletePath {
		//	_ = os.Remove(v)
		//}

		log.Println("Image Generated")
	}
}