package nicelog

import (
	"bytes"
	"io"

	"github.com/sirupsen/logrus"
)

const (
	limit = 1024 // 1kb
)

var (
	splitFlag = []byte("\n")
)

// 非并发安全，需要在每个协程内单独 New 一个实例
// 可以在下层，也就是 writer 中保证线程安全
type Writer struct {
	// 是否有数据写入
	hasWritten bool
	pending    []byte
	writer     io.Writer
}

func New(w io.Writer) *Writer {
	return &Writer{
		pending: make([]byte, 0, limit),
		writer:  w,
	}
}

func (w *Writer) Write(b []byte) (n int, err error) {
	n = len(b)
	// 如果一直收不到换行 就在内存超过 1KB 时刷新一次
	if len(w.pending) >= limit {
		w.Flush()
	}
	w.pending = append(w.pending, b...)
	if idx := bytes.LastIndex(w.pending, splitFlag); idx != -1 {
		w.hasWritten = true
		data := bytes.TrimSpace(bytes.ReplaceAll(w.pending[:idx], []byte("sh: sleep: not found"), []byte("")))
		if len(data) != 0 {
			_, err = w.writer.Write(data)
			if err != nil {
				logrus.WithField("err", err).Error("nicelog write failed")
				return
			}
		}
		w.pending = w.pending[idx+1:]
	}

	return
}

// HasWritten return true if data has been written
func (w *Writer) HasWritten() bool {
	return w.hasWritten
}

// Flush data from memory to writter
func (w *Writer) Flush() {
	if len(w.pending) > 0 {
		_, _ = w.writer.Write(w.pending)
		w.pending = w.pending[:0]
	}
}
