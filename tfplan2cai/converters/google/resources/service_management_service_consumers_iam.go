// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "fmt"

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ServiceManagementServiceConsumersIAMAssetType string = "servicemanagement.googleapis.com/ServiceConsumers"

func resourceConverterServiceManagementServiceConsumersIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         ServiceManagementServiceConsumersIAMAssetType,
		Convert:           GetServiceManagementServiceConsumersIamPolicyCaiObject,
		MergeCreateUpdate: MergeServiceManagementServiceConsumersIamPolicy,
	}
}

func resourceConverterServiceManagementServiceConsumersIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         ServiceManagementServiceConsumersIAMAssetType,
		Convert:           GetServiceManagementServiceConsumersIamBindingCaiObject,
		FetchFullResource: FetchServiceManagementServiceConsumersIamPolicy,
		MergeCreateUpdate: MergeServiceManagementServiceConsumersIamBinding,
		MergeDelete:       MergeServiceManagementServiceConsumersIamBindingDelete,
	}
}

func resourceConverterServiceManagementServiceConsumersIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         ServiceManagementServiceConsumersIAMAssetType,
		Convert:           GetServiceManagementServiceConsumersIamMemberCaiObject,
		FetchFullResource: FetchServiceManagementServiceConsumersIamPolicy,
		MergeCreateUpdate: MergeServiceManagementServiceConsumersIamMember,
		MergeDelete:       MergeServiceManagementServiceConsumersIamMemberDelete,
	}
}

func GetServiceManagementServiceConsumersIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newServiceManagementServiceConsumersIamAsset(d, config, expandIamPolicyBindings)
}

func GetServiceManagementServiceConsumersIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newServiceManagementServiceConsumersIamAsset(d, config, expandIamRoleBindings)
}

func GetServiceManagementServiceConsumersIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newServiceManagementServiceConsumersIamAsset(d, config, expandIamMemberBindings)
}

func MergeServiceManagementServiceConsumersIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeServiceManagementServiceConsumersIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeServiceManagementServiceConsumersIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeServiceManagementServiceConsumersIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeServiceManagementServiceConsumersIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newServiceManagementServiceConsumersIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//servicemanagement.googleapis.com/services/{{service_name}}/consumers/{{consumer_project}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: ServiceManagementServiceConsumersIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchServiceManagementServiceConsumersIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("service_name"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("consumer_project"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		ServiceManagementServiceConsumersIamUpdaterProducer,
		d,
		config,
		"//servicemanagement.googleapis.com/services/{{service_name}}/consumers/{{consumer_project}}",
		ServiceManagementServiceConsumersIAMAssetType,
	)
}
