resource "huggingface_endpoint" "endpoint1" {
  name = "test-endpoint-issa1"

  compute = {
    accelerator   = "cpu"
    instance_size = "x8"
    instance_type = "intel-icl"
    scaling = {
      min_replica           = 0
      max_replica           = 1
      scale_to_zero_timout  = 30
    }
  }

  model = {
    framework = "pytorch"
    image = {
      huggingface = {
        env = {}
      }
    }
    repository = "sentence-transformers/all-MiniLM-L6-v2"
    revision   = "e4ce9877abf3edfe10b0d82785e83bdcb973e22e"
    task       = "sentence-embeddings"
  }

  cloud = {
    region = "us-east-1"
    vendor = "aws"
  }

  type = "protected"
}

output "endpoint1" {
  value = huggingface_endpoint.endpoint1
}
