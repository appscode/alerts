#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eou pipefail

crd-importer \
    --input=https://github.com/prometheus-operator/prometheus-operator/raw/v0.59.1/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml \
    --out=./charts/elasticsearch-alerts/crds

crd-importer \
    --input=https://github.com/prometheus-operator/prometheus-operator/raw/v0.59.1/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml \
    --out=./charts/mariadb-alerts/crds

crd-importer \
    --input=https://github.com/prometheus-operator/prometheus-operator/raw/v0.59.1/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml \
    --out=./charts/mongodb-alerts/crds

crd-importer \
    --input=https://github.com/prometheus-operator/prometheus-operator/raw/v0.59.1/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml \
    --out=./charts/mysql-alerts/crds

crd-importer \
    --input=https://github.com/prometheus-operator/prometheus-operator/raw/v0.59.1/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml \
    --out=./charts/postgres-alerts/crds

crd-importer \
    --input=https://github.com/prometheus-operator/prometheus-operator/raw/v0.59.1/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml \
    --out=./charts/redis-alerts/crds
