// +build !ignore_autogenerated

// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallation":     schema_pkg_apis_mattermost_v1alpha1_ClusterInstallation(ref),
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec": schema_pkg_apis_mattermost_v1alpha1_ClusterInstallationSpec(ref),
	}
}

func schema_pkg_apis_mattermost_v1alpha1_ClusterInstallation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterInstallation is the Schema for the clusterinstallations API",
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
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of the Mattermost cluster. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status",
							Ref:         ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recent observed status of the Mattermost cluster. Read-only. Not included when requesting from the apiserver, only from the Mattermost Operator API itself. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status",
							Ref:         ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationStatus"},
	}
}

func schema_pkg_apis_mattermost_v1alpha1_ClusterInstallationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterInstallationSpec defines the desired state of ClusterInstallation",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image defines the ClusterInstallation Docker image.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version defines the ClusterInstallation Docker image version.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "Size defines the size of the ClusterInstallation. This is typically specified in number of users. This will set replica and resource requests/limits appropriately for the provided number of users. Accepted values are: 100users, 1000users, 5000users, 10000users, 250000users. Defaults to 5000users. Setting 'Replicas', 'Resources', 'Minio.Replicas', 'Minio.Resource', 'Database.Replicas', or 'Database.Resources' will override the values set by Size.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas defines the number of replicas to use for the Mattermost app servers. Setting this will override the number of replicas set by 'Size'.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Defines the resource requests and limits for the Mattermost app server pods.",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"ingressName": {
						SchemaProps: spec.SchemaProps{
							Description: "IngressName defines the name to be used when creating the ingress rules",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"mattermostLicenseSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret that contains the mattermost license",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified, affinity will define the pod's scheduling constraints",
							Ref:         ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"minio": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Minio"),
						},
					},
					"database": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Database"),
						},
					},
					"blueGreen": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.BlueGreen"),
						},
					},
					"elasticSearch": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ElasticSearch"),
						},
					},
					"useServiceLoadBalancer": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"serviceAnnotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"ingressAnnotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"ingressName"},
			},
		},
		Dependencies: []string{
			"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.BlueGreen", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Database", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ElasticSearch", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Minio", "k8s.io/api/core/v1.Affinity", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}
