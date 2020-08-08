package model

// 关卡图片
type ImageResource struct {
	Dir    string
	Images []string
}

// 是否存在目录
func ExistDir(list []ImageResource, dir string) bool {
	for _, v := range list {
		if v.Dir == dir {
			return true
		}
	}
	return false
}

// 添加图片
func AppendImage(list []ImageResource, dir string, image string) {
	for i, v := range list {
		if v.Dir == dir {
			list[i].Images = append(list[i].Images, image)
		}
	}
}

// slice删除
func Remove(s []ImageResource, i int) []ImageResource {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// 删除空目录记录
func FilterEmptyImages(input []ImageResource) []ImageResource {
	for i := len(input) - 1; i >= 0; i-- {
		if len(input[i].Images) == 0 {
			input = Remove(input, i)
		}
	}
	return input
}
