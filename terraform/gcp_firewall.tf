resource "google_compute_firewall" "firewall" {
  name    = "sn-project-management-v2"
  network       = google_compute_network.network.self_link
  
  allow {
    protocol = "icmp"
  }
  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443"]
  }
  target_tags = ["web"]
}
