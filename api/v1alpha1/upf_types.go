/*
Copyright 2023 nokia.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UpfSpec defines the desired state of Upf
type UpfSpec struct {
	Plmn string `json:"plmn,omitempty"`
}

// UpfStatus defines the observed state of Upf
type UpfStatus struct {
	Implementation string `json:"implementation,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Upf is the Schema for the upfs API
type Upf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpfSpec   `json:"spec,omitempty"`
	Status UpfStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UpfList contains a list of Upf
type UpfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Upf `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Upf{}, &UpfList{})
}
