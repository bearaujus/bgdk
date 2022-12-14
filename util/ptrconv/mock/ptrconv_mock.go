// Code generated by MockGen. DO NOT EDIT.
// Source: util/ptrconv/init.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPtrConv is a mock of PtrConv interface.
type MockPtrConv struct {
	ctrl     *gomock.Controller
	recorder *MockPtrConvMockRecorder
}

// MockPtrConvMockRecorder is the mock recorder for MockPtrConv.
type MockPtrConvMockRecorder struct {
	mock *MockPtrConv
}

// NewMockPtrConv creates a new mock instance.
func NewMockPtrConv(ctrl *gomock.Controller) *MockPtrConv {
	mock := &MockPtrConv{ctrl: ctrl}
	mock.recorder = &MockPtrConvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPtrConv) EXPECT() *MockPtrConvMockRecorder {
	return m.recorder
}

// ConvBool mocks base method.
func (m *MockPtrConv) ConvBool(v bool) *bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvBool", v)
	ret0, _ := ret[0].(*bool)
	return ret0
}

// ConvBool indicates an expected call of ConvBool.
func (mr *MockPtrConvMockRecorder) ConvBool(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvBool", reflect.TypeOf((*MockPtrConv)(nil).ConvBool), v)
}

// ConvByte mocks base method.
func (m *MockPtrConv) ConvByte(v byte) *byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvByte", v)
	ret0, _ := ret[0].(*byte)
	return ret0
}

// ConvByte indicates an expected call of ConvByte.
func (mr *MockPtrConvMockRecorder) ConvByte(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvByte", reflect.TypeOf((*MockPtrConv)(nil).ConvByte), v)
}

// ConvComplex128 mocks base method.
func (m *MockPtrConv) ConvComplex128(v complex128) *complex128 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvComplex128", v)
	ret0, _ := ret[0].(*complex128)
	return ret0
}

// ConvComplex128 indicates an expected call of ConvComplex128.
func (mr *MockPtrConvMockRecorder) ConvComplex128(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvComplex128", reflect.TypeOf((*MockPtrConv)(nil).ConvComplex128), v)
}

// ConvComplex64 mocks base method.
func (m *MockPtrConv) ConvComplex64(v complex64) *complex64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvComplex64", v)
	ret0, _ := ret[0].(*complex64)
	return ret0
}

// ConvComplex64 indicates an expected call of ConvComplex64.
func (mr *MockPtrConvMockRecorder) ConvComplex64(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvComplex64", reflect.TypeOf((*MockPtrConv)(nil).ConvComplex64), v)
}

// ConvFloat32 mocks base method.
func (m *MockPtrConv) ConvFloat32(v float32) *float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvFloat32", v)
	ret0, _ := ret[0].(*float32)
	return ret0
}

// ConvFloat32 indicates an expected call of ConvFloat32.
func (mr *MockPtrConvMockRecorder) ConvFloat32(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvFloat32", reflect.TypeOf((*MockPtrConv)(nil).ConvFloat32), v)
}

// ConvFloat64 mocks base method.
func (m *MockPtrConv) ConvFloat64(v float64) *float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvFloat64", v)
	ret0, _ := ret[0].(*float64)
	return ret0
}

// ConvFloat64 indicates an expected call of ConvFloat64.
func (mr *MockPtrConvMockRecorder) ConvFloat64(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvFloat64", reflect.TypeOf((*MockPtrConv)(nil).ConvFloat64), v)
}

// ConvInt mocks base method.
func (m *MockPtrConv) ConvInt(v int) *int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvInt", v)
	ret0, _ := ret[0].(*int)
	return ret0
}

// ConvInt indicates an expected call of ConvInt.
func (mr *MockPtrConvMockRecorder) ConvInt(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvInt", reflect.TypeOf((*MockPtrConv)(nil).ConvInt), v)
}

// ConvInt16 mocks base method.
func (m *MockPtrConv) ConvInt16(v int16) *int16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvInt16", v)
	ret0, _ := ret[0].(*int16)
	return ret0
}

// ConvInt16 indicates an expected call of ConvInt16.
func (mr *MockPtrConvMockRecorder) ConvInt16(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvInt16", reflect.TypeOf((*MockPtrConv)(nil).ConvInt16), v)
}

// ConvInt32 mocks base method.
func (m *MockPtrConv) ConvInt32(v int32) *int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvInt32", v)
	ret0, _ := ret[0].(*int32)
	return ret0
}

// ConvInt32 indicates an expected call of ConvInt32.
func (mr *MockPtrConvMockRecorder) ConvInt32(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvInt32", reflect.TypeOf((*MockPtrConv)(nil).ConvInt32), v)
}

// ConvInt64 mocks base method.
func (m *MockPtrConv) ConvInt64(v int64) *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvInt64", v)
	ret0, _ := ret[0].(*int64)
	return ret0
}

// ConvInt64 indicates an expected call of ConvInt64.
func (mr *MockPtrConvMockRecorder) ConvInt64(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvInt64", reflect.TypeOf((*MockPtrConv)(nil).ConvInt64), v)
}

