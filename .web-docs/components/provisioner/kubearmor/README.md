Type: `kubearmor`

  The KubeArmor provisioner will configure the build and provide the [KubeArmor Hardening Host Security Policies](https://docs.kubearmor.io/kubearmor/use-cases/hardening_guide). These hardening policies will be in the context of your workload, so you can see how they will be applied and what impact they will have on your system. This allows you to make informed decisions about which policies to apply, and helps you understand the trade-offs between security and functionality.


<!-- Provisioner Configuration Fields -->

### Required

- `policyPath` (string) - directory where security policies will be stored the inside the build (Please use the path where the user has access to read, write and create).


### Example Usage

With docker uase case :

```hcl
packer {
  required_plugins {
    docker = {
      version = ">= 0.0.7"
      source = "github.com/hashicorp/docker"
    }
    kubearmor = {
      version = ">= 0.0.1"
      source = "github.com/kubearmor/kubearmor"
    }
  }
}

source "docker" "ubuntu" {
  image  = "ubuntu:xenial"
  commit = true
}

build {
  name = "learn-packer"
  sources = [
    "source.docker.ubuntu"
  ]

  provisioner "kubearmor" {
    policyPath = "/policies"
  }
}

```

With virtualbox use case :

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
    kubearmor = {
      version = ">= 0.0.1"
      source = "github.com/kubearmor/kubearmor"
    }
  }
}

source "virtualbox-ovf" "basic-example" {
  source_path = "ubuntu20.ova"
  ssh_username = ""
  ssh_password = ""
  shutdown_command = "echo 'packer' | sudo -S shutdown -P now"
}

build {
  name = "learn-packer"
  sources = ["sources.virtualbox-ovf.basic-example"]
  
 
  provisioner "ansible" {
    playbook_file = "./ansible/conf.yml"
  }

  provisioner "kubearmor" {
    policyPath = "/home/prateek/policies"
  }
}
```
