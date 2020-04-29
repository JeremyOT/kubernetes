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

package validation

import (
	apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apivalidation "k8s.io/kubernetes/pkg/apis/core/validation"
	"k8s.io/kubernetes/pkg/apis/multicluster"
)

// ValidateEndpointSliceName can be used to check whether the given endpoint
// slice name is valid. Prefix indicates this name will be used as part of
// generation, in which case trailing dashes are allowed.
var ValidateServiceName = apimachineryvalidation.NameIsDNSSubdomain

// ValidateServiceExport validates an ServiceExport.
func ValidateServiceExport(serviceExport *multicluster.ServiceExport) field.ErrorList {
	allErrs := apivalidation.ValidateObjectMeta(&serviceExport.ObjectMeta, true, ValidateServiceName, field.NewPath("metadata"))

	return allErrs
}

// ValidateServiceExportCreate validates an ServiceExport when it is created.
func ValidateServiceExportCreate(serviceExport *multicluster.ServiceExport) field.ErrorList {
	return ValidateServiceExport(serviceExport)
}

// ValidateServiceExportUpdate validates an ServiceExport when it is updated.
func ValidateServiceExportUpdate(newServiceExport, oldServiceExport *multicluster.ServiceExport) field.ErrorList {
	allErrs := ValidateServiceExport(newServiceExport)

	return allErrs
}

// ValidateServiceImport validates an ServiceImport.
func ValidateServiceImport(serviceImport *multicluster.ServiceImport) field.ErrorList {
	allErrs := apivalidation.ValidateObjectMeta(&serviceImport.ObjectMeta, true, ValidateServiceName, field.NewPath("metadata"))

	return allErrs
}

// ValidateServiceImportCreate validates an ServiceImport when it is created.
func ValidateServiceImportCreate(serviceImport *multicluster.ServiceImport) field.ErrorList {
	return ValidateServiceImport(serviceImport)
}

// ValidateServiceImportUpdate validates an ServiceImport when it is updated.
func ValidateServiceImportUpdate(newServiceImport, oldServiceImport *multicluster.ServiceImport) field.ErrorList {
	allErrs := ValidateServiceImport(newServiceImport)

	return allErrs
}
