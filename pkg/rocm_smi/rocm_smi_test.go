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
	"testing"
)

func TestInit(t *testing.T) {
	ret := Init()
	if ret != STATUS_SUCCESS {
		t.Errorf("Init: %v", ret)
	} else {
		t.Logf("Init: %v", ret)
	}

	ret = Shutdown()
	if ret != STATUS_SUCCESS {
		t.Errorf("Shutdown: %v", ret)
	} else {
		t.Logf("Shutdown: %v", ret)
	}
}

func TestStatus(t *testing.T) {
	Init()
	defer Shutdown()

	all := map[RSMI_status]string{
		STATUS_SUCCESS: "STATUS_SUCCESS",
		STATUS_INVALID_ARGS: "STATUS_INVALID_ARGS",
		STATUS_NOT_SUPPORTED: "STATUS_NOT_SUPPORTED",
		STATUS_FILE_ERROR: "STATUS_FILE_ERROR",
		STATUS_PERMISSION: "STATUS_PERMISSION",
		STATUS_OUT_OF_RESOURCES: "STATUS_OUT_OF_RESOURCES",
		STATUS_INTERNAL_EXCEPTION: "STATUS_INTERNAL_EXCEPTION",
		STATUS_INPUT_OUT_OF_BOUNDS: "STATUS_INPUT_OUT_OF_BOUNDS",
		STATUS_INIT_ERROR: "STATUS_INIT_ERROR",
		STATUS_NOT_YET_IMPLEMENTED: "STATUS_NOT_YET_IMPLEMENTED",
		STATUS_NOT_FOUND: "STATUS_NOT_FOUND",
		STATUS_INSUFFICIENT_SIZE: "STATUS_INSUFFICIENT_SIZE",
		STATUS_INTERRUPT: "STATUS_INTERRUPT",
		STATUS_UNEXPECTED_SIZE: "STATUS_UNEXPECTED_SIZE",
		STATUS_NO_DATA: "STATUS_NO_DATA",
		STATUS_UNEXPECTED_DATA: "STATUS_UNEXPECTED_DATA",
		STATUS_BUSY: "STATUS_BUSY",
		STATUS_REFCOUNT_OVERFLOW: "STATUS_REFCOUNT_OVERFLOW",
		STATUS_UNKNOWN_ERROR: "STATUS_UNKNOWN_ERROR",
	}

	for err, str := range all {
		s, ret := StatusString(err)
		if ret != STATUS_SUCCESS {
			t.Errorf("StatusString(%s): %v", str, ret)
		} else {
			t.Logf("StatusString(%s): %v", str, ret)
			t.Logf("  str: %v", s)
		}
	}
}

func TestSysInfo(t *testing.T) {
	Init()
	defer Shutdown()

	procs, ret := ComputeProcesses()
	if ret != STATUS_SUCCESS {
		t.Errorf("ComputeProcesses: %v", ret)
	} else {
		t.Logf("ComputeProcesses: %v", ret)
		for _, p := range procs {
			t.Logf("  vram usage: %v", p.Vram_usage)
		}
	}
}

func TestSystem(t *testing.T) {
	Init()
	defer Shutdown()

	version, ret := Version()
	if ret != STATUS_SUCCESS {
		t.Errorf("Version: %v", ret)
	} else {
		t.Logf("Version: %v", ret)
		t.Logf("  version: %v.%v.%v", version.Major, version.Minor, version.Patch)
	}

	vStr, ret := ComponentVersionString(SW_COMP_DRIVER)
	if ret != STATUS_SUCCESS {
		t.Errorf("ComponentVersionString: %v", ret)
	} else {
		t.Logf("ComponentVersionString: %v", ret)
		t.Logf("  version (SW_COMP_DRIVER): %v", vStr)
	}

	count, ret := NumMonitorDevices()
	if ret != STATUS_SUCCESS {
		t.Errorf("NumMonitorDevices: %v", ret)
	} else {
		t.Logf("NumMonitorDevices: %v", ret)
		t.Logf("  count: %v", count)
	}
}

