package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// BGPAdvertisementSpec defines the desired state of BGPAdvertisement.
type BGPAdvertisementSpec struct {
	// The aggregation-length advertisement option lets you “roll up” the /32s into a
	// larger prefix. Defaults to 32. Works for IPv4 addresses.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default:=32
	// +optional
	AggregationLength *int32 `json:"aggregationLength,omitempty"`

	// The aggregation-length advertisement option lets you “roll up” the /128s into a
	// larger prefix. Defaults to 128. Works for IPv6 addresses.
	// +kubebuilder:default:=128
	// +optional
	AggregationLengthV6 *int32 `json:"aggregationLengthV6,omitempty"`

	// The BGP LOCAL_PREF attribute which is used by BGP best path algorithm,
	// Path with higher localpref is preferred over one with lower localpref.
	// +optional
	LocalPref uint32 `json:"localPref,omitempty"`

	// The BGP communities to be associated with the announcement. Each item can be a
	// community of the form 1234:1234 or the name of an alias defined in the Community CRD.
	// +optional
	Communities []string `json:"communities,omitempty"`

	// The list of IPAddressPools to advertise via this advertisement, selected by name.
	// +optional
	IPAddressPools []string `json:"ipAddressPools,omitempty"`

	// A selector for the IPAddressPools which would get advertised via this advertisement.
	// If no IPAddressPool is selected by this or by the list, the advertisement is applied to all the IPAddressPools.
	// +optional
	IPAddressPoolSelectors []metav1.LabelSelector `json:"ipAddressPoolSelectors,omitempty"`

	// NodeSelectors allows to limit the nodes to announce as next hops for the LoadBalancer IP.
	// When empty, all the nodes having  are announced as next hops.
	// +optional
	NodeSelectors []metav1.LabelSelector `json:"nodeSelectors,omitempty"`

	// Peers limits the bgppeer to advertise the ips of the selected pools to.
	// When empty, the loadbalancer IP is announced to all the BGPPeers configured.
	// +optional
	Peers []string `json:"peers,omitempty"`
}

// BGPAdvertisementStatus defines the observed state of BGPAdvertisement.
type BGPAdvertisementStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// BGPAdvertisement allows to advertise the IPs coming
// from the selected IPAddressPools via BGP, setting the parameters of the
// BGP Advertisement.
type BGPAdvertisement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BGPAdvertisementSpec   `json:"spec,omitempty"`
	Status BGPAdvertisementStatus `json:"status,omitempty"`
}

// BGPAdvertisementList contains a list of BGPAdvertisement.
type BGPAdvertisementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BGPAdvertisement `json:"items"`
}
