/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Status  int
	Message string
}

func (e *Error) Code() int {
	return e.Status
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(status int, message string) *Error {
	return &Error{
		status,
		message,
	}
}

func NewNotFound(message string) *Error {
	return NewError(http.StatusNotFound, message)
}

var InternalError = NewError(http.StatusInternalServerError, "Server Internal Error")

// Deprecated: use ctx.Err() instead
var TaskCanceled = fmt.Errorf("task got canceled")
