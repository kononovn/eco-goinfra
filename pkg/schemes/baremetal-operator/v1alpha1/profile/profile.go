package profile

import (
	"fmt"

	metal3api "github.com/openshift-kni/eco-goinfra/pkg/schemes/baremetal-operator/v1alpha1"
)

const (
	// DefaultProfileName is the default hardware profile to use when
	// no other profile matches.
	DefaultProfileName string = "unknown"

	// EmptyProfileName is the hardware profile without configuration.
	EmptyProfileName string = "empty"
)

// Profile holds the settings for a class of hardware.
type Profile struct {
	// Name holds the profile name
	Name string

	// RootDeviceHints holds the suggestions for placing the storage
	// for the root filesystem.
	RootDeviceHints metal3api.RootDeviceHints
}

var profiles = make(map[string]Profile)

func init() {
	profiles[DefaultProfileName] = Profile{
		Name: DefaultProfileName,
		RootDeviceHints: metal3api.RootDeviceHints{
			DeviceName: "/dev/sda",
		},
	}

	profiles["libvirt"] = Profile{
		Name: "libvirt",
		RootDeviceHints: metal3api.RootDeviceHints{
			DeviceName: "/dev/vda",
		},
	}

	profiles["dell"] = Profile{
		Name: "dell",
		RootDeviceHints: metal3api.RootDeviceHints{
			HCTL: "0:0:0:0",
		},
	}

	profiles["dell-raid"] = Profile{
		Name: "dell-raid",
		RootDeviceHints: metal3api.RootDeviceHints{
			HCTL: "0:2:0:0",
		},
	}

	profiles["openstack"] = Profile{
		Name: "openstack",
		RootDeviceHints: metal3api.RootDeviceHints{
			DeviceName: "/dev/vdb",
		},
	}

	profiles[EmptyProfileName] = Profile{
		Name: EmptyProfileName,
	}
}

// GetProfile returns the named profile.
func GetProfile(name string) (Profile, error) {
	profile, ok := profiles[name]
	if !ok {
		return Profile{}, fmt.Errorf("no hardware profile named %q", name)
	}
	return profile, nil
}
