# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# Details on using this Integration template can be found at https://github.com/hashicorp/integration-template
# Alternatively this metadata.hcl file can be placed under the docs/ subdirectory or any other config subdirectory that 
# makes senses for the plugin. 
integration {
  name = "packer-plugin-kubearmor"
  description = "a plugin for kubearmor which provides kubearmor hardening security policies for the build workloads"
  identifier = "packer/kubeArmor/kubearmor"
  docs {
    process_docs = true
    # We recommend using the default readme_location of just `./README.md` here
    # This projects README needs to document the interface of an integration.
    #
    # If you need a separate README from what you will display on GitHub vs
    # what is shown on HashiCorp Developer, this is totally valid, though!
    readme_location = "./README.md"
  }
  license {
    type = "MPL-2.0"
    url = "https://github.com/kubearmor/packer-plugin-kubearmor/blob/main/LICENSE"
  }
  component {
    type = "provisioner"
    name = "KubeArmor"
    slug = "KubeArmor"
  }
}
