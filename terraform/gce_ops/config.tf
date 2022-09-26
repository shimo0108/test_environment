variable "enable_test" {
  default = false
}
variable "test_subnet" {
  default = ""
}
variable "test_settings" {
  type = list(any)
  default = []
}

module "test" {
  source = "./modules/gce"

  enable      = var.enable_test
  environment = terraform.workspace
  subnet      = var.test_subnet
  settings    = var.test_settings
}