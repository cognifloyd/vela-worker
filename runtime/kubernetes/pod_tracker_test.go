// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package kubernetes

import (
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
)

func TestNewPodTracker(t *testing.T) {
	// setup types
	logger := logrus.NewEntry(logrus.StandardLogger())
	clientset := fake.NewSimpleClientset()

	tests := []struct {
		name    string
		pod     *v1.Pod
		wantErr bool
	}{
		{
			name:    "pass-with-pod",
			pod:     _pod,
			wantErr: false,
		},
		{
			name:    "error-with-nil-pod",
			pod:     nil,
			wantErr: true,
		},
		{
			name:    "error-with-empty-pod",
			pod:     &v1.Pod{},
			wantErr: true,
		},
		{
			name: "error-with-pod-without-namespace",
			pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
			},
			wantErr: true,
		},
		{
			name: "fail-with-pod",
			pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "github-octocat-1-for-some-odd-reason-this-name-is-way-too-long-and-will-cause-an-error",
					Namespace: _pod.ObjectMeta.Namespace,
					Labels:    _pod.ObjectMeta.Labels,
				},
				TypeMeta: _pod.TypeMeta,
				Spec:     _pod.Spec,
				Status:   _pod.Status,
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := newPodTracker(logger, clientset, test.pod, 0*time.Second)
			if (err != nil) != test.wantErr {
				t.Errorf("newPodTracker() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_podTracker_getTrackedPod(t *testing.T) {
	// setup types
	logger := logrus.NewEntry(logrus.StandardLogger())

	tests := []struct {
		name       string
		trackedPod string // namespace/podName
		obj        interface{}
		want       *v1.Pod
	}{
		{
			name:       "got-tracked-pod",
			trackedPod: "test/github-octocat-1",
			obj:        _pod,
			want:       _pod,
		},
		{
			name:       "wrong-pod",
			trackedPod: "test/github-octocat-2",
			obj:        _pod,
			want:       nil,
		},
		{
			name:       "invalid-type",
			trackedPod: "test/github-octocat-1",
			obj:        new(v1.PodTemplate),
			want:       nil,
		},
		{
			name:       "nil",
			trackedPod: "test/nil",
			obj:        nil,
			want:       nil,
		},
		{
			name:       "tombstone-pod",
			trackedPod: "test/github-octocat-1",
			obj: cache.DeletedFinalStateUnknown{
				Key: "test/github-octocat-1",
				Obj: _pod,
			},
			want: _pod,
		},
		{
			name:       "tombstone-nil",
			trackedPod: "test/github-octocat-1",
			obj: cache.DeletedFinalStateUnknown{
				Key: "test/github-octocat-1",
				Obj: nil,
			},
			want: nil,
		},
		{
			name:       "tombstone-invalid-type",
			trackedPod: "test/github-octocat-1",
			obj: cache.DeletedFinalStateUnknown{
				Key: "test/github-octocat-1",
				Obj: new(v1.PodTemplate),
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := podTracker{
				Logger:     logger,
				TrackedPod: test.trackedPod,
				// other fields not used by getTrackedPod
				// if they're needed, use newPodTracker
			}
			if got := p.getTrackedPod(test.obj); !reflect.DeepEqual(got, test.want) {
				t.Errorf("getTrackedPod() = %v, want %v", got, test.want)
			}
		})
	}
}