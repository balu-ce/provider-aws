/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this AccessPoint
func (mg *AccessPoint) GetTerraformResourceType() string {
	return "aws_s3_access_point"
}

// GetConnectionDetailsMapping for this AccessPoint
func (tr *AccessPoint) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AccessPoint
func (tr *AccessPoint) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AccessPoint
func (tr *AccessPoint) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AccessPoint
func (tr *AccessPoint) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AccessPoint
func (tr *AccessPoint) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AccessPoint
func (tr *AccessPoint) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AccessPoint using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AccessPoint) LateInitialize(attrs []byte) (bool, error) {
	params := &AccessPointParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AccessPoint) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AccountPublicAccessBlock
func (mg *AccountPublicAccessBlock) GetTerraformResourceType() string {
	return "aws_s3_account_public_access_block"
}

// GetConnectionDetailsMapping for this AccountPublicAccessBlock
func (tr *AccountPublicAccessBlock) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AccountPublicAccessBlock
func (tr *AccountPublicAccessBlock) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AccountPublicAccessBlock
func (tr *AccountPublicAccessBlock) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AccountPublicAccessBlock
func (tr *AccountPublicAccessBlock) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AccountPublicAccessBlock
func (tr *AccountPublicAccessBlock) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AccountPublicAccessBlock
func (tr *AccountPublicAccessBlock) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AccountPublicAccessBlock using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AccountPublicAccessBlock) LateInitialize(attrs []byte) (bool, error) {
	params := &AccountPublicAccessBlockParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AccountPublicAccessBlock) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AccessPointPolicy
func (mg *AccessPointPolicy) GetTerraformResourceType() string {
	return "aws_s3control_access_point_policy"
}

// GetConnectionDetailsMapping for this AccessPointPolicy
func (tr *AccessPointPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AccessPointPolicy
func (tr *AccessPointPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AccessPointPolicy
func (tr *AccessPointPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AccessPointPolicy
func (tr *AccessPointPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AccessPointPolicy
func (tr *AccessPointPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AccessPointPolicy
func (tr *AccessPointPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AccessPointPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AccessPointPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &AccessPointPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AccessPointPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MultiRegionAccessPoint
func (mg *MultiRegionAccessPoint) GetTerraformResourceType() string {
	return "aws_s3control_multi_region_access_point"
}

// GetConnectionDetailsMapping for this MultiRegionAccessPoint
func (tr *MultiRegionAccessPoint) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MultiRegionAccessPoint
func (tr *MultiRegionAccessPoint) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MultiRegionAccessPoint
func (tr *MultiRegionAccessPoint) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MultiRegionAccessPoint
func (tr *MultiRegionAccessPoint) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MultiRegionAccessPoint
func (tr *MultiRegionAccessPoint) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MultiRegionAccessPoint
func (tr *MultiRegionAccessPoint) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MultiRegionAccessPoint using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MultiRegionAccessPoint) LateInitialize(attrs []byte) (bool, error) {
	params := &MultiRegionAccessPointParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MultiRegionAccessPoint) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MultiRegionAccessPointPolicy
func (mg *MultiRegionAccessPointPolicy) GetTerraformResourceType() string {
	return "aws_s3control_multi_region_access_point_policy"
}

// GetConnectionDetailsMapping for this MultiRegionAccessPointPolicy
func (tr *MultiRegionAccessPointPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MultiRegionAccessPointPolicy
func (tr *MultiRegionAccessPointPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MultiRegionAccessPointPolicy
func (tr *MultiRegionAccessPointPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MultiRegionAccessPointPolicy
func (tr *MultiRegionAccessPointPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MultiRegionAccessPointPolicy
func (tr *MultiRegionAccessPointPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MultiRegionAccessPointPolicy
func (tr *MultiRegionAccessPointPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MultiRegionAccessPointPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MultiRegionAccessPointPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &MultiRegionAccessPointPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MultiRegionAccessPointPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ObjectLambdaAccessPoint
func (mg *ObjectLambdaAccessPoint) GetTerraformResourceType() string {
	return "aws_s3control_object_lambda_access_point"
}

// GetConnectionDetailsMapping for this ObjectLambdaAccessPoint
func (tr *ObjectLambdaAccessPoint) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ObjectLambdaAccessPoint
func (tr *ObjectLambdaAccessPoint) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ObjectLambdaAccessPoint
func (tr *ObjectLambdaAccessPoint) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ObjectLambdaAccessPoint
func (tr *ObjectLambdaAccessPoint) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ObjectLambdaAccessPoint
func (tr *ObjectLambdaAccessPoint) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ObjectLambdaAccessPoint
func (tr *ObjectLambdaAccessPoint) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ObjectLambdaAccessPoint using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ObjectLambdaAccessPoint) LateInitialize(attrs []byte) (bool, error) {
	params := &ObjectLambdaAccessPointParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ObjectLambdaAccessPoint) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ObjectLambdaAccessPointPolicy
func (mg *ObjectLambdaAccessPointPolicy) GetTerraformResourceType() string {
	return "aws_s3control_object_lambda_access_point_policy"
}

// GetConnectionDetailsMapping for this ObjectLambdaAccessPointPolicy
func (tr *ObjectLambdaAccessPointPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ObjectLambdaAccessPointPolicy
func (tr *ObjectLambdaAccessPointPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ObjectLambdaAccessPointPolicy
func (tr *ObjectLambdaAccessPointPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ObjectLambdaAccessPointPolicy
func (tr *ObjectLambdaAccessPointPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ObjectLambdaAccessPointPolicy
func (tr *ObjectLambdaAccessPointPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ObjectLambdaAccessPointPolicy
func (tr *ObjectLambdaAccessPointPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ObjectLambdaAccessPointPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ObjectLambdaAccessPointPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &ObjectLambdaAccessPointPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ObjectLambdaAccessPointPolicy) GetTerraformSchemaVersion() int {
	return 0
}
