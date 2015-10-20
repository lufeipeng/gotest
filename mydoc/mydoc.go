// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
This package is use to test godoc format

eg function

eg var

eg type

eg Example
*/
package mydoc

import (
	"fmt"
	"io"
)

//MyDocVar
var MyDocVar string;

//This is MyDoc init, used to init godoc package
func init() {
}

// A Compressor returns a compressing writer, writing to the
// provided writer. On Close, any pending data should be flushed.
type Compressor func(io.Writer) (io.WriteCloser, error)


// A sparseFileReader is a numBytesReader for reading sparse file data from a tar archive.
type SparseFileReader struct {
	pos int64          // keeps track of file position
	tot int64          // total size of the file
}

// Keywords for GNU sparse files in a PAX extended header
const (
	PaxGNUSparseNumBlocks = "GNU.sparse.numblocks"
	paxGNUSparseOffset    = "GNU.sparse.offset"
	paxGNUSparseNumBytes  = "GNU.sparse.numbytes"
	paxGNUSparseMap       = "GNU.sparse.map"
	paxGNUSparseName      = "GNU.sparse.name"
	paxGNUSparseMajor     = "GNU.sparse.major"
	paxGNUSparseMinor     = "GNU.sparse.minor"
	paxGNUSparseSize      = "GNU.sparse.size"
	paxGNUSparseRealSize  = "GNU.sparse.realsize"
)

/*
MyDocFunction Param

myDocFunction Return

Add Another return value
*/
func MyDocFunction() {
	fmt.Println("MyDocFunction is called");	
}