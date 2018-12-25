// +build !ignore_autogenerated

// Copyright 2018 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServerConfig) DeepCopyInto(out *APIServerConfig) {
	*out = *in
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServerConfig.
func (in *APIServerConfig) DeepCopy() *APIServerConfig {
	if in == nil {
		return nil
	}
	out := new(APIServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingSpec) DeepCopyInto(out *AlertingSpec) {
	*out = *in
	if in.Alertmanagers != nil {
		in, out := &in.Alertmanagers, &out.Alertmanagers
		*out = make([]AlertmanagerEndpoints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingSpec.
func (in *AlertingSpec) DeepCopy() *AlertingSpec {
	if in == nil {
		return nil
	}
	out := new(AlertingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alertmanager) DeepCopyInto(out *Alertmanager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(AlertmanagerStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alertmanager.
func (in *Alertmanager) DeepCopy() *Alertmanager {
	if in == nil {
		return nil
	}
	out := new(Alertmanager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerEndpoints) DeepCopyInto(out *AlertmanagerEndpoints) {
	*out = *in
	out.Port = in.Port
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerEndpoints.
func (in *AlertmanagerEndpoints) DeepCopy() *AlertmanagerEndpoints {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerEndpoints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerList) DeepCopyInto(out *AlertmanagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alertmanager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerList.
func (in *AlertmanagerList) DeepCopy() *AlertmanagerList {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerSpec) DeepCopyInto(out *AlertmanagerSpec) {
	*out = *in
	if in.PodMetadata != nil {
		in, out := &in.PodMetadata, &out.PodMetadata
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.ObjectMeta)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]core_v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PodSecurityContext)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]core_v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalPeers != nil {
		in, out := &in.AdditionalPeers, &out.AdditionalPeers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerSpec.
func (in *AlertmanagerSpec) DeepCopy() *AlertmanagerSpec {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerStatus) DeepCopyInto(out *AlertmanagerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerStatus.
func (in *AlertmanagerStatus) DeepCopy() *AlertmanagerStatus {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Password.DeepCopyInto(&out.Password)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrdKind) DeepCopyInto(out *CrdKind) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrdKind.
func (in *CrdKind) DeepCopy() *CrdKind {
	if in == nil {
		return nil
	}
	out := new(CrdKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrdKinds) DeepCopyInto(out *CrdKinds) {
	*out = *in
	out.Prometheus = in.Prometheus
	out.Alertmanager = in.Alertmanager
	out.ServiceMonitor = in.ServiceMonitor
	out.PrometheusRule = in.PrometheusRule
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrdKinds.
func (in *CrdKinds) DeepCopy() *CrdKinds {
	if in == nil {
		return nil
	}
	out := new(CrdKinds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		if *in == nil {
			*out = nil
		} else {
			*out = new(intstr.IntOrString)
			**out = **in
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.MetricRelabelConfigs != nil {
		in, out := &in.MetricRelabelConfigs, &out.MetricRelabelConfigs
		*out = make([]*RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(RelabelConfig)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.RelabelConfigs != nil {
		in, out := &in.RelabelConfigs, &out.RelabelConfigs
		*out = make([]*RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(RelabelConfig)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.ProxyURL != nil {
		in, out := &in.ProxyURL, &out.ProxyURL
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSelector) DeepCopyInto(out *NamespaceSelector) {
	*out = *in
	if in.MatchNames != nil {
		in, out := &in.MatchNames, &out.MatchNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSelector.
func (in *NamespaceSelector) DeepCopy() *NamespaceSelector {
	if in == nil {
		return nil
	}
	out := new(NamespaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(PrometheusStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus.
func (in *Prometheus) DeepCopy() *Prometheus {
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusList) DeepCopyInto(out *PrometheusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*Prometheus, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Prometheus)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusList.
func (in *PrometheusList) DeepCopy() *PrometheusList {
	if in == nil {
		return nil
	}
	out := new(PrometheusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRule) DeepCopyInto(out *PrometheusRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRule.
func (in *PrometheusRule) DeepCopy() *PrometheusRule {
	if in == nil {
		return nil
	}
	out := new(PrometheusRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleList) DeepCopyInto(out *PrometheusRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*PrometheusRule, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(PrometheusRule)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleList.
func (in *PrometheusRuleList) DeepCopy() *PrometheusRuleList {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleSpec) DeepCopyInto(out *PrometheusRuleSpec) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]RuleGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleSpec.
func (in *PrometheusRuleSpec) DeepCopy() *PrometheusRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSpec) DeepCopyInto(out *PrometheusSpec) {
	*out = *in
	if in.PodMetadata != nil {
		in, out := &in.PodMetadata, &out.PodMetadata
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.ObjectMeta)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ServiceMonitorSelector != nil {
		in, out := &in.ServiceMonitorSelector, &out.ServiceMonitorSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ServiceMonitorNamespaceSelector != nil {
		in, out := &in.ServiceMonitorNamespaceSelector, &out.ServiceMonitorNamespaceSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]core_v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.ExternalLabels != nil {
		in, out := &in.ExternalLabels, &out.ExternalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RuleSelector != nil {
		in, out := &in.RuleSelector, &out.RuleSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RuleNamespaceSelector != nil {
		in, out := &in.RuleNamespaceSelector, &out.RuleNamespaceSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Alerting != nil {
		in, out := &in.Alerting, &out.Alerting
		if *in == nil {
			*out = nil
		} else {
			*out = new(AlertingSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemoteWrite != nil {
		in, out := &in.RemoteWrite, &out.RemoteWrite
		*out = make([]RemoteWriteSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemoteRead != nil {
		in, out := &in.RemoteRead, &out.RemoteRead
		*out = make([]RemoteReadSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PodSecurityContext)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]core_v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalScrapeConfigs != nil {
		in, out := &in.AdditionalScrapeConfigs, &out.AdditionalScrapeConfigs
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AdditionalAlertRelabelConfigs != nil {
		in, out := &in.AdditionalAlertRelabelConfigs, &out.AdditionalAlertRelabelConfigs
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AdditionalAlertManagerConfigs != nil {
		in, out := &in.AdditionalAlertManagerConfigs, &out.AdditionalAlertManagerConfigs
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.APIServerConfig != nil {
		in, out := &in.APIServerConfig, &out.APIServerConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(APIServerConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Thanos != nil {
		in, out := &in.Thanos, &out.Thanos
		if *in == nil {
			*out = nil
		} else {
			*out = new(ThanosSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSpec.
func (in *PrometheusSpec) DeepCopy() *PrometheusSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusStatus) DeepCopyInto(out *PrometheusStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusStatus.
func (in *PrometheusStatus) DeepCopy() *PrometheusStatus {
	if in == nil {
		return nil
	}
	out := new(PrometheusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueConfig) DeepCopyInto(out *QueueConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueConfig.
func (in *QueueConfig) DeepCopy() *QueueConfig {
	if in == nil {
		return nil
	}
	out := new(QueueConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelabelConfig) DeepCopyInto(out *RelabelConfig) {
	*out = *in
	if in.SourceLabels != nil {
		in, out := &in.SourceLabels, &out.SourceLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelabelConfig.
func (in *RelabelConfig) DeepCopy() *RelabelConfig {
	if in == nil {
		return nil
	}
	out := new(RelabelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteReadSpec) DeepCopyInto(out *RemoteReadSpec) {
	*out = *in
	if in.RequiredMatchers != nil {
		in, out := &in.RequiredMatchers, &out.RequiredMatchers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteReadSpec.
func (in *RemoteReadSpec) DeepCopy() *RemoteReadSpec {
	if in == nil {
		return nil
	}
	out := new(RemoteReadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteWriteSpec) DeepCopyInto(out *RemoteWriteSpec) {
	*out = *in
	if in.WriteRelabelConfigs != nil {
		in, out := &in.WriteRelabelConfigs, &out.WriteRelabelConfigs
		*out = make([]RelabelConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	if in.QueueConfig != nil {
		in, out := &in.QueueConfig, &out.QueueConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(QueueConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteWriteSpec.
func (in *RemoteWriteSpec) DeepCopy() *RemoteWriteSpec {
	if in == nil {
		return nil
	}
	out := new(RemoteWriteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	out.Expr = in.Expr
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleGroup) DeepCopyInto(out *RuleGroup) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleGroup.
func (in *RuleGroup) DeepCopy() *RuleGroup {
	if in == nil {
		return nil
	}
	out := new(RuleGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitor) DeepCopyInto(out *ServiceMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitor.
func (in *ServiceMonitor) DeepCopy() *ServiceMonitor {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitorList) DeepCopyInto(out *ServiceMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*ServiceMonitor, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(ServiceMonitor)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitorList.
func (in *ServiceMonitorList) DeepCopy() *ServiceMonitorList {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitorSpec) DeepCopyInto(out *ServiceMonitorSpec) {
	*out = *in
	if in.TargetLabels != nil {
		in, out := &in.TargetLabels, &out.TargetLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodTargetLabels != nil {
		in, out := &in.PodTargetLabels, &out.PodTargetLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitorSpec.
func (in *ServiceMonitorSpec) DeepCopy() *ServiceMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.EmptyDirVolumeSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosGCSSpec) DeepCopyInto(out *ThanosGCSSpec) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosGCSSpec.
func (in *ThanosGCSSpec) DeepCopy() *ThanosGCSSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosGCSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosS3Spec) DeepCopyInto(out *ThanosS3Spec) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.SignatureVersion2 != nil {
		in, out := &in.SignatureVersion2, &out.SignatureVersion2
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.EncryptSSE != nil {
		in, out := &in.EncryptSSE, &out.EncryptSSE
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosS3Spec.
func (in *ThanosS3Spec) DeepCopy() *ThanosS3Spec {
	if in == nil {
		return nil
	}
	out := new(ThanosS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosSpec) DeepCopyInto(out *ThanosSpec) {
	*out = *in
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.SHA != nil {
		in, out := &in.SHA, &out.SHA
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.BaseImage != nil {
		in, out := &in.BaseImage, &out.BaseImage
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		if *in == nil {
			*out = nil
		} else {
			*out = new(ThanosGCSSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		if *in == nil {
			*out = nil
		} else {
			*out = new(ThanosS3Spec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosSpec.
func (in *ThanosSpec) DeepCopy() *ThanosSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosSpec)
	in.DeepCopyInto(out)
	return out
}