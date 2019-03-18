// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaBackup) DeepCopyInto(out *ArmadaBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaBackup.
func (in *ArmadaBackup) DeepCopy() *ArmadaBackup {
	if in == nil {
		return nil
	}
	out := new(ArmadaBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaBackupList) DeepCopyInto(out *ArmadaBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmadaBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaBackupList.
func (in *ArmadaBackupList) DeepCopy() *ArmadaBackupList {
	if in == nil {
		return nil
	}
	out := new(ArmadaBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaBackupSpec) DeepCopyInto(out *ArmadaBackupSpec) {
	*out = *in
	if in.ArmadaEndpoints != nil {
		in, out := &in.ArmadaEndpoints, &out.ArmadaEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BackupPolicy != nil {
		in, out := &in.BackupPolicy, &out.BackupPolicy
		*out = new(BackupPolicy)
		**out = **in
	}
	in.BackupSource.DeepCopyInto(&out.BackupSource)
	if in.Charts != nil {
		in, out := &in.Charts, &out.Charts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaBackupSpec.
func (in *ArmadaBackupSpec) DeepCopy() *ArmadaBackupSpec {
	if in == nil {
		return nil
	}
	out := new(ArmadaBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaBackupStatus) DeepCopyInto(out *ArmadaBackupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HelmResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaBackupStatus.
func (in *ArmadaBackupStatus) DeepCopy() *ArmadaBackupStatus {
	if in == nil {
		return nil
	}
	out := new(ArmadaBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChart) DeepCopyInto(out *ArmadaChart) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChart.
func (in *ArmadaChart) DeepCopy() *ArmadaChart {
	if in == nil {
		return nil
	}
	out := new(ArmadaChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaChart) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartGroup) DeepCopyInto(out *ArmadaChartGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartGroup.
func (in *ArmadaChartGroup) DeepCopy() *ArmadaChartGroup {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaChartGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartGroupList) DeepCopyInto(out *ArmadaChartGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmadaChartGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartGroupList.
func (in *ArmadaChartGroupList) DeepCopy() *ArmadaChartGroupList {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaChartGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartGroupSpec) DeepCopyInto(out *ArmadaChartGroupSpec) {
	*out = *in
	if in.Charts != nil {
		in, out := &in.Charts, &out.Charts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartGroupSpec.
func (in *ArmadaChartGroupSpec) DeepCopy() *ArmadaChartGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartGroupStatus) DeepCopyInto(out *ArmadaChartGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HelmResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartGroupStatus.
func (in *ArmadaChartGroupStatus) DeepCopy() *ArmadaChartGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartList) DeepCopyInto(out *ArmadaChartList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmadaChart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartList.
func (in *ArmadaChartList) DeepCopy() *ArmadaChartList {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaChartList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartSource) DeepCopyInto(out *ArmadaChartSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartSource.
func (in *ArmadaChartSource) DeepCopy() *ArmadaChartSource {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartSpec) DeepCopyInto(out *ArmadaChartSpec) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(ArmadaChartSource)
		**out = **in
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(ArmadaChartValues)
		**out = **in
	}
	if in.Delete != nil {
		in, out := &in.Delete, &out.Delete
		*out = new(ArmadaDelete)
		**out = **in
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(ArmadaUpgrade)
		(*in).DeepCopyInto(*out)
	}
	if in.Protected != nil {
		in, out := &in.Protected, &out.Protected
		*out = new(ArmadaProtectedRelease)
		**out = **in
	}
	if in.Test != nil {
		in, out := &in.Test, &out.Test
		*out = new(ArmadaTest)
		**out = **in
	}
	if in.Wait != nil {
		in, out := &in.Wait, &out.Wait
		*out = new(ArmadaWait)
		(*in).DeepCopyInto(*out)
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartSpec.
func (in *ArmadaChartSpec) DeepCopy() *ArmadaChartSpec {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartStatus) DeepCopyInto(out *ArmadaChartStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HelmResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartStatus.
func (in *ArmadaChartStatus) DeepCopy() *ArmadaChartStatus {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaChartValues) DeepCopyInto(out *ArmadaChartValues) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaChartValues.
func (in *ArmadaChartValues) DeepCopy() *ArmadaChartValues {
	if in == nil {
		return nil
	}
	out := new(ArmadaChartValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaDelete) DeepCopyInto(out *ArmadaDelete) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaDelete.
func (in *ArmadaDelete) DeepCopy() *ArmadaDelete {
	if in == nil {
		return nil
	}
	out := new(ArmadaDelete)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaHookActionItems) DeepCopyInto(out *ArmadaHookActionItems) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(ArmadaLabels)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaHookActionItems.
func (in *ArmadaHookActionItems) DeepCopy() *ArmadaHookActionItems {
	if in == nil {
		return nil
	}
	out := new(ArmadaHookActionItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaLabels) DeepCopyInto(out *ArmadaLabels) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaLabels.
func (in *ArmadaLabels) DeepCopy() *ArmadaLabels {
	if in == nil {
		return nil
	}
	out := new(ArmadaLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaManifest) DeepCopyInto(out *ArmadaManifest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaManifest.
func (in *ArmadaManifest) DeepCopy() *ArmadaManifest {
	if in == nil {
		return nil
	}
	out := new(ArmadaManifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaManifest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaManifestList) DeepCopyInto(out *ArmadaManifestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmadaManifest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaManifestList.
func (in *ArmadaManifestList) DeepCopy() *ArmadaManifestList {
	if in == nil {
		return nil
	}
	out := new(ArmadaManifestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaManifestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaManifestSpec) DeepCopyInto(out *ArmadaManifestSpec) {
	*out = *in
	if in.ChartGroups != nil {
		in, out := &in.ChartGroups, &out.ChartGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaManifestSpec.
func (in *ArmadaManifestSpec) DeepCopy() *ArmadaManifestSpec {
	if in == nil {
		return nil
	}
	out := new(ArmadaManifestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaManifestStatus) DeepCopyInto(out *ArmadaManifestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HelmResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaManifestStatus.
func (in *ArmadaManifestStatus) DeepCopy() *ArmadaManifestStatus {
	if in == nil {
		return nil
	}
	out := new(ArmadaManifestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaProtectedRelease) DeepCopyInto(out *ArmadaProtectedRelease) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaProtectedRelease.
func (in *ArmadaProtectedRelease) DeepCopy() *ArmadaProtectedRelease {
	if in == nil {
		return nil
	}
	out := new(ArmadaProtectedRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaRestore) DeepCopyInto(out *ArmadaRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaRestore.
func (in *ArmadaRestore) DeepCopy() *ArmadaRestore {
	if in == nil {
		return nil
	}
	out := new(ArmadaRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaRestoreList) DeepCopyInto(out *ArmadaRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArmadaRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaRestoreList.
func (in *ArmadaRestoreList) DeepCopy() *ArmadaRestoreList {
	if in == nil {
		return nil
	}
	out := new(ArmadaRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArmadaRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaRestoreSpec) DeepCopyInto(out *ArmadaRestoreSpec) {
	*out = *in
	in.RestoreSource.DeepCopyInto(&out.RestoreSource)
	if in.Charts != nil {
		in, out := &in.Charts, &out.Charts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaRestoreSpec.
func (in *ArmadaRestoreSpec) DeepCopy() *ArmadaRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(ArmadaRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaRestoreStatus) DeepCopyInto(out *ArmadaRestoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HelmResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaRestoreStatus.
func (in *ArmadaRestoreStatus) DeepCopy() *ArmadaRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(ArmadaRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaTest) DeepCopyInto(out *ArmadaTest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaTest.
func (in *ArmadaTest) DeepCopy() *ArmadaTest {
	if in == nil {
		return nil
	}
	out := new(ArmadaTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaUpgrade) DeepCopyInto(out *ArmadaUpgrade) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(ArmadaUpgradeOptions)
		**out = **in
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(ArmadaUpgradePost)
		(*in).DeepCopyInto(*out)
	}
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(ArmadaUpgradePre)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaUpgrade.
func (in *ArmadaUpgrade) DeepCopy() *ArmadaUpgrade {
	if in == nil {
		return nil
	}
	out := new(ArmadaUpgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaUpgradeOptions) DeepCopyInto(out *ArmadaUpgradeOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaUpgradeOptions.
func (in *ArmadaUpgradeOptions) DeepCopy() *ArmadaUpgradeOptions {
	if in == nil {
		return nil
	}
	out := new(ArmadaUpgradeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaUpgradePost) DeepCopyInto(out *ArmadaUpgradePost) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = make([]*ArmadaHookActionItems, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ArmadaHookActionItems)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaUpgradePost.
func (in *ArmadaUpgradePost) DeepCopy() *ArmadaUpgradePost {
	if in == nil {
		return nil
	}
	out := new(ArmadaUpgradePost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaUpgradePre) DeepCopyInto(out *ArmadaUpgradePre) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = make([]*ArmadaHookActionItems, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ArmadaHookActionItems)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Delete != nil {
		in, out := &in.Delete, &out.Delete
		*out = make([]*ArmadaHookActionItems, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ArmadaHookActionItems)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Update != nil {
		in, out := &in.Update, &out.Update
		*out = make([]*ArmadaHookActionItems, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ArmadaHookActionItems)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaUpgradePre.
func (in *ArmadaUpgradePre) DeepCopy() *ArmadaUpgradePre {
	if in == nil {
		return nil
	}
	out := new(ArmadaUpgradePre)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaWait) DeepCopyInto(out *ArmadaWait) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(ArmadaLabels)
		(*in).DeepCopyInto(*out)
	}
	if in.Native != nil {
		in, out := &in.Native, &out.Native
		*out = new(ArmadaWaitNative)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*ArmadaWaitResourcesItems, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ArmadaWaitResourcesItems)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaWait.
func (in *ArmadaWait) DeepCopy() *ArmadaWait {
	if in == nil {
		return nil
	}
	out := new(ArmadaWait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaWaitNative) DeepCopyInto(out *ArmadaWaitNative) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaWaitNative.
func (in *ArmadaWaitNative) DeepCopy() *ArmadaWaitNative {
	if in == nil {
		return nil
	}
	out := new(ArmadaWaitNative)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArmadaWaitResourcesItems) DeepCopyInto(out *ArmadaWaitResourcesItems) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(ArmadaLabels)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArmadaWaitResourcesItems.
func (in *ArmadaWaitResourcesItems) DeepCopy() *ArmadaWaitResourcesItems {
	if in == nil {
		return nil
	}
	out := new(ArmadaWaitResourcesItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicy) DeepCopyInto(out *BackupPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicy.
func (in *BackupPolicy) DeepCopy() *BackupPolicy {
	if in == nil {
		return nil
	}
	out := new(BackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSource) DeepCopyInto(out *BackupSource) {
	*out = *in
	if in.Offsite != nil {
		in, out := &in.Offsite, &out.Offsite
		*out = new(OffsiteBackupSource)
		**out = **in
	}
	if in.Ceph != nil {
		in, out := &in.Ceph, &out.Ceph
		*out = new(CephBackupSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSource.
func (in *BackupSource) DeepCopy() *BackupSource {
	if in == nil {
		return nil
	}
	out := new(BackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephBackupSource) DeepCopyInto(out *CephBackupSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephBackupSource.
func (in *CephBackupSource) DeepCopy() *CephBackupSource {
	if in == nil {
		return nil
	}
	out := new(CephBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephRestoreSource) DeepCopyInto(out *CephRestoreSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephRestoreSource.
func (in *CephRestoreSource) DeepCopy() *CephRestoreSource {
	if in == nil {
		return nil
	}
	out := new(CephRestoreSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevision) DeepCopyInto(out *ControllerRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Data.DeepCopyInto(&out.Data)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevision.
func (in *ControllerRevision) DeepCopy() *ControllerRevision {
	if in == nil {
		return nil
	}
	out := new(ControllerRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevisionList) DeepCopyInto(out *ControllerRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevisionList.
func (in *ControllerRevisionList) DeepCopy() *ControllerRevisionList {
	if in == nil {
		return nil
	}
	out := new(ControllerRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmResourceCondition) DeepCopyInto(out *HelmResourceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmResourceCondition.
func (in *HelmResourceCondition) DeepCopy() *HelmResourceCondition {
	if in == nil {
		return nil
	}
	out := new(HelmResourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmResourceConditionListHelper) DeepCopyInto(out *HelmResourceConditionListHelper) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmResourceConditionListHelper.
func (in *HelmResourceConditionListHelper) DeepCopy() *HelmResourceConditionListHelper {
	if in == nil {
		return nil
	}
	out := new(HelmResourceConditionListHelper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmResourceStatusHelper) DeepCopyInto(out *HelmResourceStatusHelper) {
	*out = *in
	if in.Cond != nil {
		in, out := &in.Cond, &out.Cond
		*out = new(HelmResourceCondition)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmResourceStatusHelper.
func (in *HelmResourceStatusHelper) DeepCopy() *HelmResourceStatusHelper {
	if in == nil {
		return nil
	}
	out := new(HelmResourceStatusHelper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OffsiteBackupSource) DeepCopyInto(out *OffsiteBackupSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OffsiteBackupSource.
func (in *OffsiteBackupSource) DeepCopy() *OffsiteBackupSource {
	if in == nil {
		return nil
	}
	out := new(OffsiteBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OffsiteRestoreSource) DeepCopyInto(out *OffsiteRestoreSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OffsiteRestoreSource.
func (in *OffsiteRestoreSource) DeepCopy() *OffsiteRestoreSource {
	if in == nil {
		return nil
	}
	out := new(OffsiteRestoreSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreSource) DeepCopyInto(out *RestoreSource) {
	*out = *in
	if in.Offsite != nil {
		in, out := &in.Offsite, &out.Offsite
		*out = new(OffsiteRestoreSource)
		**out = **in
	}
	if in.Ceph != nil {
		in, out := &in.Ceph, &out.Ceph
		*out = new(CephRestoreSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreSource.
func (in *RestoreSource) DeepCopy() *RestoreSource {
	if in == nil {
		return nil
	}
	out := new(RestoreSource)
	in.DeepCopyInto(out)
	return out
}
