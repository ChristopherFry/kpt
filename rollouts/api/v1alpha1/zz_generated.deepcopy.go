//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRef) DeepCopyInto(out *ClusterRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRef.
func (in *ClusterRef) DeepCopy() *ClusterRef {
	if in == nil {
		return nil
	}
	out := new(ClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	out.PackageStatus = in.PackageStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTargetSelector) DeepCopyInto(out *ClusterTargetSelector) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTargetSelector.
func (in *ClusterTargetSelector) DeepCopy() *ClusterTargetSelector {
	if in == nil {
		return nil
	}
	out := new(ClusterTargetSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubSelector) DeepCopyInto(out *GitHubSelector) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubSelector.
func (in *GitHubSelector) DeepCopy() *GitHubSelector {
	if in == nil {
		return nil
	}
	out := new(GitHubSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubSource) DeepCopyInto(out *GitHubSource) {
	*out = *in
	out.Selector = in.Selector
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubSource.
func (in *GitHubSource) DeepCopy() *GitHubSource {
	if in == nil {
		return nil
	}
	out := new(GitHubSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitInfo) DeepCopyInto(out *GitInfo) {
	*out = *in
	out.Period = in.Period
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitInfo.
func (in *GitInfo) DeepCopy() *GitInfo {
	if in == nil {
		return nil
	}
	out := new(GitInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageStatus) DeepCopyInto(out *PackageStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageStatus.
func (in *PackageStatus) DeepCopy() *PackageStatus {
	if in == nil {
		return nil
	}
	out := new(PackageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageToClusterMatcher) DeepCopyInto(out *PackageToClusterMatcher) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageToClusterMatcher.
func (in *PackageToClusterMatcher) DeepCopy() *PackageToClusterMatcher {
	if in == nil {
		return nil
	}
	out := new(PackageToClusterMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackagesConfig) DeepCopyInto(out *PackagesConfig) {
	*out = *in
	out.GitHub = in.GitHub
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackagesConfig.
func (in *PackagesConfig) DeepCopy() *PackagesConfig {
	if in == nil {
		return nil
	}
	out := new(PackagesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PauseAfterWave) DeepCopyInto(out *PauseAfterWave) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PauseAfterWave.
func (in *PauseAfterWave) DeepCopy() *PauseAfterWave {
	if in == nil {
		return nil
	}
	out := new(PauseAfterWave)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutStrategy) DeepCopyInto(out *ProgressiveRolloutStrategy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutStrategy.
func (in *ProgressiveRolloutStrategy) DeepCopy() *ProgressiveRolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProgressiveRolloutStrategy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutStrategyList) DeepCopyInto(out *ProgressiveRolloutStrategyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProgressiveRolloutStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutStrategyList.
func (in *ProgressiveRolloutStrategyList) DeepCopy() *ProgressiveRolloutStrategyList {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutStrategyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProgressiveRolloutStrategyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutStrategySpec) DeepCopyInto(out *ProgressiveRolloutStrategySpec) {
	*out = *in
	if in.Waves != nil {
		in, out := &in.Waves, &out.Waves
		*out = make([]Wave, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutStrategySpec.
func (in *ProgressiveRolloutStrategySpec) DeepCopy() *ProgressiveRolloutStrategySpec {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutStrategySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutStrategyStatus) DeepCopyInto(out *ProgressiveRolloutStrategyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutStrategyStatus.
func (in *ProgressiveRolloutStrategyStatus) DeepCopy() *ProgressiveRolloutStrategyStatus {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutStrategyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRootSync) DeepCopyInto(out *RemoteRootSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRootSync.
func (in *RemoteRootSync) DeepCopy() *RemoteRootSync {
	if in == nil {
		return nil
	}
	out := new(RemoteRootSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteRootSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRootSyncList) DeepCopyInto(out *RemoteRootSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RemoteRootSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRootSyncList.
func (in *RemoteRootSyncList) DeepCopy() *RemoteRootSyncList {
	if in == nil {
		return nil
	}
	out := new(RemoteRootSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteRootSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRootSyncSpec) DeepCopyInto(out *RemoteRootSyncSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(RootSyncInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRootSyncSpec.
func (in *RemoteRootSyncSpec) DeepCopy() *RemoteRootSyncSpec {
	if in == nil {
		return nil
	}
	out := new(RemoteRootSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRootSyncStatus) DeepCopyInto(out *RemoteRootSyncStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRootSyncStatus.
func (in *RemoteRootSyncStatus) DeepCopy() *RemoteRootSyncStatus {
	if in == nil {
		return nil
	}
	out := new(RemoteRootSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rollout) DeepCopyInto(out *Rollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rollout.
func (in *Rollout) DeepCopy() *Rollout {
	if in == nil {
		return nil
	}
	out := new(Rollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutList) DeepCopyInto(out *RolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutList.
func (in *RolloutList) DeepCopy() *RolloutList {
	if in == nil {
		return nil
	}
	out := new(RolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutSpec) DeepCopyInto(out *RolloutSpec) {
	*out = *in
	out.Packages = in.Packages
	in.Targets.DeepCopyInto(&out.Targets)
	out.PackageToTargetMatcher = in.PackageToTargetMatcher
	in.Strategy.DeepCopyInto(&out.Strategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutSpec.
func (in *RolloutSpec) DeepCopy() *RolloutSpec {
	if in == nil {
		return nil
	}
	out := new(RolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStatus) DeepCopyInto(out *RolloutStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterStatuses != nil {
		in, out := &in.ClusterStatuses, &out.ClusterStatuses
		*out = make([]ClusterStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStatus.
func (in *RolloutStatus) DeepCopy() *RolloutStatus {
	if in == nil {
		return nil
	}
	out := new(RolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStrategy) DeepCopyInto(out *RolloutStrategy) {
	*out = *in
	if in.AllAtOnce != nil {
		in, out := &in.AllAtOnce, &out.AllAtOnce
		*out = new(StrategyAllAtOnce)
		**out = **in
	}
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(StrategyRollingUpdate)
		**out = **in
	}
	if in.Progressive != nil {
		in, out := &in.Progressive, &out.Progressive
		*out = new(StrategyProgressive)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStrategy.
func (in *RolloutStrategy) DeepCopy() *RolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(RolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootSyncInfo) DeepCopyInto(out *RootSyncInfo) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(RootSyncSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootSyncInfo.
func (in *RootSyncInfo) DeepCopy() *RootSyncInfo {
	if in == nil {
		return nil
	}
	out := new(RootSyncInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootSyncSpec) DeepCopyInto(out *RootSyncSpec) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitInfo)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootSyncSpec.
func (in *RootSyncSpec) DeepCopy() *RootSyncSpec {
	if in == nil {
		return nil
	}
	out := new(RootSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference.
func (in *SecretReference) DeepCopy() *SecretReference {
	if in == nil {
		return nil
	}
	out := new(SecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrategyAllAtOnce) DeepCopyInto(out *StrategyAllAtOnce) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrategyAllAtOnce.
func (in *StrategyAllAtOnce) DeepCopy() *StrategyAllAtOnce {
	if in == nil {
		return nil
	}
	out := new(StrategyAllAtOnce)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrategyProgressive) DeepCopyInto(out *StrategyProgressive) {
	*out = *in
	out.PauseAfterWave = in.PauseAfterWave
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrategyProgressive.
func (in *StrategyProgressive) DeepCopy() *StrategyProgressive {
	if in == nil {
		return nil
	}
	out := new(StrategyProgressive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StrategyRollingUpdate) DeepCopyInto(out *StrategyRollingUpdate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StrategyRollingUpdate.
func (in *StrategyRollingUpdate) DeepCopy() *StrategyRollingUpdate {
	if in == nil {
		return nil
	}
	out := new(StrategyRollingUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wave) DeepCopyInto(out *Wave) {
	*out = *in
	in.Targets.DeepCopyInto(&out.Targets)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wave.
func (in *Wave) DeepCopy() *Wave {
	if in == nil {
		return nil
	}
	out := new(Wave)
	in.DeepCopyInto(out)
	return out
}
