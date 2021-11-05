terraform {
  required_providers {
    paths = {
      source  = "toowoxx/paths"
    }
  }
}

provider "paths" {}

data "paths_components" "components" {
  path = "/a/b/c.txt"
}

data "paths_parents" "parents" {
  path = "/1/2/3.txt"
}

output "components" {
  value = data.paths_components.components.components
}

output "parents" {
  value = data.paths_parents.parents.parents
}
