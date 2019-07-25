package tools

const (
	// ChartsTempl template for generate Chart.yaml
	ChartsTempl = `apiVersion: v1
appVersion: "1.0"
description: A Helm chart for Kubernetes
name: {{.Name}}
version: 0.1.0`

	// DeploymentTempl template for generate template/deployment.yaml
	/*
			DeploymentTempl = `apiVersion: apps/v1
		kind: Deployment
		metadata:
		  name: {{ .Release.Name }}
		  namespace: {{ .Release.Namespace }}
		  labels:
		    app.kubernetes.io/name: {{ .Release.Name }}
		    helm.sh/chart: {{ .Release.Name }}
		    app.kubernetes.io/instance: {{ .Release.Name }}
		    app.kubernetes.io/managed-by: {{ .Release.Service }}
		spec:
		  replicas: {{ .Values.replicaCount }}
		  selector:
		    matchLabels:
		      app.kubernetes.io/name: {{ .Release.Name }}
		      app.kubernetes.io/instance: {{ .Release.Name }}
		  template:
		    metadata:
		      labels:
		        app.kubernetes.io/name: {{ .Release.Name }}
		        app.kubernetes.io/instance: {{ .Release.Name }}
		    spec:
		      containers:
		        - name: {{ .Chart.Name }}
		          image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
		          imagePullPolicy: {{ .Values.image.pullPolicy }}
		          ports:
		            - name: http
		              containerPort: 8080
		              protocol: TCP

		          {{- if .Values.env}}
		          env:
		          {{range $e := .Values.env}}
		          - name: "{{ $e.name }}"
		            value: "{{ $e.value }}"
		          {{- end}}
		          {{- end}}
		          resources:
		            {{- toYaml .Values.resources | nindent 12 }}
		      {{- with .Values.nodeSelector }}
		      nodeSelector:
		        {{- toYaml . | nindent 8 }}
		      {{- end }}
		    {{- with .Values.affinity }}
		      affinity:
		        {{- toYaml . | nindent 8 }}
		    {{- end }}
		    {{- with .Values.tolerations }}
		      tolerations:
		        {{- toYaml . | nindent 8 }}
		    {{- end }}
		    {{- if .Values.imagePullSecrets}}
		      imagePullSecrets:
		      - name: {{ .Values.imagePullSecrets }}
		    {{- end }}
		`
	*/
	DeploymentTempl = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Values.name }}
  namespace: {{ .Values.namespace}}
  labels:
    app.kubernetes.io/name: {{ .Values.name }}
    helm.sh/chart: {{ .Values.name }}
    app.kubernetes.io/instance: {{ .Values.name }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app.kubernetes.io/name: {{ .Values.name }}
      app.kubernetes.io/instance: {{ .Values.name }}
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      labels:
        app.kubernetes.io/name: {{ .Values.name }}
        app.kubernetes.io/instance: {{ .Values.name }}
    spec:
      containers:
      - name: {{ .Chart.Name }}
        image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
        imagePullPolicy: {{ .Values.image.pullPolicy }}
        ports:
        - name: http
          containerPort: 8080
          protocol: TCP
        {{- if .Values.env}}
        env:
        {{- range $e := .Values.env}}
        - name: {{ $e.name }}
          value: {{ $e.value | quote }}
        {{- end}}
        {{- end}}
        resources:
          limits:
            cpu: {{ .Values.resources.limits.cpu }}
            memory: {{ .Values.resources.limits.memory }}
          requests:
            cpu: {{ .Values.resources.requests.cpu }}
            memory: {{ .Values.resources.requests.memory }}
        lifecycle:
          preStop:
            exec: 
              command:
              - /gracefulstop.sh
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /{{.Values.name}}/metrics/healthcheck
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 300
          periodSeconds: 120
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /{{.Values.name}}/metrics/healthcheck
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        volumeMounts:
        - mountPath: /apps/configmap/{{.Values.name}}.conf
          name: configs
          subPath: {{.Values.name}}.conf
        - mountPath: /apps/configmap/url
          name: configs
          subPath: url
      - name: debug
        image: "{{.Values.DebugImage.repository}}:{{.Values.DebugImage.tag}}"
        imagePullPolicy: {{.Values.DebugImage.pullPolicy}}
        ports: 
        - name: http
          containerPort: 8090
          protocol: TCP
        resources:
          limits:
            cpu: 200m
            memory: 200M
          requests:
            cpu: 50m
            memory: 50M
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- if .Values.imagePullSecrets}}
      imagePullSecrets:
      - name: {{ .Values.imagePullSecrets }}
      {{- end }}
      imagePullSecrets:
      - name: myregistrykey
      restartPolicy: Always
      securityContext: {}
      terminationGracePeriodSeconds: 30
      volumes:
      - configMap:
          defaultMode: 511
          items:
          - key: {{.Values.name}}
            path: {{.Values.name}}.conf
          - key: url
            path: url
          name: {{.Values.name}}
        name: configs`

	// ServiceTempl template for generate template/service.yaml
	ServiceTempl = `apiVersion: v1
