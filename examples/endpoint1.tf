resource "huggingface_endpoint" "enhanced_search_embedding" {
  name = "enh-srch-embed-staging-us-cpu"

  compute = {
    accelerator   = "cpu"
    instance_size = "x1"
    instance_type = "intel-icl"
    scaling = {
      min_replica = 0
      max_replica = 1
      measure = {
        hardware_usage = 1.23424
      }
    }
  }

  model = {
    framework = "pytorch"
    image = {
      custom = {
        url = "ghcr.io/huggingface/text-embeddings-inference:cpu-0.6.0"
      }
    }
    env = {
      MAX_BATCH_TOKENS        = "1000000"
      MAX_CONCURRENT_REQUESTS = "512"
      MODEL_ID                = "/repository"
    }
    repository = "avsolatorio/GIST-Embedding-v0"
    task       = "sentence-embeddings"
    revision   = "025ccf7d0a8f03dbd7cead428899acfdf6636432"
  }

  cloud = {
    region = "us-east-1"
    vendor = "aws"
  }

  type = "protected"
}

output "enhanced_search_embedding" {
  value = huggingface_endpoint.enhanced_search_embedding
}
