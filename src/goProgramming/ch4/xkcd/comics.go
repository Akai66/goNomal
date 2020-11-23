//获取漫画并生成离线索引
package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const XkcdURL = "https://xkcd.com/%d/info.0.json"

type Comic struct {
	Num   int
	Year  string
	Month string
	Title string
	Img   string
}

func GetComics(num int) ([]Comic, error) {
	var comics []Comic
	for i := 1; i <= num; i++ {
		url := fmt.Sprintf(XkcdURL, i)
		var cm Comic
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("get url failed,url:%s,err:%v\n", url, err)
		}
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&cm); err != nil {
			return nil, fmt.Errorf("json decode failed,url:%s,err:%v\n", url, err)
		}
		comics = append(comics, cm)
	}
	return comics, nil
}
