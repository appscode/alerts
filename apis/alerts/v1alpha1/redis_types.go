/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "kubepack.dev/lib-app/api/v1alpha1"
)

const (
	ResourceKindRedis = "Redis"
	ResourceRedis     = "redis"
	ResourceRediss    = "rediss"
)

// Redis defines the schama for KubeDB Ops Manager Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=rediss,singular=redis,categories={kubedb,appscode}
type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisSpec `json:"spec,omitempty"`
}

// RedisSpec is the schema for kubedb-autoscaler chart values file
type RedisSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Form         RedisSpecForm `json:"form"`
}

type RedisSpecForm struct {
	Alert RedisAlert `json:"alert"`
}

type RedisAlert struct {
	Enabled bool              `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string `json:"additionalRuleLabels"`
	Groups               RedisAlertGroups  `json:"groups"`
}

type RedisAlertGroups struct {
	Database    RedisDatabaseAlert `json:"database"`
	Provisioner ProvisionerAlert   `json:"provisioner"`
	OpsManager  OpsManagerAlert    `json:"opsManager"`
	Stash       StashAlert         `json:"stash"`
}

type RedisDatabaseAlert struct {
	Enabled bool                    `json:"enabled"`
	Rules   RedisDatabaseAlertRules `json:"rules"`
}

type RedisDatabaseAlertRules struct {
	RedisDown                FixedAlert  `json:"redisDown"`
	RedisMissingMaster       IntValAlert `json:"redisMissingMaster"`
	RedisTooManyMasters      IntValAlert `json:"redisTooManyMasters"`
	RedisDisconnectedSlaves  IntValAlert `json:"redisDisconnectedSlaves"`
	RedisTooManyConnections  IntValAlert `json:"redisTooManyConnections"`
	RedisRejectedConnections IntValAlert `json:"redisRejectedConnections"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisList is a list of Rediss
type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Redis CRD objects
	Items []Redis `json:"items,omitempty"`
}
