//
// Copyright (c) 2018 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	pb "github.com/kata-containers/agent/protocols/grpc"
	"github.com/stretchr/testify/assert"
)

var (
	testCtrPath = "test-ctr-path"
)

func createFakeDevicePath() (string, error) {
	f, err := ioutil.TempFile("", "fake-dev-path")
	if err != nil {
		return "", err
	}
	path := f.Name()
	f.Close()

	return path, nil
}

func testVirtioBlkDeviceHandlerFailure(t *testing.T, device pb.Device, spec *pb.Spec) {
	devPath, err := createFakeDevicePath()
	assert.Nil(t, err, "Fake device path creation failed: %v", err)
	defer os.RemoveAll(devPath)

	device.VmPath = devPath
	device.ContainerPath = "some-not-empty-path"

	err = virtioBlkDeviceHandler(device, spec)
	assert.NotNil(t, err, "blockDeviceHandler() should have failed")
}

func TestVirtioBlkDeviceHandlerEmptyContainerPath(t *testing.T) {
	spec := &pb.Spec{}
	device := pb.Device{
		ContainerPath: testCtrPath,
	}

	testVirtioBlkDeviceHandlerFailure(t, device, spec)
}

func TestVirtioBlkDeviceHandlerNilLinuxSpecFailure(t *testing.T) {
	spec := &pb.Spec{}
	device := pb.Device{
		ContainerPath: testCtrPath,
	}

	testVirtioBlkDeviceHandlerFailure(t, device, spec)
}

func TestVirtioBlkDeviceHandlerEmptyLinuxDevicesSpecFailure(t *testing.T) {
	spec := &pb.Spec{
		Linux: &pb.Linux{},
	}
	device := pb.Device{
		ContainerPath: testCtrPath,
	}

	testVirtioBlkDeviceHandlerFailure(t, device, spec)
}

