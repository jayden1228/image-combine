package walk

import (
	"image-combine/model"
	"log"
	"os"
	"path/filepath"
)

// 遍历目录
func GetAllFileIncludeSubFolder(folder string) ([]model.ImageResource, error) {
	var result []model.ImageResource

	filepath.Walk(folder, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}

		if fi.IsDir() {
			if !model.ExistDir(result, path) {
				result = append(result, model.ImageResource{
					Dir: path,
				})
			}
		} else {
			//如果想要忽略这个目录，请返回filepath.SkipDir，即：
			//return filepath.SkipDir
			ext := filepath.Ext(path)
			if ext == ".png" || ext == ".jpg" {
				model.AppendImage(result, filepath.Dir(path), path)
			} else {
				_ = os.Remove(path)
			}
		}

		return nil
	})

	return result, nil
}
