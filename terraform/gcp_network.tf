resource "google_compute_subnetwork" "sub-network" {
  name          = "sn-project-management-v2-sub-network"
  ip_cidr_range = "10.2.0.0/16"
  region        = "${var.gcp_region}"
  network       = google_compute_network.network.self_link
}

resource "google_compute_network" "network" {
  name                    = "sn-project-management-v2-network"
  auto_create_subnetworks = false
}
