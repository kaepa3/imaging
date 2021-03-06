package effector

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_Effect(t *testing.T) {
	// ファイル読み込み
	in := inputFile("sampleimage/test.jpg")
	if nil == in {
		t.Error("error read file")
	}

	// outputFile("mono", Monochrome(in))
	// outputFile("revcon", ReverseDensity(in))
	// outputFile("fourtone", FourTone(in, 4))
	// outputFile("linerden", LinearDensity(in, 0x10, 0xFF00))
	// outputFile("unden", UnlinearDensity(in, 0.5))
	// outputFile("contrast", ContrastImprovement(in, 0.5))
	// outputFile("ave", AverageHistogram(in))
	// outputFile("avefil", AverageFilter(in, 1.2))
	// outputFile("media", MedianFilter(in))
	// outputFile("prewitt", PrewittFilter(in))
	// outputFile("vlinefil", VirticalLineFilter(in, 1.2, true))
	// outputFile("hlinefil", HorizontalLineFilter(in, 1.2, true))
	// outputFile("laplacian", LaplacianFilter(in))
	// outputFile("sharpe", SharpeningFilter(in))
	// outputFile("sizekin", ChangeSizeKin(in, 0.8, 0.8))
	// outputFile("sizesen", ChangeSizeSen(in, 0.7, 0.7))
	// outputFile("sizesen", ChangeSizeSen(in, 0.7, 0.7))
	// outputFile("histgram_org", Histogram(in, "test.png hist", "nodo", "dosu"))
}
func Test_EffectTwoVal(t *testing.T) {
	// ファイル読み込み
	in := inputFile("sampleimage/test2.jpg")
	if nil == in {
		t.Error("error read file")
	}
	// outputFile("static_th", StaticThreshold(in, 0xFFFF/2))
	// outputFile("vari_th", VariableThreshold(in, 40, 0x33, false))

	in = inputFile("sampleimage/test_vari_th.png")
	if nil == in {
		t.Error("error read file")
	}

	outputFile("thinning", Thinning(in))
	outputFile("boundary", BoundaryTracking(in))
}

func inputFile(path string) image.Image {
	// ファイル読み込み
	inputFile, err := os.Open(path)
	if nil != err {
		fmt.Println(err)
		return nil
	}
	// decodeの実施
	inputImage, _, err := image.Decode(inputFile)
	if nil != err {
		fmt.Println(err)
		return nil
	}
	inputFile.Close()
	return inputImage
}

func outputFile(append string, outputImage image.Image) bool {
	outputFile, err := os.Create("sampleimage/test_" + append + ".png")
	if nil != err {
		fmt.Println(err)
		return false
	}
	err = png.Encode(outputFile, outputImage) // エンコード

	if nil != err {
		fmt.Println(err)
		return false
	}
	defer outputFile.Close()
	return true
}
