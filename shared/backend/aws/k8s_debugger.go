package aws

type K8sDebugSignal struct {
	Kind             string
	Reason           string
	MessageSignature string
	Description      string
	Remediation      string
	RunbookUrl       string
}

var K8sDebugSignals = []K8sDebugSignal{
	{
		Kind:             "Pod",
		Reason:           "Unhealthy",
		MessageSignature: "Readiness probe failed",
		Description:      "readiness probe is failing (server did not respond with a success code, there was a timeout etc)",
		Remediation:      "Make sure the application is healthy and responds to HTTP requests, there's no port mismatch, no timeouts, and http code returned is withing 200-399 range",
		RunbookUrl:       "https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/",
	},
	{
		Kind:             "Pod",
		Reason:           "Unhealthy",
		MessageSignature: "Liveness probe failed",
		Description:      "liveness probe is failing (server did not respond with a success code, there was a timeout etc)",
		Remediation:      "Make sure the application is healthy and responds to HTTP requests, there's no port mismatch, no timeouts, and http code returned is withing 200-399 range",
		RunbookUrl:       "https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/",
	},
	{
		Kind:             "Pod",
		Reason:           "BackOff",
		MessageSignature: "",
		Description:      "the application has repeatedly failed to start, it most likely crashes on start or doesn't pass the health checks",
		Remediation:      "Run 'happy logs <STACK_NAME> <SERVICE_NAME>' to determine the root cause",
		RunbookUrl:       "https://containersolutions.github.io/runbooks/posts/kubernetes/crashloopbackoff/",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedCreatePodSandBox",
		MessageSignature: "failed to assign an IP address to container",
		Description:      "private subnet IP address pool is exhausted",
		Remediation:      "",
		RunbookUrl:       "https://repost.aws/knowledge-center/eks-failed-create-pod-sandbox",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedCreatePodSandBox",
		MessageSignature: "failed to setup network for sandbox",
		Description:      "EKS network plugin is unhealthy",
		Remediation:      "",
		RunbookUrl:       "https://repost.aws/knowledge-center/eks-failed-create-pod-sandbox",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedCreatePodSandBox",
		MessageSignature: "Error while dialing dial tcp 127.0.0.1:50051",
		Description:      "aws-node pod is not running, node is not healthy",
		Remediation:      "Refresh cluster nodes",
		RunbookUrl:       "https://repost.aws/knowledge-center/eks-failed-create-pod-sandbox",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedCreatePodSandBox",
		MessageSignature: "",
		Description:      "underlying node is unhealthy",
		Remediation:      "Refresh cluster nodes",
		RunbookUrl:       "https://repost.aws/knowledge-center/eks-failed-create-pod-sandbox",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedScheduling",
		MessageSignature: "",
		Description:      "kubernetes cluster doesn't have sufficient capacity",
		Remediation:      "If autoscaler is working properly, and maximum capacity has not been reached, this issue will eventually resolve.",
		RunbookUrl:       "https://repost.aws/knowledge-center/eks-failed-create-pod-sandbox",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedMount",
		MessageSignature: "configmap",
		Description:      "configmap is missing",
		Remediation:      "Create a missing onfigmap",
		RunbookUrl:       "",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedMount",
		MessageSignature: "secret",
		Description:      "secret is missing",
		Remediation:      "Create a missing secret",
		RunbookUrl:       "",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedMount",
		MessageSignature: "",
		Description:      "unable to mount a disk volume, secret, or configmap",
		Remediation:      "Create a missing volume, secret or configmap",
		RunbookUrl:       "",
	},
	{
		Kind:             "Pod",
		Reason:           "FailedMount",
		MessageSignature: "",
		Description:      "secret or configmap is missing, or unable to create a disk volume",
		Remediation:      "Create a missing secret or configmap",
		RunbookUrl:       "",
	},
	{
		Kind:             "Pod",
		Reason:           "NodeNotReady",
		MessageSignature: "",
		Description:      "underlying node is unhealthy",
		Remediation:      "Refresh cluster nodes",
		RunbookUrl:       "https://czi.atlassian.net/wiki/spaces/SI/pages/1786743552/InfraEng+Kubernetes#InfraEng%7CKubernetes-Replacing%2FRotatingAllNodes",
	},
	{
		Kind:             "HorizontalPodAutoscaler",
		Reason:           "FailedGetResourceMetric",
		MessageSignature: "",
		Description:      "failed to collect resource metrics",
		Remediation:      "Make sure metrics-server running and resource requests/limits are set on a corresponding deployment",
		RunbookUrl:       "https://aptakube.com/blog/how-to-fix-failedgeteesourcemetric-hpa",
	},
	{
		Kind:             "HorizontalPodAutoscaler",
		Reason:           "FailedComputeMetricsReplicas",
		MessageSignature: "",
		Description:      "failed to compute the resource the number of replicas based on metrics",
		Remediation:      "Make sure metrics-server running and resource requests/limits are set on a corresponding deployment",
		RunbookUrl:       "https://aptakube.com/blog/how-to-fix-failedgeteesourcemetric-hpa",
	},
	{
		Kind:             "TargetGroupBinding",
		Reason:           "BackendNotFound",
		MessageSignature: "",
		Description:      "cannot create an ALB route because backend service is missing",
		Remediation:      "Remove all resources that are in orphaned state and try again",
		RunbookUrl:       "https://aptakube.com/blog/how-to-fix-failedgeteesourcemetric-hpa",
	},
	{
		Kind:             "Ingress",
		Reason:           "FailedBuildModel",
		MessageSignature: "InvalidGroup.NotFound",
		Description:      "unable to create a load balancer",
		Remediation:      "security group doesn't exist",
		RunbookUrl:       "",
	},
	{
		Kind:             "Ingress",
		Reason:           "FailedBuildModel",
		MessageSignature: "InvalidGroup.Duplicate",
		Description:      "unable to create a load balancer",
		Remediation:      "security group already exists exist",
		RunbookUrl:       "",
	},
}
