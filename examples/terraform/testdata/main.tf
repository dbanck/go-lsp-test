resource "random_pet" "application" {
  count = var.fooo
  keepers = {
    unique = "unique"
  }
}

variable "fooo" {
  type = number
  default = 3
}

output "pet_count" {
  value = var.fooo
}