// ConvInt8 mocks base method.
func (m *MockPtrConv) ConvInt8(v int8) *int8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvInt8", v)
	ret0, _ := ret[0].(*int8)
	return ret0
}

// ConvInt8 indicates an expected call of ConvInt8.
func (mr *MockPtrConvMockRecorder) ConvInt8(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvInt8", reflect.TypeOf((*MockPtrConv)(nil).ConvInt8), v)
}

// ConvPtrBool mocks base method.
func (m *MockPtrConv) ConvPtrBool(v *bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrBool", v)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ConvPtrBool indicates an expected call of ConvPtrBool.
func (mr *MockPtrConvMockRecorder) ConvPtrBool(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrBool", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrBool), v)
}

// ConvPtrByte mocks base method.
func (m *MockPtrConv) ConvPtrByte(v *byte) byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrByte", v)
	ret0, _ := ret[0].(byte)
	return ret0
}

// ConvPtrByte indicates an expected call of ConvPtrByte.
func (mr *MockPtrConvMockRecorder) ConvPtrByte(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrByte", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrByte), v)
}

// ConvPtrComplex128 mocks base method.
func (m *MockPtrConv) ConvPtrComplex128(v *complex128) complex128 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrComplex128", v)
	ret0, _ := ret[0].(complex128)
	return ret0
}

// ConvPtrComplex128 indicates an expected call of ConvPtrComplex128.
func (mr *MockPtrConvMockRecorder) ConvPtrComplex128(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrComplex128", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrComplex128), v)
}

// ConvPtrComplex64 mocks base method.
func (m *MockPtrConv) ConvPtrComplex64(v *complex64) complex64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrComplex64", v)
	ret0, _ := ret[0].(complex64)
	return ret0
}

// ConvPtrComplex64 indicates an expected call of ConvPtrComplex64.
func (mr *MockPtrConvMockRecorder) ConvPtrComplex64(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrComplex64", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrComplex64), v)
}

// ConvPtrFloat32 mocks base method.
func (m *MockPtrConv) ConvPtrFloat32(v *float32) float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrFloat32", v)
	ret0, _ := ret[0].(float32)
	return ret0
}

// ConvPtrFloat32 indicates an expected call of ConvPtrFloat32.
func (mr *MockPtrConvMockRecorder) ConvPtrFloat32(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrFloat32", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrFloat32), v)
}

// ConvPtrFloat64 mocks base method.
func (m *MockPtrConv) ConvPtrFloat64(v *float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrFloat64", v)
	ret0, _ := ret[0].(float64)
	return ret0
}

// ConvPtrFloat64 indicates an expected call of ConvPtrFloat64.
func (mr *MockPtrConvMockRecorder) ConvPtrFloat64(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrFloat64", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrFloat64), v)
}

// ConvPtrInt mocks base method.
func (m *MockPtrConv) ConvPtrInt(v *int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrInt", v)
	ret0, _ := ret[0].(int)
	return ret0
}

// ConvPtrInt indicates an expected call of ConvPtrInt.
func (mr *MockPtrConvMockRecorder) ConvPtrInt(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrInt", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrInt), v)
}

// ConvPtrInt16 mocks base method.
func (m *MockPtrConv) ConvPtrInt16(v *int16) int16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrInt16", v)
	ret0, _ := ret[0].(int16)
	return ret0
}

// ConvPtrInt16 indicates an expected call of ConvPtrInt16.
func (mr *MockPtrConvMockRecorder) ConvPtrInt16(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrInt16", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrInt16), v)
}

// ConvPtrInt32 mocks base method.
func (m *MockPtrConv) ConvPtrInt32(v *int32) int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrInt32", v)
	ret0, _ := ret[0].(int32)
	return ret0
}

// ConvPtrInt32 indicates an expected call of ConvPtrInt32.
func (mr *MockPtrConvMockRecorder) ConvPtrInt32(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrInt32", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrInt32), v)
}

// ConvPtrInt64 mocks base method.
func (m *MockPtrConv) ConvPtrInt64(v *int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrInt64", v)
	ret0, _ := ret[0].(int64)
	return ret0
}

// ConvPtrInt64 indicates an expected call of ConvPtrInt64.
func (mr *MockPtrConvMockRecorder) ConvPtrInt64(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrInt64", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrInt64), v)
}

// ConvPtrInt8 mocks base method.
func (m *MockPtrConv) ConvPtrInt8(v *int8) int8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrInt8", v)
	ret0, _ := ret[0].(int8)
	return ret0
}

// ConvPtrInt8 indicates an expected call of ConvPtrInt8.
func (mr *MockPtrConvMockRecorder) ConvPtrInt8(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrInt8", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrInt8), v)
}

// ConvPtrRune mocks base method.
func (m *MockPtrConv) ConvPtrRune(v *rune) rune {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrRune", v)
	ret0, _ := ret[0].(rune)
	return ret0
}

// ConvPtrRune indicates an expected call of ConvPtrRune.
func (mr *MockPtrConvMockRecorder) ConvPtrRune(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrRune", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrRune), v)
}

