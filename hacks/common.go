// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package hacks close the file
package hacks

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// CloseCheckErr close file
func CloseCheckErr(f *os.File, fname string) {
	err := f.Close()
	if err != nil {
		log.WithFields(log.Fields{
			"file": fname,
		}).Error("close file failed")
	}
}
