package net

import (
	"bytes"
	"net"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockConn mocks net.Conn
type MockConn struct {
	mock.Mock

	// ReadBuffer can be used to mock data reads. If this buffer has any data,
	// calls to Read will read from this buffer.
	ReadBuffer bytes.Buffer

	// ReadError is the error returned from Read.
	ReadError error

	// WriteBuffer will successfully write data if WriteError is nil.
	WriteBuffer bytes.Buffer

	// WriteError will be returned from Write if it is not nil.
	WriteError error

	// CloseError is returned from Close.
	CloseError error

	// Local mocks calls to LocalAddr.
	Local net.Addr

	// Remote mocks calls to RemoteAddr.
	Remote net.Addr

	// DeadlineError is returned from SetDeadline.
	DeadlineError error

	// ReadDeadlineError is returned from SetReadDeadline.
	ReadDeadlineError error

	// WriteDeadlineError is returned from SetWriteDeadline.
	WriteDeadlineError error
}

// Read mocks calls to conn.Read.
func (m *MockConn) Read(b []byte) (n int, err error) {
	m.Called(b)

	// Read from test buffer if it has data
	if m.ReadBuffer.Len() > 0 {
		return m.ReadBuffer.Read(b)
	}

	// Return read error
	return 0, m.ReadError
}

// Write mocks calls to Write.
func (m *MockConn) Write(b []byte) (n int, err error) {
	m.Called(b)

	// Return error if exists
	if m.WriteError != nil {
		return 0, m.WriteError
	}

	// Write into buffer
	return m.WriteBuffer.Write(b)
}

// Close mocks connection.Close.
func (m *MockConn) Close() error {
	m.Called()
	return m.CloseError
}

// LocalAddr mocks conn.LocalAddr.
func (m *MockConn) LocalAddr() net.Addr {
	m.Called()
	return m.Local
}

// RemoteAddr mocks conn.RemtoeAddr.
func (m *MockConn) RemoteAddr() net.Addr {
	m.Called()
	return m.Remote
}

// SetDeadline mocks calls to conn.SetDeadline.
func (m *MockConn) SetDeadline(t time.Time) error {
	m.Called(t)
	return m.DeadlineError
}

// SetReadDeadline mocks calls to conn.SetReadDeadline.
func (m *MockConn) SetReadDeadline(t time.Time) error {
	m.Called(t)
	return m.ReadDeadlineError
}

// SetWriteDeadline mocks calls to conn.SetWriteDeadline.
func (m *MockConn) SetWriteDeadline(t time.Time) error {
	m.Called(t)
	return m.WriteDeadlineError
}
