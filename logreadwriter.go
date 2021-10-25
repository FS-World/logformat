package logformat

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
	"github.com/zeebo/blake3"
)

// SerializedLogWriter writes Serializable messages to a logfile and hashes the logfile in the background.
type SerializedLogWriter struct {
	w *msgp.Writer   // w is the writer instance
	f *os.File       // f is the file we read from
	h *blake3.Hasher // h is the hash function
}

// SerializedLogReader is a reader for log files with Serializable messages. It hashes the input file in the background.
type SerializedLogReader struct {
	r        *msgp.Reader   // r is the reader instance
	f        *os.File       // f is the file we read from
	dataType Serializable   // dataType is the type of logdata we want to decode to
	h        *blake3.Hasher // h is the hash function
}

// NewSerializedLogWriter creates a new LogWriter
func NewSerializedLogWriter(filename string) (*SerializedLogWriter, error) {
	w := new(SerializedLogWriter)
	var err error
	w.f, err = os.Create(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file %s for writing", filename)
	}
	w.h = blake3.New()
	w.w = msgp.NewWriter(io.MultiWriter(w.f, w.h))
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

// Close syncs and closes the logfile; returns the hash of the file
func (w *SerializedLogWriter) Close() ([]byte, error) {
	err := w.Sync()
	if err != nil {
		return nil, errors.Wrap(err, "error syncing file before close")
	}
	return w.h.Sum(nil), errors.Wrap(w.f.Close(), "error closing file")
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
	r.h = blake3.New()
	r.r = msgp.NewReader(io.TeeReader(r.f, r.h))
	return r, nil
}

// ReadAll reads all messages from the log; returns the data and blake3 hash on success
func (r *SerializedLogReader) ReadAll() ([]Serializable, []byte, error) {
	var msgs []Serializable
	for {
		err := r.dataType.DecodeMsg(r.r)
		if err == msgp.WrapError(io.EOF) {
			return msgs, r.h.Sum(nil), nil
		} else if err != nil {
			return msgs, nil, errors.Wrap(err, "error decoding msg")
		}
		msgs = append(msgs, r.dataType)
	}
}

// Close closes the underlying file handle
func (r *SerializedLogReader) Close() error {
	return r.f.Close()
}
