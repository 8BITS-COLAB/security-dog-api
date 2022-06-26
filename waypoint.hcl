project = "security-dog-api"

app "security-dog-api" {
  labels = {
    "service" = "security-dog-api",
    "env"     = "dev"
  }

  build {
    use "pack" {}
  }

  deploy {
    use "docker" {}
  }
}