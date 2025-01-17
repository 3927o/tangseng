package index

import (
	"fmt"

	"github.com/CocaineCong/tangseng/app/search_engine/engine"
	inputData "github.com/CocaineCong/tangseng/app/search_engine/inputdata"
	"github.com/CocaineCong/tangseng/app/search_engine/segment"
	"github.com/CocaineCong/tangseng/app/search_engine/types"
	"github.com/CocaineCong/tangseng/config"
	logs "github.com/CocaineCong/tangseng/pkg/logger"
)

// IndexEngine 构建索引的引擎
type IndexEngine struct {
	*engine.Engine
	*engine.Meta
}

// NewIndexEngine init
func NewIndexEngine(meta *engine.Meta) *IndexEngine {
	return &IndexEngine{
		Engine: engine.NewTangSengEngine(meta, segment.IndexMode),
		Meta:   meta,
	}
}

// AddDoc 读取配置文件，进行doc文件转成struct
func AddDoc(in *IndexEngine) {
	// TODO: 后续配置文件改成多选择的
	docList := inputData.ReadFiles([]string{config.Conf.SeConfig.SourceWuKoFile})
	go in.Scheduler.Merge()
	// wg := new(sync.WaitGroup) // TODO: 后续改成并发的，稍微留意一下map的一些结构体字段
	for _, item := range docList[1:] {
		// wg.Add(1)
		// go func(item string) {
		doc, err := inputData.Doc2Struct(item)
		if err != nil {
			logs.LogrusObj.Errorf("index addDoc doc2Struct: %v", err)
		}

		err = in.AddDocument(doc)
		// }(item)
	}
	// wg.Wait()
	// 读取结束 写入磁盘
	err := in.FlushDict(true)
	if err != nil {
		logs.LogrusObj.Errorf("AddDoc-FlushDict: %v", err)
		return
	}

	err = in.FlushInvertedIndex(true)
	if err != nil {
		logs.LogrusObj.Errorf("AddDoc-FlushInvertedIndex: %v", err)
		return
	}
}

// AddDocument 添加文档
func (in *IndexEngine) AddDocument(doc *types.Document) (err error) {
	if doc == nil || doc.DocId <= 0 || doc.Title == "" {
		return fmt.Errorf("doc err: doc || doc_id || title")
	}
	err = in.AddForwardIndex(doc)
	if err != nil {
		logs.LogrusObj.Errorf("forward doc add err:%v", err)
		return
	}

	err = in.Text2PostingsLists(doc.Body, doc.DocId)
	if err != nil {
		logs.LogrusObj.Errorf("Text2PostingsLists:%v", err)
		return
	}

	return
}

// Close --
func (in *IndexEngine) Close() {
	in.Engine.Close()
}
