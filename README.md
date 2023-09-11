# Packer Plugin KubeArmor

A plugin for Packer which provides [KubeArmor Hardening Host Security policies](https://docs.kubearmor.io/kubearmor/use-cases/hardening_guide) for the build workloads.

> [KubeArmor](https://docs.kubearmor.io/kubearmor/) is a security solution for the Kubernetes and cloud native platforms that helps protect your workloads from attacks and threats. It does this by providing a set of hardening policies that are based on industry-leading compliance and attack frameworks such as CIS, MITRE, NIST-800-53, and STIGs. These policies are designed to help you secure your workloads in a way that is compliant with these frameworks and recommended best practices.

## Installation

Since we do not have a release currently so we have to build the provisioner and use it.
1. Clone the repo.
2. Run `make build`.
3. Configure the Packer_PLUGIN_PATH - `export PACKER_PLUGIN_PATH=~/go/src/github.com/packer-plugin-kubearmor`.
4. Run `packer init .`
5. Run `packer build <file>.hcl`.

## Requirements

In order to use the provided KubeArmor Host Security Policies by the provisioner the build also needs to have KubeArmor and other tools and dependencies required. For installing KubeArmor and ensuring all the dependencies are met we are providing an ansible playbook which could be used with ansible provisioner. 

## Example

```
packer {
  required_plugins {
    virtualbox-ovf = {
      source  = "github.com/hashicorp/virtualbox"
      version = "~> 1"
    }
    ansible = {
      source  = "github.com/hashicorp/ansible"
      version = "~> 1"
    }
  }
}

source "virtualbox-ovf" "basic-example" {
  source_path = "ubuntu20.ova"
  ssh_username = "prateek"
  ssh_password = "1234"
  shutdown_command = "echo 'packer' | sudo -S shutdown -P now"
}

build {
  name = "learn-packer"
  sources = ["sources.virtualbox-ovf.basic-example"]
  
 
  provisioner "ansible" {
    playbook_file = "./ansible/conf.yml"
  }

  provisioner "image" {
    policyPath = "/home/prateek/policies"
  }
}
```