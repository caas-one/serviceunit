package nginx

import (
	"os"
	"testing"
	"text/template"
)

var nginx = Nginx{
	Upstreams: []Upstream{
		Upstream{
			Name: "default-ldap",
			UpstreamServers: []UpstreamServer{
				UpstreamServer{
					Address: "10.151.160.80:26379",
				},
				UpstreamServer{
					Address: "10.151.160.81:26379",
				},
				UpstreamServer{
					Address: "10.151.160.83:26379",
				},
			},
		},
		Upstream{
			Name: "default-redisoperator",
			UpstreamServers: []UpstreamServer{
				UpstreamServer{
					Address: "10.151.160.80:32379",
				},
				UpstreamServer{
					Address: "10.151.160.81:32379",
				},
				UpstreamServer{
					Address: "10.151.160.83:32379",
				},
			},
		},
		Upstream{
			Name: "default-rfs-redisfailover",
			UpstreamServers: []UpstreamServer{
				UpstreamServer{
					Address: "10.151.160.80:30379",
				},
				UpstreamServer{
					Address: "10.151.160.81:30379",
				},
				UpstreamServer{
					Address: "10.151.160.83:30379",
				},
			},
		},
		Upstream{
			Name: "yce-crasher",
			UpstreamServers: []UpstreamServer{
				UpstreamServer{
					Address: "10.151.160.84:32308",
				},
				UpstreamServer{
					Address: "10.151.160.85:32308",
				},
			},
		},
		Upstream{
			Name: "yce-nginx-apiserver",
			UpstreamServers: []UpstreamServer{
				UpstreamServer{
					Address: "10.151.160.84:32319",
				},
				UpstreamServer{
					Address: "10.151.160.85:32319",
				},
			},
		},
		Upstream{
			Name: "helm-test-yce-backend",
			UpstreamServers: []UpstreamServer{
				UpstreamServer{
					Address: "10.151.160.84:32320",
				},
				UpstreamServer{
					Address: "10.151.160.87:32320",
				},
			},
		},
	},
	Servers: []Server{
		Server{
			Namespace: "default",
			Locations: []Location{
				Location{
					Path:         "ldap",
					UpstreamName: "default-ldap",
				},
				Location{
					Path:         "redisoperator",
					UpstreamName: "default-redisoperator",
				},
				Location{
					Path:         "rfs-redisfailover",
					UpstreamName: "default-rfs-redisfailover",
				},
			},
		},
		Server{
			Namespace: "yce",
			Locations: []Location{
				Location{
					Path:         "crasher",
					UpstreamName: "yce-crasher",
				},
				Location{
					Path:         "nginx-apiserver",
					UpstreamName: "yce-nginx-apiserver",
				},
			},
		},
		Server{
			Namespace: "helm-test",
			Locations: []Location{
				Location{
					Path:         "yce-backend",
					UpstreamName: "helm-test-yce-backend",
				},
			},
		},
	},
}

func Test_NginxRender(t *testing.T) {
	temp := template.Must(template.ParseFiles("../nginx.tmpl"))
	err := temp.Execute(os.Stdout, nginx)
	if err != nil {
		t.Errorf("template Execute error: err=%s\n", err)
		return
	}
}
