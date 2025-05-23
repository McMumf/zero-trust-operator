/*
Copyright 2025 Carson McCaffrey.

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

package ingress

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	networkingv1 "k8s.io/api/networking/v1"
)

// IngressControllerReconciler reconciles a Ingress object
type IngressControllerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=get
func (r *IngressControllerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	log.Info("Watching Ingress Thing")
	ingress := &networkingv1.Ingress{}
	if err := r.Get(ctx, req.NamespacedName, ingress); err != nil {
		log.Error(err, "Unabled to fetch Ingress")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	ingressCname := ingress.Name + "-cname"

	// Check if CNAME for ingress already exists
	err := r.Get(ctx, types.NamespacedName{Name: ingressCname, Namespace: ingress.Namespace}, nil)
	if err == nil {
		log.Info("CNAME already exists, skipping creation")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IngressControllerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkingv1.Ingress{}).
		Watches(
			new(networkingv1.Ingress), // Watch ingress resources
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {
				// Check if the Busybox resource has the label 'backup-needed: "true"'
				if val, ok := obj.GetAnnotations()[""]; ok && val == "true" {
					// If the label is present and set to "true", trigger reconciliation for BackupBusybox
					return []reconcile.Request{
						{
							NamespacedName: types.NamespacedName{
								Name:      "ingress",          // Reconcile the associated BackupBusybox resource
								Namespace: obj.GetNamespace(), // Use the namespace of the changed Busybox
							},
						},
					}
				}
				// If the label is not present or doesn't match, don't trigger reconciliation
				return []reconcile.Request{}
			}),
		).
		Complete(r)
}
