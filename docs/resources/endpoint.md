---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "huggingface_endpoint Resource - huggingface"
subcategory: ""
description: |-
  
---

# huggingface_endpoint (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud` (Attributes) (see [below for nested schema](#nestedatt--cloud))
- `compute` (Attributes) (see [below for nested schema](#nestedatt--compute))
- `model` (Attributes) (see [below for nested schema](#nestedatt--model))
- `name` (String)
- `type` (String)

### Optional

- `account_id` (String)

<a id="nestedatt--cloud"></a>
### Nested Schema for `cloud`

Required:

- `region` (String)
- `vendor` (String)


<a id="nestedatt--compute"></a>
### Nested Schema for `compute`

Required:

- `accelerator` (String)
- `instance_size` (String)
- `instance_type` (String)
- `scaling` (Attributes) (see [below for nested schema](#nestedatt--compute--scaling))

<a id="nestedatt--compute--scaling"></a>
### Nested Schema for `compute.scaling`

Required:

- `max_replica` (Number)
- `min_replica` (Number)

Optional:

- `scale_to_zero_timeout` (Number)



<a id="nestedatt--model"></a>
### Nested Schema for `model`

Required:

- `framework` (String)
- `image` (Attributes) (see [below for nested schema](#nestedatt--model--image))
- `repository` (String)
- `task` (String)

Optional:

- `revision` (String)

<a id="nestedatt--model--image"></a>
### Nested Schema for `model.image`

Optional:

- `custom` (Attributes) (see [below for nested schema](#nestedatt--model--image--custom))
- `huggingface` (Attributes) (see [below for nested schema](#nestedatt--model--image--huggingface))

<a id="nestedatt--model--image--custom"></a>
### Nested Schema for `model.image.custom`

Required:

- `url` (String)

Optional:

- `credentials` (Attributes) (see [below for nested schema](#nestedatt--model--image--custom--credentials))
- `env` (Map of String)
- `health_route` (String)
- `port` (Number)

<a id="nestedatt--model--image--custom--credentials"></a>
### Nested Schema for `model.image.custom.credentials`

Required:

- `password` (String)
- `username` (String)



<a id="nestedatt--model--image--huggingface"></a>
### Nested Schema for `model.image.huggingface`

Optional:

- `env` (Map of String)
