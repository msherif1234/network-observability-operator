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

package v1alpha1

import (
	"fmt"

	"github.com/netobserv/network-observability-operator/api/v1beta1"
	utilconversion "github.com/netobserv/network-observability-operator/pkg/conversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this v1alpha1 FlowCollector to its v1beta1 equivalent (the conversion Hub)
// https://book.kubebuilder.io/multiversion-tutorial/conversion.html
func (r *FlowCollector) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.FlowCollector)

	if err := Convert_v1alpha1_FlowCollector_To_v1beta1_FlowCollector(r, dst, nil); err != nil {
		return fmt.Errorf("copying v1alpha1.FlowCollector into v1beta1.FlowCollector: %w", err)
	}
	dst.Status.Conditions = make([]v1.Condition, len(r.Status.Conditions))
	copy(dst.Status.Conditions, r.Status.Conditions)

	// Manually restore data.
	restored := &v1beta1.FlowCollector{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	dst.Spec.Processor.LogTypes = restored.Spec.Processor.LogTypes

	if restored.Spec.Processor.ConversationHeartbeatInterval != nil {
		dst.Spec.Processor.ConversationHeartbeatInterval = restored.Spec.Processor.ConversationHeartbeatInterval
	}

	if restored.Spec.Processor.ConversationEndTimeout != nil {
		dst.Spec.Processor.ConversationEndTimeout = restored.Spec.Processor.ConversationEndTimeout
	}

	if restored.Spec.Processor.Metrics.DisableAlerts != nil {
		dst.Spec.Processor.Metrics.DisableAlerts = restored.Spec.Processor.Metrics.DisableAlerts
	}

	dst.Spec.Loki.StatusTLS = restored.Spec.Loki.StatusTLS
	dst.Spec.Agent.EBPF.ProfilePort = restored.Spec.Agent.EBPF.ProfilePort

	return nil
}

// ConvertFrom converts the hub version v1beta1 FlowCollector object to v1alpha1
func (r *FlowCollector) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.FlowCollector)

	if err := Convert_v1beta1_FlowCollector_To_v1alpha1_FlowCollector(src, r, nil); err != nil {
		return fmt.Errorf("copying v1beta1.FlowCollector into v1alpha1.FlowCollector: %w", err)
	}
	r.Status.Conditions = make([]v1.Condition, len(src.Status.Conditions))
	copy(r.Status.Conditions, src.Status.Conditions)

	// Preserve Hub data on down-conversion except for metadata
	if err := utilconversion.MarshalData(src, r); err != nil {
		return err
	}
	return nil
}

func (r *FlowCollectorList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.FlowCollectorList)
	return Convert_v1alpha1_FlowCollectorList_To_v1beta1_FlowCollectorList(r, dst, nil)
}

func (r *FlowCollectorList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.FlowCollectorList)
	return Convert_v1beta1_FlowCollectorList_To_v1alpha1_FlowCollectorList(src, r, nil)
}

// This function need to be manually created because conversion-gen not able to create it intentionally because
// we have new defined fields in v1beta1 not in v1alpha1
// nolint:golint,stylecheck,revive
func Convert_v1beta1_FlowCollectorFLP_To_v1alpha1_FlowCollectorFLP(in *v1beta1.FlowCollectorFLP, out *FlowCollectorFLP, s apiconversion.Scope) error {
	return autoConvert_v1beta1_FlowCollectorFLP_To_v1alpha1_FlowCollectorFLP(in, out, s)
}

// This function need to be manually created because conversion-gen not able to create it intentionally because
// we have new defined fields in v1beta1 not in v1alpha1
// nolint:golint,stylecheck,revive
func Convert_v1beta1_FLPMetrics_To_v1alpha1_FLPMetrics(in *v1beta1.FLPMetrics, out *FLPMetrics, s apiconversion.Scope) error {
	return autoConvert_v1beta1_FLPMetrics_To_v1alpha1_FLPMetrics(in, out, s)
}

// This function need to be manually created because conversion-gen not able to create it intentionally because
// we have new defined fields in v1beta1 not in v1alpha1
// nolint:golint,stylecheck,revive
func Convert_v1beta1_FlowCollectorLoki_To_v1alpha1_FlowCollectorLoki(in *v1beta1.FlowCollectorLoki, out *FlowCollectorLoki, s apiconversion.Scope) error {
	return autoConvert_v1beta1_FlowCollectorLoki_To_v1alpha1_FlowCollectorLoki(in, out, s)
}

// This function need to be manually created because conversion-gen not able to create it intentionally because
// we have new defined fields in v1beta1 not in v1alpha1
// nolint:golint,stylecheck,revive
func Convert_v1beta1_FlowCollectorEBPF_To_v1alpha1_FlowCollectorEBPF(in *v1beta1.FlowCollectorEBPF, out *FlowCollectorEBPF, s apiconversion.Scope) error {
	return autoConvert_v1beta1_FlowCollectorEBPF_To_v1alpha1_FlowCollectorEBPF(in, out, s)
}
