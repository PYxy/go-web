package option

// EtcdOptions defines options for etcd cluster.
type EtcdOptions struct {
	Endpoints            []string `json:"endpoints"               mapstructure:"endpoints"`
	Timeout              int      `json:"timeout"                 mapstructure:"timeout"`
	RequestTimeout       int      `json:"request-timeout"         mapstructure:"request-timeout"`
	LeaseExpire          int      `json:"lease-expire"            mapstructure:"lease-expire"`
	Username             string   `json:"username"                mapstructure:"username"`
	Password             string   `json:"password"                mapstructure:"password"`
	UseTLS               bool     `json:"use-tls"                 mapstructure:"use-tls"`
	CaCert               string   `json:"ca-cert"                 mapstructure:"ca-cert"`
	Cert                 string   `json:"cert"                    mapstructure:"cert"`
	Key                  string   `json:"key"                     mapstructure:"key"`
	HealthBeatPathPrefix string   `json:"health_beat_path_prefix" mapstructure:"health_beat_path_prefix"`
	HealthBeatIFaceName  string   `json:"health_beat_iface_name"  mapstructure:"health_beat_iface_name"`
	Namespace            string   `json:"namespace"               mapstructure:"namespace"`
}

// NewEtcdOptions create a `zero` value instance.
func NewEtcdOptions() *EtcdOptions {
	return &EtcdOptions{
		Timeout:        5,
		RequestTimeout: 2,
		LeaseExpire:    5,
	}
}
