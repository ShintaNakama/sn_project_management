// Configure the Google Cloud provider
provider "google" {
 credentials = file("/Users/nakamashinta/.gcp/sn-project-management-v1-f746a8a5aa3a.json")
 project     = "sn-project-management-v1"
 region      = "us-west1"
}
provider "google-beta" {
 credentials = file("/Users/nakamashinta/.gcp/sn-project-management-v1-f746a8a5aa3a.json")
 project     = "sn-project-management-v1"
 region = "us-central1"
 zone   = "us-central1-a"
}
