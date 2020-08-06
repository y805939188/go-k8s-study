/*


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

package v1alpha666

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DingShin888TypeSpec defines the desired state of DingShin888Type
type DingShin888TypeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DingShin888Type. Edit DingShin888Type_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DingShin888TypeStatus defines the observed state of DingShin888Type
type DingShin888TypeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// DingShin888Type is the Schema for the dingshin888types API
type DingShin888Type struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DingShin888TypeSpec   `json:"spec,omitempty"`
	Status DingShin888TypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DingShin888TypeList contains a list of DingShin888Type
type DingShin888TypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DingShin888Type `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DingShin888Type{}, &DingShin888TypeList{})
}