func TestDevice(t *testing.T) {
	Init()
	defer Shutdown()

	deviceCount, ret := NumMonitorDevices()
	if ret != STATUS_SUCCESS {
		t.Errorf("NumMonitorDevices: %v", ret)
	} else {
		t.Logf("NumMonitorDevices: %v", ret)
		t.Logf("  count: %v", deviceCount)
	}

	if deviceCount == 0 {
		t.Skip("Skipping test with no Units.")
	}

	devHandle, ret := DeviceGetHandleByIndex(0)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetHandleByIndex: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetHandleByIndex: %v", ret)
	} else {
		t.Logf("DeviceGetHandleByIndex: %v", ret)
		t.Logf("  handle: %v %v", devHandle.Index(), devHandle.ID())
	}

	brandName, ret := DeviceGetBrand(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetBrand: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetBrand: %v", ret)
	} else {
		t.Logf("DeviceGetBrand: %v", ret)
		t.Logf("  brand: %v", brandName)
	}

	name, ret := DeviceGetName(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetName: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetName: %v", ret)
	} else {
		t.Logf("DeviceGetName: %v", ret)
		t.Logf("  name: %v", name)
	}

	vendId, ret := DeviceGetVendorId(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVendorId: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVendorId: %v", ret)
	} else {
		t.Logf("DeviceGetVendorId: %v", ret)
		t.Logf("  name: %v", vendId)
	}

	vendName, ret := DeviceGetVendorName(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVendorName: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVendorName: %v", ret)
	} else {
		t.Logf("DeviceGetVendorName: %v", ret)
		t.Logf("  vendor name: %v", vendName)
	}

	vramName, ret := DeviceGetVramVendor(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVramVendor: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVramVendor: %v", ret)
	} else {
		t.Logf("DeviceGetVramVendor: %v", ret)
		t.Logf("  vram name: %v", vramName)
	}

	serial, ret := DeviceGetSerialNumber(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetSerialNumber: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetSerialNumber: %v", ret)
	} else {
		t.Logf("DeviceGetSerialNumber: %v", ret)
		t.Logf("  serial: %v", serial)
	}

	sku, ret := DeviceGetSku(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetSku: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetSku: %v", ret)
	} else {
		t.Logf("DeviceGetSku: %v", ret)
		t.Logf("  sku: %v", sku)
	}

	subId, ret := DeviceGetSubsystemId(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetSubsystemId: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetSubsystemId: %v", ret)
	} else {
		t.Logf("DeviceGetSubsystemId: %v", ret)
		t.Logf("  subsystem id: %v", subId)
	}

	subSys, ret := DeviceGetSubsystemName(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetSubsystemName: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetSubsystemName: %v", ret)
	} else {
		t.Logf("DeviceGetSubsystemName: %v", ret)
		t.Logf("  subsystem: %v", subSys)
	}

	drmMinor, ret := DeviceGetDrmRenderMinor(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetDrmRenderMinor: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetDrmRenderMinor: %v", ret)
	} else {
		t.Logf("DeviceGetDrmRenderMinor: %v", ret)
		t.Logf("  DRM minor: %v", drmMinor)
	}

	pciId, ret := DeviceGetPciId(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPciId: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciId: %v", ret)
	} else {
		t.Logf("DeviceGetPciId: %v", ret)
		t.Logf("  pci id: 0x%X", pciId)
	}

	pciInfo, ret := DeviceGetPciInfo(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPciInfo: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciInfo: %v", ret)
	} else {
		t.Logf("DeviceGetPciInfo: %v", ret)
		t.Logf("  pci domain: 0x%X", pciInfo.Domain)
		t.Logf("  pci bus: 0x%X", pciInfo.Bus)
		t.Logf("  pci device: 0x%X", pciInfo.Device)
		t.Logf("  pci function: 0x%X", pciInfo.Function)
	}

	pciBand, ret := DeviceGetPciBandwidth(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPciBandwidth: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciBandwidth: %v", ret)
	} else {
		t.Logf("DeviceGetPciBandwidth: %v", ret)
		t.Logf("  pci bandwidth rate num_supported: %v", pciBand.Rate.Supported)
		for i, r := range pciBand.Rate.Frequency {
			if i >= int(pciBand.Rate.Supported) {
				break
			}
			if i == int(pciBand.Rate.Current) {
				t.Logf("  pci bandwidth rate: %v (current)", r)
			} else {
				t.Logf("  pci bandwidth rate: %v", r)
			}
		}
		for i, l := range pciBand.Lanes {
			if i >= int(pciBand.Rate.Supported) {
				break
			}
			if i == int(pciBand.Rate.Current) {
				t.Logf("  pci bandwidth lanes: %v (current)", l)
			} else {
				t.Logf("  pci bandwidth lanes: %v", l)
			}
		}
	}

	sent, recv, max_pkts_size, ret := DeviceGetPciThroughput(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPciThroughput: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciThroughput: %v", ret)
	} else {
		t.Logf("DeviceGetPciThroughput: %v", ret)
		t.Logf("  pci throughput sent: %v", sent)
		t.Logf("  pci throughput recv: %v", recv)
		t.Logf("  pci throughput max_pkts_size: %v", max_pkts_size)
	}

	replay, ret := DeviceGetPciReplayCounter(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPciReplayCounter: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciReplayCounter: %v", ret)
	} else {
		t.Logf("DeviceGetPciReplayCounter: %v", ret)
		t.Logf("  replay counter: %v", replay)
	}

	numaAffinity, ret := DeviceGetNumaAffinity(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetNumaAffinity: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetNumaAffinity: %v", ret)
	} else {
		t.Logf("DeviceGetNumaAffinity: %v", ret)
		t.Logf("  numa node: %v", numaAffinity)
	}

	avgPower, ret := DeviceGetPowerAverage(devHandle, 0)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPowerAverage: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPowerAverage: %v", ret)
	} else {
		t.Logf("DeviceGetPowerAverage: %v", ret)
		t.Logf("  average power (sensor 0): %v [microwatts]", avgPower)
	}

	capPower, ret := DeviceGetPowerCap(devHandle, 0)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPowerCap: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPowerCap: %v", ret)
	} else {
		t.Logf("DeviceGetPowerCap: %v", ret)
		t.Logf("  power capping (sensor 0): %v [microwatts]", capPower)
	}

	defCapPower, ret := DeviceGetDefaultPowerCap(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetDefaultPowerCap: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetDefaultPowerCap: %v", ret)
	} else {
		t.Logf("DeviceGetDefaultPowerCap: %v", ret)
		t.Logf("  power capping (default): %v [microwatts]", defCapPower)
	}

	capPowerMax, capPowerMin, ret := DeviceGetPowerCapRange(devHandle, 0)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPowerCapRange: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPowerCapRange: %v", ret)
	} else {
		t.Logf("DeviceGetPowerCapRange: %v", ret)
		t.Logf("  power capping (sensor 0): %v - %v [microwatts]", capPowerMin, capPowerMax)
	}

	energy, resolution, timestamp, ret := DeviceGetEnergyCount(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetEnergyCount: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetEnergyCount: %v", ret)
	} else {
		t.Logf("DeviceGetEnergyCount: %v", ret)
		t.Logf("  energy: %v [microJoules]", energy)
		t.Logf("  resolution: %v [microJoules]", resolution)
		t.Logf("  timestamp: %v [ns]", timestamp)
	}

	totalMemFirst, ret := DeviceGetTotalMemory(devHandle, MEM_TYPE_FIRST)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetTotalMemory: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetTotalMemory: %v", ret)
	} else {
		t.Logf("DeviceGetTotalMemory: %v", ret)
		t.Logf("  total memory (MEM_TYPE_FIRST): %v", totalMemFirst)
	}
	totalMemLast, ret := DeviceGetTotalMemory(devHandle, MEM_TYPE_LAST)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetTotalMemory: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetTotalMemory: %v", ret)
	} else {
		t.Logf("DeviceGetTotalMemory: %v", ret)
		t.Logf("  total memory (MEM_TYPE_LAST): %v", totalMemLast)
	}

	usedMemFirst, ret := DeviceGetUsedMemory(devHandle, MEM_TYPE_FIRST)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetUsedMemory: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetUsedMemory: %v", ret)
	} else {
		t.Logf("DeviceGetUsedMemory: %v", ret)
		t.Logf("  used memory (MEM_TYPE_FIRST): %v", usedMemFirst)
	}
	usedMemLast, ret := DeviceGetUsedMemory(devHandle, MEM_TYPE_LAST)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetUsedMemory: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetUsedMemory: %v", ret)
	} else {
		t.Logf("DeviceGetUsedMemory: %v", ret)
		t.Logf("  used memory (MEM_TYPE_LAST): %v", usedMemLast)
	}

	memUtil, ret := DeviceGetMemoryUtilization(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetMemoryUtilization: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetMemoryUtilization: %v", ret)
	} else {
		t.Logf("DeviceGetMemoryUtilization: %v", ret)
		t.Logf("  memory utilization: %v [%%]", memUtil)
	}

	fanRpm, ret := DeviceGetFanRpms(devHandle, 0)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetFanRpms: %v", ret)
		t.Logf("  fan speed (sensor 0): %v [rpm]", fanRpm)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetFanRpms: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetFanRpms: %v", ret)
	}

	fanSpeed, ret := DeviceGetFanSpeed(devHandle, 0)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetFanSpeed: %v", ret)
		t.Logf("  fan speed (sensor 0): %v", fanSpeed)
		t.Logf("  fan speed (sensor 0): %v [%%]", float64(fanSpeed)/MAX_FAN_SPEED)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetFanSpeed: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetFanSpeed: %v", ret)
	}

	maxFanSpeed, ret := DeviceGetMaxFanSpeed(devHandle, 0)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetMaxFanSpeed: %v", ret)
		t.Logf("  max fan speed (sensor 0): %v", maxFanSpeed)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetMaxFanSpeed: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetMaxFanSpeed: %v", ret)
	}

	tempCurMem, ret := DeviceGetTemperatureMetric(devHandle, TEMP_TYPE_EDGE, TEMP_CURRENT)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetTemperatureMetric: %v", ret)
		t.Logf("  temperature (memory, current): %v [millidegrees Celcius]", tempCurMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetTemperatureMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetTemperatureMetric: %v", ret)
	}
	tempMinMem, ret := DeviceGetTemperatureMetric(devHandle, TEMP_TYPE_EDGE, TEMP_MIN)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetTemperatureMetric: %v", ret)
		t.Logf("  temperature (memory, min): %v [millidegrees Celcius]", tempMinMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetTemperatureMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetTemperatureMetric: %v", ret)
	}
	tempMaxMem, ret := DeviceGetTemperatureMetric(devHandle, TEMP_TYPE_EDGE, TEMP_MAX)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetTemperatureMetric: %v", ret)
		t.Logf("  temperature (memory, max): %v [millidegrees Celcius]", tempMaxMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetTemperatureMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetTemperatureMetric: %v", ret)
	}

	voltCurMem, ret := DeviceGetVoltageMetric(devHandle, VOLT_TYPE_VDDGFX, VOLT_FIRST)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetVoltageMetric: %v", ret)
		t.Logf("  voltage (vddgfx, current): %v [millivolts]", voltCurMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVoltageMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetVoltageMetric: %v", ret)
	}
	voltMinMem, ret := DeviceGetVoltageMetric(devHandle, VOLT_TYPE_VDDGFX, VOLT_MIN)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetVoltageMetric: %v", ret)
		t.Logf("  voltage (vddgfx, min): %v [millivolts]", voltMinMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVoltageMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetVoltageMetric: %v", ret)
	}
	voltMaxMem, ret := DeviceGetVoltageMetric(devHandle, VOLT_TYPE_VDDGFX, VOLT_MAX)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetVoltageMetric: %v", ret)
		t.Logf("  voltage (vddgfx, max): %v [millivolts]", voltMaxMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVoltageMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetVoltageMetric: %v", ret)
	}

	busy, ret := DeviceGetBusyPercent(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetBusyPercent: %v", ret)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetBusyPercent: %v (NOT SUPPORTED)", ret)
	} else {
		t.Logf("DeviceGetBusyPercent: %v", ret)
		t.Logf("  utilization: %v [%%]", busy)
	}

	perfLevel, ret := DeviceGetPerfLevel(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPerfLevel: %v", ret)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPerfLevel: %v (NOT SUPPORTED)", ret)
	} else {
		t.Logf("DeviceGetPerfLevel: %v", ret)
		t.Logf("  perf level: %v ", int(perfLevel))
	}

	overdriveLevel, ret := DeviceGetOverdriveLevel(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetOverdriveLevel: %v", ret)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetOverdriveLevel: %v (NOT SUPPORTED)", ret)
	} else {
		t.Logf("DeviceGetOverdriveLevel: %v", ret)
		t.Logf("  overdrive level: %v ", overdriveLevel)
	}

	vbiosVer, ret := DeviceGetVbiosVersionString(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVbiosVersionString: %v", ret)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVbiosVersionString: %v (NOT SUPPORTED)", ret)
	} else {
		t.Logf("DeviceGetVbiosVersionString: %v", ret)
		t.Logf("  VBIOS version: %v", vbiosVer)
	}

	asdVer, ret := DeviceGetFirmwareVersion(devHandle, FW_BLOCK_ASD)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetFirmwareVersion: %v", ret)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetFirmwareVersion: %v (NOT SUPPORTED)", ret)
	} else {
		t.Logf("DeviceGetFirmwareVersion: %v", ret)
		t.Logf("  FW_BLOCK_ASD firmware version: %v", asdVer)
	}

	eccCounts, ret := DeviceGetEccCount(devHandle, GPU_BLOCK_FIRST)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetEccCount: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetEccCount: %v", ret)
	} else {
		t.Logf("DeviceGetEccCount: %v", ret)
		t.Logf("  ECC errors correctable: %v", eccCounts.Correctable_err)
		t.Logf("  ECC errors uncorrectable: %v", eccCounts.Uncorrectable_err)
	}

	eccMask, ret := DeviceGetEccMask(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetEccMask: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetEccMask: %v", ret)
	} else {
		t.Logf("DeviceGetEccMask: %v", ret)
		t.Logf("  ECC mask: %v", eccMask)
	}

	eccStatus, ret := DeviceGetEccStatus(devHandle, GPU_BLOCK_FIRST)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetEccStatus: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetEccStatus: %v", ret)
	} else {
		t.Logf("DeviceGetEccStatus: %v", ret)
		t.Logf("  ECC status: %v", eccStatus)
	}

	reservedPages, ret := DeviceGetMemoryReservedPages(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetMemoryReservedPages: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetMemoryReservedPages: %v", ret)
	} else {
		t.Logf("DeviceGetMemoryReservedPages: %v", ret)
		t.Logf("  Num Reserved Pages: %v", len(reservedPages))
		for _, r := range reservedPages {
			t.Logf("  Page Addr: %v Page Size: %v Page Status: %v", r.Address, r.Size, r.Status)
		}
	}

	clock, ret := DeviceGetClockFrequency(devHandle, CLK_TYPE_FIRST)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetClockFrequency: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetClockFrequency: %v", ret)
	} else {
		t.Logf("DeviceGetClockFrequency: %v", ret)
		t.Logf("  Clock supported: %v", clock.Supported)
		for i := 0; i < int(clock.Supported); i++ {
			if i != int(clock.Current) {
				t.Logf("  Clock current: %v", clock.Frequency[i])
			} else {
				t.Logf("  Clock current: %v (*)", clock.Frequency[i])
			}
		}
	}

	utilCounts, utilStamp, ret := DeviceGetUtilizationCounters(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetUtilizationCounters: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetUtilizationCounters: %v", ret)
	} else {
		t.Logf("DeviceGetUtilizationCounters: %v", ret)
		t.Logf("  Timestamp: %v", utilStamp)
		for _, counter := range utilCounts {
			t.Logf("  Utilization (Type %v): %v", counter.Type, counter.Value)
		}
	}

	freqVoltCurve, ret := DeviceGetVoltageFrequencyCurve(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVoltageFrequencyCurve: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVoltageFrequencyCurve: %v", ret)
	} else {
		t.Logf("DeviceGetVoltageFrequencyCurve: %v", ret)
		t.Logf("  NumRegions: %v", freqVoltCurve.Num_regions)
		t.Logf("  Current SCLK: %v - %v (Limits: %v - %v)", freqVoltCurve.Curr_sclk_range.Lower_bound, freqVoltCurve.Curr_sclk_range.Upper_bound, freqVoltCurve.Sclk_freq_limits.Lower_bound, freqVoltCurve.Sclk_freq_limits.Upper_bound)
		t.Logf("  Current MCLK: %v - %v (Limits: %v - %v)", freqVoltCurve.Curr_mclk_range.Lower_bound, freqVoltCurve.Curr_mclk_range.Upper_bound, freqVoltCurve.Mclk_freq_limits.Lower_bound, freqVoltCurve.Mclk_freq_limits.Upper_bound)
	}

	freqVoltCurveRegions, ret := DeviceGetVoltageFrequencyCurveRegions(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetVoltageFrequencyCurveRegions: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVoltageFrequencyCurveRegions: %v", ret)
	} else {
		t.Logf("DeviceGetVoltageFrequencyCurveRegions: %v", ret)
		t.Logf("  NumRegions: %v", len(freqVoltCurveRegions))
		for i, region := range freqVoltCurveRegions {
			t.Logf("  Region %v Freq (%v - %v) Voltage (%v - %v)", i, region.Freq_range.Lower_bound, region.Freq_range.Upper_bound, region.Volt_range.Lower_bound, region.Volt_range.Upper_bound)
		}
	}

	metrics, ret := DeviceGetMetrics(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetMetrics: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetMetrics: %v", ret)
	} else {
		t.Logf("DeviceGetMetrics: %v", ret)
		t.Logf("  Average GFX Activity: %v", metrics.Average_gfx_activity)
		t.Logf("  Average UMC Activity: %v", metrics.Average_umc_activity)
		t.Logf("  Average MM Activity: %v", metrics.Average_mm_activity)
		t.Logf("  Average Socket Power: %v", metrics.Average_socket_power)
	}

	node, ret := DeviceGetNumaNode(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetNumaNode: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetNumaNode: %v", ret)
	} else {
		t.Logf("DeviceGetNumaNode: %v", ret)
		t.Logf("  NUMA node: %v", node)
	}

	xgmiStatus, ret := DeviceXgmiErrorStatus(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceXgmiErrorStatus: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceXgmiErrorStatus: %v", ret)
	} else {
		t.Logf("DeviceXgmiErrorStatus: %v", ret)
		t.Logf("  XGMI error status: %v", xgmiStatus)
	}

	xgmiHive, ret := DeviceXgmiHiveId(devHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceXgmiHiveId: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceXgmiHiveId: %v", ret)
	} else {
		t.Logf("DeviceXgmiHiveId: %v", ret)
		t.Logf("  XGMI hive id: %v", xgmiHive)
	}

	powerProf, ret := DeviceGetPowerProfile(devHandle, 0)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPowerProfile: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPowerProfile: %v", ret)
	} else {
		t.Logf("DeviceGetPowerProfile: %v", ret)
		t.Logf("  current power profile: %v", powerProf.Current)
		t.Logf("  available power profiles: %v", powerProf.Available_profiles)
		t.Logf("  number of power profiles: %v", powerProf.Num_profiles)
	}
}

