package helpers

import (
	"errors"
	_ "errors"
	"fmt"
	_ "io/ioutil"
	_ "net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
)

// 存储类型

// 更多存储类型有待扩展
const (
	Version           = "1.0"
	StoreLocal string = "local"
	StoreOss   string = "oss"
)

var (
	BasePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	StoreType   = beego.AppConfig.String("store_type") // 存储类型
)

// 评分处理
func ScoreFloat(score int) string {
	return fmt.Sprintf("%1.1f", float32(score)/10.0)
}

// 操作图片显示
// 如果用的是oss存储，这style是avatar、cover可选项
func ShowImg(img string, style ...string) (url string) {
	if strings.HasPrefix(img, "https://") || strings.HasPrefix(img, "http://") {
		return img
	}
	img = "/" + strings.TrimLeft(img, "./")
	switch StoreType {
	case StoreOss:
		s := ""
		if len(style) > 0 && strings.TrimSpace(style[0]) != "" {
			s = "/" + style[0]
		}
		url = strings.TrimRight(beego.AppConfig.String("oss::Domain"), "/ ") + img + s
	case StoreLocal:
		url = img
	}
	fmt.Println(img)
	fmt.Println(url)
	return url
}

// Substr returns the substr from start to length.
func Substr(s string, length int) string {
	bt := []rune(s)
	start := 0
	dot := false

	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
		dot = true
	}

	str := string(bt[start:end])
	if dot {
		str = str + "..."
	}
	return str
}

// 判断数据是否在map中
func InMap(maps map[int]bool, key int) (ret bool) {
	if _, ok := maps[key]; ok {
		return true
	}
	return
}

// 判断 value 数据是否在 target 中，target支持的类型 array,slice,map
func IsContain(value interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == value {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(value)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

func IsContainArray(value interface{}, traget interface{}) {

}
