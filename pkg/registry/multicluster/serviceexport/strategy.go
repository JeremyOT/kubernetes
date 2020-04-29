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

package serviceexport

import (
	"context"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/multicluster"
	"k8s.io/kubernetes/pkg/apis/multicluster/validation"
)

// serviceExportStrategy implements verification logic for Replication.
type serviceExportStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// Strategy is the default logic that applies when creating and updating Replication ServiceExport objects.
var Strategy = serviceExportStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

// NamespaceScoped returns true because all ServiceExports need to be within a namespace.
func (serviceExportStrategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate clears the status of an ServiceExport before creation.
func (serviceExportStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	serviceExport := obj.(*multicluster.ServiceExport)
	serviceExport.Generation = 1
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (serviceExportStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	newEPS := obj.(*multicluster.ServiceExport)
	oldEPS := old.(*multicluster.ServiceExport)

	// Increment generation if anything other than meta changed
	// This needs to be changed if a status attribute is added to ServiceExport
	ogNewMeta := newEPS.ObjectMeta
	ogOldMeta := oldEPS.ObjectMeta
	newEPS.ObjectMeta = v1.ObjectMeta{}
	oldEPS.ObjectMeta = v1.ObjectMeta{}

	if !apiequality.Semantic.DeepEqual(newEPS, oldEPS) {
		ogNewMeta.Generation = ogOldMeta.Generation + 1
	}

	newEPS.ObjectMeta = ogNewMeta
	oldEPS.ObjectMeta = ogOldMeta
}

// Validate validates a new ServiceExport.
func (serviceExportStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	serviceExport := obj.(*multicluster.ServiceExport)
	err := validation.ValidateServiceExportCreate(serviceExport)
	return err
}

// Canonicalize normalizes the object after validation.
func (serviceExportStrategy) Canonicalize(obj runtime.Object) {
}

// AllowCreateOnUpdate is false for ServiceExport; this means POST is needed to create one.
func (serviceExportStrategy) AllowCreateOnUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for an end user.
func (serviceExportStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	newEPS := new.(*multicluster.ServiceExport)
	oldEPS := old.(*multicluster.ServiceExport)
	return validation.ValidateServiceExportUpdate(newEPS, oldEPS)
}

// AllowUnconditionalUpdate is the default update policy for ServiceExport objects.
func (serviceExportStrategy) AllowUnconditionalUpdate() bool {
	return true
}
