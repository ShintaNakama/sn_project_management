resource "google_compute_firewall" "sn-project-management-v1-firewall" {
  name    = "sn-project-management-v1"
  network       = google_compute_network.sn-project-management-v1-network.self_link
  
  allow {
    protocol = "icmp"
  }
  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443"]
  }
  target_tags = ["web"]
}
