package nicelog

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

var testData []byte

func TestWriter_Write(t *testing.T) {
	tbs := [][]byte{
		[]byte("abc\n"),
		[]byte("ab\nc"),
		[]byte("a\nbc"),
		[]byte("abc"),
		[]byte("abc\n"),
		[]byte("abc"),
	}

	for _, tb := range tbs {
		test(tb)
	}

	fmt.Println(string(testData))
	testData = testData[:0]
}

func test(temp []byte) {
	testData = append(testData, temp...)
	// if bytes.Contains(testData, splitFlag) {
	// 	bs := bytes.SplitAfter(testData, splitFlag)
	// 	fmt.Println(string(bs[0]))
	// 	testData = testData[len(bs[0]):]
	// }
}

func TestWriter_AAA(t *testing.T) {
	a := []byte("asfsafa\nsgdzvcx\n")
	t.Error(bytes.IndexByte(a, '\n'))
}

func TestWriter_Write1(t *testing.T) {
	a := New(os.Stdout)
	a.Write([]byte("abc"))
	a.Write([]byte("ab\nc"))
	a.Write([]byte("a\nbc"))
	a.Write([]byte("abc"))
	a.Write([]byte("abc\n"))
	a.Write([]byte("abc"))

}
