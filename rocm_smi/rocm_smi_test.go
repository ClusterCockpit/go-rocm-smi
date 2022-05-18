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

	success, ret := StatusString(STATUS_SUCCESS)
	if ret != STATUS_SUCCESS {
		t.Errorf("StatusString(STATUS_SUCCESS): %v", ret)
	} else {
		t.Logf("StatusString(STATUS_SUCCESS): %v", ret)
		t.Logf("  str: %v", success)
	}

	invalArgs, ret := StatusString(STATUS_INVALID_ARGS)
	if ret != STATUS_SUCCESS {
		t.Errorf("StatusString(STATUS_INVALID_ARGS): %v", ret)
	} else {
		t.Logf("StatusString(STATUS_INVALID_ARGS): %v", ret)
		t.Logf("  str: %v", invalArgs)
	}

	notSupp, ret := StatusString(STATUS_NOT_SUPPORTED)
	if ret != STATUS_SUCCESS {
		t.Errorf("StatusString(STATUS_NOT_SUPPORTED): %v", ret)
	} else {
		t.Logf("StatusString(STATUS_NOT_SUPPORTED): %v", ret)
		t.Logf("  str: %v", notSupp)
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

func TestUnit(t *testing.T) {
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
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetHandleByIndex: %v", ret)
	} else {
		t.Logf("DeviceGetHandleByIndex: %v", ret)
		t.Logf("  unit: %v", devHandle)
	}

	brandName, ret := DeviceGetBrand(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetBrand: %v", ret)
	} else {
		t.Logf("DeviceGetBrand: %v", ret)
		t.Logf("  brand: %v", brandName)
	}

	name, ret := DeviceGetName(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetName: %v", ret)
	} else {
		t.Logf("DeviceGetName: %v", ret)
		t.Logf("  name: %v", name)
	}

	vendId, ret := DeviceGetVendorId(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVendorId: %v", ret)
	} else {
		t.Logf("DeviceGetVendorId: %v", ret)
		t.Logf("  name: %v", vendId)
	}

	vendName, ret := DeviceGetVendorName(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVendorName: %v", ret)
	} else {
		t.Logf("DeviceGetVendorName: %v", ret)
		t.Logf("  vendor name: %v", vendName)
	}

	vramName, ret := DeviceGetVramVendor(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVramVendor: %v", ret)
	} else {
		t.Logf("DeviceGetVramVendor: %v", ret)
		t.Logf("  vram name: %v", vramName)
	}

	serial, ret := DeviceGetSerial(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetSerial: %v", ret)
	} else {
		t.Logf("DeviceGetSerial: %v", ret)
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
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetSubsystemId: %v", ret)
	} else {
		t.Logf("DeviceGetSubsystemId: %v", ret)
		t.Logf("  subsystem id: %v", subId)
	}

	subSys, ret := DeviceGetSubsystemName(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetSubsystemName: %v", ret)
	} else {
		t.Logf("DeviceGetSubsystemName: %v", ret)
		t.Logf("  subsystem: %v", subSys)
	}

	pciId, ret := DeviceGetPciId(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciId: %v", ret)
	} else {
		t.Logf("DeviceGetPciId: %v", ret)
		t.Logf("  pci id: 0x%X", pciId)
	}

	pciInfo, ret := DeviceGetPciInfo(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciId: %v", ret)
	} else {
		t.Logf("DeviceGetPciId: %v", ret)
		t.Logf("  pci domain: 0x%X", pciInfo.Domain)
		t.Logf("  pci bus: 0x%X", pciInfo.Bus)
		t.Logf("  pci device: 0x%X", pciInfo.Device)
		t.Logf("  pci function: 0x%X", pciInfo.Function)
	}

	pciBand, ret := DeviceGetPciBandwidth(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciBandwidth: %v", ret)
	} else {
		t.Logf("DeviceGetPciBandwidth: %v", ret)
		t.Logf("  pci band: %v", pciBand)
	}

	sent, recv, max_pkts_size, ret := DeviceGetPciThroughput(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciThroughput: %v", ret)
	} else {
		t.Logf("DeviceGetPciThroughput: %v", ret)
		t.Logf("  pci throughput sent: %v", sent)
		t.Logf("  pci throughput recv: %v", recv)
		t.Logf("  pci throughput max_pkts_size: %v", max_pkts_size)
	}

	replay, ret := DeviceGetPciReplayCounter(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPciReplayCounter: %v", ret)
	} else {
		t.Logf("DeviceGetPciReplayCounter: %v", ret)
		t.Logf("  replay counter: %v", replay)
	}

	numaAffinity, ret := DeviceGetNumaAffinity(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetNumaAffinity: %v", ret)
	} else {
		t.Logf("DeviceGetNumaAffinity: %v", ret)
		t.Logf("  numa node: %v", numaAffinity)
	}

	avgPower, ret := DeviceGetPowerAverage(devHandle, 0)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPowerAverage: %v", ret)
	} else {
		t.Logf("DeviceGetPowerAverage: %v", ret)
		t.Logf("  average power (sensor 0): %v [microwatts]", avgPower)
	}

	capPower, ret := DeviceGetPowerCap(devHandle, 0)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPowerCap: %v", ret)
	} else {
		t.Logf("DeviceGetPowerCap: %v", ret)
		t.Logf("  power capping (sensor 0): %v [microwatts]", capPower)
	}

	defCapPower, ret := DeviceGetDefaultPowerCap(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetDefaultPowerCap: %v", ret)
	} else {
		t.Logf("DeviceGetDefaultPowerCap: %v", ret)
		t.Logf("  power capping (default): %v [microwatts]", defCapPower)
	}

	capPowerMax, capPowerMin, ret := DeviceGetPowerCapRange(devHandle, 0)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPowerCapRange: %v", ret)
	} else {
		t.Logf("DeviceGetPowerCapRange: %v", ret)
		t.Logf("  power capping (sensor 0): %v - %v [microwatts]", capPowerMin, capPowerMax)
	}

	energy, resolution, timestamp, ret := DeviceGetEnergyCount(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetEnergyCount: %v", ret)
	} else {
		t.Logf("DeviceGetEnergyCount: %v", ret)
		t.Logf("  energy: %v [microJoules]", energy)
		t.Logf("  resolution: %v [microJoules]", resolution)
		t.Logf("  timestamp: %v [ns]", timestamp)
	}

	totalMemFirst, ret := DeviceGetTotalMemory(devHandle, MEM_TYPE_FIRST)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetTotalMemory: %v", ret)
	} else {
		t.Logf("DeviceGetTotalMemory: %v", ret)
		t.Logf("  total memory (MEM_TYPE_FIRST): %v", totalMemFirst)
	}
	totalMemLast, ret := DeviceGetTotalMemory(devHandle, MEM_TYPE_LAST)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetTotalMemory: %v", ret)
	} else {
		t.Logf("DeviceGetTotalMemory: %v", ret)
		t.Logf("  total memory (MEM_TYPE_LAST): %v", totalMemLast)
	}

	usedMemFirst, ret := DeviceGetUsedMemory(devHandle, MEM_TYPE_FIRST)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetUsedMemory: %v", ret)
	} else {
		t.Logf("DeviceGetUsedMemory: %v", ret)
		t.Logf("  used memory (MEM_TYPE_FIRST): %v", usedMemFirst)
	}
	usedMemLast, ret := DeviceGetUsedMemory(devHandle, MEM_TYPE_LAST)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetUsedMemory: %v", ret)
	} else {
		t.Logf("DeviceGetUsedMemory: %v", ret)
		t.Logf("  used memory (MEM_TYPE_LAST): %v", usedMemLast)
	}

	memUtil, ret := DeviceGetMemoryUtilization(devHandle)
	if ret != STATUS_SUCCESS {
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

	tempCurMem, ret := DeviceGetTemperatureMetric(devHandle, TEMP_TYPE_MEMORY, TEMP_CURRENT)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetTemperatureMetric: %v", ret)
		t.Logf("  temperature (memory, current): %v [millidegrees Celcius]", tempCurMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetTemperatureMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetTemperatureMetric: %v", ret)
	}
	tempMinMem, ret := DeviceGetTemperatureMetric(devHandle, TEMP_TYPE_MEMORY, TEMP_MIN)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetTemperatureMetric: %v", ret)
		t.Logf("  temperature (memory, min): %v [millidegrees Celcius]", tempMinMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetTemperatureMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetTemperatureMetric: %v", ret)
	}
	tempMaxMem, ret := DeviceGetTemperatureMetric(devHandle, TEMP_TYPE_MEMORY, TEMP_MAX)
	if ret == STATUS_SUCCESS {
		t.Logf("DeviceGetTemperatureMetric: %v", ret)
		t.Logf("  temperature (memory, max): %v [millidegrees Celcius]", tempMaxMem)
	} else if ret == STATUS_NOT_SUPPORTED {
		t.Logf("DeviceGetTemperatureMetric: %v (NOT SUPPORTED)", ret)
	} else {
		t.Errorf("DeviceGetTemperatureMetric: %v", ret)
	}

	voltCurMem, ret := DeviceGetVoltageMetric(devHandle, VOLT_TYPE_VDDGFX, VOLT_CURRENT)
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
	} else {
		t.Logf("DeviceGetBusyPercent: %v", ret)
		t.Logf("  utilization: %v [%%]", busy)
	}

	perfLevel, ret := DeviceGetPerfLevel(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetPerfLevel: %v", ret)
	} else {
		t.Logf("DeviceGetPerfLevel: %v", ret)
		t.Logf("  perf level: %v ", int(perfLevel))
	}

	overdriveLevel, ret := DeviceGetOverdriveLevel(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetOverdriveLevel: %v", ret)
	} else {
		t.Logf("DeviceGetOverdriveLevel: %v", ret)
		t.Logf("  overdrive level: %v ", overdriveLevel)
	}

	vbiosVer, ret := DeviceGetVbiosVersionString(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetVbiosVersionString: %v", ret)
	} else {
		t.Logf("DeviceGetVbiosVersionString: %v", ret)
		t.Logf("  VBIOS version: %v", vbiosVer)
	}

	asdVer, ret := DeviceGetFirmwareVersion(devHandle, FW_BLOCK_ASD)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetFirmwareVersion: %v", ret)
	} else {
		t.Logf("DeviceGetFirmwareVersion: %v", ret)
		t.Logf("  FW_BLOCK_ASD firmware version: %v", asdVer)
	}

	eccCounts, ret := DeviceGetEccCount(devHandle, GPU_BLOCK_FIRST)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetEccCount: %v", ret)
	} else {
		t.Logf("DeviceGetEccCount: %v", ret)
		t.Logf("  ECC errors correctable: %v", eccCounts.Correctable_err)
		t.Logf("  ECC errors uncorrectable: %v", eccCounts.Uncorrectable_err)
	}

	eccMask, ret := DeviceGetEccMask(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetEccMask: %v", ret)
	} else {
		t.Logf("DeviceGetEccMask: %v", ret)
		t.Logf("  ECC mask: %v", eccMask)
	}

	eccStatus, ret := DeviceGetEccStatus(devHandle, GPU_BLOCK_FIRST)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetEccStatus: %v", ret)
	} else {
		t.Logf("DeviceGetEccStatus: %v", ret)
		t.Logf("  ECC status: %v", eccStatus)
	}

	reservedPages, ret := DeviceGetMemoryReservedPages(devHandle)
	if ret != STATUS_SUCCESS {
		t.Errorf("DeviceGetMemoryReservedPages: %v", ret)
	} else {
		t.Logf("DeviceGetMemoryReservedPages: %v", ret)
		t.Logf("  Num Reserved Pages: %v", len(reservedPages))
		for _, r := range reservedPages {
			t.Logf("  Page Addr: %v Page Size: %v Page Status: %v", r.Address, r.Size, r.Status)
		}
	}

	clock, ret := DeviceGetClockFrequency(devHandle, CLK_TYPE_FIRST)
	if ret != STATUS_SUCCESS {
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
}
