/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TrafficSplitSpec defines the desired state of TrafficSplit
type TrafficSplitSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Service  string                `json:"service,omitempty"`
	Backends []TrafficSplitBackend `json:"backends,omitempty"`
}

// TrafficSplitBackend defines a backend
type TrafficSplitBackend struct {
	Service string             `json:"service,omitempty"`
	Weight  *resource.Quantity `json:"weight,omitempty"`
}

// TrafficSplitStatus defines the observed state of TrafficSplit
type TrafficSplitStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TrafficSplit is the Schema for the trafficsplits API
type TrafficSplit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficSplitSpec   `json:"spec,omitempty"`
	Status TrafficSplitStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TrafficSplitList contains a list of TrafficSplit
type TrafficSplitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficSplit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficSplit{}, &TrafficSplitList{})
}
