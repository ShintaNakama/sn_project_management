resource "google_compute_instance" "sn-project-management-v1-web-a" {
  name         = "sn-project-management-v1-wab-a"
  machine_type = "f1-micro"
  zone         = "${var.gcp_region_zone}"
  description  = "sn-project-management-v1 web-a"
  tags         = ["web"]

  boot_disk {
    initialize_params {
      size  = 30
      image = "centos-cloud/centos-7"
    }
  }

  // Local SSD disk
  //scratch_disk {
  //  interface = "SCSI"
  //}

  network_interface {
    access_config {
      // Ephemeral IP
    }
  subnetwork = google_compute_subnetwork.sn-project-management-v1-sub-network.self_link
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro", "bigquery", "monitoring"]
  }
  
  scheduling {
    on_host_maintenance = "MIGRATE"
    automatic_restart   = true
  }
}
