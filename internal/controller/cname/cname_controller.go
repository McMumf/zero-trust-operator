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

package cname

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	cnamev1alpha1 "mcmumf.dev/zero-trust-operator/api/v1alpha1"
)

// CnameReconciler reconciles a Cname object
type CnameReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cname.mcmumf.dev,resources=cnames,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cname.mcmumf.dev,resources=cnames/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cname.mcmumf.dev,resources=cnames/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Cname object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.4/pkg/reconcile
func (r *CnameReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	log.Info("Watching CNAME")
	cnameObj := &cnamev1alpha1.Cname{}
	if err := r.Get(ctx, req.NamespacedName, cnameObj); err != nil {
		log.Error(err, "unable to fetch CNAME")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CnameReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cnamev1alpha1.Cname{}).
		Named("cname").
		Complete(r)
}
