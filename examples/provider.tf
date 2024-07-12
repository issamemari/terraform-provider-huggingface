terraform {
  required_providers {
    huggingface = {
      source  = "issamemari/huggingface"
      version = "1.1.7"
    }
  }
}

provider "huggingface" {
  host      = "https://api.endpoints.huggingface.cloud/v2/endpoint"
  namespace = "issamemari"
  token     = ""
}