func TestScanSCSIBus(t *testing.T) {
	testDir, err := ioutil.TempDir("", "kata-agent-tmp-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testDir)

	scsiHostPath = filepath.Join(testDir, "scsi_host")
	os.RemoveAll(scsiHostPath)

	defer os.RemoveAll(scsiHostPath)

	scsiAddr := "1"

	err = scanSCSIBus(scsiAddr)
	assert.NotNil(t, err, "scanSCSIBus() should have failed")

	if err := os.MkdirAll(scsiHostPath, mountPerm); err != nil {
		t.Fatal(err)
	}

	scsiAddr = "1:1"
	err = scanSCSIBus(scsiAddr)
	assert.Nil(t, err, "scanSCSIBus() failed: %v", err)

	host := filepath.Join(scsiHostPath, "host0")
	if err := os.MkdirAll(host, mountPerm); err != nil {
		t.Fatal(err)
	}

	err = scanSCSIBus(scsiAddr)
	assert.Nil(t, err, "scanSCSIBus() failed: %v", err)

	scanPath := filepath.Join(host, "scan")
	_, err = os.Stat(scanPath)
	assert.Nil(t, err, "os.Stat() %s failed: %v", scanPath, err)
}

func testAddDevicesSuccessful(t *testing.T, devices []*pb.Device, spec *pb.Spec) {
	err := addDevices(devices, spec)
	assert.Nil(t, err, "addDevices() failed: %v", err)
}

func TestAddDevicesEmptyDevicesSuccessful(t *testing.T) {
	var devices []*pb.Device
	spec := &pb.Spec{}

	testAddDevicesSuccessful(t, devices, spec)
}

func TestAddDevicesNilMountsSuccessful(t *testing.T) {
	devices := []*pb.Device{
		nil,
	}

	spec := &pb.Spec{}

	testAddDevicesSuccessful(t, devices, spec)
}

func noopDeviceHandlerReturnNil(device pb.Device, spec *pb.Spec) error {
	return nil
}

func noopDeviceHandlerReturnError(device pb.Device, spec *pb.Spec) error {
	return fmt.Errorf("Noop handler failure")
}

func TestAddDevicesNoopHandlerSuccessful(t *testing.T) {
	noopHandlerTag := "noop"
	deviceHandlerList = map[string]deviceHandler{
		noopHandlerTag: noopDeviceHandlerReturnNil,
	}

	devices := []*pb.Device{
		{
			Type: noopHandlerTag,
		},
	}

	spec := &pb.Spec{}

	testAddDevicesFailure(t, devices, spec)
}

func testAddDevicesFailure(t *testing.T, devices []*pb.Device, spec *pb.Spec) {
	err := addDevices(devices, spec)
	assert.NotNil(t, err, "addDevices() should have failed")
}

func TestAddDevicesUnknownHandlerFailure(t *testing.T) {
	deviceHandlerList = map[string]deviceHandler{}

	devices := []*pb.Device{
		{
			Type: "unknown",
		},
	}

	spec := &pb.Spec{}

	testAddDevicesFailure(t, devices, spec)
}

func TestAddDevicesNoopHandlerFailure(t *testing.T) {
	noopHandlerTag := "noop"
	deviceHandlerList = map[string]deviceHandler{
		noopHandlerTag: noopDeviceHandlerReturnError,
	}

	devices := []*pb.Device{
		{
			Type: noopHandlerTag,
		},
	}

	spec := &pb.Spec{}

	testAddDevicesFailure(t, devices, spec)
}

func TestAddDevice(t *testing.T) {
	assert := assert.New(t)

	emptySpec := &pb.Spec{}

	// Use a dummy handler so that addDevice() will be successful
	// if the Device itself is valid.
	noopHandlerTag := "noop"
	deviceHandlerList = map[string]deviceHandler{
		noopHandlerTag: noopDeviceHandlerReturnNil,
	}

	type testData struct {
		device      *pb.Device
		spec        *pb.Spec
		expectError bool
	}

	data := []testData{
		{
			device:      nil,
			spec:        nil,
			expectError: true,
		},
		{
			device:      &pb.Device{},
			spec:        emptySpec,
			expectError: true,
		},
		{
			device: &pb.Device{
				Id: "foo",
			},
			spec:        emptySpec,
			expectError: true,
		},
		{
			device: &pb.Device{
				Id: "foo",
			},
			spec:        emptySpec,
			expectError: true,
		},
		{
			device: &pb.Device{
				// Missing type
				VmPath:        "/foo",
				ContainerPath: "/foo",
			},
			spec:        emptySpec,
			expectError: true,
		},
		{
			device: &pb.Device{
				// Missing VmPath
				Type:          noopHandlerTag,
				ContainerPath: "/foo",
			},
			spec:        emptySpec,
			expectError: true,
		},
		{
			device: &pb.Device{
				// Missing ContainerPath
				Type:   noopHandlerTag,
				VmPath: "/foo",
			},
			spec:        emptySpec,
			expectError: true,
		},
		{
			device: &pb.Device{
				// Id is optional
				Type:          noopHandlerTag,
				VmPath:        "/foo",
				ContainerPath: "/foo",
				Options:       []string{},
			},
			spec:        emptySpec,
			expectError: false,
		},
		{
			device: &pb.Device{
				// Options are... optional ;)
				Id:            "foo",
				Type:          noopHandlerTag,
				VmPath:        "/foo",
				ContainerPath: "/foo",
			},
			spec:        emptySpec,
			expectError: false,
		},
		{
			device: &pb.Device{
				Id:            "foo",
				Type:          noopHandlerTag,
				VmPath:        "/foo",
				ContainerPath: "/foo",
				Options:       []string{},
			},
			spec:        emptySpec,
			expectError: false,
		},
		{
			device: &pb.Device{
				Type:          noopHandlerTag,
				VmPath:        "/foo",
				ContainerPath: "/foo",
			},
			spec:        emptySpec,
			expectError: false,
		},
	}

	for i, d := range data {
		err := addDevice(d.device, d.spec)
		if d.expectError {
			assert.Errorf(err, "test %d (%+v)", i, d)
		} else {
			assert.NoErrorf(err, "test %d (%+v)", i, d)
		}
	}
}
