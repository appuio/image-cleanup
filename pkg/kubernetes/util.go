package kubernetes

import (
	"strings"

	"k8s.io/client-go/dynamic"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type (
	// Kubernetes defines the interface to interact with K8s
	Kubernetes interface {
		ResourceContains(namespace, value string, resource schema.GroupVersionResource) (bool, error)
	}
	// kubernetesImpl is an implementation of the interface. (Better name? introduced for better testing support)
	kubernetesImpl struct {
		client dynamic.Interface
	}
)

// New creates a new Kubernetes instance
func New() Kubernetes {
	return &kubernetesImpl{}
}

// ResourceContains evaluates if a given resource contains a given string
func (k *kubernetesImpl) ResourceContains(namespace, value string, resource schema.GroupVersionResource) (bool, error) {
	err := k.initClient()
	if err != nil {
		return false, err
	}
	objectlist, err := k.client.Resource(resource).Namespace(namespace).List(metav1.ListOptions{})
	if err != nil {
		return false, err
	}

	return UnstructuredListContains(objectlist, value), nil
}

func (k *kubernetesImpl) initClient() error {
	if k.client == nil {
		client, err := NewDynamicClient()
		if err != nil {
			return err
		}
		k.client = client
	}
	return nil
}

// ObjectContains evaluates if a Kubernetes object contains a certain string
func ObjectContains(genericObject interface{}, value string) bool {
	switch (genericObject).(type) {
	case map[string]interface{}:
		objects := (genericObject).(map[string]interface{})

		for key := range objects {
			object := objects[key]
			if ObjectContains(object, value) {
				return true
			}
		}

		return false

	case []interface{}:
		for _, object := range (genericObject).([]interface{}) {
			if ObjectContains(object, value) {
				return true
			}
		}

		return false

	case string:
		return strings.Contains(genericObject.(string), value)

	default:
		return false
	}
}

func UnstructuredListContains(unstructuredList *unstructured.UnstructuredList, value string) bool {
	for _, item := range unstructuredList.Items {
		if ObjectContains(item.Object, value) {
			return true
		}
	}
	return false
}
