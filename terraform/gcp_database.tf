//resource "google_compute_network" "private_network" {
//  provider = google-beta
//
//  name = "private-network"
//}

//resource "google_compute_subnetwork" "sn-project-management-v1-db-network" {
//  name          = "sn-project-management-v1-db-network"
//  ip_cidr_range = "10.2.0.0/16"
//  region        = "${var.gcp_region}"
//  network       = google_compute_network.sn-project-management-v1-network.self_link
//}

resource "google_compute_global_address" "private_ip_address" {
  provider = google-beta

  name          = "private-ip-address"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.sn-project-management-v1-network.self_link
}

resource "google_service_networking_connection" "private_vpc_connection" {
  provider = google-beta
  network                 = google_compute_network.sn-project-management-v1-network.self_link
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_address.name]
}

resource "random_id" "db_name_suffix" {
  byte_length = 4
}

resource "google_sql_database_instance" "instance" {
  provider = google-beta

  name             = "sn-project-management-v1-${random_id.db_name_suffix.hex}"
  region           = "${var.gcp_region}"
  database_version = "MYSQL_5_7"

  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.sn-project-management-v1-network.self_link
    }
  }
}

resource "google_sql_database" "database" {
  name       = "sn_project_management"
  instance   = google_sql_database_instance.instance.name
  charset    = "utf8mb4"
  collation  = "utf8mb4_general_ci"
}

