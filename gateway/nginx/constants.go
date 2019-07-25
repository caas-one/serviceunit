package nginx

const (
	// Weight sets the weight of the server, by default, 1.
	Weight = "1"

	// MaxConns in nginx upstream: limits the maximum number of simultaneous active connections to the proxied server
	MaxConns = "200"

	// MaxFails in nginx upstream: sets the number of unsuccessful attempts to communicate with the server that should happen in the duration set by the fail_timeout parameter to consider the server unavailable for a duration also set by the fail_timeout parameter.
	MaxFails = "2"

	// FailTimeout in nginx upstream:sets the time during which the specified number of unsuccessful attempts to communicate with the server should happen to consider the server unavailable;
	FailTimeout = "10s"

	// NginxConfigFile
	NginxConfigFile = "/etc/nginx/conf.d/ingress.conf"
	// NginxConfigFile = "ingress.conf"
)