func TestPerfLevel(t *testing.T) {
	Init()
	defer Shutdown()

	deviceCount, ret := NumMonitorDevices()
	if ret != STATUS_SUCCESS {
		t.Errorf("NumMonitorDevices: %v", ret)
	} else {
		t.Logf("NumMonitorDevices: %v", ret)
		t.Logf("  count: %v", deviceCount)
	}
	if deviceCount == 0 {
		t.Skip("Skipping test because it requires one GPU.")
	}
	devHandle, ret := DeviceGetHandleByIndex(0)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetHandleByIndex(0): %v", ret)
	} else {
		t.Logf("DeviceGetHandleByIndex(0): %v", ret)
	}

	perfLevelBefore, ret := DeviceGetPerfLevel(devHandle)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetPerfLevel: %v", ret)		
		t.Logf("  perf level (before): %v ", int(perfLevelBefore))
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPerfLevel: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetPerfLevel: %v", ret)
	}
	perfLevelTest := DEV_PERF_LEVEL_HIGH
	if (perfLevelBefore == perfLevelTest) {
		perfLevelTest = DEV_PERF_LEVEL_LOW
	}

	ret = DeviceSetPerfLevel(devHandle, perfLevelTest)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceSetPerfLevel: %v", ret)		
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceSetPerfLevel: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceSetPerfLevel: %v", ret)
	}

	perfLevelRead, ret := DeviceGetPerfLevel(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPerfLevel: %v", ret)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPerfLevel: %v (NOT SUPPORTED)", ret)
	} else {
		t.Logf("DeviceGetPerfLevel: %v", ret)
		t.Logf("  new perf level: %v", int(perfLevelRead))
	}

	ret = DeviceSetPerfLevel(devHandle, perfLevelBefore)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceSetPerfLevel: %v", ret)		
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceSetPerfLevel: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceSetPerfLevel: %v", ret)
	}

	perfLevelRead2, ret := DeviceGetPerfLevel(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPerfLevel: %v", ret)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetPerfLevel: %v (NOT SUPPORTED)", ret)
	} else {
		t.Logf("DeviceGetPerfLevel: %v", ret)
		t.Logf("  reset perf level: %v", int(perfLevelRead2))
	}

	if perfLevelBefore != perfLevelRead2 {
		t.Errorf("Failed to reset device to: %v", int(perfLevelBefore))
	}
}

