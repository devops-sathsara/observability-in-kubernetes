resource "google_service_account" "sa" {
  project    = var.env_project_id
  account_id = var.service_account_name
}

resource "google_project_iam_member" "project" {
  project = var.env_project_id
  role    = "roles/storage.admin"
  member  = "serviceAccount:${google_service_account.sa.email}"
}

module "oidc" {
  source      = "terraform-google-modules/github-actions-runners/google//modules/gh-oidc"
  project_id  = var.env_project_id
  pool_id     = "gh-pool"
  provider_id = "gh-provider"
  sa_mapping = {
    (google_service_account.sa.account_id) = {
      sa_name   = google_service_account.sa.name
      attribute = "attribute.repository/user/repo"
    }
  }
}