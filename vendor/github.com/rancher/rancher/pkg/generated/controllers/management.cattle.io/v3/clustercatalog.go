/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/v3/pkg/generic"
)

// ClusterCatalogController interface for managing ClusterCatalog resources.
type ClusterCatalogController interface {
	generic.ControllerInterface[*v3.ClusterCatalog, *v3.ClusterCatalogList]
}

// ClusterCatalogClient interface for managing ClusterCatalog resources in Kubernetes.
type ClusterCatalogClient interface {
	generic.ClientInterface[*v3.ClusterCatalog, *v3.ClusterCatalogList]
}

// ClusterCatalogCache interface for retrieving ClusterCatalog resources in memory.
type ClusterCatalogCache interface {
	generic.CacheInterface[*v3.ClusterCatalog]
}
