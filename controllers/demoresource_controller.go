/*
Copyright 2021 The Partner Lifecycle Engineering Team.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	toolsv1 "github.com/opdev/simple-demo-operator/api/v1"
)

// DemoResourceReconciler reconciles a DemoResource object
type DemoResourceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=tools.opdev.io,resources=demoresources,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tools.opdev.io,resources=demoresources/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tools.opdev.io,resources=demoresources/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *DemoResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	key := req.NamespacedName
	l := log.FromContext(ctx).WithValues("NamespacedName", key)

	l.Info("This is the reconciliation loop of the simple-demo-operator.")

	var rsrc toolsv1.DemoResource
	if err := r.Get(ctx, key, &rsrc); err != nil {
		if apierrors.IsNotFound(err) {
			l.Info("Whoops! We couldn't find the resource associated with this event.")
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	rsrc.Status.SpecMessage = rsrc.Spec.Message
	l.Info("Let's update the status with the message from the spec!")
	if err := r.Status().Update(ctx, &rsrc, &client.UpdateOptions{}); err != nil {
		l.Info("Uh oh! I ran into something unexpected.")
		return ctrl.Result{}, err
	}

	l.Info("All finished!")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DemoResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&toolsv1.DemoResource{}).
		Complete(r)
}
