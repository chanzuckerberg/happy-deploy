{{/*
Expand the name of the chart.
*/}}
{{- define "stack.name" -}}
{{- default .Values.stackName | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "service.name" -}}
{{- default .name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "stack.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "stack.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "stack.labels" -}}
helm.sh/chart: {{ include "stack.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: happy
app.kubernetes.io/part-of: {{ include "stack.name" . }}
stack: {{ include "stack.name" . }}
{{- end }}

/* service labels */
{{- define "service.labels" -}}
app: {{.}}
app.kubernetes.io/component: {{.}}
{{- end}}

/* service labels */
{{- define "happy.intSecretVolumeMount" -}}
anchor:
- mountPath: /var/happy
  name: integration-secret
  readOnly: true
{{- end}}

{{- define "ingress.base.annotations" -}}
kubernetes.io/ingress.class: alb
alb.ingress.kubernetes.io/healthcheck-interval-seconds: "300"
alb.ingress.kubernetes.io/healthcheck-path: {{ .service.healthCheck.path | quote }}
alb.ingress.kubernetes.io/listen-ports: '[{"HTTPS":443},{"HTTP":80}]'
alb.ingress.kubernetes.io/actions.redirect: '{"RedirectConfig":{"Port":"443","Protocol":"HTTPS","StatusCode":"HTTP_301"},"Type":"redirect"}'
alb.ingress.kubernetes.io/scheme: internet-facing
alb.ingress.kubernetes.io/success-codes: {{ .service.routing.successCodes | quote }}
alb.ingress.kubernetes.io/target-group-attributes: deregistration_delay.timeout_seconds=60
alb.ingress.kubernetes.io/target-type: instance
alb.ingress.kubernetes.io/group.name: stack-along3-electric-dragon # TODO
alb.ingress.kubernetes.io/group.order: "1" # TODO
alb.ingress.kubernetes.io/certificate-arn: {{ .service.routing.certificateArn }}
alb.ingress.kubernetes.io/ssl-policy: ELBSecurityPolicy-TLS-1-2-2017-01    
alb.ingress.kubernetes.io/load-balancer-attributes: {{ join "," .service.routing.loadBalancerAttributes | quote }}
alb.ingress.kubernetes.io/healthcheck-protocol: {{if eq .Values.serviceMesh.enabled true}}HTTPS{{else}}{{ .service.routing.scheme }}{{end}}
alb.ingress.kubernetes.io/backend-protocol: {{if eq .Values.serviceMesh.enabled true}}HTTPS{{else}}{{ .service.routing.scheme }}{{end}}
alb.ingress.kubernetes.io/subnets: {{ required ".Values.aws.cloudEnv.publicSubnets is required" (join "," .Values.aws.cloudEnv.publicSubnets) | quote }}
alb.ingress.kubernetes.io/tags: env={{.Values.deploymentStage}},happy_env={{.Values.deploymentStage}},happy_last_applied={{ now | date "20060102150405" }},happy_region={{ .Values.aws.region }},happy_stack_name={{ include "stack.name" . }},managedBy=happy,owner={{ .Values.aws.tags.owner }},project={{ .Values.aws.tags.project }},service={{ .Values.aws.tags.service }}
{{- end}}

{{/*
Create the name of the service account to use
*/}}
{{- define "stack.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "stack.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
