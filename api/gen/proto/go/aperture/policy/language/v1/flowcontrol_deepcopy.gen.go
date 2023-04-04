// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package languagev1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using FlowSelector within kubernetes types, where deepcopy-gen is used.
func (in *FlowSelector) DeepCopyInto(out *FlowSelector) {
	p := proto.Clone(in).(*FlowSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSelector. Required by controller-gen.
func (in *FlowSelector) DeepCopy() *FlowSelector {
	if in == nil {
		return nil
	}
	out := new(FlowSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FlowSelector. Required by controller-gen.
func (in *FlowSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServiceSelector within kubernetes types, where deepcopy-gen is used.
func (in *ServiceSelector) DeepCopyInto(out *ServiceSelector) {
	p := proto.Clone(in).(*ServiceSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSelector. Required by controller-gen.
func (in *ServiceSelector) DeepCopy() *ServiceSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSelector. Required by controller-gen.
func (in *ServiceSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FlowMatcher within kubernetes types, where deepcopy-gen is used.
func (in *FlowMatcher) DeepCopyInto(out *FlowMatcher) {
	p := proto.Clone(in).(*FlowMatcher)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowMatcher. Required by controller-gen.
func (in *FlowMatcher) DeepCopy() *FlowMatcher {
	if in == nil {
		return nil
	}
	out := new(FlowMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FlowMatcher. Required by controller-gen.
func (in *FlowMatcher) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FlowControlResources within kubernetes types, where deepcopy-gen is used.
func (in *FlowControlResources) DeepCopyInto(out *FlowControlResources) {
	p := proto.Clone(in).(*FlowControlResources)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowControlResources. Required by controller-gen.
func (in *FlowControlResources) DeepCopy() *FlowControlResources {
	if in == nil {
		return nil
	}
	out := new(FlowControlResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FlowControlResources. Required by controller-gen.
func (in *FlowControlResources) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FluxMeter within kubernetes types, where deepcopy-gen is used.
func (in *FluxMeter) DeepCopyInto(out *FluxMeter) {
	p := proto.Clone(in).(*FluxMeter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter. Required by controller-gen.
func (in *FluxMeter) DeepCopy() *FluxMeter {
	if in == nil {
		return nil
	}
	out := new(FluxMeter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter. Required by controller-gen.
func (in *FluxMeter) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FluxMeter_StaticBuckets within kubernetes types, where deepcopy-gen is used.
func (in *FluxMeter_StaticBuckets) DeepCopyInto(out *FluxMeter_StaticBuckets) {
	p := proto.Clone(in).(*FluxMeter_StaticBuckets)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter_StaticBuckets. Required by controller-gen.
func (in *FluxMeter_StaticBuckets) DeepCopy() *FluxMeter_StaticBuckets {
	if in == nil {
		return nil
	}
	out := new(FluxMeter_StaticBuckets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter_StaticBuckets. Required by controller-gen.
func (in *FluxMeter_StaticBuckets) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FluxMeter_LinearBuckets within kubernetes types, where deepcopy-gen is used.
func (in *FluxMeter_LinearBuckets) DeepCopyInto(out *FluxMeter_LinearBuckets) {
	p := proto.Clone(in).(*FluxMeter_LinearBuckets)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter_LinearBuckets. Required by controller-gen.
func (in *FluxMeter_LinearBuckets) DeepCopy() *FluxMeter_LinearBuckets {
	if in == nil {
		return nil
	}
	out := new(FluxMeter_LinearBuckets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter_LinearBuckets. Required by controller-gen.
func (in *FluxMeter_LinearBuckets) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FluxMeter_ExponentialBuckets within kubernetes types, where deepcopy-gen is used.
func (in *FluxMeter_ExponentialBuckets) DeepCopyInto(out *FluxMeter_ExponentialBuckets) {
	p := proto.Clone(in).(*FluxMeter_ExponentialBuckets)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter_ExponentialBuckets. Required by controller-gen.
func (in *FluxMeter_ExponentialBuckets) DeepCopy() *FluxMeter_ExponentialBuckets {
	if in == nil {
		return nil
	}
	out := new(FluxMeter_ExponentialBuckets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter_ExponentialBuckets. Required by controller-gen.
func (in *FluxMeter_ExponentialBuckets) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FluxMeter_ExponentialBucketsRange within kubernetes types, where deepcopy-gen is used.
func (in *FluxMeter_ExponentialBucketsRange) DeepCopyInto(out *FluxMeter_ExponentialBucketsRange) {
	p := proto.Clone(in).(*FluxMeter_ExponentialBucketsRange)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter_ExponentialBucketsRange. Required by controller-gen.
func (in *FluxMeter_ExponentialBucketsRange) DeepCopy() *FluxMeter_ExponentialBucketsRange {
	if in == nil {
		return nil
	}
	out := new(FluxMeter_ExponentialBucketsRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FluxMeter_ExponentialBucketsRange. Required by controller-gen.
func (in *FluxMeter_ExponentialBucketsRange) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Classifier within kubernetes types, where deepcopy-gen is used.
func (in *Classifier) DeepCopyInto(out *Classifier) {
	p := proto.Clone(in).(*Classifier)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Classifier. Required by controller-gen.
func (in *Classifier) DeepCopy() *Classifier {
	if in == nil {
		return nil
	}
	out := new(Classifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Classifier. Required by controller-gen.
func (in *Classifier) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule within kubernetes types, where deepcopy-gen is used.
func (in *Rule) DeepCopyInto(out *Rule) {
	p := proto.Clone(in).(*Rule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule. Required by controller-gen.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule. Required by controller-gen.
func (in *Rule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule_Rego within kubernetes types, where deepcopy-gen is used.
func (in *Rule_Rego) DeepCopyInto(out *Rule_Rego) {
	p := proto.Clone(in).(*Rule_Rego)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule_Rego. Required by controller-gen.
func (in *Rule_Rego) DeepCopy() *Rule_Rego {
	if in == nil {
		return nil
	}
	out := new(Rule_Rego)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule_Rego. Required by controller-gen.
func (in *Rule_Rego) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rego within kubernetes types, where deepcopy-gen is used.
func (in *Rego) DeepCopyInto(out *Rego) {
	p := proto.Clone(in).(*Rego)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rego. Required by controller-gen.
func (in *Rego) DeepCopy() *Rego {
	if in == nil {
		return nil
	}
	out := new(Rego)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rego. Required by controller-gen.
func (in *Rego) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rego_LabelProperties within kubernetes types, where deepcopy-gen is used.
func (in *Rego_LabelProperties) DeepCopyInto(out *Rego_LabelProperties) {
	p := proto.Clone(in).(*Rego_LabelProperties)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rego_LabelProperties. Required by controller-gen.
func (in *Rego_LabelProperties) DeepCopy() *Rego_LabelProperties {
	if in == nil {
		return nil
	}
	out := new(Rego_LabelProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rego_LabelProperties. Required by controller-gen.
func (in *Rego_LabelProperties) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Extractor within kubernetes types, where deepcopy-gen is used.
func (in *Extractor) DeepCopyInto(out *Extractor) {
	p := proto.Clone(in).(*Extractor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extractor. Required by controller-gen.
func (in *Extractor) DeepCopy() *Extractor {
	if in == nil {
		return nil
	}
	out := new(Extractor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Extractor. Required by controller-gen.
func (in *Extractor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using JSONExtractor within kubernetes types, where deepcopy-gen is used.
func (in *JSONExtractor) DeepCopyInto(out *JSONExtractor) {
	p := proto.Clone(in).(*JSONExtractor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONExtractor. Required by controller-gen.
func (in *JSONExtractor) DeepCopy() *JSONExtractor {
	if in == nil {
		return nil
	}
	out := new(JSONExtractor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new JSONExtractor. Required by controller-gen.
func (in *JSONExtractor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AddressExtractor within kubernetes types, where deepcopy-gen is used.
func (in *AddressExtractor) DeepCopyInto(out *AddressExtractor) {
	p := proto.Clone(in).(*AddressExtractor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressExtractor. Required by controller-gen.
func (in *AddressExtractor) DeepCopy() *AddressExtractor {
	if in == nil {
		return nil
	}
	out := new(AddressExtractor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AddressExtractor. Required by controller-gen.
func (in *AddressExtractor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using JWTExtractor within kubernetes types, where deepcopy-gen is used.
func (in *JWTExtractor) DeepCopyInto(out *JWTExtractor) {
	p := proto.Clone(in).(*JWTExtractor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTExtractor. Required by controller-gen.
func (in *JWTExtractor) DeepCopy() *JWTExtractor {
	if in == nil {
		return nil
	}
	out := new(JWTExtractor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new JWTExtractor. Required by controller-gen.
func (in *JWTExtractor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PathTemplateMatcher within kubernetes types, where deepcopy-gen is used.
func (in *PathTemplateMatcher) DeepCopyInto(out *PathTemplateMatcher) {
	p := proto.Clone(in).(*PathTemplateMatcher)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PathTemplateMatcher. Required by controller-gen.
func (in *PathTemplateMatcher) DeepCopy() *PathTemplateMatcher {
	if in == nil {
		return nil
	}
	out := new(PathTemplateMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PathTemplateMatcher. Required by controller-gen.
func (in *PathTemplateMatcher) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using FlowControl within kubernetes types, where deepcopy-gen is used.
func (in *FlowControl) DeepCopyInto(out *FlowControl) {
	p := proto.Clone(in).(*FlowControl)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowControl. Required by controller-gen.
func (in *FlowControl) DeepCopy() *FlowControl {
	if in == nil {
		return nil
	}
	out := new(FlowControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new FlowControl. Required by controller-gen.
func (in *FlowControl) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiter within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiter) DeepCopyInto(out *RateLimiter) {
	p := proto.Clone(in).(*RateLimiter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter. Required by controller-gen.
func (in *RateLimiter) DeepCopy() *RateLimiter {
	if in == nil {
		return nil
	}
	out := new(RateLimiter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter. Required by controller-gen.
func (in *RateLimiter) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiter_Parameters within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiter_Parameters) DeepCopyInto(out *RateLimiter_Parameters) {
	p := proto.Clone(in).(*RateLimiter_Parameters)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_Parameters. Required by controller-gen.
func (in *RateLimiter_Parameters) DeepCopy() *RateLimiter_Parameters {
	if in == nil {
		return nil
	}
	out := new(RateLimiter_Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_Parameters. Required by controller-gen.
func (in *RateLimiter_Parameters) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiter_Parameters_LazySync within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiter_Parameters_LazySync) DeepCopyInto(out *RateLimiter_Parameters_LazySync) {
	p := proto.Clone(in).(*RateLimiter_Parameters_LazySync)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_Parameters_LazySync. Required by controller-gen.
func (in *RateLimiter_Parameters_LazySync) DeepCopy() *RateLimiter_Parameters_LazySync {
	if in == nil {
		return nil
	}
	out := new(RateLimiter_Parameters_LazySync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_Parameters_LazySync. Required by controller-gen.
func (in *RateLimiter_Parameters_LazySync) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiter_Override within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiter_Override) DeepCopyInto(out *RateLimiter_Override) {
	p := proto.Clone(in).(*RateLimiter_Override)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_Override. Required by controller-gen.
func (in *RateLimiter_Override) DeepCopy() *RateLimiter_Override {
	if in == nil {
		return nil
	}
	out := new(RateLimiter_Override)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_Override. Required by controller-gen.
func (in *RateLimiter_Override) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiter_DynamicConfig within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiter_DynamicConfig) DeepCopyInto(out *RateLimiter_DynamicConfig) {
	p := proto.Clone(in).(*RateLimiter_DynamicConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_DynamicConfig. Required by controller-gen.
func (in *RateLimiter_DynamicConfig) DeepCopy() *RateLimiter_DynamicConfig {
	if in == nil {
		return nil
	}
	out := new(RateLimiter_DynamicConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_DynamicConfig. Required by controller-gen.
func (in *RateLimiter_DynamicConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimiter_Ins within kubernetes types, where deepcopy-gen is used.
func (in *RateLimiter_Ins) DeepCopyInto(out *RateLimiter_Ins) {
	p := proto.Clone(in).(*RateLimiter_Ins)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_Ins. Required by controller-gen.
func (in *RateLimiter_Ins) DeepCopy() *RateLimiter_Ins {
	if in == nil {
		return nil
	}
	out := new(RateLimiter_Ins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiter_Ins. Required by controller-gen.
func (in *RateLimiter_Ins) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConcurrencyLimiter within kubernetes types, where deepcopy-gen is used.
func (in *ConcurrencyLimiter) DeepCopyInto(out *ConcurrencyLimiter) {
	p := proto.Clone(in).(*ConcurrencyLimiter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConcurrencyLimiter. Required by controller-gen.
func (in *ConcurrencyLimiter) DeepCopy() *ConcurrencyLimiter {
	if in == nil {
		return nil
	}
	out := new(ConcurrencyLimiter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConcurrencyLimiter. Required by controller-gen.
func (in *ConcurrencyLimiter) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Scheduler within kubernetes types, where deepcopy-gen is used.
func (in *Scheduler) DeepCopyInto(out *Scheduler) {
	p := proto.Clone(in).(*Scheduler)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler. Required by controller-gen.
func (in *Scheduler) DeepCopy() *Scheduler {
	if in == nil {
		return nil
	}
	out := new(Scheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler. Required by controller-gen.
func (in *Scheduler) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Scheduler_Workload within kubernetes types, where deepcopy-gen is used.
func (in *Scheduler_Workload) DeepCopyInto(out *Scheduler_Workload) {
	p := proto.Clone(in).(*Scheduler_Workload)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler_Workload. Required by controller-gen.
func (in *Scheduler_Workload) DeepCopy() *Scheduler_Workload {
	if in == nil {
		return nil
	}
	out := new(Scheduler_Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler_Workload. Required by controller-gen.
func (in *Scheduler_Workload) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Scheduler_Workload_Parameters within kubernetes types, where deepcopy-gen is used.
func (in *Scheduler_Workload_Parameters) DeepCopyInto(out *Scheduler_Workload_Parameters) {
	p := proto.Clone(in).(*Scheduler_Workload_Parameters)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler_Workload_Parameters. Required by controller-gen.
func (in *Scheduler_Workload_Parameters) DeepCopy() *Scheduler_Workload_Parameters {
	if in == nil {
		return nil
	}
	out := new(Scheduler_Workload_Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler_Workload_Parameters. Required by controller-gen.
func (in *Scheduler_Workload_Parameters) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Scheduler_Parameters within kubernetes types, where deepcopy-gen is used.
func (in *Scheduler_Parameters) DeepCopyInto(out *Scheduler_Parameters) {
	p := proto.Clone(in).(*Scheduler_Parameters)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler_Parameters. Required by controller-gen.
func (in *Scheduler_Parameters) DeepCopy() *Scheduler_Parameters {
	if in == nil {
		return nil
	}
	out := new(Scheduler_Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler_Parameters. Required by controller-gen.
func (in *Scheduler_Parameters) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Scheduler_Outs within kubernetes types, where deepcopy-gen is used.
func (in *Scheduler_Outs) DeepCopyInto(out *Scheduler_Outs) {
	p := proto.Clone(in).(*Scheduler_Outs)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler_Outs. Required by controller-gen.
func (in *Scheduler_Outs) DeepCopy() *Scheduler_Outs {
	if in == nil {
		return nil
	}
	out := new(Scheduler_Outs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler_Outs. Required by controller-gen.
func (in *Scheduler_Outs) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadActuator within kubernetes types, where deepcopy-gen is used.
func (in *LoadActuator) DeepCopyInto(out *LoadActuator) {
	p := proto.Clone(in).(*LoadActuator)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator. Required by controller-gen.
func (in *LoadActuator) DeepCopy() *LoadActuator {
	if in == nil {
		return nil
	}
	out := new(LoadActuator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator. Required by controller-gen.
func (in *LoadActuator) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadActuator_DynamicConfig within kubernetes types, where deepcopy-gen is used.
func (in *LoadActuator_DynamicConfig) DeepCopyInto(out *LoadActuator_DynamicConfig) {
	p := proto.Clone(in).(*LoadActuator_DynamicConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator_DynamicConfig. Required by controller-gen.
func (in *LoadActuator_DynamicConfig) DeepCopy() *LoadActuator_DynamicConfig {
	if in == nil {
		return nil
	}
	out := new(LoadActuator_DynamicConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator_DynamicConfig. Required by controller-gen.
func (in *LoadActuator_DynamicConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadActuator_Ins within kubernetes types, where deepcopy-gen is used.
func (in *LoadActuator_Ins) DeepCopyInto(out *LoadActuator_Ins) {
	p := proto.Clone(in).(*LoadActuator_Ins)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator_Ins. Required by controller-gen.
func (in *LoadActuator_Ins) DeepCopy() *LoadActuator_Ins {
	if in == nil {
		return nil
	}
	out := new(LoadActuator_Ins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadActuator_Ins. Required by controller-gen.
func (in *LoadActuator_Ins) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AIMDConcurrencyController within kubernetes types, where deepcopy-gen is used.
func (in *AIMDConcurrencyController) DeepCopyInto(out *AIMDConcurrencyController) {
	p := proto.Clone(in).(*AIMDConcurrencyController)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIMDConcurrencyController. Required by controller-gen.
func (in *AIMDConcurrencyController) DeepCopy() *AIMDConcurrencyController {
	if in == nil {
		return nil
	}
	out := new(AIMDConcurrencyController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AIMDConcurrencyController. Required by controller-gen.
func (in *AIMDConcurrencyController) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AIMDConcurrencyController_Ins within kubernetes types, where deepcopy-gen is used.
func (in *AIMDConcurrencyController_Ins) DeepCopyInto(out *AIMDConcurrencyController_Ins) {
	p := proto.Clone(in).(*AIMDConcurrencyController_Ins)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIMDConcurrencyController_Ins. Required by controller-gen.
func (in *AIMDConcurrencyController_Ins) DeepCopy() *AIMDConcurrencyController_Ins {
	if in == nil {
		return nil
	}
	out := new(AIMDConcurrencyController_Ins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AIMDConcurrencyController_Ins. Required by controller-gen.
func (in *AIMDConcurrencyController_Ins) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AIMDConcurrencyController_Outs within kubernetes types, where deepcopy-gen is used.
func (in *AIMDConcurrencyController_Outs) DeepCopyInto(out *AIMDConcurrencyController_Outs) {
	p := proto.Clone(in).(*AIMDConcurrencyController_Outs)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIMDConcurrencyController_Outs. Required by controller-gen.
func (in *AIMDConcurrencyController_Outs) DeepCopy() *AIMDConcurrencyController_Outs {
	if in == nil {
		return nil
	}
	out := new(AIMDConcurrencyController_Outs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AIMDConcurrencyController_Outs. Required by controller-gen.
func (in *AIMDConcurrencyController_Outs) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}