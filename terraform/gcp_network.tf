resource "google_compute_subnetwork" "sn-project-management-v1-sub-network" {
  name          = "sn-project-management-v1-sub-network"
  ip_cidr_range = "10.2.0.0/16"
  region        = "${var.gcp_region}"
  network       = google_compute_network.sn-project-management-v1-network.self_link
}

resource "google_compute_network" "sn-project-management-v1-network" {
  name                    = "sn-project-management-v1-network"
  auto_create_subnetworks = false
}
