/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	multiclusterv1alpha1 "k8s.io/api/multicluster/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/features"
)

var defaultIPFamily = v1.IPv4Protocol

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_ServiceImport(obj *multiclusterv1alpha1.ServiceImport) {
	for i := range obj.Spec.Ports {
		sp := &obj.Spec.Ports[i]
		if sp.Protocol == "" {
			sp.Protocol = v1.ProtocolTCP
		}
		if sp.TargetPort == intstr.FromInt(0) || sp.TargetPort == intstr.FromString("") {
			sp.TargetPort = intstr.FromInt(int(sp.Port))
		}
	}
	if utilfeature.DefaultFeatureGate.Enabled(features.IPv6DualStack) &&
		obj.Spec.IPFamily == nil {
		obj.Spec.IPFamily = &defaultIPFamily
	}
}
