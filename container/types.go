package container

// The container spec
// VolumeMounts is not included for the security concern.
//
// The reason that we don't want to import engine-api is that:
// We don't want user to specify an option that will not be enabled or supported.
// import "github.com/docker/engine-api/types/container"
type Config struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Args    []string `json:"args"`
	Command []string `json:"command"`
	Env     []EnvVar `json:"env"`
	Ports   []Port   `json:"ports"`
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
