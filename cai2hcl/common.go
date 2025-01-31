package cai2hcl

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	tpg "github.com/hashicorp/terraform-provider-google/google"
	"github.com/zclconf/go-cty/cty"
)

// Converter interface for resources.
type Converter interface {
	// Convert turns assets into hcl blocks.
	Convert(asset []*caiasset.Asset) ([]*HCLResourceBlock, error)
}

// HCLResourceBlock identifies the HCL block's labels and content.
type HCLResourceBlock struct {
	Labels []string
	Value  cty.Value
}

// converterNames map key is the CAI Asset type, value is the TF resource name.
var converterNames = map[string]string{
	ComputeInstanceAssetType: "google_compute_instance",
	ProjectAssetType:         "google_project",
	ProjectBillingAssetType:  "google_project",
}

// converterMap initializes converters by their TF resource name.
var converterMap = map[string]Converter{
	"google_compute_instance": NewComputeInstanceConverter(),
	"google_project":          NewProjectConverter(),
}

// schemaProvider has schemas for all resources.
var schemaProvider = tpg.Provider()
