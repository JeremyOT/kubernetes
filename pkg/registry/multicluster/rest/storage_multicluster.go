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

package rest

import (
	multiclusterv1alpha1 "k8s.io/api/multicluster/v1alpha1"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/multicluster"
	serviceexportstorage "k8s.io/kubernetes/pkg/registry/multicluster/serviceexport/storage"
	serviceimportstorage "k8s.io/kubernetes/pkg/registry/multicluster/serviceimport/storage"
)

// RESTStorageProvider is a REST storage provider for multicluster.k8s.io.
type RESTStorageProvider struct{}

// NewRESTStorage returns a new storage provider.
func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool, error) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(multicluster.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
	// If you add a version here, be sure to add an entry in `k8s.io/kubernetes/cmd/kube-apiserver/app/aggregator.go with specific priorities.
	// TODO refactor the plumbing to provide the information in the APIGroupInfo

	if apiResourceConfigSource.VersionEnabled(multiclusterv1alpha1.SchemeGroupVersion) {
		storageMap, err := p.v1alpha1Storage(apiResourceConfigSource, restOptionsGetter)
		if err != nil {
			return genericapiserver.APIGroupInfo{}, false, err
		}
		apiGroupInfo.VersionedResourcesStorageMap[multiclusterv1alpha1.SchemeGroupVersion.Version] = storageMap
	}

	return apiGroupInfo, true, nil
}

func (p RESTStorageProvider) v1alpha1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (map[string]rest.Storage, error) {
	storage := map[string]rest.Storage{}

	serviceExportStorage, err := serviceexportstorage.NewREST(restOptionsGetter)
	if err != nil {
		return storage, err
	}

	storage["serviceexports"] = serviceExportStorage

	serviceImportStorage, err := serviceimportstorage.NewREST(restOptionsGetter)
	if err != nil {
		return storage, err
	}

	storage["serviceimports"] = serviceImportStorage
	return storage, err
}

// GroupName is the group name for the storage provider.
func (p RESTStorageProvider) GroupName() string {
	return multicluster.GroupName
}
