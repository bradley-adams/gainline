terraform {
  backend "gcs" {
    bucket = "gainline-tfstate"
    prefix = "dev"
  }
}