//func TestOverdriveLevel(t *testing.T) {
//	Init()
//	defer Shutdown()

//	deviceCount, ret := NumMonitorDevices()
//	if ret != STATUS_SUCCESS {
//		t.Errorf("NumMonitorDevices: %v", ret)
//	} else {
//		t.Logf("NumMonitorDevices: %v", ret)
//		t.Logf("  count: %v", deviceCount)
//	}
//	if deviceCount == 0 {
//		t.Skip("Skipping test because it requires one GPU.")
//	}
//	devHandle, ret := DeviceGetHandleByIndex(0)
//	if ret != STATUS_SUCCESS {
//		t.Errorf("DeviceGetHandleByIndex(0): %v", ret)
//	} else {
//		t.Logf("DeviceGetHandleByIndex(0): %v", ret)
//	}

//	overdriveBefore, ret := DeviceGetOverdriveLevel(devHandle)
//	if ret == STATUS_SUCCESS {
//		t.Logf("DeviceGetOverdriveLevel: %v", ret)		
//		t.Logf("  perf overdrive (before): %v ", int(overdriveBefore))
//	} else if ret == STATUS_NOT_SUPPORTED {
//		t.Logf("DeviceGetOverdriveLevel: %v (NOT SUPPORTED)", ret)
//	} else {
//		t.Errorf("DeviceGetOverdriveLevel: %v", ret)
//	}
//	overdriveTest := overdriveBefore/2
//	if (overdriveBefore == overdriveTest) {
//		overdriveTest = overdriveBefore*2
//	}

