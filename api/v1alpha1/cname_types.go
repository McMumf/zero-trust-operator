/*
Copyright 2025 Carson McCaffrey.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CnameSpec defines the desired state of Cname.
type CnameSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	CNAME string `json:"cname,omitempty"`
}

// CnameStatus defines the observed state of Cname.
type CnameStatus struct {
	// +nullable
	// The time and date the cname was fetched and tunnel provider updated
	RefreshTime metav1.Time `json:"refreshTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Cname is the Schema for the cnames API.
type Cname struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CnameSpec   `json:"spec,omitempty"`
	Status CnameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CnameList contains a list of Cname.
type CnameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cname `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cname{}, &CnameList{})
}
