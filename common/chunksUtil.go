package common

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"os"
)

// 分割文档
/**
dirFile 文件路径
chunkSize 块大小
chunkOverlap 块重叠大小
*/
func FileToChunks(dirFile string, chunkSize, chunkOverlap int) ([]schema.Document, error) {
	file, err := os.Open(dirFile)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	fileType, err := GetFileType(file.Name())
	if err != nil {
		return nil, err
	}
	fInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	// 创建一个新的递归字符文本分割器
	split := textsplitter.NewRecursiveCharacter()
	// 设置块大小
	split.ChunkSize = chunkSize
	// 设置块重叠大小
	split.ChunkOverlap = chunkOverlap

	// 创建一个新的文本文档加载器
	if fileType == "pdf" {
		docLoaded := documentloaders.NewPDF(file, fInfo.Size())
		// 加载并分割文档
		docs, err := docLoaded.LoadAndSplit(context.Background(), split)
		if err != nil {
			return nil, err
		}
		return docs, nil
	} else if fileType == "txt" {
		docLoaded := documentloaders.NewText(file)
		// 加载并分割文档
		docs, err := docLoaded.LoadAndSplit(context.Background(), split)
		if err != nil {
			return nil, err
		}
		return docs, nil
	} else if fileType == "html" {
		docLoaded := documentloaders.NewHTML(file)
		// 加载并分割文档
		docs, err := docLoaded.LoadAndSplit(context.Background(), split)
		if err != nil {
			return nil, err
		}
		return docs, nil
	} else {
		return nil, fmt.Errorf("fileType error，current type is " + fileType)
	}

}
