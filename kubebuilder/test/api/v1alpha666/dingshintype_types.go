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

// DingShinTypeSpec defines the desired state of DingShinType
type DingShinTypeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DingShinType. Edit DingShinType_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DingShinTypeStatus defines the observed state of DingShinType
type DingShinTypeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// DingShinType is the Schema for the dingshintypes API
type DingShinType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DingShinTypeSpec   `json:"spec,omitempty"`
	Status DingShinTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DingShinTypeList contains a list of DingShinType
type DingShinTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DingShinType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DingShinType{}, &DingShinTypeList{})
}
