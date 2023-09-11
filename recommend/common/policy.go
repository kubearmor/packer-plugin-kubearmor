// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package common contains object types used by multiple packages
package common

import (
	pol "github.com/kubearmor/KubeArmor/pkg/KubeArmorController/api/security.kubearmor.com/v1"
)

// Handler interface
var Handler interface{}

// MatchSpec spec to match for defining policy
type MatchSpec struct {
	Name         string                      `json:"name" yaml:"name"`
	Precondition []string                    `json:"precondition" yaml:"precondition"`
	Description  Description                 `json:"description" yaml:"description"`
	Yaml         string                      `json:"yaml" yaml:"yaml"`
	Spec         pol.KubeArmorHostPolicySpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// Ref for the policy rules
type Ref struct {
	Name string   `json:"name" yaml:"name"`
	URL  []string `json:"url" yaml:"url"`
}

// Description detailed description for the policy rule
type Description struct {
	Refs     []Ref  `json:"refs" yaml:"refs"`
	Tldr     string `json:"tldr" yaml:"tldr"`
	Detailed string `json:"detailed" yaml:"detailed"`
}

// Options
type Options struct {
	OutDir string
}
