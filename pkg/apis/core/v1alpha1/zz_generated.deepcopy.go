// +build !ignore_autogenerated

// Copyright 2019 The Kubernetes Authors.
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

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Advisor) DeepCopyInto(out *Advisor) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(AdvisorType)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int16)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Advisor.
func (in *Advisor) DeepCopy() *Advisor {
	if in == nil {
		return nil
	}
	out := new(Advisor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deployer) DeepCopyInto(out *Deployer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deployer.
func (in *Deployer) DeepCopy() *Deployer {
	if in == nil {
		return nil
	}
	out := new(Deployer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Deployer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerList) DeepCopyInto(out *DeployerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Deployer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerList.
func (in *DeployerList) DeepCopy() *DeployerList {
	if in == nil {
		return nil
	}
	out := new(DeployerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeployerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerSet) DeepCopyInto(out *DeployerSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerSet.
func (in *DeployerSet) DeepCopy() *DeployerSet {
	if in == nil {
		return nil
	}
	out := new(DeployerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeployerSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerSetList) DeepCopyInto(out *DeployerSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeployerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerSetList.
func (in *DeployerSetList) DeepCopy() *DeployerSetList {
	if in == nil {
		return nil
	}
	out := new(DeployerSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeployerSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerSetSpec) DeepCopyInto(out *DeployerSetSpec) {
	*out = *in
	if in.Deployers != nil {
		in, out := &in.Deployers, &out.Deployers
		*out = make([]DeployerSpecDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerSetSpec.
func (in *DeployerSetSpec) DeepCopy() *DeployerSetSpec {
	if in == nil {
		return nil
	}
	out := new(DeployerSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerSetStatus) DeepCopyInto(out *DeployerSetStatus) {
	*out = *in
	if in.Deployers != nil {
		in, out := &in.Deployers, &out.Deployers
		*out = make([]DeployerStatusDescriptor, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerSetStatus.
func (in *DeployerSetStatus) DeepCopy() *DeployerSetStatus {
	if in == nil {
		return nil
	}
	out := new(DeployerSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerSpec) DeepCopyInto(out *DeployerSpec) {
	*out = *in
	if in.PlacementTarget != nil {
		in, out := &in.PlacementTarget, &out.PlacementTarget
		*out = new(v1.GroupVersionResource)
		**out = **in
	}
	if in.OperatorRef != nil {
		in, out := &in.OperatorRef, &out.OperatorRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]rbacv1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerSpec.
func (in *DeployerSpec) DeepCopy() *DeployerSpec {
	if in == nil {
		return nil
	}
	out := new(DeployerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerSpecDescriptor) DeepCopyInto(out *DeployerSpecDescriptor) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerSpecDescriptor.
func (in *DeployerSpecDescriptor) DeepCopy() *DeployerSpecDescriptor {
	if in == nil {
		return nil
	}
	out := new(DeployerSpecDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerStatus) DeepCopyInto(out *DeployerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerStatus.
func (in *DeployerStatus) DeepCopy() *DeployerStatus {
	if in == nil {
		return nil
	}
	out := new(DeployerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerStatusDescriptor) DeepCopyInto(out *DeployerStatusDescriptor) {
	*out = *in
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerStatusDescriptor.
func (in *DeployerStatusDescriptor) DeepCopy() *DeployerStatusDescriptor {
	if in == nil {
		return nil
	}
	out := new(DeployerStatusDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRule) DeepCopyInto(out *PlacementRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRule.
func (in *PlacementRule) DeepCopy() *PlacementRule {
	if in == nil {
		return nil
	}
	out := new(PlacementRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlacementRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRuleList) DeepCopyInto(out *PlacementRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PlacementRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRuleList.
func (in *PlacementRuleList) DeepCopy() *PlacementRuleList {
	if in == nil {
		return nil
	}
	out := new(PlacementRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlacementRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRuleSpec) DeepCopyInto(out *PlacementRuleSpec) {
	*out = *in
	if in.DeployerType != nil {
		in, out := &in.DeployerType, &out.DeployerType
		*out = new(string)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.TargetLabels != nil {
		in, out := &in.TargetLabels, &out.TargetLabels
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.DecisionWeight != nil {
		in, out := &in.DecisionWeight, &out.DecisionWeight
		*out = new(int16)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int16)
		**out = **in
	}
	if in.Advisors != nil {
		in, out := &in.Advisors, &out.Advisors
		*out = make([]Advisor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRuleSpec.
func (in *PlacementRuleSpec) DeepCopy() *PlacementRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PlacementRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementRuleStatus) DeepCopyInto(out *PlacementRuleStatus) {
	*out = *in
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
	if in.Candidates != nil {
		in, out := &in.Candidates, &out.Candidates
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Eliminators != nil {
		in, out := &in.Eliminators, &out.Eliminators
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Recommendations != nil {
		in, out := &in.Recommendations, &out.Recommendations
		*out = make(map[string]Recommendation, len(*in))
		for key, val := range *in {
			var outVal []ScoredObjectReference
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Recommendation, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Decisions != nil {
		in, out := &in.Decisions, &out.Decisions
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementRuleStatus.
func (in *PlacementRuleStatus) DeepCopy() *PlacementRuleStatus {
	if in == nil {
		return nil
	}
	out := new(PlacementRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Recommendation) DeepCopyInto(out *Recommendation) {
	{
		in := &in
		*out = make(Recommendation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recommendation.
func (in Recommendation) DeepCopy() Recommendation {
	if in == nil {
		return nil
	}
	out := new(Recommendation)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoredObjectReference) DeepCopyInto(out *ScoredObjectReference) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(int16)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoredObjectReference.
func (in *ScoredObjectReference) DeepCopy() *ScoredObjectReference {
	if in == nil {
		return nil
	}
	out := new(ScoredObjectReference)
	in.DeepCopyInto(out)
	return out
}
