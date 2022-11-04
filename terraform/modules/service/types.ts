export type ECSComputeLimit = { cpu: ".25 vcpu", mem: "512" } |
{ cpu: ".25 vcpu", mem: "1 GB" } |
{ cpu: ".25 vcpu", mem: "2 GB" } |
{ cpu: ".5 vcpu", mem: "1 GB" } |
{ cpu: ".5 vcpu", mem: "2 GB" } |
{ cpu: ".5 vcpu", mem: "3 GB" } |
{ cpu: ".5 vcpu", mem: "4 GB" } |
{ cpu: "1 vcpu", mem: "2 GB" } |
{ cpu: "1 vcpu", mem: "3 GB" } |
{ cpu: "1 vcpu", mem: "4 GB" } |
{ cpu: "1 vcpu", mem: "5 GB" } |
{ cpu: "1 vcpu", mem: "6 GB" } |
{ cpu: "1 vcpu", mem: "7 GB" } |
{ cpu: "1 vcpu", mem: "8 GB" } |
{ cpu: "2 vcpu", mem: "4 GB" } |
{ cpu: "2 vcpu", mem: "5 GB" } |
{ cpu: "2 vcpu", mem: "6 GB" } |
{ cpu: "2 vcpu", mem: "7 GB" } |
{ cpu: "2 vcpu", mem: "8 GB" } |
{ cpu: "2 vcpu", mem: "9 GB" } |
{ cpu: "2 vcpu", mem: "10 GB" } |
{ cpu: "2 vcpu", mem: "11 GB" } |
{ cpu: "2 vcpu", mem: "12 GB" } |
{ cpu: "2 vcpu", mem: "13 GB" } |
{ cpu: "2 vcpu", mem: "14 GB" } |
{ cpu: "2 vcpu", mem: "15 GB" } |
{ cpu: "2 vcpu", mem: "16 GB" } |
{ cpu: "1 vcpu", mem: "4 GB" }

export type ServiceType = "PRIVATE" | "INTERNAL" | "EXTERNAL"

export type Environment = "rdev" | "dev" | "staging" | "prod"

export type AWSRegion = "us-east-1" | "us-east-2" | "us-west-1" | "us-west-2"

export interface ECSServiceDefinition {
    name: string
    // the number of replicas
    desiredCount: number
    // the port to expose
    port: number
    // the image to use in the container
    image: string
    // deployment size (only valid combinations allowed)
    computeLimits: ECSComputeLimit
    // if the service is on the internet, protected by Okta, or only exposed within the cluster
    serviceType: ServiceType
    healthCheckPath?: string
    environment?: EnvironmentVariables
}

export interface EnvironmentVariables { name: string, value: string }[]

export interface AWSLogConfig {
    logDriver: "awslogs",
    options: {
        "awslogs-stream-prefix": string,
        "awslogs-group": string,
        "awslogs-region": AWSRegion,
    }
}

export interface ContainerDefinition {
    name: string,
    image: string,
    memory: string,
    portMappings: {
        containerPort: number,
        hostPort?: number,
        protocol?: string
    }[]
    logConfiguration?: AWSLogConfig
    essential?: boolean,
    environment?: EnvironmentVariables
    command?: string
}

export interface HappyServiceMeta {
    env: Environment
    stackName: string
    region: AWSRegion
    serviceDef: ECSServiceDefinition
}
