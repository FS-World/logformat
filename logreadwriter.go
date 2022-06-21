package logformat

import (
	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
	"github.com/zeebo/blake3"
	"io"
	"os"
)

// SerializedLogWriter writes Serializable messages to a logfile and hashes the logfile in the background.
type SerializedLogWriter struct {
	w *msgp.Writer   // w is the writer instance
	f *os.File       // f is the file we read from
	h *blake3.Hasher // h is the hash function
}

// SerializedLogReader is a reader for log files with Serializable messages. It hashes the input file in the background.
type SerializedLogReader struct {
	r *msgp.Reader   // r is the reader instance
	f *os.File       // f is the file we read from
	h *blake3.Hasher // h is the hash function
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
	w.w = msgp.NewWriter(io.MultiWriter(w.h, w.f))
	return w, nil
}

// Write writes a message to the logfile
func (w *SerializedLogWriter) Write(msg msgp.Encodable) error {
	err := msg.EncodeMsg(w.w)
	if err != nil {
		return err
	}
	err = w.w.Flush()
	return err
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
func NewSerializedLogReader(filename string) (*SerializedLogReader, error) {
	r := new(SerializedLogReader)
	var err error
	r.f, err = os.Open(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file %s for reading", filename)
	}
	r.h = blake3.New()
	r.r = msgp.NewReader(io.TeeReader(r.f, r.h))
	return r, nil
}

// ReadAllCanFrames reads messages of type CANFrame from logfile.
// return slice of all messages, b3 hash of logfile and error
func (r *SerializedLogReader) ReadAllCanFrames() ([]CANFrame, []byte, error) {
	var msgs []CANFrame
	for {
		var data CANFrame
		err := data.DecodeMsg(r.r)
		if err == msgp.WrapError(io.EOF) {
			return msgs, r.h.Sum(nil), nil
		} else if err != nil {
			return msgs, nil, errors.Wrap(err, "error decoding msg")
		}
		msgs = append(msgs, data)
	}
}

// ReadAllIvtBufferedMessages reads messages of type IvtBufferedMeasurement from logfile.
// return slice of all messages, b3 hash of logfile and error
func (r *SerializedLogReader) ReadAllIvtBufferedMessages() ([]IvtBufferedMeasurement, []byte, error) {
	var msgs []IvtBufferedMeasurement
	for {
		var data IvtBufferedMeasurement
		err := data.DecodeMsg(r.r)
		if err == msgp.WrapError(io.EOF) {
			return msgs, r.h.Sum(nil), nil
		} else if err != nil {
			return msgs, nil, errors.Wrap(err, "error decoding msg")
		}
		msgs = append(msgs, data)
	}
}

// Close closes the underlying file handle
func (r *SerializedLogReader) Close() error {
	return r.f.Close()
}
