package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Endpoint struct {
	AccountId *string `tfsdk:"account_id"`
	Compute   Compute `tfsdk:"compute"`
	Model     Model   `tfsdk:"model"`
	Name      string  `tfsdk:"name"`
	Cloud     Cloud   `tfsdk:"cloud"`
	Status    *Status `tfsdk:"status"`
	Type      string  `tfsdk:"type"`
}

type Compute struct {
	Accelerator  string  `tfsdk:"accelerator"`
	InstanceSize string  `tfsdk:"instance_size"`
	InstanceType string  `tfsdk:"instance_type"`
	Scaling      Scaling `tfsdk:"scaling"`
}

type Scaling struct {
	MaxReplica         int         `tfsdk:"max_replica"`
	MinReplica         int         `tfsdk:"min_replica"`
	ScaleToZeroTimeout types.Int64 `tfsdk:"scale_to_zero_timeout"`
}

type Model struct {
	Framework  string       `tfsdk:"framework"`
	Image      Image        `tfsdk:"image"`
	Repository string       `tfsdk:"repository"`
	Revision   types.String `tfsdk:"revision"`
	Task       types.String `tfsdk:"task"`
}

type Image struct {
	Huggingface *Huggingface `tfsdk:"huggingface"`
	Custom      *Custom      `tfsdk:"custom"`
	Tei         *Tei         `tfsdk:"tei"`
	Tgi         *Tgi         `tfsdk:"tgi"`
	TgiTpu      *TgiTpu      `tfsdk:"tgi_tpu"`
	TgiNeuron   *TgiNeuron   `tfsdk:"tgi_neuron"`
	Llamacpp    *Llamacpp    `tfsdk:"llamacpp"`
}

type Tei struct {
	HealthRoute           *string     `json:"health_route,omitempty"`
	Port                  types.Int64 `json:"port,omitempty"`
	URL                   string      `json:"url"`
	MaxBatchTokens        *int        `json:"maxBatchTokens,omitempty"`
	MaxConcurrentRequests *int        `json:"maxConcurrentRequests,omitempty"`
	Pooling               *string     `json:"pooling,omitempty"`
}

type Llamacpp struct {
	HealthRoute *string     `tfsdk:"health_route"`
	Port        types.Int64 `tfsdk:"port"`
	URL         string      `tfsdk:"url"`
	CtxSize     *int        `tfsdk:"ctxSize"`
	Embeddings  *bool       `tfsdk:"embeddings"`
	ModelPath   string      `tfsdk:"modelPath"`
	NParallel   *int        `tfsdk:"nParallel"`
	ThreadsHttp *int        `tfsdk:"threadsHttp"`
}

type TgiNeuron struct {
	HealthRoute           *string     `tfsdk:"health_route"`
	Port                  types.Int64 `tfsdk:"port"`
	URL                   string      `tfsdk:"url"`
	MaxBatchPrefillTokens *int        `tfsdk:"maxBatchPrefillTokens"`
	MaxBatchTotalTokens   *int        `tfsdk:"maxBatchTotalTokens"`
	MaxInputLength        *int        `tfsdk:"maxInputLength"`
	MaxTotalTokens        *int        `tfsdk:"maxTotalTokens"`
	HfAutoCastType        *string     `tfsdk:"hfAutoCastType"`
	HfNumCores            *int        `tfsdk:"hfNumCores"`
}

type TgiTpu struct {
	HealthRoute           *string     `tfsdk:"health_route"`
	Port                  types.Int64 `tfsdk:"port"`
	URL                   string      `tfsdk:"url"`
	MaxBatchPrefillTokens *int        `tfsdk:"maxBatchPrefillTokens"`
	MaxBatchTotalTokens   *int        `tfsdk:"maxBatchTotalTokens"`
	MaxInputLength        *int        `tfsdk:"maxInputLength"`
	MaxTotalTokens        *int        `tfsdk:"maxTotalTokens"`
	DisableCustomKernels  *bool       `tfsdk:"disableCustomKernels"`
	Quantize              *string     `tfsdk:"quantize"`
}

type Tgi struct {
	HealthRoute           *string     `tfsdk:"health_route"`
	Port                  types.Int64 `tfsdk:"port"`
	URL                   string      `tfsdk:"url"`
	MaxBatchPrefillTokens *int        `tfsdk:"maxBatchPrefillTokens"`
	MaxBatchTotalTokens   *int        `tfsdk:"maxBatchTotalTokens"`
	MaxInputLength        *int        `tfsdk:"maxInputLength"`
	MaxTotalTokens        *int        `tfsdk:"maxTotalTokens"`
	DisableCustomKernels  *bool       `tfsdk:"disableCustomKernels"`
	Quantize              *string     `tfsdk:"quantize"`
}

type Custom struct {
	Credentials *Credentials      `tfsdk:"credentials"`
	Env         map[string]string `tfsdk:"env"`
	HealthRoute *string           `tfsdk:"health_route"`
	Port        types.Int64       `tfsdk:"port"`
	URL         string            `tfsdk:"url"`
}

type Credentials struct {
	Password string `tfsdk:"password"`
	Username string `tfsdk:"username"`
}

type Huggingface struct {
	Env map[string]string `tfsdk:"env"`
}

type Cloud struct {
	Region string `tfsdk:"region"`
	Vendor string `tfsdk:"vendor"`
}

type Status struct {
	CreatedAt     string  `tfsdk:"created_at"`
	CreatedBy     User    `tfsdk:"created_by"`
	ErrorMessage  string  `tfsdk:"error_message"`
	Message       string  `tfsdk:"message"`
	Private       Private `tfsdk:"private"`
	ReadyReplica  int     `tfsdk:"ready_replica"`
	State         string  `tfsdk:"state"`
	TargetReplica int     `tfsdk:"target_replica"`
	UpdatedAt     string  `tfsdk:"updated_at"`
	UpdatedBy     User    `tfsdk:"updated_by"`
	URL           string  `tfsdk:"url"`
}

type User struct {
	ID   string `tfsdk:"id"`
	Name string `tfsdk:"name"`
}

type Private struct {
	ServiceName string `tfsdk:"service_name"`
}
