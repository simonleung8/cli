package plugin

type Apps struct {
	Name string
}

type Resource interface {
	GetApps() []Apps
	SetApp(string)
}

type resource struct {
	apps []Apps
}

func NewResource() Resource {
	return &resource{}
}

func (r *resource) SetApp(appName string) {
	r.apps = append(r.apps, Apps{appName})
}

func (r *resource) GetApps() []Apps {
	return r.apps
}
