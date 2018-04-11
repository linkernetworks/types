package container

import (
	"gopkg.in/mgo.v2/bson"

	v1 "k8s.io/api/core/v1"
)

// The container spec
// VolumeMounts is not included for the security concern.
//
// The reason that we don't want to import engine-api is that:
// We don't want user to specify an option that will not be enabled or supported.
// import "github.com/docker/engine-api/types/container"
type Config struct {
	Name            string        `json:"name"`
	Image           string        `json:"image"`
	ImagePullPolicy v1.PullPolicy `json:"imagePullPolicy"`
	Args            []string      `json:"args"`
	Command         []string      `json:"command"`
	WorkingDir      string        `json:"workingDir"`
	Env             []EnvVar      `json:"env"`
	Ports           []Port        `json:"ports"`
	VolumeMounts    []VolumeMount `json:"volumeMounts"`
	Privileged      bool          `json:"securityContext"`

	ResourceRequirements v1.ResourceRequirements `json:"resourceRequirements"`

	// the port name to be exposed to the external service
	// for example, we use "notebook" for notebooks
	ExposePortName string `json:"exposePortName"`

	NodeSelector map[string]string `json:"nodeSelector"`
}

func (c *Config) Copy() Config {
	var n = *c
	copy(n.Args, c.Args)
	copy(n.Command, c.Command)
	copy(n.Env, c.Env)
	copy(n.Ports, c.Ports)
	copy(n.VolumeMounts, c.VolumeMounts)
	return n
}

func (c *Config) GetKubernetesContainerPorts() (containerPorts []v1.ContainerPort) {
	for _, port := range c.Ports {
		containerPorts = append(containerPorts, v1.ContainerPort{
			Name:          port.Name,
			ContainerPort: port.ContainerPort,

			// TODO: pull out this option when we have udp protocol needed.
			Protocol: v1.ProtocolTCP,
		})
	}
	return containerPorts
}

// GetKubernetesVolumeMounts converts the container volume mount definition to
// the kubernetes volume mount definition.
func (c *Config) GetKubernetesVolumeMounts() (volumeMounts []v1.VolumeMount) {
	for _, mount := range c.VolumeMounts {
		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      mount.Name,
			SubPath:   mount.SubPath,
			MountPath: mount.MountPath,
		})
	}
	return volumeMounts
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Port struct {
	ContainerPort int32  `json:"containerPort"`
	HostPort      int32  `json:"hostPort"`
	Name          string `json:"name"`

	// TCP or UDP
	Protocol string `json:"protocol"`
}

type VolumeMount struct {
	Name string `bson:"name" json:"name"`

	// path to mount
	MountPath string `bson:"mountPath" json:"mountPath"`

	// subpath to limit the volume access (optional)
	SubPath string `bson:"subPath" json:"subPath"`
}

//FIXME we should containet the Volume and VolumeMount, and the Volume should contianas the VolumeSource
type Volume struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"id"`

	// persistent volume claim for mount
	ClaimName string `bson:"claimName" json:"claimName"`

	// where to mount
	VolumeMount VolumeMount `bson:"volumeMount" json:"volumeMount"`
}
