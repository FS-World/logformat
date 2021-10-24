package logformat

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

type SerializedLogWriter struct {
	w *msgp.Writer
	f *os.File
}

// SerializedLogReader is a reader for log files
type SerializedLogReader struct {
	r        *msgp.Reader // r is the reader instance
	f        *os.File     // f is the file we read from
	dataType Serializable // dataType is the type of logdata we want to decode to
}

// NewSerializedLogWriter creates a new LogWriter
func NewSerializedLogWriter(filename string) (*SerializedLogWriter, error) {
	w := new(SerializedLogWriter)
	var err error
	w.f, err = os.Create(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file %s for writing", filename)
	}
	w.w = msgp.NewWriter(w.f)
	return w, nil
}

// Write writes a message to the logfile
func (w *SerializedLogWriter) Write(msg Serializable) error {
	return msg.EncodeMsg(w.w)
}

// Sync flushes the internal buffers and syncs the logfile to disk
func (w *SerializedLogWriter) Sync() error {
	err := w.w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing buffers")
	}
	return errors.Wrap(w.f.Sync(), "error syncing file to disk")
}

// Close syncs and closes the logfile
func (w *SerializedLogWriter) Close() error {
	err := w.Sync()
	if err != nil {
		return errors.Wrap(err, "error syncing file before close")
	}
	return errors.Wrap(w.f.Close(), "error closing file")
}

// NewSerializedLogReader creates a new SerializedLogReader from the provided file for the provided dataType
func NewSerializedLogReader(filename string, dataType Serializable) (*SerializedLogReader, error) {
	r := new(SerializedLogReader)
	r.dataType = dataType
	var err error
	r.f, err = os.Open(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file %s for reading", filename)
	}
	r.r = msgp.NewReader(r.f)
	return r, nil
}

// ReadAll reads all messages from the log
func (r *SerializedLogReader) ReadAll() ([]Serializable, error) {
	var msgs []Serializable
	for {
		err := r.dataType.DecodeMsg(r.r)
		if err == msgp.WrapError(io.EOF) {
			return msgs, nil
		} else if err != nil {
			return msgs, errors.Wrap(err, "error decoding msg")
		}
		msgs = append(msgs, r.dataType)
	}
}

// Close closes the underlying file handle
func (r *SerializedLogReader) Close() error {
	return r.f.Close()
}
