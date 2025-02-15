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
	"github.com/AYaro/prototool/internal/strs"
	"github.com/AYaro/prototool/internal/text"
	"github.com/emicklei/proto"
)

var messageFieldNamesLowerSnakeCaseLinter = NewLinter(
	"MESSAGE_FIELD_NAMES_LOWER_SNAKE_CASE",
	"Verifies that all message field names are lower_snake_case.",
	checkMessageFieldNamesLowerSnakeCase,
)

func checkMessageFieldNamesLowerSnakeCase(add func(*text.Failure), dirPath string, descriptors []*FileDescriptor) error {
	return runVisitor(messageFieldNamesLowerSnakeCaseVisitor{baseAddVisitor: newBaseAddVisitor(add)}, descriptors)
}

type messageFieldNamesLowerSnakeCaseVisitor struct {
	baseAddVisitor
}

func (v messageFieldNamesLowerSnakeCaseVisitor) VisitMessage(message *proto.Message) {
	for _, element := range message.Elements {
		element.Accept(v)
	}
}

func (v messageFieldNamesLowerSnakeCaseVisitor) VisitOneof(oneof *proto.Oneof) {
	for _, element := range oneof.Elements {
		element.Accept(v)
	}
}

func (v messageFieldNamesLowerSnakeCaseVisitor) VisitNormalField(field *proto.NormalField) {
	if !strs.IsLowerSnakeCase(field.Name) {
		v.AddFailuref(field.Position, "Field name %q must be lower_snake_case.", field.Name)
	}
}

func (v messageFieldNamesLowerSnakeCaseVisitor) VisitOneofField(field *proto.OneOfField) {
	if !strs.IsLowerSnakeCase(field.Name) {
		v.AddFailuref(field.Position, "Field name %q must be lower_snake_case.", field.Name)
	}
}

func (v messageFieldNamesLowerSnakeCaseVisitor) VisitMapField(field *proto.MapField) {
	if !strs.IsLowerSnakeCase(field.Name) {
		v.AddFailuref(field.Position, "Field name %q must be lower_snake_case.", field.Name)
	}
}
