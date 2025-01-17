package inputData

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/CocaineCong/tangseng/app/search_engine/types"
	"github.com/CocaineCong/tangseng/pkg/util/stringutils"
)

// Doc2Struct 从csv读取数据 TODO：后续区分一下输入源，如果是爬虫那边的数据，处理不一样
func Doc2Struct(docStr string) (*types.Document, error) {
	docStr = strings.Replace(docStr, "\"", "", -1)
	d := strings.Split(docStr, ",")
	if len(d) <= 5 { // just fix the stupid data
		for i := 0; i < 6-len(d); i++ {
			d = append(d, "a")
		}
	}
	doc := &types.Document{
		DocId: cast.ToInt64(d[0]),
		Title: d[1],
		Body:  stringutils.StrConcat([]string{d[2], d[3], d[4]}),
	}

	return doc, nil
}
