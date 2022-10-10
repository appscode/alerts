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
	ResourceKindPerconaxtradb = "Perconaxtradb"
	ResourcePerconaxtradb     = "perconaxtradb"
	ResourcePerconaxtradbs    = "perconaxtradbs"
)

// Perconaxtradb defines the schama for KubeDB Ops Manager Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=perconaxtradbs,singular=perconaxtradb,categories={kubedb,appscode}
type Perconaxtradb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PerconaxtradbSpec `json:"spec,omitempty"`
}

// PerconaxtradbSpec is the schema for kubedb-autoscaler chart values file
type PerconaxtradbSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Form         PerconaxtradbSpecForm `json:"form"`
}

type PerconaxtradbSpecForm struct {
	Alert PerconaXtraDBAlert `json:"alert"`
}

type PerconaXtraDBAlert struct {
	Enabled bool              `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string        `json:"additionalRuleLabels"`
	Groups               PerconaXtraDBAlertGroups `json:"groups"`
}

type PerconaXtraDBAlertGroups struct {
	Database      PerconaXtraDBDatabaseAlert `json:"database"`
	Cluster       PerconaXtraDBClusterAlert  `json:"cluster"`
	Provisioner   ProvisionerAlert           `json:"provisioner"`
	OpsManager    OpsManagerAlert            `json:"opsManager"`
	Stash         StashAlert                 `json:"stash"`
	SchemaManager SchemaManagerAlert         `json:"schemaManager"`
}

type PerconaXtraDBDatabaseAlert struct {
	Enabled bool                            `json:"enabled"`
	Rules   PerconaXtraDBDatabaseAlertRules `json:"rules"`
}

type PerconaXtraDBDatabaseAlertRules struct {
	MySQLInstanceDown       FixedAlert  `json:"mySQLInstanceDown"`
	MySQLServiceDown        FixedAlert  `json:"mySQLServiceDown"`
	MySQLTooManyConnections IntValAlert `json:"mySQLTooManyConnections"`
	MySQLHighThreadsRunning IntValAlert `json:"mySQLHighThreadsRunning"`
	MySQLSlowQueries        FixedAlert  `json:"mySQLSlowQueries"`
	MySQLInnoDBLogWaits     IntValAlert `json:"mySQLInnoDBLogWaits"`
	MySQLRestarted          IntValAlert `json:"mySQLRestarted"`
	MySQLHighQPS            IntValAlert `json:"mySQLHighQPS"`
	MySQLHighIncomingBytes  IntValAlert `json:"mySQLHighIncomingBytes"`
	MySQLHighOutgoingBytes  IntValAlert `json:"mySQLHighOutgoingBytes"`
	MySQLTooManyOpenFiles   IntValAlert `json:"mySQLTooManyOpenFiles"`
}

type PerconaXtraDBClusterAlert struct {
	Enabled bool                           `json:"enabled"`
	Rules   PerconaXtraDBClusterAlertRules `json:"rules"`
}

type PerconaXtraDBClusterAlertRules struct {
	GaleraReplicationLatencyTooLong FloatValAlertConfig `json:"galeraReplicationLatencyTooLong"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerconaxtradbList is a list of Perconaxtradbs
type PerconaxtradbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Perconaxtradb CRD objects
	Items []Perconaxtradb `json:"items,omitempty"`
}
