// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package manifest

import (
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kustomize/api/hasher"
	"sigs.k8s.io/kustomize/api/ifc"
)

// kustHash computes a hash of an unstructured object.
type kustHash struct{}

// NewKustHash returns a kustHash object
func NewKustHash() *kustHash {
	return &kustHash{}
}

// Hash returns a hash of either a ConfigMap or a Secret
func (h *kustHash) Hash(m ifc.Kunstructured) (string, error) {
	u := unstructured.Unstructured{
		Object: m.Map(),
	}
	cm, err := unstructuredToConfigmap(u)
	if err != nil {
		return "", err
	}
	return configMapHash(cm)
}

// configMapHash returns a hash of the ConfigMap.
// The Data, Kind, and Name are taken into account.
func configMapHash(cm *v1.ConfigMap) (string, error) {
	encoded, err := encodeConfigMap(cm)
	if err != nil {
		return "", err
	}
	h, err := hasher.Encode(hasher.Hash(encoded))
	if err != nil {
		return "", err
	}
	return h, nil
}

// encodeConfigMap encodes a ConfigMap.
// Data, Kind, and Name are taken into account.
func encodeConfigMap(cm *v1.ConfigMap) (string, error) {
	// json.Marshal sorts the keys in a stable order in the encoding
	m := map[string]interface{}{"kind": "ConfigMap", "name": cm.Name, "data": cm.Data}
	if len(cm.BinaryData) > 0 {
		m["binaryData"] = cm.BinaryData
	}
	data, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func unstructuredToConfigmap(u unstructured.Unstructured) (*v1.ConfigMap, error) {
	marshaled, err := json.Marshal(u.Object)
	if err != nil {
		return nil, err
	}
	var out v1.ConfigMap
	err = json.Unmarshal(marshaled, &out)
	return &out, err
}