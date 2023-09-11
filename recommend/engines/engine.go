// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package engines provides interfaces and implementations for policy generation
package engines

import (
	"github.com/packer-plugin-kubearmor/recommend/common"
	"github.com/packer-plugin-kubearmor/recommend/image"
)

// Engine interface used by policy generators to generate policies
type Engine interface {
	Init() error
	Scan(img *image.Info, options common.Options) (map[string][]byte, map[string]interface{}, error)
}
