package common

import (
	"github.com/gotoeasy/glang/cmn"
	"image"
	"image/color"
	"image/draw"
)

// 图片二值化
func SetImageBinarization(originalImage image.Image) image.Image {
	//// 解码图片
	//originalImage, err := jpeg.Decode(file)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 创建一个新的二值图像
	binaryImage := image.NewGray(originalImage.Bounds())
	draw.Draw(binaryImage, binaryImage.Bounds(), originalImage, image.Point{}, draw.Src)
	// 计算图像的灰度平均值
	var totalGray uint64
	for y := 0; y < binaryImage.Bounds().Dy(); y++ {
		for x := 0; x < binaryImage.Bounds().Dx(); x++ {
			gray := binaryImage.GrayAt(x, y).Y
			totalGray += uint64(gray)
		}
	}
	// 计算平均值
	avgGray := totalGray / uint64(binaryImage.Bounds().Dx()*binaryImage.Bounds().Dy())
	// 设置阈值
	threshold := uint8(avgGray)
	//黑底阈值小于128 需要黑白反转
	cmn.Info("阈值：", threshold)
	// 对每个像素进行阈值化处理
	bounds := originalImage.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldPixel := originalImage.At(x, y)
			grayPixel := color.GrayModel.Convert(oldPixel).(color.Gray)
			newPixel := color.Gray{}
			//阈值大的
			if threshold < 128 {
				if grayPixel.Y > threshold {
					newPixel.Y = 0
				} else {
					newPixel.Y = 255
				}
			} else {
				if grayPixel.Y > threshold {
					newPixel.Y = 255
				} else {
					newPixel.Y = 0
				}
			}

			binaryImage.Set(x, y, newPixel)
		}
	}
	return binaryImage
	//jpeg.Encode(file, binaryImage, &jpeg.Options{Quality: 300})
}
