terraform {
    backend "gcs" {
        bucket = "test-terraform-state"
        prefix = "test"
    }
}
