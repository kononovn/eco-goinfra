package networkpolicy

import (
	"context"
	"fmt"
	"github.com/k8snetworkplumbingwg/multi-networkpolicy/pkg/apis/k8s.cni.cncf.io/v1beta1"

	"github.com/golang/glog"
	"github.com/openshift-kni/eco-goinfra/pkg/clients"
	"github.com/openshift-kni/eco-goinfra/pkg/msg"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MultiNetworkPolicyBuilder provides struct for multiNetworkPolicy object.
type MultiNetworkPolicyBuilder struct {
	// MultiNetworkPolicy definition. Used to create MultiNetworkPolicy object with minimum set of required elements.
	Definition *v1beta1.MultiNetworkPolicy
	// Created MultiNetworkPolicy object on the cluster.
	Object *v1beta1.MultiNetworkPolicy
	// api client to interact with the cluster.
	apiClient *clients.Settings
	// errorMsg is processed before NetworkPolicy object is created.
	errorMsg string
}

// NewMultiNetworkPolicyBuilder method creates new instance of builder.
func NewMultiNetworkPolicyBuilder(apiClient *clients.Settings, name, nsname string) *MultiNetworkPolicyBuilder {
	glog.V(100).Infof(
		"Initializing new MultiNetworkPolicyBuilder structure with the following params: name: %s, namespace: %s",
		name, nsname)

	builder := &MultiNetworkPolicyBuilder{
		apiClient: apiClient,
		Definition: &v1beta1.MultiNetworkPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: nsname,
			},
		},
	}

	if name == "" {
		glog.V(100).Infof("The name of the multiNetworkPolicy is empty")

		builder.errorMsg = "The multiNetworkPolicy 'name' cannot be empty"
	}

	if nsname == "" {
		glog.V(100).Infof("The namespace of the multiNetworkPolicy is empty")

		builder.errorMsg = "The multiNetworkPolicy 'namespace' cannot be empty"
	}

	return builder
}

// Pull loads an existing multiNetworkPolicy into the Builder struct.
func Pull(apiClient *clients.Settings, name, nsname string) (*MultiNetworkPolicyBuilder, error) {
	glog.V(100).Infof("Pulling existing multiNetworkPolicy name: %s namespace:%s", name, nsname)

	builder := MultiNetworkPolicyBuilder{
		apiClient: apiClient,
		Definition: &v1beta1.MultiNetworkPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: nsname,
			},
		},
	}

	if name == "" {
		glog.V(100).Infof("The name of the multiNetworkPolicy is empty")

		builder.errorMsg = "multiNetworkPolicy 'name' cannot be empty"
	}

	if nsname == "" {
		glog.V(100).Infof("The namespace of the multiNetworkPolicy is empty")

		builder.errorMsg = "multiNetworkPolicy 'namespace' cannot be empty"
	}

	if builder.errorMsg != "" {
		return nil, fmt.Errorf("failed to pull multiNetworkPolicy object due to the following error: %s",
			builder.errorMsg)
	}

	if !builder.Exists() {
		glog.V(100).Infof(
			"Failed to pull multiNetworkPolicy object %s from namespace %s. Object doesn't exist",
			name, nsname)

		return nil, fmt.Errorf("multiNetworkPolicy object %s doesn't exist in namespace %s", name, nsname)
	}

	builder.Definition = builder.Object

	return &builder, nil
}

// Create makes a multiNetworkPolicy in cluster and stores the created object in struct.
func (builder *MultiNetworkPolicyBuilder) Create() (*MultiNetworkPolicyBuilder, error) {
	if valid, err := builder.validate(); !valid {
		return builder, err
	}

	glog.V(100).Infof("Creating the networkPolicy %s in %s namespace",
		builder.Definition.Name, builder.Definition.Namespace)

	var err error
	if !builder.Exists() {
		builder.Object, err = builder.apiClient.NetworkPolicies(builder.Definition.Namespace).Create(
			context.TODO(), builder.Definition, metav1.CreateOptions{})
	}

	return builder, err
}

// Exists checks whether the given MultiNetworkPolicy exists.
func (builder *MultiNetworkPolicyBuilder) Exists() bool {
	if valid, _ := builder.validate(); !valid {
		return false
	}

	glog.V(100).Infof("Checking if multiNetworkPolicy %s exists in namespace %s",
		builder.Definition.Name, builder.Definition.Namespace)

	var err error
	builder.Object, err = builder.apiClient.NetworkPolicies(builder.Definition.Namespace).Get(
		context.Background(), builder.Definition.Name, metav1.GetOptions{})

	return err == nil || !k8serrors.IsNotFound(err)
}

// Delete removes a multiNetworkPolicy object from a cluster.
func (builder *MultiNetworkPolicyBuilder) Delete() error {
	if valid, err := builder.validate(); !valid {
		return err
	}

	glog.V(100).Infof("Deleting the multiNetworkPolicy object %s from %s namespace",
		builder.Definition.Name, builder.Definition.Namespace)

	if !builder.Exists() {
		return fmt.Errorf("multiNetworkPolicy cannot be deleted because it does not exist")
	}

	err := builder.apiClient.NetworkPolicies(builder.Definition.Namespace).Delete(
		context.TODO(), builder.Definition.Name, metav1.DeleteOptions{})

	if err != nil {
		return fmt.Errorf("cannot delete MachineConfig: %w", err)
	}

	builder.Object = nil

	return err
}

// Update renovates the existing multiNetworkPolicy object with networkPolicy definition in builder.
func (builder *MultiNetworkPolicyBuilder) Update() (*MultiNetworkPolicyBuilder, error) {
	if valid, err := builder.validate(); !valid {
		return builder, err
	}

	glog.V(100).Infof("Updating multiNetworkPolicy %s in %s namespace ",
		builder.Definition.Name, builder.Definition.Namespace)

	var err error
	builder.Object, err = builder.apiClient.NetworkPolicies(builder.Definition.Namespace).Update(
		context.TODO(), builder.Definition, metav1.UpdateOptions{})

	return builder, err
}

// validate will check that the builder and builder definition are properly initialized before
// accessing any member fields.
func (builder *MultiNetworkPolicyBuilder) validate() (bool, error) {
	resourceCRD := "NetworkPolicy"

	if builder == nil {
		glog.V(100).Infof("The %s builder is uninitialized", resourceCRD)

		return false, fmt.Errorf("error: received nil %s builder", resourceCRD)
	}

	if builder.Definition == nil {
		glog.V(100).Infof("The %s is undefined", resourceCRD)

		builder.errorMsg = msg.UndefinedCrdObjectErrString(resourceCRD)
	}

	if builder.apiClient == nil {
		glog.V(100).Infof("The %s builder apiclient is nil", resourceCRD)

		builder.errorMsg = fmt.Sprintf("%s builder cannot have nil apiClient", resourceCRD)
	}

	if builder.errorMsg != "" {
		glog.V(100).Infof("The %s builder has error message: %s", resourceCRD, builder.errorMsg)

		return false, fmt.Errorf(builder.errorMsg)
	}

	return true, nil
}
