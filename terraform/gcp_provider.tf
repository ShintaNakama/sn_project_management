// Configure the Google Cloud provider
provider "google" {
 credentials = file("/Users/nakamashinta/.gcp/sn-project-management-v2-30a6bbc88daa.json")
 project     = "sn-project-management-v2"
 region      = "asia-northeast1"
}
provider "google-beta" {
 credentials = file("/Users/nakamashinta/.gcp/sn-project-management-v2-30a6bbc88daa.json")
 project     = "sn-project-management-v2"
 region      = "asia-northeast1"
 zone        = "asia-northeast1-a"
}
