terraform {
  required_providers {
    huggingface = {
      source  = "jesus/huggingface"
      #version = "1.1.4"
    }
  }
}

provider "huggingface" {
  host      = "https://api.endpoints.huggingface.cloud/v2/endpoint"
  namespace = "issamemari"
  token     = ""
}
