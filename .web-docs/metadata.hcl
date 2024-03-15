# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# Details on using this Integration template can be found at https://github.com/hashicorp/integration-template
# Alternatively this metadata.hcl file can be placed under the docs/ subdirectory or any other config subdirectory that 
# makes senses for the plugin. 
integration {
  name = "KubeArmor"
  description = "The KubeArmor plugin which provides kubearmor hardening security policies for the build workloads"
  identifier = "packer/hashicorp/kubearmor"
  component {
    type = "provisioner"
    name = "KubeArmor"
    slug = "KubeArmor"
  }
}
