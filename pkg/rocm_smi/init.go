// Copyright (c) 2019 RRZE, University Erlangen-Nuremberg

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package rocm_smi

import (
	"fmt"
	"github.com/NVIDIA/go-nvml/pkg/dl"
	"runtime"
)

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#cgo CFLAGS:
#include <stdint.h>
#include "rocm_smi/rocm_smi.h"
#include "rocm_smi/kfd_ioctl.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	rocmSmiLibraryName      = "librocm_smi64.so.5.0.50100"
	rocmSmiLibraryLoadFlags = dl.RTLD_LAZY | dl.RTLD_GLOBAL
)

var rocm_smi_lib *dl.DynamicLibrary

type DeviceHandle struct {
	handle uint16
	index  uint32
}
type DeviceIndex uint32

// rocm_smi.Init()
func Init() RSMI_status {
	lib := dl.New(rocmSmiLibraryName, rocmSmiLibraryLoadFlags)
	if lib == nil {
		panic(fmt.Sprintf("error instantiating DynamicLibrary for %s", rocmSmiLibraryName))
	}

	err := lib.Open()
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", rocmSmiLibraryName, err))
	}

	rocm_smi_lib = lib
	updateFunctionPointers()
	return rsmi_init(0)
}

// rocm_smi.InitWithFlags()
func InitWithFlags(Flags uint32) RSMI_status {
	lib := dl.New(rocmSmiLibraryName, rocmSmiLibraryLoadFlags)
	if lib == nil {
		panic(fmt.Sprintf("error instantiating DynamicLibrary for %s", rocmSmiLibraryName))
	}

	err := lib.Open()
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", rocmSmiLibraryName, err))
	}

	rocm_smi_lib = lib
	updateFunctionPointers()
	return rsmi_init(Flags)
}

// rocm_smi.Shutdown()
func Shutdown() RSMI_status {
	ret := rsmi_shut_down()
	if ret != STATUS_SUCCESS {
		return ret
	}

	err := rocm_smi_lib.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing %s: %v", rocmSmiLibraryName, err))
	}

	return ret
}

func Version() (RSMI_version, RSMI_status) {
	var v RSMI_version
	ret := rsmi_version_get(&v)
	return v, ret
}

func ComponentVersionString(Component RSMI_sw_component) (string, RSMI_status) {
	var version []byte = make([]byte, 100)
	vptr := &version[0]
	ret := rsmi_version_str_get(Component, vptr, 100)
	return string(version), ret
}

// Created manually since the the c-for-go parser does not generate a version with &cStr
func StatusString(Status RSMI_status) (string, RSMI_status) {
	var cStr *C.char
	cStatus, cStatusAllocMap := (C.rsmi_status_t)(Status), cgoAllocsUnknown
	__ret := C.rsmi_status_string(cStatus, &cStr)
	runtime.KeepAlive(cStatusAllocMap)
	__v := (RSMI_status)(__ret)
	return C.GoString(cStr), __v
}

// Created manually since the the c-for-go parser does not generate a version with &cStr
// Version without returing RSMI_status for simpler usage
func StatusStringNoError(Status RSMI_status) string {
	var cStr *C.char
	cStatus, cStatusAllocMap := (C.rsmi_status_t)(Status), cgoAllocsUnknown
	C.rsmi_status_string(cStatus, &cStr)
	runtime.KeepAlive(cStatusAllocMap)
	return C.GoString(cStr)
}

// Check whether some symbols are defined by the rocm_smi library and update
// functions pointers accordingly
func updateFunctionPointers() {
	var err error
	err = rocm_smi_lib.Lookup("rsmi_dev_sku_get")
	if err == nil {
		DeviceGetSku = DeviceGetSkuReal
	}

	err = rocm_smi_lib.Lookup("rsmi_dev_perf_level_set")
	if err == nil {
		DeviceSetPerfLevel = DeviceSetPerfLevel_v2
	}

	err = rocm_smi_lib.Lookup("rsmi_dev_overdrive_level_set")
	if err == nil {
		DeviceSetOverdriveLevel = DeviceSetOverdriveLevel_v2
	}
}
