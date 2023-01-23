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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UpfImplementationSpec defines the desired state of UpfImplementation
type UpfImplementationSpec struct {
	Implementation string `json:"implementation,omitempty"`
}

// UpfImplementationStatus defines the observed state of UpfImplementation
type UpfImplementationStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UpfImplementation is the Schema for the upfimplementations API
type UpfImplementation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpfImplementationSpec   `json:"spec,omitempty"`
	Status UpfImplementationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UpfImplementationList contains a list of UpfImplementation
type UpfImplementationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpfImplementation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UpfImplementation{}, &UpfImplementationList{})
}

var (
	UpfImplementationKind = reflect.TypeOf(UpfImplementation{}).Name()
)
