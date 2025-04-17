// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, ref
func (_m *Interface) Get(ctx context.Context, ref *v1alpha1.ArtifactRepositoryRefStatus) (*v1alpha1.ArtifactRepository, error) {
	ret := _m.Called(ctx, ref)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *v1alpha1.ArtifactRepository
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ArtifactRepositoryRefStatus) (*v1alpha1.ArtifactRepository, error)); ok {
		return rf(ctx, ref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ArtifactRepositoryRefStatus) *v1alpha1.ArtifactRepository); ok {
		r0 = rf(ctx, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ArtifactRepository)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.ArtifactRepositoryRefStatus) error); ok {
		r1 = rf(ctx, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Resolve provides a mock function with given fields: ctx, ref, workflowNamespace
func (_m *Interface) Resolve(ctx context.Context, ref *v1alpha1.ArtifactRepositoryRef, workflowNamespace string) (*v1alpha1.ArtifactRepositoryRefStatus, error) {
	ret := _m.Called(ctx, ref, workflowNamespace)

	if len(ret) == 0 {
		panic("no return value specified for Resolve")
	}

	var r0 *v1alpha1.ArtifactRepositoryRefStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ArtifactRepositoryRef, string) (*v1alpha1.ArtifactRepositoryRefStatus, error)); ok {
		return rf(ctx, ref, workflowNamespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.ArtifactRepositoryRef, string) *v1alpha1.ArtifactRepositoryRefStatus); ok {
		r0 = rf(ctx, ref, workflowNamespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ArtifactRepositoryRefStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.ArtifactRepositoryRef, string) error); ok {
		r1 = rf(ctx, ref, workflowNamespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
