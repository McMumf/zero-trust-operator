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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	networkingv1 "k8s.io/api/networking/v1"
)

var _ = Describe("IngressThing Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()
		ingressClassName := "nginx"
		pathType := "Prefix"

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default", // TODO(user):Modify as needed
		}
		ingressController := &networkingv1.Ingress{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind IngressThing")
			err := k8sClient.Get(ctx, typeNamespacedName, ingressController)
			if err != nil && errors.IsNotFound(err) {
				resource := &networkingv1.Ingress{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: networkingv1.IngressSpec{
						IngressClassName: &ingressClassName,
						Rules: []networkingv1.IngressRule{
							{
								IngressRuleValue: networkingv1.IngressRuleValue{
									HTTP: &networkingv1.HTTPIngressRuleValue{
										Paths: []networkingv1.HTTPIngressPath{
											{
												Path:     "/ass",
												PathType: (*networkingv1.PathType)(&pathType),
												Backend: networkingv1.IngressBackend{
													Service: &networkingv1.IngressServiceBackend{
														Name: "svc-name",
														Port: networkingv1.ServiceBackendPort{
															Number: 8188,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					// TODO(user): Specify other spec details if needed.
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			// TODO(user): Cleanup logic after each test, like removing the resource instance.
			resource := &networkingv1.Ingress{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance Cname")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})
		It("should successfully reconcile the resource", func() {
			By("Reconciling the created resource")
			controllerReconciler := &IngressControllerReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())
			// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
			// Example: If you expect a certain status condition after reconciliation, verify it here.
		})
	})
})
