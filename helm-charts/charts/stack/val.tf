locals {
  ym = <<EOF
stackName: along3
deploymentStage: rdev
aws:
  region: us-west-2
  tags:
    env: "ADDTAGS"
    owner: "ADDTAGS"
    project: "ADDTAGS"
    service: "ADDTAGS"
    managedBy: "ADDTAGS"
  cloudEnv:
    publicSubnets: ["subnet-a1", "subnet-a2"]
    privateSubnets: []
    databaseSubnets: []
    databaseSubnetGroup: ""
    vpcId: ""
    vpcCidrBlock: ""
  wafAclArn: ""
  dnsZone: edu-platform.rdev.si.czi.technology
serviceMesh:
  enabled: true
datadog:
  enabled: true
services:
- name: "service2"
  image:
    repository: "blalbhal"
    tag: "tag1"
    tagMutability: true
    scanOnPush: false
    platformArchitecture: "amd64"
    pullPolicy: "IfNotPresent"
  cmd: []
  args: []
  resources:
    limits:
      cpu: "100m"
      memory: "100Mi"
    requests:
      cpu: "10m"
      memory: "10Mi"
  scaling:
    desiredCount: 2
    maxCount: 2
    cpuThresholdPercentage: 80
  env:
    additionalEnvVars: []
    # additionalEnvVars: [{name: "balh", value: "blah"}]
    #
    #
    additionalEnvVarsFromConfigMaps: []
    #  additionalEnvVarsFromConfigMaps: [
    # {configMapRef: {name: "balh"}, prefix: "blah"}
    #]
    additionalEnvVarsFromSecrets: []
    #  additionalEnvVarsFromSecrets: [
    # {secretRef: {name: "balh"}, prefix: "blah"}
    #]
  volumes:
    additionalVolumesFromSecrets: [{mountPath: "blah2", readOnly: true, name: "blah2"}]
    #  additionalEnvVarsFromSecrets: [
    # {mountPath: "blah", readOnly: true, name: "blah"}
    # - name: config-vol
    configMap:
      name: log-config
      items:
      - key: log_level
        path: log_level
    #]
    additionalVolumesFromConfigMaps: [{mountPath: "blah", readOnly: true, name: "blah"}]
    #  additionalEnvVarsFromSecrets: [
    # {mountPath: "blah", readOnly: true, name: "blah"}
    #]
  stackPrefix: ""
  datadog:
    createDashboard: false
  skipConfigInjection: false
  waitForSteadyState: true
  certificateArn: "blahblahbs"
  serviceEndpoints: {}
  healthCheck:
    path: "/"
    periodSeconds: 3
    initialDelaySeconds: 30
  awsIam:
    roleArn: arn:aws:iam::00000000000:role/zzz/zzz
  routing:
    method: "DOMAIN"
    hostMatch: ""
    groupName: ""
    alb:
      loadBalancerAttributes:
      - idle_timeout.timeout_seconds=60
      securityGroup: sg-123
      targetGroupArn: arn:aws:elasticloadbalancing:us-west-2:00000000000:targetgroup/zzz/zzz
      targetGroup: group1
    priority: 4
    path: "/*"
    serviceName: ""
    port: 3000
    scheme: "HTTP"
    successCodes: "200-499"
    serviceType: "EXTERNAL"
    oidcConfig:
      issuer: ""
      authorizationEndpoint: ""
      tokenEndpoint: ""
      userInfoEndpoint: ""
      secretName: ""
    bypasses:
    - field: http-request-method
      httpRequestMethodConfig:
        Values:
        - GET
        - OPTIONS
    - field: path-pattern
      pathPatternConfig:
        Values:
        - /blah
        - /test/skip

  # sidecars:
  #   sidecar1:
  #     image: "sidecar-image-1"
  #     tag: "1.0.0"
  #     port: 8080
  #     scheme: "HTTP"
  #     memory: "256Mi"
  #     cpu: "250m"
  #     imagePullPolicy: "IfNotPresent"
  #     healthCheckPath: "/health"
  #     initialDelaySeconds: 15
  #     periodSeconds: 5
  # sidecars: {}
  sidecars:
  - name: sidecar1
    image:
      repository: "blalbhal"
      tag: "tag1"
    routing:
      port: 8080
      scheme: "HTTP"
    resources:
      limits:
        cpu: "100m"
        memory: "100Mi"
      requests:
        cpu: "10m"
        memory: "10Mi"
    imagePullPolicy: "IfNotPresent"
    healthCheck:
      path: "/health"
      periodSeconds: 3
      initialDelaySeconds: 30
    initialDelaySeconds: 15
    periodSeconds: 5
  regionalWafv2Arn: null
  additionalNodeSelectors: {}
  additionalPodLabels: {}
  serviceMesh:
    allowServices:
    - service: "service1"
      stack: "stack1"
      serviceAccountName: "sa1"
tasks:
- name: migrate
  suspend: true
  schedule: "0 0 1 1 *"
  cmd: ["./manage.py", "migrate"]
  resources:
    limits:
      cpu: "100m"
      memory: "100Mi"
    requests:
      cpu: "10m"
      memory: "10Mi"
  image:
    repository: "blalbhal"
    tag: "tag1"
    platformArchitecture: "amd64"
    pullPolicy: "IfNotPresent"
  awsIam:
    roleArn: arn:aws:iam::00000000000:role/zzz/zzz
  env:
    additionalEnvVars: []
    # additionalEnvVars: [{name: "balh", value: "blah"}]
    #
    #
    additionalEnvVarsFromConfigMaps: []
    #  additionalEnvVarsFromConfigMaps: [
    # {configMapRef: {name: "balh"}, prefix: "blah"}
    #]
    additionalEnvVarsFromSecrets: []
    #  additionalEnvVarsFromSecrets: [
    # {secretRef: {name: "balh"}, prefix: "blah"}
    #]
  volumes:
    additionalVolumesFromSecrets: [{mountPath: "blah2", readOnly: true, name: "blah2"}]
    #  additionalEnvVarsFromSecrets: [
    # {mountPath: "blah", readOnly: true, name: "blah"}
    # - name: config-vol
    configMap:
      name: log-config
      items:
      - key: log_level
        path: log_level
    #]
    additionalVolumesFromConfigMaps: [{mountPath: "blah", readOnly: true, name: "blah"}]
    #  additionalEnvVarsFromSecrets: [
    # {mountPath: "blah", readOnly: true, name: "blah"}
    #]
  additionalNodeSelectors: {}
  additionalPodLabels: {}
  EOF
  // Convert to hcl
    hcl = yamldecode(local.ym)
}