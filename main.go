package main

import (
	"context"
	"os"

	"github.com/henderiw-k8s-lcnc/fn-sdk/go/fn"
	upfv1alpha1 "github.com/henderiw-k8s-lcnc/fn-upf/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ fn.Runner = &upf{}

type upf struct{}

func main() {
	ctx := context.TODO()
	if err := fn.AsMain(fn.WithContext(ctx, &upf{})); err != nil {
		os.Exit(1)
	}
}

func (r *upf) Run(ctx *fn.Context, functionConfig map[string]runtime.RawExtension, resources *fn.Resources, results *fn.Results) bool {

	impl := &upfv1alpha1.UpfImplementation{
		TypeMeta: metav1.TypeMeta{
			APIVersion: upfv1alpha1.GroupVersion.String(),
			Kind:       upfv1alpha1.UpfImplementationKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "dummy",
		},
		Spec: upfv1alpha1.UpfImplementationSpec{
			Implementation: "a",
		},
	}

	resources.AddResource(impl, nil)
	return true
}
