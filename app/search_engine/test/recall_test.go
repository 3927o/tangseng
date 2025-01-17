package test

import (
	"fmt"
	"testing"

	"github.com/CocaineCong/tangseng/app/search_engine/analyzer"
	"github.com/CocaineCong/tangseng/app/search_engine/recall"
	"github.com/CocaineCong/tangseng/config"
	log "github.com/CocaineCong/tangseng/pkg/logger"
)

func TestMain(m *testing.M) {
	// 这个文件相对于config.yaml的位置
	re := config.ConfigReader{FileName: "../../../config/config.yaml"}
	config.InitConfigForTest(&re)
	analyzer.InitSeg()
	log.InitLog()
	fmt.Println("Write tests on values: ", config.Conf)
	m.Run()
}

func TestRecall(t *testing.T) {
	q := "国家,西游记"
	searchItem, err := recall.SearchRecall(q)
	if err != nil {
		fmt.Println(err)
	}
	for i := range searchItem {
		if searchItem[i].DocId == 0 {
			continue
		}
		fmt.Println(searchItem[i].Score, searchItem[i].DocId, searchItem[i].Content)
	}
}
