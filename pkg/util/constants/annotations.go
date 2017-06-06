package constants

const (
	ZalandoDNSNameAnnotation           = "external-dns.alpha.kubernetes.io/hostname"
	ElbTimeoutAnnotationName           = "service.beta.kubernetes.io/aws-load-balancer-connection-idle-timeout"
	ElbTimeoutAnnotationValue          = "3600"
	KubeIAmAnnotation                  = "iam.amazonaws.com/role"
	VolumeStorateProvisionerAnnotation = "pv.kubernetes.io/provisioned-by"
)
