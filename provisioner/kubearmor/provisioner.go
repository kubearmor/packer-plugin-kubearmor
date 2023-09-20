// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package kubearmor

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/packer-plugin-kubearmor/recommend"
	"github.com/packer-plugin-kubearmor/recommend/common"
	genericpolicies "github.com/packer-plugin-kubearmor/recommend/engines/generic_policies"
)

type Config struct {
	PolicyPath string `mapstructure:"policyPath"`
}

type Provisioner struct {
	config Config
}

func (p *Provisioner) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{}, raws...)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, comm packer.Communicator, generatedData map[string]interface{}) error {
	var stdout1 bytes.Buffer

	cmd1 := &packer.RemoteCmd{
		Command: "find / -type f",
		Stdout:  &stdout1,
	}
	err := comm.Start(context.Background(), cmd1)
	if err != nil {
		return err
	}

	cmd1.Wait()

	files := strings.Split(stdout1.String(), "\n")

	var stdout2 bytes.Buffer

	cmd2 := &packer.RemoteCmd{
		Command: "uname",
		Stdout:  &stdout2,
	}
	err = comm.Start(context.Background(), cmd2)
	if err != nil {
		return fmt.Errorf("non-linux platforms are not supported, yet. error : %s", err)
	}

	cmd2.Wait()

	os := stdout2.String()
	if os != "Linux\n" {
		return fmt.Errorf("non-linux platforms are not supported, yet")
	}

	options := common.Options{
		OutDir: "out",
	}

	err = recommend.Recommend(options, files, genericpolicies.GenericPolicy{})
	if err != nil {
		log.Printf("err : %s", err)
		return err
	}

	err = comm.UploadDir(p.config.PolicyPath, "./out", nil)
	if err != nil {
		log.Printf("err : %s", err)
		return err
	}
	var stdout3 bytes.Buffer

	cmd3 := &packer.RemoteCmd{
		Command: "find / -type f",
		Stdout:  &stdout3,
	}
	err = comm.Start(context.Background(), cmd3)
	if err != nil {
		return err
	}

	cmd3.Wait()

	files = strings.Split(stdout3.String(), "\n")
	for _, file := range files {
		if strings.HasPrefix(file, p.config.PolicyPath) {
			ui.Say(fmt.Sprintf("file uploaded: %s", file))
			continue
		}
	}
	return nil
}
