// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package image scan and provide image info
package image

import (
	_ "embed" // need for embedding
	"fmt"
	"os"
	"path/filepath"

	"github.com/clarketm/json"
	pol "github.com/kubearmor/KubeArmor/pkg/KubeArmorController/api/security.kubearmor.com/v1"
	"github.com/packer-plugin-kubearmor/recommend/common"
	log "github.com/sirupsen/logrus"
)

// Info contains image information
type Info struct {
	FileList []string
	TempDir  string
}

func addPolicyRule(policy *pol.KubeArmorHostPolicy, r pol.KubeArmorHostPolicySpec) {

	if len(r.File.MatchDirectories) != 0 || len(r.File.MatchPaths) != 0 {
		policy.Spec.File = r.File
	}
	if len(r.Process.MatchDirectories) != 0 || len(r.Process.MatchPaths) != 0 {
		policy.Spec.Process = r.Process
	}
	if len(r.Network.MatchProtocols) != 0 {
		policy.Spec.Network = r.Network
	}
}

func (img *Info) createPolicy(ms common.MatchSpec) (pol.KubeArmorHostPolicy, error) {
	policy := pol.KubeArmorHostPolicy{
		Spec: pol.KubeArmorHostPolicySpec{
			Severity: 1, // by default
			NodeSelector: pol.NodeSelectorType{
				MatchLabels: map[string]string{}},
		},
	}
	policy.APIVersion = "security.kubearmor.com/v1"
	policy.Kind = "KubeArmorHostPolicy"

	policy.ObjectMeta.Name = ms.Name

	policy.Spec.Action = ms.Spec.Action
	policy.Spec.Severity = ms.Spec.Severity
	if ms.Spec.Message != "" {
		policy.Spec.Message = ms.Spec.Message
	}
	if len(ms.Spec.Tags) > 0 {
		policy.Spec.Tags = ms.Spec.Tags
	}

	addPolicyRule(&policy, ms.Spec)
	return policy, nil
}

// GetPolicy - creates policy and return back
func (img *Info) GetPolicy(ms common.MatchSpec, options common.Options) ([]byte, string) {
	policy, err := img.createPolicy(ms)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"image": img, "spec": ms,
		}).Error("create policy failed, skipping")
	}

	arr, _ := json.Marshal(policy)
	outFile := filepath.Join(options.OutDir, ms.Name)
	err = os.MkdirAll(filepath.Dir(outFile), 0750)
	if err != nil {
		log.WithError(err).Error("failed to create directory")
	}
	_, err = os.Create(filepath.Clean(outFile))
	if err != nil {
		log.WithError(err).Error(fmt.Sprintf("create file %s failed", outFile))
	}

	return arr, outFile
}
