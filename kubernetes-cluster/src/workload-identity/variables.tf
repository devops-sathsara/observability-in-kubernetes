variable "env_project_id" {
  type        = string
  description = "The project id to create WIF pool and SA"
}

variable "service_account_name" {
  type        = string
  description = "The service account name"
}

variable "github_username" {
  type        = string
  description = "GitHub username to grant access"
}

variable "github_repo_name" {
  type        = string
  description = "GitHub repo name to grant access"
}