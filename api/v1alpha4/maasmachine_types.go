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

package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/cluster-api/errors"
)

const (
	// MachineFinalizer allows MaasMachineReconciler to clean up resources associated with MaasMachine before
	// removing it from the apiserver.
	MachineFinalizer = "maascluster.infrastructure.cluster.x-k8s.io"
)

// MaasMachineSpec defines the desired state of MaasMachine
type MaasMachineSpec struct {
	// ProviderID will be the container name in ProviderID format (maas://<zone>/system_id)
	// +optional
	ProviderID *string `json:"providerID,omitempty"`

	// ProviderID will be the container name in ProviderID format (maas://<zone>/system_id)
	// +optional
	SystemID *string `json:"systemID,omitempty"`
}

// MaasMachineStatus defines the observed state of MaasMachine
type MaasMachineStatus struct {

	// Ready denotes that the machine (maas container) is ready
	// +kubebuilder:default=false
	Ready bool `json:"ready"`

	// InstanceState is the state of the AWS instance for this machine.
	MachineState *MachineState `json:"instanceState,omitempty"`

	// MachinePowered is if the machine is "Powered" on
	MachinePowered bool `json:"machinePowered,omitempty"`

	// LoadBalancerConfigured denotes that the machine has been
	// added to the load balancer
	LoadBalancerConfigured bool `json:"loadBalancerConfigured,omitempty"`

	// Addresses contains the associated addresses for the maas machine.
	Addresses []clusterv1.MachineAddress `json:"addresses,omitempty"`

	// Conditions defines current service state of the MaasMachine.
	Conditions clusterv1.Conditions `json:"conditions,omitempty"`

	// FailureReason will be set in the event that there is a terminal problem
	// reconciling the Machine and will contain a succinct value suitable
	// for machine interpretation.
	FailureReason *errors.MachineStatusError `json:"failureReason,omitempty"`

	// FailureMessage will be set in the event that there is a terminal problem
	// reconciling the Machine and will contain a more verbose string suitable
	// for logging and human consumption.
	FailureMessage *string `json:"failureMessage,omitempty"`
}

// +kubebuilder:resource:path=maasmachines,scope=Namespaced,categories=cluster-api
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// MaasMachine is the Schema for the maasmachines API
type MaasMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MaasMachineSpec   `json:"spec,omitempty"`
	Status MaasMachineStatus `json:"status,omitempty"`
}

func (c *MaasMachine) GetConditions() clusterv1.Conditions {
	return c.Status.Conditions
}

func (c *MaasMachine) SetConditions(conditions clusterv1.Conditions) {
	c.Status.Conditions = conditions
}

// +kubebuilder:object:root=true

// MaasMachineList contains a list of MaasMachine
type MaasMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaasMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MaasMachine{}, &MaasMachineList{})
}
