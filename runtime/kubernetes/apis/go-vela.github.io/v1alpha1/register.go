package v1alpha1

import (
	"github.com/go-vela/worker/runtime/kubernetes/apis/go-vela.github.io"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: go_vela_github_io.GroupName, Version: "v1alpha1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

//var (
//	// SchemeBuilder initializes a scheme builder
//	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
//	// AddToScheme is a global function that registers this API group & version to a scheme
//	AddToScheme = SchemeBuilder.AddToScheme
//)

// Adds the list of known types to Scheme.
//func addKnownTypes(scheme *runtime.Scheme) error {
//	scheme.AddKnownTypes(SchemeGroupVersion,
//		&PipelinePod{},
//		&PipelinePodList{},
//	)
//	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
//	return nil
//}
