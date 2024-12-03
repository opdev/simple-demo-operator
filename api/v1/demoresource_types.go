/*
Copyright 2021 The Partner Lifecycle Engineering Team.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DemoResourceSpec defines the desired state of DemoResource
type DemoResourceSpec struct {
	// Message is a random string.
	Message string `json:"message"`
}

// DemoResourceStatus defines the observed state of DemoResource
type DemoResourceStatus struct {
	// SpecMessage reflects the message passed by the user in the spec.
	SpecMessage string `json:"specMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DemoResource is the Schema for the demoresources API
// +kubebuilder:printcolumn:name="Message",type=string,JSONPath=`.status.specMessage`
type DemoResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoResourceSpec   `json:"spec,omitempty"`
	Status DemoResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DemoResourceList contains a list of DemoResource
type DemoResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoResource{}, &DemoResourceList{})
}
