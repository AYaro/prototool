// Copyright (c) 2020 Uber Technologies, Inc.
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

package lint

import (
	"strings"

	"github.com/AYaro/prototool/internal/text"
	"github.com/emicklei/proto"
)

var (
	gogoNotImportedLinter = NewLinter(
		"GOGO_NOT_IMPORTED",
		`Verifies that the "gogo.proto" file from gogo/protobuf is not imported.`,
		checkGogoNotImported,
	)
)

func checkGogoNotImported(add func(*text.Failure), dirPath string, descriptors []*FileDescriptor) error {
	return runVisitor(&gogoNotImportedVisitor{baseAddVisitor: newBaseAddVisitor(add)}, descriptors)
}

type gogoNotImportedVisitor struct {
	baseAddVisitor
}

func (v gogoNotImportedVisitor) VisitImport(element *proto.Import) {
	if strings.HasSuffix(element.Filename, "gogo.proto") {
		v.AddFailuref(element.Position, `Importing "gogo.proto" is not allowed.`)
	}
}
