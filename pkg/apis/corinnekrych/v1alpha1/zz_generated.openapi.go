// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.ArchDecisionRecord":       schema_pkg_apis_corinnekrych_v1alpha1_ArchDecisionRecord(ref),
		"github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.ArchDecisionRecordSpec":   schema_pkg_apis_corinnekrych_v1alpha1_ArchDecisionRecordSpec(ref),
		"github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.ArchDecisionRecordStatus": schema_pkg_apis_corinnekrych_v1alpha1_ArchDecisionRecordStatus(ref),
	}
}

func schema_pkg_apis_corinnekrych_v1alpha1_ArchDecisionRecord(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArchDecisionRecord is the Schema for the archdecisionrecords API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.ArchDecisionRecordSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.ArchDecisionRecordStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.ArchDecisionRecordSpec", "github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.ArchDecisionRecordStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_corinnekrych_v1alpha1_ArchDecisionRecordSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArchDecisionRecordSpec defines the desired state of ArchDecisionRecord",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Container image use to build",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"source": {
						SchemaProps: spec.SchemaProps{
							Description: "Location of the source ir: github url, where the ADR is located",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"image", "source"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_corinnekrych_v1alpha1_ArchDecisionRecordStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArchDecisionRecordStatus defines the observed state of ArchDecisionRecord",
				Properties: map[string]spec.Schema{
					"steps": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.Step"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1.Step"},
	}
}
