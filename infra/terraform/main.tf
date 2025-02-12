# Mock Terraform Configuration

terraform {
  required_version = ">= 1.0.0"

  backend "local" { # Stores Terraform state locally (change for cloud storage)
    path = "terraform.tfstate"
  }
}

provider "null" {
  # The "null" provider is a placeholder that does nothing.
}

resource "null_resource" "mock" {
  triggers = {
    example = "This is a mock Terraform file for saving purposes."
  }
}

output "mock_output" {
  value = "Terraform mock setup complete."
}