//	ret = DeviceSetOverdriveLevel(devHandle, overdriveTest)
//	if ret == STATUS_SUCCESS {
//		t.Logf("DeviceSetOverdriveLevel: %v", ret)		
//	} else if ret == STATUS_NOT_SUPPORTED {
//		t.Logf("DeviceSetOverdriveLevel: %v (NOT SUPPORTED)", ret)
//	} else {
//		t.Errorf("DeviceSetOverdriveLevel: %v", ret)
//	}

//	overdriveRead, ret := DeviceGetOverdriveLevel(devHandle)
//	if ret != STATUS_SUCCESS {
//		t.Errorf("DeviceGetOverdriveLevel: %v", ret)
//	} else if ret == STATUS_NOT_SUPPORTED {
//		t.Logf("DeviceGetOverdriveLevel: %v (NOT SUPPORTED)", ret)
//	} else {
//		t.Logf("DeviceGetOverdriveLevel: %v", ret)
//		t.Logf("  new overdrive level: %v", int(overdriveRead))
//	}

//	ret = DeviceSetOverdriveLevel(devHandle, overdriveBefore)
//	if ret == STATUS_SUCCESS {
//		t.Logf("DeviceSetOverdriveLevel: %v", ret)		
//	} else if ret == STATUS_NOT_SUPPORTED {
//		t.Logf("DeviceSetOverdriveLevel: %v (NOT SUPPORTED)", ret)
//	} else {
//		t.Errorf("DeviceSetOverdriveLevel: %v", ret)
//	}