kind: Service
metadata:
  name: {{ .Values.name }}-svc
  namespace: {{ .Values.namespace}}
  labels:
    app.kubernetes.io/name: {{ .Values.name }}
    helm.sh/chart: {{ .Values.name }}
    app.kubernetes.io/instance: {{ .Values.name }}
    app.kubernetes.io/managed-by: {{ .Values.name }}
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: {{ .Values.service.port }}
      nodePort: {{ .Values.service.nodePort }}
      protocol: TCP
      name: http
  selector:
    app.kubernetes.io/name: {{ .Values.name }}
    app.kubernetes.io/instance: {{ .Values.name }}`

	// IngressTempl template for generate template/ingress.yaml
	IngressTempl = `{{- if .Values.ingress.enabled -}}
{{- $fullName := "{{" .Values.name "}}" -}}
{{- $ingressPaths := .Values.ingress.paths -}}
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: {{ $fullName }}
  labels:
    app.kubernetes.io/name: {{ .Values.name }}
    helm.sh/chart: {{ .Values.name }}
    app.kubernetes.io/instance: {{ .Values.name }}
    app.kubernetes.io/managed-by: {{ .Values.Service }}
  {{- with .Values.ingress.annotations }}
  annotations:
    {{- toYaml . | nindent 4 }}
  {{- end }}
spec:
{{- if .Values.ingress.tls }}
  tls:
  {{- range .Values.ingress.tls }}
    - hosts:
      {{- range .hosts }}
        - {{ . | quote }}
      {{- end }}
      secretName: {{ .secretName }}
  {{- end }}
{{- end }}
  rules:
  {{- range .Values.ingress.hosts }}
    - host: {{ . | quote }}
      http:
        paths:
	{{- range $ingressPaths }}
          - path: {{ . }}
            backend:
              serviceName: {{ $fullName }}
              servicePort: http
	{{- end }}
  {{- end }}
{{- end }}`

	// NotesTxt
	NotesTxt = `1. Get the application URL by running these commands:
{{- if .Values.ingress.enabled }}
{{- range $host := .Values.ingress.hosts }}
  {{- range $.Values.ingress.paths }}
  http{{ if $.Values.ingress.tls }}s{{ end }}://{{ $host }}{{ . }}
  {{- end }}
{{- end }}
{{- else if contains "NodePort" .Values.service.type }}
  export NODE_PORT=$(kubectl get --namespace {{ .Values.namespace }} -o jsonpath="{.spec.ports[0].nodePort}" services {{ .Values.name }})
  export NODE_IP=$(kubectl get nodes --namespace {{ .Values.namespace }} -o jsonpath="{.items[0].status.addresses[0].address}")
  echo http://$NODE_IP:$NODE_PORT
{{- else if contains "LoadBalancer" .Values.service.type }}
     NOTE: It may take a few minutes for the LoadBalancer IP to be available.
           You can watch the status of by running 'kubectl get svc -w {{ .Values.name }}'
  export SERVICE_IP=$(kubectl get svc --namespace {{ .Values.namespace }} {{ .Values.name }} -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
  echo http://$SERVICE_IP:{{ .Values.service.port }}
{{- else if contains "ClusterIP" .Values.service.type }}
  export POD_NAME=$(kubectl get pods --namespace {{ .Values.namespace }} -l "app.kubernetes.io/name={{ .Values.name }},app.kubernetes.io/instance={{ .Values.name }}" -o jsonpath="{.items[0].metadata.name}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl port-forward $POD_NAME 8080:80
{{- end }}`

	// HelperTempl
	HelperTempl = `{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "yce.name" -}}
{{- default .Chart.name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "yce.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Values.name -}}
{{- .Values.name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Values.name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "yce.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}`

	ValuesTempl = `
namespace: {{.Namespace}}
name: {{.AppName}}
replicaCount: {{.ReplicaCount}}

image:
  repository: {{.Image.Repository}}
  tag: "{{ .Image.Tag }}"
  pullPolicy: {{.Image.PullPolicy}}

DebugImage:
  repository: {{.DebugImage.Repository}}
  tag: "{{ .DebugImage.Tag }}"
  pullPolicy: {{.DebugImage.PullPolicy}}

imagePullSecrets: {{.ImagePullSecrets}}
nameOverride: {{.AppName}}
fullnameOverride: ""

service:
  type: {{.Service.Type}}
  port: {{.Service.Port}}
  {{- if .Service.NodePort}}
  nodePort: {{.Service.NodePort}}
  {{- end }}

env:
{{range $env := .Env}}
- name: {{$env.Name}}
  value: {{$env.Value}}
{{end -}}

ingress:
  enabled: false
  annotations: {}
  paths: []
  hosts:
    - chart-example.local
  tls: []

resources:
   limits:
    cpu: {{.Resources.Limits.CPU}}
    memory: {{.Resources.Limits.Memory}}
   requests:
    cpu: {{.Resources.Requests.CPU}}
    memory: {{.Resources.Requests.Memory}}
nodeSelector: {}

tolerations: []

affinity: {}
`
)