// ConvPtrString mocks base method.
func (m *MockPtrConv) ConvPtrString(v *string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrString", v)
	ret0, _ := ret[0].(string)
	return ret0
}

// ConvPtrString indicates an expected call of ConvPtrString.
func (mr *MockPtrConvMockRecorder) ConvPtrString(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrString", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrString), v)
}

// ConvPtrUint mocks base method.
func (m *MockPtrConv) ConvPtrUint(v *uint) uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrUint", v)
	ret0, _ := ret[0].(uint)
	return ret0
}

// ConvPtrUint indicates an expected call of ConvPtrUint.
func (mr *MockPtrConvMockRecorder) ConvPtrUint(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrUint", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrUint), v)
}

// ConvPtrUint16 mocks base method.
func (m *MockPtrConv) ConvPtrUint16(v *uint16) uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrUint16", v)
	ret0, _ := ret[0].(uint16)
	return ret0
}

// ConvPtrUint16 indicates an expected call of ConvPtrUint16.
func (mr *MockPtrConvMockRecorder) ConvPtrUint16(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrUint16", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrUint16), v)
}

// ConvPtrUint32 mocks base method.
func (m *MockPtrConv) ConvPtrUint32(v *uint32) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrUint32", v)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// ConvPtrUint32 indicates an expected call of ConvPtrUint32.
func (mr *MockPtrConvMockRecorder) ConvPtrUint32(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrUint32", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrUint32), v)
}

// ConvPtrUint64 mocks base method.
func (m *MockPtrConv) ConvPtrUint64(v *uint64) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrUint64", v)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ConvPtrUint64 indicates an expected call of ConvPtrUint64.
func (mr *MockPtrConvMockRecorder) ConvPtrUint64(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrUint64", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrUint64), v)
}

// ConvPtrUint8 mocks base method.
func (m *MockPtrConv) ConvPtrUint8(v *uint8) uint8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvPtrUint8", v)
	ret0, _ := ret[0].(uint8)
	return ret0
}

// ConvPtrUint8 indicates an expected call of ConvPtrUint8.
func (mr *MockPtrConvMockRecorder) ConvPtrUint8(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvPtrUint8", reflect.TypeOf((*MockPtrConv)(nil).ConvPtrUint8), v)
}

// ConvRune mocks base method.
func (m *MockPtrConv) ConvRune(v rune) *rune {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvRune", v)
	ret0, _ := ret[0].(*rune)
	return ret0
}

// ConvRune indicates an expected call of ConvRune.
func (mr *MockPtrConvMockRecorder) ConvRune(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvRune", reflect.TypeOf((*MockPtrConv)(nil).ConvRune), v)
}

// ConvString mocks base method.
func (m *MockPtrConv) ConvString(v string) *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvString", v)
	ret0, _ := ret[0].(*string)
	return ret0
}

// ConvString indicates an expected call of ConvString.
func (mr *MockPtrConvMockRecorder) ConvString(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvString", reflect.TypeOf((*MockPtrConv)(nil).ConvString), v)
}

// ConvUint mocks base method.
func (m *MockPtrConv) ConvUint(v uint) *uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvUint", v)
	ret0, _ := ret[0].(*uint)
	return ret0
}

// ConvUint indicates an expected call of ConvUint.
func (mr *MockPtrConvMockRecorder) ConvUint(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvUint", reflect.TypeOf((*MockPtrConv)(nil).ConvUint), v)
}

// ConvUint16 mocks base method.
func (m *MockPtrConv) ConvUint16(v uint16) *uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvUint16", v)
	ret0, _ := ret[0].(*uint16)
	return ret0
}

// ConvUint16 indicates an expected call of ConvUint16.
func (mr *MockPtrConvMockRecorder) ConvUint16(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvUint16", reflect.TypeOf((*MockPtrConv)(nil).ConvUint16), v)
}

// ConvUint32 mocks base method.
func (m *MockPtrConv) ConvUint32(v uint32) *uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvUint32", v)
	ret0, _ := ret[0].(*uint32)
	return ret0
}

// ConvUint32 indicates an expected call of ConvUint32.
func (mr *MockPtrConvMockRecorder) ConvUint32(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvUint32", reflect.TypeOf((*MockPtrConv)(nil).ConvUint32), v)
}

// ConvUint64 mocks base method.
func (m *MockPtrConv) ConvUint64(v uint64) *uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvUint64", v)
	ret0, _ := ret[0].(*uint64)
	return ret0
}

// ConvUint64 indicates an expected call of ConvUint64.
func (mr *MockPtrConvMockRecorder) ConvUint64(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvUint64", reflect.TypeOf((*MockPtrConv)(nil).ConvUint64), v)
}

// ConvUint8 mocks base method.
func (m *MockPtrConv) ConvUint8(v uint8) *uint8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvUint8", v)
	ret0, _ := ret[0].(*uint8)
	return ret0
}

// ConvUint8 indicates an expected call of ConvUint8.
func (mr *MockPtrConvMockRecorder) ConvUint8(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvUint8", reflect.TypeOf((*MockPtrConv)(nil).ConvUint8), v)
}
