package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ArchDecisionRecordSpec defines the desired state of ArchDecisionRecord
// +k8s:openapi-gen=true
type ArchDecisionRecordSpec struct {
    // Container image use to build
	Image string	`json:"image"`
	// Location of the source ir: github url, where the ADR is located
	Source string   `json:"source"`
}

// ArchDecisionRecordStatus defines the observed state of ArchDecisionRecord
// +k8s:openapi-gen=true
type ArchDecisionRecordStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Steps []Step		`json:"steps,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArchDecisionRecord is the Schema for the archdecisionrecords API
// +k8s:openapi-gen=true
type ArchDecisionRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArchDecisionRecordSpec   `json:"spec,omitempty"`
	Status ArchDecisionRecordStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArchDecisionRecordList contains a list of ArchDecisionRecord
type ArchDecisionRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArchDecisionRecord `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArchDecisionRecord{}, &ArchDecisionRecordList{})
}
type Step struct {
	Name StepName `json:"name,omitempty"`
	Phase Phase `json:"phase,omitempty"`
}
type Phase string
const (
	Creating Phase = "Creating"
	Created Phase = "Created"
)
type StepName string
const (
	ImageStreamCreate StepName = "ImageStreamCreate"
)