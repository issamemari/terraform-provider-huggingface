resource "huggingface_endpoint" "product_ident_reranker_vllm" {
  name = "product-ident-reranker-vllm-prod"
  compute = {
    accelerator   = "gpu"
    instance_size = "x1"
    instance_type = "nvidia-l4"
    scaling = {
      min_replica = 1
      max_replica = 2
      measure = {
        pending_requests = 10
      }
    }
  }
  model = {
    framework = "pytorch"
    image = {
      vllm = {
        port                 = 8000
        url                  = "vllm/vllm-openai:gptoss"
        tensor_parallel_size = 1
        kv_cache_dtype       = "auto"
      }
    }
    env        = {}
    repository = "sentence-transformers/stsb-roberta-base"
    task       = "text-classification"
  }
  cloud = {
    region = "us-east-1"
    vendor = "aws"
  }
  type = "private"
}
output "product_ident_reranker_vllm" {
  description = "Product identification reranker vLLM"
  value       = huggingface_endpoint.product_ident_reranker_vllm
}