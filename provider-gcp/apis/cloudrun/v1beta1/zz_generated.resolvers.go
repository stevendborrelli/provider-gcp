/*
Copyright 2021 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DomainMapping.
func (mg *DomainMapping) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Metadata); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Metadata[i3].Namespace),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Metadata[i3].NamespaceRef,
			Selector:     mg.Spec.ForProvider.Metadata[i3].NamespaceSelector,
			To: reference.To{
				List:    &v1beta1.ProjectList{},
				Managed: &v1beta1.Project{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Metadata[i3].Namespace")
		}
		mg.Spec.ForProvider.Metadata[i3].Namespace = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Metadata[i3].NamespaceRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Spec[i3].RouteName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Spec[i3].RouteNameRef,
			Selector:     mg.Spec.ForProvider.Spec[i3].RouteNameSelector,
			To: reference.To{
				List:    &ServiceList{},
				Managed: &Service{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].RouteName")
		}
		mg.Spec.ForProvider.Spec[i3].RouteName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Spec[i3].RouteNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Metadata); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Metadata[i3].Namespace),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Metadata[i3].NamespaceRef,
			Selector:     mg.Spec.ForProvider.Metadata[i3].NamespaceSelector,
			To: reference.To{
				List:    &v1beta1.ProjectList{},
				Managed: &v1beta1.Project{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Metadata[i3].Namespace")
		}
		mg.Spec.ForProvider.Metadata[i3].Namespace = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Metadata[i3].NamespaceRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ServiceIAMBinding.
func (mg *ServiceIAMBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectRef,
		Selector:     mg.Spec.ForProvider.ProjectSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Service),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceRef,
		Selector:     mg.Spec.ForProvider.ServiceSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Service")
	}
	mg.Spec.ForProvider.Service = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceIAMMember.
func (mg *ServiceIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectRef,
		Selector:     mg.Spec.ForProvider.ProjectSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Service),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceRef,
		Selector:     mg.Spec.ForProvider.ServiceSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Service")
	}
	mg.Spec.ForProvider.Service = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceIAMPolicy.
func (mg *ServiceIAMPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectRef,
		Selector:     mg.Spec.ForProvider.ProjectSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Service),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceRef,
		Selector:     mg.Spec.ForProvider.ServiceSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Service")
	}
	mg.Spec.ForProvider.Service = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceRef = rsp.ResolvedReference

	return nil
}