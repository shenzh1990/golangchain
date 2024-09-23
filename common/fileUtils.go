package common

import (
	"github.com/h2non/filetype"
	"io/ioutil"
	"os"
)

func GetFileType(filename string) (string, error) {
	fb, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	kind, err := filetype.Match(fb)
	if err != nil {
		return "", err
	}
	return kind.Extension, nil

}
func IsTempFileEmpty(filePath string) (bool, error) {
	// 打开临时文件
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// 获取文件信息
	stat, err := file.Stat()
	if err != nil {
		return false, err
	}
	// 如果文件大小为0，则文件为空
	return stat.Size() == 0, nil
}

//tempfilename := tempfile.Name()
//file, _ := os.Open(tempfile.Name())
//defer file.Close()
//file_type, _ := common.GetFileType(tempfile.Name())
//cmn.Info("file_type:", file_type)
//if file_type == "jpg" {
//originalImage, err := jpeg.Decode(file)
//if err == nil {
//tempbinary, err := ioutil.TempFile("", "ocrserver-binary-*.jpg")
//defer func() {
//tempbinary.Close()
//os.Remove(tempbinary.Name())
//}()
//if err == nil {
//err = jpeg.Encode(tempbinary, common.SetImageBinarization(originalImage), &jpeg.Options{Quality: 300}) // 使用jpeg.Encode保存JPEG格式
//if err == nil {
//tempfilename = tempbinary.Name()
//}
//}
//}
//} else if file_type == "png" {
//originalImage, err := png.Decode(file)
//if err == nil {
//tempbinary, err := ioutil.TempFile("", "ocrserver-binary-*.png")
//defer func() {
//tempbinary.Close()
//os.Remove(tempbinary.Name())
//}()
//if err == nil {
//err = png.Encode(tempbinary, common.SetImageBinarization(originalImage))
//if err == nil {
//tempfilename = tempbinary.Name()
//}
//}
//}
//}