//	overdriveRead2, ret := DeviceGetOverdriveLevel(devHandle)
//	if ret != STATUS_SUCCESS {
//		t.Errorf("DeviceGetOverdriveLevel: %v", ret)
//	} else if ret == STATUS_NOT_SUPPORTED {
//		t.Logf("DeviceGetOverdriveLevel: %v (NOT SUPPORTED)", ret)
//	} else {
//		t.Logf("DeviceGetOverdriveLevel: %v", ret)
//		t.Logf("  reset overdrive level: %v", int(overdriveRead2))
//	}

//	if overdriveBefore != overdriveRead2 {
//		t.Errorf("Failed to reset device to: %v", int(overdriveBefore))
//	}
//}

func TestMultiDevice(t *testing.T) {
	Init()
	defer Shutdown()

	deviceCount, ret := NumMonitorDevices()
	if ret != STATUS_SUCCESS {
		t.Errorf("NumMonitorDevices: %v", ret)
	} else {
		t.Logf("NumMonitorDevices: %v", ret)
		t.Logf("  count: %v", deviceCount)
	}

	if deviceCount < 2 {
		t.Skip("Skipping test because it requires at least two GPUs.")
	}

	firstHandle, ret := DeviceGetHandleByIndex(0)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetHandleByIndex(0): %v", ret)
	} else {
		t.Logf("DeviceGetHandleByIndex(0): %v", ret)
	}

	secondHandle, ret := DeviceGetHandleByIndex(1)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetHandleByIndex(1): %v", ret)
	} else {
		t.Logf("DeviceGetHandleByIndex(1): %v", ret)
	}

	linkWidth, linkType, ret := DeviceGetLinkType(firstHandle, secondHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetLinkType: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetLinkType: %v", ret)
	} else {
		t.Logf("DeviceGetLinkType: %v", ret)
		t.Logf("  link width: %v", linkWidth)
		t.Logf("  link type: %v", linkType)
	}

	linkWeight, ret := DeviceGetLinkWeight(firstHandle, secondHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetLinkWeight: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetLinkWeight: %v", ret)
	} else {
		t.Logf("DeviceGetLinkWeight: %v", ret)
		t.Logf("  link weight: %v", linkWeight)
	}

	p2p, ret := DeviceIsP2PAccessible(firstHandle, secondHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceIsP2PAccessible: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceIsP2PAccessible: %v", ret)
	} else {
		t.Logf("DeviceIsP2PAccessible: %v", ret)
		t.Logf("  P2P accessible: %v", p2p)
	}

	minBand, maxBand, ret := DeviceGetMinMaxBandwidth(firstHandle, secondHandle)
	if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetMinMaxBandwidth: %v (NOT SUPPORTED)", ret)
	} else if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetMinMaxBandwidth: %v", ret)
	} else {
		t.Logf("DeviceGetMinMaxBandwidth: %v", ret)
		t.Logf("  minimal bandwidth: %v", minBand)
		t.Logf("  maximal bandwidth: %v", maxBand)
	}
}

