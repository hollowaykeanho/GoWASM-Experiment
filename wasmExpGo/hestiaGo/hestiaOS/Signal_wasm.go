// Copyright 2022 "Holloway" Chew, Kean Ho <hollowaykeanho@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

//go:build wasm
// +build wasm

package hestiaOS

import (
	"hestiaGo/hestiaError"
)

func _signalInit(sig *Signal, bufferSize int) hestiaError.Error {
	if sig == nil {
		return hestiaError.ENOENT
	}

	if bufferSize == 0 {
		bufferSize = 3
	}

	sig.channel = make(chan uint16, bufferSize)

	return hestiaError.OK
}

func _signalSend(sig *Signal, data uint16) hestiaError.Error {
	if sig == nil {
		return hestiaError.ENOENT
	}

	sig.channel <- data

	return hestiaError.OK
}

func _signalWait(sig *Signal) (value uint16) {
	if sig == nil {
		return value
	}

	value = <-sig.channel
	return value
}
