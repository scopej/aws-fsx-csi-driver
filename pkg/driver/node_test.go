/*
Copyright 2019 The Kubernetes Authors.

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

package driver

import (
	"context"
	"fmt"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/aws-fsx-csi-driver/pkg/driver/mocks"
)

// {
// 	name: "success: normal with flock mount options",
// 	testFunc: func(t *testing.T) {
// 		mockCtrl := gomock.NewController(t)
// 		mockMounter := mocks.NewMockMounter(mockCtrl)
// 		driver := &Driver{
// 			endpoint: endpoint,
// 			nodeID:   nodeID,
// 			mounter:  mockMounter,
// 		}

// 		ctx := context.Background()
// 		req := &csi.NodePublishVolumeRequest{
// 			VolumeId: "volumeId",
// 			VolumeContext: map[string]string{
// 				volumeContextDnsName:   dnsname,
// 				volumeContextMountName: mountname,
// 			},
// 			VolumeCapability: &csi.VolumeCapability{
// 				AccessType: &csi.VolumeCapability_Mount{
// 					Mount: &csi.VolumeCapability_MountVolume{
// 						MountFlags: []string{"flock"},
// 					},
// 				},
// 				AccessMode: &csi.VolumeCapability_AccessMode{
// 					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
// 				},
// 			},
// 			TargetPath:        targetPath,
// 			StagingTargetPath: stagingTargetPath,
// 		}
// 		// TODO: flock on bind mount?
// 		mockMounter.EXPECT().MakeDir(gomock.Eq(stagingTargetPath)).Return(nil)
// 		mockMounter.EXPECT().MakeDir(gomock.Eq(targetPath)).Return(nil)
// 		mockMounter.EXPECT().Mount(gomock.Eq(stagingTargetPath), gomock.Eq(targetPath), gomock.Eq(""), gomock.Eq([]string{"bind"})).Return(nil)
// 		_, err := driver.NodePublishVolume(ctx, req)
// 		if err != nil {
// 			t.Fatalf("NodePublishVolume is failed: %v", err)
// 		}

// 		mockCtrl.Finish()
// 	},
// },
// {
// 	name: "fail: missing dns name",
// 	testFunc: func(t *testing.T) {
// 		mockCtrl := gomock.NewController(t)
// 		mockMounter := mocks.NewMockMounter(mockCtrl)
// 		driver := &Driver{
// 			endpoint: endpoint,
// 			nodeID:   nodeID,
// 			mounter:  mockMounter,
// 		}

// 		ctx := context.Background()
// 		req := &csi.NodePublishVolumeRequest{
// 			VolumeId: "volumeId",
// 			VolumeContext: map[string]string{
// 				volumeContextMountName: mountname,
// 			},
// 			VolumeCapability:  stdVolCap,
// 			TargetPath:        targetPath,
// 			StagingTargetPath: stagingTargetPath,
// 		}

// 		_, err := driver.NodePublishVolume(ctx, req)
// 		if err == nil {
// 			t.Fatalf("NodePublishVolume is not failed: %v", err)
// 		}

// 		mockCtrl.Finish()
// 	},
// },

// {
// 	name: "success: normal with flock mount options",
// 	testFunc: func(t *testing.T) {
// 		mockCtrl := gomock.NewController(t)
// 		mockMounter := mocks.NewMockMounter(mockCtrl)
// 		driver := &Driver{
// 			endpoint: endpoint,
// 			nodeID:   nodeID,
// 			mounter:  mockMounter,
// 		}

// 		ctx := context.Background()
// 		req := &csi.NodePublishVolumeRequest{
// 			VolumeId: "volumeId",
// 			VolumeContext: map[string]string{
// 				volumeContextDnsName:   dnsname,
// 				volumeContextMountName: mountname,
// 			},
// 			VolumeCapability: &csi.VolumeCapability{
// 				AccessType: &csi.VolumeCapability_Mount{
// 					Mount: &csi.VolumeCapability_MountVolume{
// 						MountFlags: []string{"flock"},
// 					},
// 				},
// 				AccessMode: &csi.VolumeCapability_AccessMode{
// 					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
// 				},
// 			},
// 			TargetPath:        targetPath,
// 			StagingTargetPath: stagingTargetPath,
// 		}
// 		// TODO: flock on bind mount?
// 		mockMounter.EXPECT().MakeDir(gomock.Eq(stagingTargetPath)).Return(nil)
// 		mockMounter.EXPECT().MakeDir(gomock.Eq(targetPath)).Return(nil)
// 		mockMounter.EXPECT().Mount(gomock.Eq(stagingTargetPath), gomock.Eq(targetPath), gomock.Eq(""), gomock.Eq([]string{"bind"})).Return(nil)
// 		_, err := driver.NodePublishVolume(ctx, req)
// 		if err != nil {
// 			t.Fatalf("NodePublishVolume is failed: %v", err)
// 		}

// 		mockCtrl.Finish()
// 	},
// },
// {
// 	name: "fail: missing dns name",
// 	testFunc: func(t *testing.T) {
// 		mockCtrl := gomock.NewController(t)
// 		mockMounter := mocks.NewMockMounter(mockCtrl)
// 		driver := &Driver{
// 			endpoint: endpoint,
// 			nodeID:   nodeID,
// 			mounter:  mockMounter,
// 		}

// 		ctx := context.Background()
// 		req := &csi.NodePublishVolumeRequest{
// 			VolumeId: "volumeId",
// 			VolumeContext: map[string]string{
// 				volumeContextMountName: mountname,
// 			},
// 			VolumeCapability:  stdVolCap,
// 			TargetPath:        targetPath,
// 			StagingTargetPath: stagingTargetPath,
// 		}

// 		_, err := driver.NodePublishVolume(ctx, req)
// 		if err == nil {
// 			t.Fatalf("NodePublishVolume is not failed: %v", err)
// 		}

// 		mockCtrl.Finish()
// 	},
// },
// {
// 	name: "fail: missing volume capability",
// 	testFunc: func(t *testing.T) {
// 		mockCtrl := gomock.NewController(t)
// 		mockMounter := mocks.NewMockMounter(mockCtrl)
// 		driver := &Driver{
// 			endpoint: endpoint,
// 			nodeID:   nodeID,
// 			mounter:  mockMounter,
// 		}

// 		ctx := context.Background()
// 		req := &csi.NodePublishVolumeRequest{
// 			VolumeId: "volumeId",
// 			VolumeContext: map[string]string{
// 				volumeContextDnsName:   dnsname,
// 				volumeContextMountName: mountname,
// 			},
// 			TargetPath:        targetPath,
// 			StagingTargetPath: stagingTargetPath,
// 		}

// 		_, err := driver.NodePublishVolume(ctx, req)
// 		if err == nil {
// 			t.Fatalf("NodePublishVolume is not failed: %v", err)
// 		}

// 		mockCtrl.Finish()
// 	},
// },
// {
// 	name: "fail: unsupported volume capability",
// 	testFunc: func(t *testing.T) {
// 		mockCtrl := gomock.NewController(t)
// 		mockMounter := mocks.NewMockMounter(mockCtrl)
// 		driver := &Driver{
// 			endpoint: endpoint,
// 			nodeID:   nodeID,
// 			mounter:  mockMounter,
// 		}

// 		ctx := context.Background()
// 		req := &csi.NodePublishVolumeRequest{
// 			VolumeId: "volumeId",
// 			VolumeContext: map[string]string{
// 				volumeContextDnsName:   dnsname,
// 				volumeContextMountName: mountname,
// 			},
// 			VolumeCapability: &csi.VolumeCapability{
// 				AccessType: &csi.VolumeCapability_Mount{
// 					Mount: &csi.VolumeCapability_MountVolume{},
// 				},
// 				AccessMode: &csi.VolumeCapability_AccessMode{
// 					Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY,
// 				},
// 			},
// 			TargetPath:        targetPath,
// 			StagingTargetPath: stagingTargetPath,
// 		}

// 		_, err := driver.NodePublishVolume(ctx, req)
// 		if err == nil {
// 			t.Fatalf("NodePublishVolume is not failed: %v", err)
// 		}

// 		mockCtrl.Finish()
// 	},
// },

func TestNodePublishVolume(t *testing.T) {

	var (
		endpoint          = "endpoint"
		nodeID            = "nodeID"
		dnsname           = "fs-0a2d0632b5ff567e9.fsx.us-west-2.amazonaws.com"
		mountname         = "random"
		targetPath        = "/target/path"
		stagingTargetPath = "/staging/target/path"
		stdVolCap         = &csi.VolumeCapability{
			AccessType: &csi.VolumeCapability_Mount{
				Mount: &csi.VolumeCapability_MountVolume{},
			},
			AccessMode: &csi.VolumeCapability_AccessMode{
				Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
			},
		}
	)

	mockDriver := func(mockCtrl *gomock.Controller) (*Driver, *mocks.MockMounter) {
		mockMounter := mocks.NewMockMounter(mockCtrl)
		driver := &Driver{
			endpoint: endpoint,
			nodeID:   nodeID,
			mounter:  mockMounter,
		}
		return driver, mockMounter
	}

	successfulDriverWithOptions := func(mountOptions []string) func(*gomock.Controller) *Driver {
		return func(mockCtrl *gomock.Controller) *Driver {
			driver, mockMounter := mockDriver(mockCtrl)
			mockMounter.EXPECT().MakeDir(gomock.Eq(targetPath)).Return(nil)
			mockMounter.EXPECT().MakeDir(gomock.Eq(stagingTargetPath)).Return(nil)
			mockMounter.EXPECT().Mount(gomock.Eq(stagingTargetPath), gomock.Eq(targetPath), gomock.Eq(""), gomock.Eq(mountOptions)).Return(nil)
			return driver
		}
	}

	standardRequest := func() *csi.NodePublishVolumeRequest {
		return &csi.NodePublishVolumeRequest{
			VolumeId: "volumeId",
			VolumeContext: map[string]string{
				volumeContextDnsName:   dnsname,
				volumeContextMountName: mountname,
			},
			VolumeCapability:  stdVolCap,
			TargetPath:        targetPath,
			StagingTargetPath: stagingTargetPath,
		}
	}

	testCases := []struct {
		name        string
		driver      func(mockCtrl *gomock.Controller) *Driver
		request     func() *csi.NodePublishVolumeRequest
		expectError bool
	}{
		{
			name:   "success: normal",
			driver: successfulDriverWithOptions([]string{"bind"}),
			request: func() *csi.NodePublishVolumeRequest {
				req := standardRequest()
				return req
			},
		},
		{
			name:   "success: normal with read only mount",
			driver: successfulDriverWithOptions([]string{"bind", "ro"}),
			request: func() *csi.NodePublishVolumeRequest {
				req := standardRequest()
				req.Readonly = true
				return req
			},
		},
		{
			name: "fail: missing target path",
			driver: func(mockCtrl *gomock.Controller) *Driver {
				driver, _ := mockDriver(mockCtrl)
				return driver
			},
			request: func() *csi.NodePublishVolumeRequest {
				req := standardRequest()
				req.TargetPath = ""
				return req
			},
			expectError: true,
		},
		{
			name: "fail: missing staging target path",
			driver: func(mockCtrl *gomock.Controller) *Driver {
				driver, _ := mockDriver(mockCtrl)
				return driver
			},
			request: func() *csi.NodePublishVolumeRequest {
				req := standardRequest()
				req.StagingTargetPath = ""
				return req
			},
			expectError: true,
		},
		{
			name: "fail: mounter failed to MakeDir",
			driver: func(mockCtrl *gomock.Controller) *Driver {
				driver, mockMounter := mockDriver(mockCtrl)
				err := fmt.Errorf("failed to MakeDir")
				mockMounter.EXPECT().MakeDir(gomock.Eq(targetPath)).Return(err)

				return driver
			},
			request:     standardRequest,
			expectError: true,
		},
		{
			name: "fail: mounter failed to Mount",
			driver: func(mockCtrl *gomock.Controller) *Driver {
				driver, mockMounter := mockDriver(mockCtrl)
				err := fmt.Errorf("failed to Mount")
				mockMounter.EXPECT().MakeDir(gomock.Eq(targetPath)).Return(nil)
				mockMounter.EXPECT().MakeDir(gomock.Eq(stagingTargetPath)).Return(nil)
				mockMounter.EXPECT().Mount(gomock.Eq(stagingTargetPath), gomock.Eq(targetPath), gomock.Eq(""), gomock.Eq([]string{"bind"})).Return(err)

				return driver
			},
			request:     standardRequest,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				mockCtrl = gomock.NewController(t)
				driver   = tc.driver(mockCtrl)
				req      = tc.request()
				ctx      = context.Background()
			)
			_, err := driver.NodePublishVolume(ctx, req)
			if tc.expectError {
				if err == nil {
					t.Fatalf("NodePublishVolume is not failed: %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("NodePublishVolume is failed: %v", err)
				}
			}
			mockCtrl.Finish()
		})
	}
}

func TestNodeUnpublishVolume(t *testing.T) {

	var (
		endpoint   = "endpoint"
		nodeID     = "nodeID"
		targetPath = "/target/path"
	)

	testCases := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{
			name: "success: normal",
			testFunc: func(t *testing.T) {
				mockCtrl := gomock.NewController(t)
				mockMounter := mocks.NewMockMounter(mockCtrl)
				driver := &Driver{
					endpoint: endpoint,
					nodeID:   nodeID,
					mounter:  mockMounter,
				}

				ctx := context.Background()
				req := &csi.NodeUnpublishVolumeRequest{
					VolumeId:   "volumeId",
					TargetPath: targetPath,
				}

				mockMounter.EXPECT().Unmount(gomock.Eq(targetPath)).Return(nil)

				_, err := driver.NodeUnpublishVolume(ctx, req)
				if err != nil {
					t.Fatalf("NodeUnpublishVolume is failed: %v", err)
				}
			},
		},
		{
			name: "fail: targetPath is missing",
			testFunc: func(t *testing.T) {
				mockCtrl := gomock.NewController(t)
				mockMounter := mocks.NewMockMounter(mockCtrl)
				driver := &Driver{
					endpoint: endpoint,
					nodeID:   nodeID,
					mounter:  mockMounter,
				}

				ctx := context.Background()
				req := &csi.NodeUnpublishVolumeRequest{
					VolumeId: "volumeId",
				}

				_, err := driver.NodeUnpublishVolume(ctx, req)
				if err == nil {
					t.Fatalf("NodeUnpublishVolume is not failed: %v", err)
				}
			},
		},
		{
			name: "fail: mounter failed to umount",
			testFunc: func(t *testing.T) {
				mockCtrl := gomock.NewController(t)
				mockMounter := mocks.NewMockMounter(mockCtrl)
				driver := &Driver{
					endpoint: endpoint,
					nodeID:   nodeID,
					mounter:  mockMounter,
				}

				ctx := context.Background()
				req := &csi.NodeUnpublishVolumeRequest{
					VolumeId:   "volumeId",
					TargetPath: targetPath,
				}

				mountErr := fmt.Errorf("Unmount failed")
				mockMounter.EXPECT().Unmount(gomock.Eq(targetPath)).Return(mountErr)

				_, err := driver.NodeUnpublishVolume(ctx, req)
				if err == nil {
					t.Fatalf("NodeUnpublishVolume is not failed: %v", err)
				}
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, tc.testFunc)
	}
}
