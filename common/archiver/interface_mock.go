// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by mockery v1.0.0. DO NOT EDIT.

package archiver

import context "context"
import mock "github.com/stretchr/testify/mock"

// HistoryArchiverMock is an autogenerated mock type for the HistoryArchiver type
type HistoryArchiverMock struct {
	mock.Mock
}

// Archive provides a mock function with given fields: ctx, uri, request, opts
func (_m *HistoryArchiverMock) Archive(ctx context.Context, uri URI, request *ArchiveHistoryRequest, opts ...ArchiveOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, uri, request)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, URI, *ArchiveHistoryRequest, ...ArchiveOption) error); ok {
		r0 = rf(ctx, uri, request, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, uri, request
func (_m *HistoryArchiverMock) Get(ctx context.Context, uri URI, request *GetHistoryRequest) (*GetHistoryResponse, error) {
	ret := _m.Called(ctx, uri, request)

	var r0 *GetHistoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, URI, *GetHistoryRequest) *GetHistoryResponse); ok {
		r0 = rf(ctx, uri, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetHistoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, URI, *GetHistoryRequest) error); ok {
		r1 = rf(ctx, uri, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateURI provides a mock function with given fields: uri
func (_m *HistoryArchiverMock) ValidateURI(uri URI) error {
	ret := _m.Called(uri)

	var r0 error
	if rf, ok := ret.Get(0).(func(URI) error); ok {
		r0 = rf(uri)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VisibilityArchiverMock is an autogenerated mock type for the VisibilityArchiver type
type VisibilityArchiverMock struct {
	mock.Mock
}

// ValidateURI provides a mock function with given fields: uri
func (_m *VisibilityArchiverMock) ValidateURI(uri URI) error {
	ret := _m.Called(uri)

	var r0 error
	if rf, ok := ret.Get(0).(func(URI) error); ok {
		r0 = rf(uri)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}