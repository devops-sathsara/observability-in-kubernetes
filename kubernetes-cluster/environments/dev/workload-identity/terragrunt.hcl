# Find root terragrunt.hcl and inherit its configurations
include {
  path = find_in_parent_folders()
}

terraform {
  source = "git::https://github.com/devops-sathsara/observability-in-kubernetes.git//kubernetes-cluster/src/workload-identity?ref=main"
}


locals {
  common_vars = read_terragrunt_config(find_in_parent_folders("inputs.hcl"))
}
 
inputs = merge(
  local.common_vars.inputs,
  {
    service_account_name   = "gh-workflow-sa"
  }
)
