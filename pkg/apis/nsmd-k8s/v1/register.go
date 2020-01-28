package v1

import (
	nsmd_k8s "github.com/dtornow/cnns-nsr/pkg/apis/nsmd-k8s"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   nsmd_k8s.GroupName,
	Version: "v1alpha1",
}

var scheme = runtime.NewSchemeBuilder(addKnownTypes)
var AddToScheme = scheme.AddToScheme

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds our types to the API scheme by registering
// NetworkService and NetworkServiceList
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&NetworkService{},
		&NetworkServiceList{},
		&NetworkServiceEndpoint{},
		&NetworkServiceEndpointList{},
		&NetworkServiceManager{},
		&NetworkServiceManagerList{},
	)

	// register the type in the scheme
	metaV1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
