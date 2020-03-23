// +build !ignore_autogenerated

/* Copyright (C) Couchbase, Inc 2020 - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfig) DeepCopyInto(out *CouchbaseServiceBrokerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfig.
func (in *CouchbaseServiceBrokerConfig) DeepCopy() *CouchbaseServiceBrokerConfig {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CouchbaseServiceBrokerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigBinding) DeepCopyInto(out *CouchbaseServiceBrokerConfigBinding) {
	*out = *in
	if in.ServiceInstance != nil {
		in, out := &in.ServiceInstance, &out.ServiceInstance
		*out = new(CouchbaseServiceBrokerTemplateList)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceBinding != nil {
		in, out := &in.ServiceBinding, &out.ServiceBinding
		*out = new(CouchbaseServiceBrokerTemplateList)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigBinding.
func (in *CouchbaseServiceBrokerConfigBinding) DeepCopy() *CouchbaseServiceBrokerConfigBinding {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigList) DeepCopyInto(out *CouchbaseServiceBrokerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CouchbaseServiceBrokerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigList.
func (in *CouchbaseServiceBrokerConfigList) DeepCopy() *CouchbaseServiceBrokerConfigList {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CouchbaseServiceBrokerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigSpec) DeepCopyInto(out *CouchbaseServiceBrokerConfigSpec) {
	*out = *in
	if in.Catalog != nil {
		in, out := &in.Catalog, &out.Catalog
		*out = new(ServiceCatalog)
		(*in).DeepCopyInto(*out)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]CouchbaseServiceBrokerConfigTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]CouchbaseServiceBrokerConfigBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigSpec.
func (in *CouchbaseServiceBrokerConfigSpec) DeepCopy() *CouchbaseServiceBrokerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplate) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplate) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]CouchbaseServiceBrokerConfigTemplateParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplate.
func (in *CouchbaseServiceBrokerConfigTemplate) DeepCopy() *CouchbaseServiceBrokerConfigTemplate {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameter) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameter) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(CouchbaseServiceBrokerConfigTemplateParameterDefault)
		(*in).DeepCopyInto(*out)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(CouchbaseServiceBrokerConfigTemplateParameterSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]CouchbaseServiceBrokerConfigTemplateParameterDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameter.
func (in *CouchbaseServiceBrokerConfigTemplateParameter) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameter {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterDefault) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterDefault) {
	*out = *in
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = new(string)
		**out = **in
	}
	if in.Bool != nil {
		in, out := &in.Bool, &out.Bool
		*out = new(bool)
		**out = **in
	}
	if in.Int != nil {
		in, out := &in.Int, &out.Int
		*out = new(int)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterDefault.
func (in *CouchbaseServiceBrokerConfigTemplateParameterDefault) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterDefault {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterDefault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterDestination) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterDestination) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterDestination.
func (in *CouchbaseServiceBrokerConfigTemplateParameterDestination) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterDestination {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSource) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSource) {
	*out = *in
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(CouchbaseServiceBrokerConfigTemplateParameterSourceFormat)
		(*in).DeepCopyInto(*out)
	}
	if in.GeneratePassword != nil {
		in, out := &in.GeneratePassword, &out.GeneratePassword
		*out = new(CouchbaseServiceBrokerConfigTemplateParameterSourceGeneratePassword)
		(*in).DeepCopyInto(*out)
	}
	if in.GenerateKey != nil {
		in, out := &in.GenerateKey, &out.GenerateKey
		*out = new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateKey)
		(*in).DeepCopyInto(*out)
	}
	if in.GenerateCertificate != nil {
		in, out := &in.GenerateCertificate, &out.GenerateCertificate
		*out = new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificate)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSource.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSource) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSource {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceFormat) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSourceFormat) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSourceFormat.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceFormat) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSourceFormat {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSourceFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter) {
	*out = *in
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificate) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificate) {
	*out = *in
	in.Key.DeepCopyInto(&out.Key)
	out.Name = in.Name
	out.Lifetime = in.Lifetime
	if in.AlternativeNames != nil {
		in, out := &in.AlternativeNames, &out.AlternativeNames
		*out = new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateAltNames)
		(*in).DeepCopyInto(*out)
	}
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateCA)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificate.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificate) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificate {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateAltNames) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateAltNames) {
	*out = *in
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = make([]CouchbaseServiceBrokerConfigTemplateParameterSourceFormatParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateAltNames.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateAltNames) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateAltNames {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateAltNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateCA) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateCA) {
	*out = *in
	in.Key.DeepCopyInto(&out.Key)
	in.Certificate.DeepCopyInto(&out.Certificate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateCA.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateCA) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateCA {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateCA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateName) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateName.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateName) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateName {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateCertificateName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateKey) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateKey) {
	*out = *in
	if in.Bits != nil {
		in, out := &in.Bits, &out.Bits
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateKey.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateKey) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateKey {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSourceGenerateKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGeneratePassword) DeepCopyInto(out *CouchbaseServiceBrokerConfigTemplateParameterSourceGeneratePassword) {
	*out = *in
	if in.Dictionary != nil {
		in, out := &in.Dictionary, &out.Dictionary
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerConfigTemplateParameterSourceGeneratePassword.
func (in *CouchbaseServiceBrokerConfigTemplateParameterSourceGeneratePassword) DeepCopy() *CouchbaseServiceBrokerConfigTemplateParameterSourceGeneratePassword {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerConfigTemplateParameterSourceGeneratePassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CouchbaseServiceBrokerTemplateList) DeepCopyInto(out *CouchbaseServiceBrokerTemplateList) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]CouchbaseServiceBrokerConfigTemplateParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CouchbaseServiceBrokerTemplateList.
func (in *CouchbaseServiceBrokerTemplateList) DeepCopy() *CouchbaseServiceBrokerTemplateList {
	if in == nil {
		return nil
	}
	out := new(CouchbaseServiceBrokerTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardClient) DeepCopyInto(out *DashboardClient) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardClient.
func (in *DashboardClient) DeepCopy() *DashboardClient {
	if in == nil {
		return nil
	}
	out := new(DashboardClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputParamtersSchema) DeepCopyInto(out *InputParamtersSchema) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputParamtersSchema.
func (in *InputParamtersSchema) DeepCopy() *InputParamtersSchema {
	if in == nil {
		return nil
	}
	out := new(InputParamtersSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceInfo) DeepCopyInto(out *MaintenanceInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceInfo.
func (in *MaintenanceInfo) DeepCopy() *MaintenanceInfo {
	if in == nil {
		return nil
	}
	out := new(MaintenanceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schemas) DeepCopyInto(out *Schemas) {
	*out = *in
	if in.ServiceInstance != nil {
		in, out := &in.ServiceInstance, &out.ServiceInstance
		*out = new(ServiceInstanceSchema)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceBinding != nil {
		in, out := &in.ServiceBinding, &out.ServiceBinding
		*out = new(ServiceBindingSchema)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schemas.
func (in *Schemas) DeepCopy() *Schemas {
	if in == nil {
		return nil
	}
	out := new(Schemas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingSchema) DeepCopyInto(out *ServiceBindingSchema) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(InputParamtersSchema)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingSchema.
func (in *ServiceBindingSchema) DeepCopy() *ServiceBindingSchema {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceCatalog) DeepCopyInto(out *ServiceCatalog) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServiceOffering, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceCatalog.
func (in *ServiceCatalog) DeepCopy() *ServiceCatalog {
	if in == nil {
		return nil
	}
	out := new(ServiceCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInstanceSchema) DeepCopyInto(out *ServiceInstanceSchema) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(InputParamtersSchema)
		(*in).DeepCopyInto(*out)
	}
	if in.Update != nil {
		in, out := &in.Update, &out.Update
		*out = new(InputParamtersSchema)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInstanceSchema.
func (in *ServiceInstanceSchema) DeepCopy() *ServiceInstanceSchema {
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceOffering) DeepCopyInto(out *ServiceOffering) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Requires != nil {
		in, out := &in.Requires, &out.Requires
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.DashboardClient != nil {
		in, out := &in.DashboardClient, &out.DashboardClient
		*out = new(DashboardClient)
		**out = **in
	}
	if in.Plans != nil {
		in, out := &in.Plans, &out.Plans
		*out = make([]ServicePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceOffering.
func (in *ServiceOffering) DeepCopy() *ServiceOffering {
	if in == nil {
		return nil
	}
	out := new(ServiceOffering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePlan) DeepCopyInto(out *ServicePlan) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Bindable != nil {
		in, out := &in.Bindable, &out.Bindable
		*out = new(bool)
		**out = **in
	}
	if in.PlanUpdatable != nil {
		in, out := &in.PlanUpdatable, &out.PlanUpdatable
		*out = new(bool)
		**out = **in
	}
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = new(Schemas)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceInfo != nil {
		in, out := &in.MaintenanceInfo, &out.MaintenanceInfo
		*out = new(MaintenanceInfo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePlan.
func (in *ServicePlan) DeepCopy() *ServicePlan {
	if in == nil {
		return nil
	}
	out := new(ServicePlan)
	in.DeepCopyInto(out)
	return out
}