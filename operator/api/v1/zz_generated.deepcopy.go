//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseCluster) DeepCopyInto(out *HbaseCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseCluster.
func (in *HbaseCluster) DeepCopy() *HbaseCluster {
	if in == nil {
		return nil
	}
	out := new(HbaseCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HbaseCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterConfiguration) DeepCopyInto(out *HbaseClusterConfiguration) {
	*out = *in
	if in.HbaseConfig != nil {
		in, out := &in.HbaseConfig, &out.HbaseConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HadoopConfig != nil {
		in, out := &in.HadoopConfig, &out.HadoopConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HbaseTenantConfig != nil {
		in, out := &in.HbaseTenantConfig, &out.HbaseTenantConfig
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.HadoopTenantConfig != nil {
		in, out := &in.HadoopTenantConfig, &out.HadoopTenantConfig
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterConfiguration.
func (in *HbaseClusterConfiguration) DeepCopy() *HbaseClusterConfiguration {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterContainer) DeepCopyInto(out *HbaseClusterContainer) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]HbaseClusterContainerPort, len(*in))
		copy(*out, *in)
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]HbaseClusterVolumeMount, len(*in))
		copy(*out, *in)
	}
	out.SecurityContext = in.SecurityContext
	in.LivenessProbe.DeepCopyInto(&out.LivenessProbe)
	in.ReadinessProbe.DeepCopyInto(&out.ReadinessProbe)
	in.StartupProbe.DeepCopyInto(&out.StartupProbe)
	in.Lifecycle.DeepCopyInto(&out.Lifecycle)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterContainer.
func (in *HbaseClusterContainer) DeepCopy() *HbaseClusterContainer {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterContainerPort) DeepCopyInto(out *HbaseClusterContainerPort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterContainerPort.
func (in *HbaseClusterContainerPort) DeepCopy() *HbaseClusterContainerPort {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterContainerPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterDeployment) DeepCopyInto(out *HbaseClusterDeployment) {
	*out = *in
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
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]HbaseClusterContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SideCarContainers != nil {
		in, out := &in.SideCarContainers, &out.SideCarContainers
		*out = make([]HbaseClusterSideCarContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]HbaseClusterInitContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaims != nil {
		in, out := &in.VolumeClaims, &out.VolumeClaims
		*out = make([]HbaseClusterVolumeClaim, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]HbaseClusterVolume, len(*in))
		copy(*out, *in)
	}
	if in.DNSConfig != nil {
		in, out := &in.DNSConfig, &out.DNSConfig
		*out = new(corev1.PodDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]corev1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterDeployment.
func (in *HbaseClusterDeployment) DeepCopy() *HbaseClusterDeployment {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterDeployments) DeepCopyInto(out *HbaseClusterDeployments) {
	*out = *in
	in.Zookeeper.DeepCopyInto(&out.Zookeeper)
	in.Journalnode.DeepCopyInto(&out.Journalnode)
	in.Namenode.DeepCopyInto(&out.Namenode)
	in.Datanode.DeepCopyInto(&out.Datanode)
	in.Hmaster.DeepCopyInto(&out.Hmaster)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterDeployments.
func (in *HbaseClusterDeployments) DeepCopy() *HbaseClusterDeployments {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterDeployments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterInitContainer) DeepCopyInto(out *HbaseClusterInitContainer) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]HbaseClusterVolumeMount, len(*in))
		copy(*out, *in)
	}
	out.SecurityContext = in.SecurityContext
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterInitContainer.
func (in *HbaseClusterInitContainer) DeepCopy() *HbaseClusterInitContainer {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterInitContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterLifecycle) DeepCopyInto(out *HbaseClusterLifecycle) {
	*out = *in
	if in.PostStart != nil {
		in, out := &in.PostStart, &out.PostStart
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreStop != nil {
		in, out := &in.PreStop, &out.PreStop
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterLifecycle.
func (in *HbaseClusterLifecycle) DeepCopy() *HbaseClusterLifecycle {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterLifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterList) DeepCopyInto(out *HbaseClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HbaseCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterList.
func (in *HbaseClusterList) DeepCopy() *HbaseClusterList {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HbaseClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterProbe) DeepCopyInto(out *HbaseClusterProbe) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterProbe.
func (in *HbaseClusterProbe) DeepCopy() *HbaseClusterProbe {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterSecurity) DeepCopyInto(out *HbaseClusterSecurity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterSecurity.
func (in *HbaseClusterSecurity) DeepCopy() *HbaseClusterSecurity {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterSideCarContainer) DeepCopyInto(out *HbaseClusterSideCarContainer) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.SecurityContext = in.SecurityContext
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]HbaseClusterVolumeMount, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterSideCarContainer.
func (in *HbaseClusterSideCarContainer) DeepCopy() *HbaseClusterSideCarContainer {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterSideCarContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterSpec) DeepCopyInto(out *HbaseClusterSpec) {
	*out = *in
	in.Deployments.DeepCopyInto(&out.Deployments)
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.TenantNamespaces != nil {
		in, out := &in.TenantNamespaces, &out.TenantNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterSpec.
func (in *HbaseClusterSpec) DeepCopy() *HbaseClusterSpec {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterStatus) DeepCopyInto(out *HbaseClusterStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterStatus.
func (in *HbaseClusterStatus) DeepCopy() *HbaseClusterStatus {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterVolume) DeepCopyInto(out *HbaseClusterVolume) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterVolume.
func (in *HbaseClusterVolume) DeepCopy() *HbaseClusterVolume {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterVolumeClaim) DeepCopyInto(out *HbaseClusterVolumeClaim) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterVolumeClaim.
func (in *HbaseClusterVolumeClaim) DeepCopy() *HbaseClusterVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseClusterVolumeMount) DeepCopyInto(out *HbaseClusterVolumeMount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseClusterVolumeMount.
func (in *HbaseClusterVolumeMount) DeepCopy() *HbaseClusterVolumeMount {
	if in == nil {
		return nil
	}
	out := new(HbaseClusterVolumeMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseStandalone) DeepCopyInto(out *HbaseStandalone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseStandalone.
func (in *HbaseStandalone) DeepCopy() *HbaseStandalone {
	if in == nil {
		return nil
	}
	out := new(HbaseStandalone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HbaseStandalone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseStandaloneConfiguration) DeepCopyInto(out *HbaseStandaloneConfiguration) {
	*out = *in
	if in.HbaseConfig != nil {
		in, out := &in.HbaseConfig, &out.HbaseConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseStandaloneConfiguration.
func (in *HbaseStandaloneConfiguration) DeepCopy() *HbaseStandaloneConfiguration {
	if in == nil {
		return nil
	}
	out := new(HbaseStandaloneConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseStandaloneList) DeepCopyInto(out *HbaseStandaloneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HbaseStandalone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseStandaloneList.
func (in *HbaseStandaloneList) DeepCopy() *HbaseStandaloneList {
	if in == nil {
		return nil
	}
	out := new(HbaseStandaloneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HbaseStandaloneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseStandaloneSpec) DeepCopyInto(out *HbaseStandaloneSpec) {
	*out = *in
	in.Standalone.DeepCopyInto(&out.Standalone)
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseStandaloneSpec.
func (in *HbaseStandaloneSpec) DeepCopy() *HbaseStandaloneSpec {
	if in == nil {
		return nil
	}
	out := new(HbaseStandaloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseStandaloneStatus) DeepCopyInto(out *HbaseStandaloneStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseStandaloneStatus.
func (in *HbaseStandaloneStatus) DeepCopy() *HbaseStandaloneStatus {
	if in == nil {
		return nil
	}
	out := new(HbaseStandaloneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseTenant) DeepCopyInto(out *HbaseTenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseTenant.
func (in *HbaseTenant) DeepCopy() *HbaseTenant {
	if in == nil {
		return nil
	}
	out := new(HbaseTenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HbaseTenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseTenantList) DeepCopyInto(out *HbaseTenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HbaseTenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseTenantList.
func (in *HbaseTenantList) DeepCopy() *HbaseTenantList {
	if in == nil {
		return nil
	}
	out := new(HbaseTenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HbaseTenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseTenantSpec) DeepCopyInto(out *HbaseTenantSpec) {
	*out = *in
	in.Datanode.DeepCopyInto(&out.Datanode)
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseTenantSpec.
func (in *HbaseTenantSpec) DeepCopy() *HbaseTenantSpec {
	if in == nil {
		return nil
	}
	out := new(HbaseTenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HbaseTenantStatus) DeepCopyInto(out *HbaseTenantStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HbaseTenantStatus.
func (in *HbaseTenantStatus) DeepCopy() *HbaseTenantStatus {
	if in == nil {
		return nil
	}
	out := new(HbaseTenantStatus)
	in.DeepCopyInto(out)
	return out
}
