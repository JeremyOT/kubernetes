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

package util

// ServiceImportName returns an import name for the supplied  service,
// formatted to avoid conflicts with same-named regular services. This
// name is not a valid K8s Service name and is for internal use only.
func ServiceImportName(serviceName string) string {
	return "import:" + serviceName
}
