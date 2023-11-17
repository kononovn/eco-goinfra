package mlbtypes

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type MatchExpression struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	// +kubebuilder:validation:MinItems:=1
	Values []string `json:"values"`
}

type NodeSelector struct {
	// +optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	// +optional
	MatchExpressions []MatchExpression `json:"matchExpressions,omitempty"`
}

// BGPPeerSpec defines the desired state of Peer.
type BGPPeerSpec struct {
	// AS number to use for the local end of the session.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	MyASN uint32 `json:"myASN"`

	// AS number to expect from the remote end of the session.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4294967295
	ASN uint32 `json:"peerASN"`

	// Address to dial when establishing the session.
	Address string `json:"peerAddress"`

	// Source address to use when establishing the session.
	// +optional
	SrcAddress string `json:"sourceAddress,omitempty"`

	// Port to dial when establishing the session.
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=16384
	Port uint16 `json:"peerPort,omitempty"`

	// Requested BGP hold time, per RFC4271.
	// +optional
	HoldTime metav1.Duration `json:"holdTime,omitempty"`

	// Requested BGP keepalive time, per RFC4271.
	// +optional
	KeepaliveTime metav1.Duration `json:"keepaliveTime,omitempty"`

	// BGP router ID to advertise to the peer
	// +optional
	RouterID string `json:"routerID,omitempty"`

	// Only connect to this peer on nodes that match one of these
	// selectors.
	// +optional
	NodeSelectors []NodeSelector `json:"nodeSelectors,omitempty"`

	// Authentication password for routers enforcing TCP MD5 authenticated sessions
	// +optional
	Password string `json:"password,omitempty"`

	BFDProfile string `json:"bfdProfile,omitempty"`

	// EBGP peer is multi-hops away
	EBGPMultiHop bool `json:"ebgpMultiHop,omitempty"`
	// Add future BGP configuration here
}

// BGPPeerStatus defines the observed state of Peer.
type BGPPeerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// BGPPeer is the Schema for the peers API.
type BGPPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BGPPeerSpec   `json:"spec,omitempty"`
	Status BGPPeerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// PeerList contains a list of Peer.
type BGPPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BGPPeer `json:"items"`
}
