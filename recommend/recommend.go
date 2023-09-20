// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package recommend provides policies by policy generators
package recommend

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/packer-plugin-kubearmor/recommend/common"
	"github.com/packer-plugin-kubearmor/recommend/engines"
	"github.com/packer-plugin-kubearmor/recommend/image"

	"github.com/fatih/color"
	"sigs.k8s.io/yaml"

	log "github.com/sirupsen/logrus"
)

func createOutDir(dir string) error {
	if dir == "" {
		return nil
	}
	_, err := os.Stat(dir)
	if errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(dir, 0750)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func writePolicyFile(policMap map[string][]byte, msMap map[string]interface{}) {
	for outFile, policy := range policMap {
		fmt.Println("writePolicyFile")
		f, err := os.OpenFile(filepath.Clean(outFile), os.O_RDWR, 0)
		if err != nil {
			log.WithError(err).Error(fmt.Sprintf("create file %s failed", outFile))
		}

		yamlPolicy, _ := yaml.JSONToYAML(policy)
		if _, err = f.WriteString(string(yamlPolicy)); err != nil {
			log.WithError(err).Error("WriteString failed")
		}
		if err = f.Sync(); err != nil {
			log.WithError(err).Error("file sync failed")
		}
		if err = f.Close(); err != nil {
			log.WithError(err).Error("file close failed")
		}
		color.Green("created policy %s ...", outFile)
	}
}

// Recommend handler for karmor cli tool
func Recommend(o common.Options, files []string, policyGenerators ...engines.Engine) error {
	var policyMap map[string][]byte
	var msMap map[string]interface{}
	var err error

	if err = createOutDir(o.OutDir); err != nil {
		return err
	}

	for _, gen := range policyGenerators {
		if err := gen.Init(); err != nil {
			log.WithError(err).Error("policy generator init failed")
		}

		img := image.Info{}
		img.FileList = files
		if policyMap, msMap, err = gen.Scan(&img, o); err != nil {
			log.WithError(err).Error("policy generator scan failed")
		}
		writePolicyFile(policyMap, msMap)
	}

	return nil
}
