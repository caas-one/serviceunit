package nginx

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"time"

	mlog "github.com/maxwell92/log"
	"gitlab.yeepay.com/yce/nodeport/k8s/node"
	"k8s.io/client-go/kubernetes"
)

var log = mlog.Log

var NginxTemplate = `{{- range $upstream := .Upstreams}}
upstream {{$upstream.Name}} {
	{{- range $server := $upstream.UpstreamServers}}
	server {{$server.Address}};
    {{- end }}
}
{{- end }}

{{- range $server := .Servers}}
server {
	listen 80;
	server_name {{$server.Namespace}}.k8s.yp;
	{{- range $location := $server.Locations}}
	location /{{$location.Path}} {
		proxy_redirect off;
		proxy_set_header Host $host;
		proxy_pass http://{{$location.UpstreamName}};
	}
	{{- end }}
}
{{- end }}`

// Upstream defines the upstream of nginx
type Upstream struct {
	Name            string           `json:"name"`
	UpstreamServers []UpstreamServer `json:"server"`
}

// NewUpstream create a new a empty upstream instance
func NewUpstream() *Upstream {
	return &Upstream{
		UpstreamServers: make([]UpstreamServer, 0),
	}
}

// NewUpstreamReal create a real upstream instance
func NewUpstreamReal(name string, list node.List, nodeport string) *Upstream {
	servers := make([]UpstreamServer, 0)
	for _, node := range list {
		addr := node.IP + ":" + nodeport
		server := NewUpstreamServer(addr)
		servers = append(servers, *server)
	}
	return &Upstream{
		Name:            name,
		UpstreamServers: servers,
	}
}

// UpstreamServer defines the server block of upstream
type UpstreamServer struct {
	Address     string `json:"address"`
	Weight      string `json:"weight"`
	MaxConns    string `json:"maxconn"`
	MaxFails    string `json:"maxFails"`
	FailTimeout string `json:"failTimeout"`
}

// NewUpstreamServer create a new Server instance
func NewUpstreamServer(addr string) *UpstreamServer {
	return &UpstreamServer{
		Address: addr,
		// Weight: Weight,
		// MaxConns: MaxConns,
		// MaxFails: MaxFails,
		// FailTimeout: FailTimeout,
	}
}

// Server define the server block in nginx config
type Server struct {
	Namespace string     `json:"namespace"`
	Locations []Location `json:"locations"`
}

// NewServer create a new&empty Instance
func NewServer() *Server {
	return &Server{
		Locations: make([]Location, 0),
	}
}

// Location routes the path(namespace-deployments) to the appropriate upstream
type Location struct {
	Path         string `json:"path"`     // equal to the content-path/appname/deployment
	UpstreamName string `json:"upstream"` // namespace-deployment
}

// NewLocation create a new&empty Location instance
func NewLocation() *Location {
	return &Location{}
}

// Nginx used to rander the template
type Nginx struct {
	Upstreams []Upstream `json:"upstreams"`
	Servers   []Server   `json:"servers"`
}

// NewNginx return a new&empty Nginx instance
func NewNginx() *Nginx {
	return &Nginx{
		Upstreams: make([]Upstream, 0),
		Servers:   make([]Server, 0),
	}
}

// NginxController
type NginxController struct {
	Lazy   bool
	Chan   chan struct{}
	Client *kubernetes.Clientset
}

// NewNginxController func create a new&empty NginxController instance
func NewNginxController(client *kubernetes.Clientset, ch chan struct{}) *NginxController {
	return &NginxController{
		Lazy:   false,
		Chan:   ch,
		Client: client,
	}
}

// Run func is a controll-loop
func (nc *NginxController) Run() {
	defer func() { recover() }()
	log.Infof("Nginx controller start.....")
	for {
		select {
		case <-nc.Chan:
			nc.Lazy = true
		case <-time.After(time.Second * 2):
			if nc.Lazy {
				nc.Lazy = false
				// render conf file
				nc.Render()
				// nginx reload
				nc.Reload()
			}
		}
	}
}

// Reload func reload nginx process
func (nc *NginxController) Reload() error {
	if err := nc.shellOut("nginx -t"); err != nil {
		return fmt.Errorf("Invalid nginx configuration detected, not reloading: %s", err)
	}
	if err := nc.shellOut("nginx -s reload"); err != nil {
		return fmt.Errorf("Reloading NGINX failed: %s", err)
	}
	return nil
}

// Render func to render the nginx template
func (nc *NginxController) Render() {
	// render
	nginx := transform(nc.Client)
	err := nc.render(nginx, NginxConfigFile)
	if err != nil {
		log.Errorf("NginxController render error: err=%s", err)
	}
}

func (nc *NginxController) render(nginx *Nginx, file string) error {
	f, err := os.Create(file)
	defer f.Close()

	if err != nil {
		log.Fatalf("os.Create file error: file=%s, err=%s", file, err)
	}

	writer := bufio.NewWriter(f)
	temp := template.Must(template.New("nginx").Parse(NginxTemplate))
	err = temp.Execute(writer, nginx)
	if err != nil {
		log.Errorf("template Execute error: err=%s\n", err)
		return err
	}
	log.Infof("Go template render ok ....")
	writer.Flush()
	return nil
}

func (nc *NginxController) shellOut(cmd string) (err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	log.Infof("executing %s", cmd)

	command := exec.Command("sh", "-c", cmd)
	command.Stdout = &stdout
	command.Stderr = &stderr

	err = command.Start()
	if err != nil {
		return fmt.Errorf("Failed to execute %v, err: %v", cmd, err)
	}

	err = command.Wait()
	if err != nil {
		return fmt.Errorf("Command %v stdout: %q\nstderr: %q\nfinished with error: %v", cmd,
			stdout.String(), stderr.String(), err)
	}
	return nil
}
