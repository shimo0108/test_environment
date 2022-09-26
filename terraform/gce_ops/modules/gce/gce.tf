variable "enable" {
  type = bool
  default = true
}
variable "environment" {
  type = string
}
variable "subnet" {
  type = string
}
variable "settings" {
  type = list(any)
  default = []
}
data "google_compute_image" "coreos" {
  project = "cos-cloud"
  family = "cos-stable"
}
resource "google_compute_instance" "instance" {
  count = var.enable ? length(var.settings) : 0

  name = "${lookup(var.settings[count.index], "name")}-${var.environment}"
  description = lookup(var.settings[count.index], "description", null)
  machine_type = lookup(var.settings[count.index], "machine_type", null)
  zone = lookup(var.settings[count.index], "zone", null)

  boot_disk {
    initialize_params {
      image = data.google_compute_image.coreos.self_link
      size = lookup(var.settings[count.index], "disk_size")
    }
  }

  network_interface {
    subnetwork = var.subnet
    network_ip = lookup(var.settings[count.index], "ip_addr")

    access_config {}
  }

  service_account {
    email = lookup(var.settings[count.index], "service_account_email")
    scopes = [
      "https://www.googleapis.com/auth/devstrage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring.write",
      "https://www.googleapis.com/auth/pubsub",
      "https://www.googleapis.com/auth/service.management.readonly",
      "https://www.googleapis.com/auth/servicecontrol",
      "https://www.googleapis.com/auth/trace.append",
      "https://www.googleapis.com/auth/cloud-platform",
      "https://www.googleapis.com/auth/spanner.data",
    ]
  }

  metadata = {
    gce_container_declaration = yamlencode(lookup(var.settings[count.index], "container_declaration"))
    google-logging-enabled = "true"
  }

  tags = concat(
    ["iap-ssh"],
    lookup(var.settings[count.index], "tags"),
  )
  allow_stopping_for_update = true

}