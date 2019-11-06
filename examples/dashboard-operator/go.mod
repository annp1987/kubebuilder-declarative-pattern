module dashboard-operator

go 1.12

require (
	github.com/go-logr/logr v0.1.0
	k8s.io/apimachinery v0.0.0-20191020214737-6c8691705fc5
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.0-beta.2
	sigs.k8s.io/kubebuilder-declarative-pattern v0.0.0-20190624171758-3bfb5869c8b7
)

replace sigs.k8s.io/kubebuilder-declarative-pattern => ../../
