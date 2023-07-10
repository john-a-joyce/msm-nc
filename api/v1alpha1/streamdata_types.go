package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StreamdataSpec defines the desired state of Streamdata
type StreamdataSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Streamdata. Edit streamdata_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// StreamdataStatus defines the observed state of Streamdata
type StreamdataStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Streamdata is the Schema for the streamdata API
type Streamdata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StreamdataSpec   `json:"spec,omitempty"`
	Status StreamdataStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StreamdataList contains a list of Streamdata
type StreamdataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Streamdata `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Streamdata{}, &StreamdataList{})
}
