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

package serviceimport

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

// serviceImportStrategy implements verification logic for Replication.
type serviceImportStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// Strategy is the default logic that applies when creating and updating Replication ServiceImport objects.
var Strategy = serviceImportStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

// NamespaceScoped returns true because all ServiceImports need to be within a namespace.
func (serviceImportStrategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate clears the status of an ServiceImport before creation.
func (serviceImportStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	serviceImport := obj.(*multicluster.ServiceImport)
	serviceImport.Generation = 1
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (serviceImportStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	newEPS := obj.(*multicluster.ServiceImport)
	oldEPS := old.(*multicluster.ServiceImport)

	// Increment generation if anything other than meta changed
	// This needs to be changed if a status attribute is added to ServiceImport
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

// Validate validates a new ServiceImport.
func (serviceImportStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	serviceImport := obj.(*multicluster.ServiceImport)
	err := validation.ValidateServiceImportCreate(serviceImport)
	return err
}

// Canonicalize normalizes the object after validation.
func (serviceImportStrategy) Canonicalize(obj runtime.Object) {
}

// AllowCreateOnUpdate is false for ServiceImport; this means POST is needed to create one.
func (serviceImportStrategy) AllowCreateOnUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for an end user.
func (serviceImportStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	newEPS := new.(*multicluster.ServiceImport)
	oldEPS := old.(*multicluster.ServiceImport)
	return validation.ValidateServiceImportUpdate(newEPS, oldEPS)
}

// AllowUnconditionalUpdate is the default update policy for ServiceImport objects.
func (serviceImportStrategy) AllowUnconditionalUpdate() bool {
	return true
}
