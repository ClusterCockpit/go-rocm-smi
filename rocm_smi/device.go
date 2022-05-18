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

//import "fmt"

// rocm_smi.NumMonitorDevices()
func NumMonitorDevices() (int, RSMI_status) {
	var DeviceCount uint32
	ret := rsmi_num_monitor_devices(&DeviceCount)
	return int(DeviceCount), ret
}

// rocm_smi.DeviceGetHandleByIndex()
func DeviceGetHandleByIndex(Index int) (DeviceHandle, RSMI_status) {
	var device uint16
	var index uint32 = uint32(Index)
	ret := rsmi_dev_id_get(index, &device)
	return DeviceHandle{
		handle: device,
		index:  index,
	}, ret
}

// rocm_smi.DeviceGetBrand()
func DeviceGetBrand(Device DeviceHandle) (string, RSMI_status) {
	var brand []byte = make([]byte, 100)
	bptr := &brand[0]
	ret := rsmi_dev_brand_get(Device.index, bptr, 100)
	return string(brand), ret
}

func (Device DeviceHandle) GetBrand() (string, RSMI_status) {
	return DeviceGetBrand(Device)
}

// rocm_smi.DeviceGetName()
func DeviceGetName(Device DeviceHandle) (string, RSMI_status) {
	var name []byte = make([]byte, 100)
	nptr := &name[0]
	ret := rsmi_dev_name_get(Device.index, nptr, 100)
	return string(name), ret
}

func (Device DeviceHandle) GetName() (string, RSMI_status) {
	return DeviceGetName(Device)
}

// rocm_smi.DeviceGetName()

var DeviceGetSku = DeviceGetSkuFake

func DeviceGetSkuReal(Device DeviceHandle) (string, RSMI_status) {
	var Sku []byte = make([]byte, 100)
	sptr := &Sku[0]
	ret := rsmi_dev_sku_get(Device.index, sptr)
	return string(Sku), ret
}

func DeviceGetSkuFake(Device DeviceHandle) (string, RSMI_status) {
	return "NA", STATUS_NOT_SUPPORTED
}

func (Device DeviceHandle) GetSku() (string, RSMI_status) {
	return DeviceGetSku(Device)
}

// rocm_smi.DeviceGetVendorName()
func DeviceGetVendorName(Device DeviceHandle) (string, RSMI_status) {
	var Name []byte = make([]byte, 100)
	nptr := &Name[0]
	ret := rsmi_dev_vendor_name_get(Device.index, nptr, 100)
	return string(Name), ret
}

func (Device DeviceHandle) GetVendorName() (string, RSMI_status) {
	return DeviceGetVendorName(Device)
}

// rocm_smi.DeviceGetVendorId()
func DeviceGetVendorId(Device DeviceHandle) (uint16, RSMI_status) {
	var id uint16
	ret := rsmi_dev_vendor_id_get(Device.index, &id)
	return id, ret
}

func (Device DeviceHandle) GetVendorId() (uint16, RSMI_status) {
	return DeviceGetVendorId(Device)
}

// rocm_smi.DeviceGetVramVendor()
func DeviceGetVramVendor(Device DeviceHandle) (string, RSMI_status) {
	var Name []byte = make([]byte, 100)
	nptr := &Name[0]
	ret := rsmi_dev_vram_vendor_get(Device.index, nptr, 100)
	return string(Name), ret
}

func (Device DeviceHandle) GetVramVendor() (string, RSMI_status) {
	return DeviceGetVramVendor(Device)
}

// rocm_smi.DeviceGetSerial()
func DeviceGetSerial(Device DeviceHandle) (string, RSMI_status) {
	var Serial []byte = make([]byte, 100)
	sptr := &Serial[0]
	ret := rsmi_dev_serial_number_get(Device.index, sptr, 100)
	return string(Serial), ret
}

func (Device DeviceHandle) DeviceGetSerial() (string, RSMI_status) {
	return DeviceGetSerial(Device)
}

// rocm_smi.DeviceGetSubsystem()
func DeviceGetSubsystemName(Device DeviceHandle) (string, RSMI_status) {
	var Name []byte = make([]byte, 100)
	nptr := &Name[0]
	ret := rsmi_dev_subsystem_name_get(Device.index, nptr, 100)
	return string(Name), ret
}

func (Device DeviceHandle) GetSubsystemName() (string, RSMI_status) {
	return DeviceGetSubsystemName(Device)
}

// rocm_smi.DeviceGetSubsystemId()
func DeviceGetSubsystemId(Device DeviceHandle) (uint16, RSMI_status) {
	var id uint16
	ret := rsmi_dev_subsystem_id_get(Device.index, &id)
	return id, ret
}

func (Device DeviceHandle) GetSubsystemId() (uint16, RSMI_status) {
	return DeviceGetSubsystemId(Device)
}

// rocm_smi.DeviceGetUniqueId()
func DeviceGetUniqueId(Device DeviceHandle) (uint32, RSMI_status) {
	var id uint32
	ret := rsmi_dev_unique_id_get(Device.index, &id)
	return id, ret
}

func (Device DeviceHandle) GetUniqueId() (uint32, RSMI_status) {
	return DeviceGetUniqueId(Device)
}

// rocm_smi.DeviceGetPciId()
func DeviceGetPciId(Device DeviceHandle) (uint32, RSMI_status) {
	var id uint32
	ret := rsmi_dev_pci_id_get(Device.index, &id)
	return id, ret
}

func (Device DeviceHandle) GetPciId() (uint32, RSMI_status) {
	return DeviceGetPciId(Device)
}

type Pci_info struct {
	Domain   uint32
	Bus      uint8
	Device   uint8
	Function uint8
}

// rocm_smi.DeviceGetPciInfo()
// own addition
func DeviceGetPciInfo(Device DeviceHandle) (Pci_info, RSMI_status) {
	var id uint32
	info := Pci_info{
		Domain:   0,
		Bus:      0,
		Device:   0,
		Function: 0,
	}
	ret := rsmi_dev_pci_id_get(Device.index, &id)
	if ret == STATUS_SUCCESS {
		info.Domain = (id >> 32) & 0xffffffff
		info.Bus = uint8((id >> 8) & 0xff)
		info.Device = uint8((id >> 3) & 0x1f)
		info.Function = uint8(id & 0x7)
	}
	return info, ret
}

func (Device DeviceHandle) GetPciInfo() (Pci_info, RSMI_status) {
	return DeviceGetPciInfo(Device)
}

// rocm_smi.DeviceGetPciBandwidth()
func DeviceGetPciBandwidth(Device DeviceHandle) (RSMI_pcie_bandwidth, RSMI_status) {
	var info RSMI_pcie_bandwidth
	ret := rsmi_dev_pci_bandwidth_get(Device.index, &info)
	return info, ret
}

func (Device DeviceHandle) GetPciBandwidth() (RSMI_pcie_bandwidth, RSMI_status) {
	return DeviceGetPciBandwidth(Device)
}

// rocm_smi.DeviceSetPciBandwidth()
func DeviceSetPciBandwidth(Device DeviceHandle, Mask uint32) RSMI_status {
	ret := rsmi_dev_pci_bandwidth_set(Device.index, Mask)
	return ret
}

func (Device DeviceHandle) SetPciBandwidth(Mask uint32) RSMI_status {
	return DeviceSetPciBandwidth(Device, Mask)
}

// rocm_smi.DeviceGetPciThroughput()
func DeviceGetPciThroughput(Device DeviceHandle) (uint32, uint32, uint32, RSMI_status) {
	var sent uint32
	var recv uint32
	var max_pkts_size uint32
	ret := rsmi_dev_pci_throughput_get(Device.index, &sent, &recv, &max_pkts_size)
	return sent, recv, max_pkts_size, ret
}

func (Device DeviceHandle) GetPciThroughput() (uint32, uint32, uint32, RSMI_status) {
	return DeviceGetPciThroughput(Device)
}

// rocm_smi.DeviceGetPciReplayCounter()
func DeviceGetPciReplayCounter(Device DeviceHandle) (uint32, RSMI_status) {
	var counter uint32
	ret := rsmi_dev_pci_replay_counter_get(Device.index, &counter)
	return counter, ret
}

func (Device DeviceHandle) GetPciReplayCounter() (uint32, RSMI_status) {
	return DeviceGetPciReplayCounter(Device)
}

// rocm_smi.DeviceGetNumaAffinity()
func DeviceGetNumaAffinity(Device DeviceHandle) (uint32, RSMI_status) {
	var id uint32
	ret := rsmi_topo_numa_affinity_get(Device.index, &id)
	return id, ret
}

func (Device DeviceHandle) GetNumaAffinity() (uint32, RSMI_status) {
	return DeviceGetNumaAffinity(Device)
}

// rocm_smi.DeviceGetPowerAverage()
func DeviceGetPowerAverage(Device DeviceHandle, Sensor uint32) (uint32, RSMI_status) {
	var power uint32
	ret := rsmi_dev_power_ave_get(Device.index, Sensor, &power)
	return power, ret
}

func (Device DeviceHandle) GetPowerAverage(Sensor uint32) (uint32, RSMI_status) {
	return DeviceGetPowerAverage(Device, Sensor)
}

// rocm_smi.DeviceGetPowerCap()
func DeviceGetPowerCap(Device DeviceHandle, Sensor uint32) (uint32, RSMI_status) {
	var power uint32
	ret := rsmi_dev_power_cap_get(Device.index, Sensor, &power)
	return power, ret
}

func (Device DeviceHandle) GetPowerCap(Sensor uint32) (uint32, RSMI_status) {
	return DeviceGetPowerCap(Device, Sensor)
}

// rocm_smi.DeviceGetDefaultPowerCap()
func DeviceGetDefaultPowerCap(Device DeviceHandle) (uint32, RSMI_status) {
	var power uint32
	ret := rsmi_dev_power_cap_default_get(Device.index, &power)
	return power, ret
}

func (Device DeviceHandle) GetDefaultPowerCap() (uint32, RSMI_status) {
	return DeviceGetDefaultPowerCap(Device)
}

// rocm_smi.DeviceGetPowerCapRange()
func DeviceGetPowerCapRange(Device DeviceHandle, Sensor uint32) (uint32, uint32, RSMI_status) {
	var mini uint32
	var maxi uint32
	ret := rsmi_dev_power_cap_range_get(Device.index, Sensor, &maxi, &mini)
	return maxi, mini, ret
}

func (Device DeviceHandle) GetPowerCapRange(Sensor uint32) (uint32, uint32, RSMI_status) {
	return DeviceGetPowerCapRange(Device, Sensor)
}

// rocm_smi.DeviceGetEnergyCount()
func DeviceGetEnergyCount(Device DeviceHandle) (uint32, float32, uint32, RSMI_status) {
	var power uint32
	var resolution float32
	var timestamp uint32
	ret := rsmi_dev_energy_count_get(Device.index, &power, &resolution, &timestamp)
	return power, resolution, timestamp, ret
}

func (Device DeviceHandle) GetEnergyCount() (uint32, float32, uint32, RSMI_status) {
	return DeviceGetEnergyCount(Device)
}

// rocm_smi.DeviceSetPowerCap()
func DeviceSetPowerCap(Device DeviceHandle, Sensor uint32, Mask uint32) RSMI_status {
	ret := rsmi_dev_power_cap_set(Device.index, Sensor, Mask)
	return ret
}

func (Device DeviceHandle) SetPowerCap(Sensor uint32, Mask uint32) RSMI_status {
	return DeviceSetPowerCap(Device, Sensor, Mask)
}

// rocm_smi.DeviceSetPowerProfile
func DeviceSetPowerProfile(Device DeviceHandle, Reserved uint32, Preset RSMI_power_profile_preset_masks) RSMI_status {
	ret := rsmi_dev_power_profile_set(Device.index, Reserved, Preset)
	return ret
}

func (Device DeviceHandle) SetPowerProfile(Reserved uint32, Preset RSMI_power_profile_preset_masks) RSMI_status {
	return DeviceSetPowerProfile(Device, Reserved, Preset)
}

// rocm_smi.DeviceGetTotalMemory()
func DeviceGetTotalMemory(Device DeviceHandle, Type RSMI_memory_type) (uint32, RSMI_status) {
	var size uint32
	ret := rsmi_dev_memory_total_get(Device.index, Type, &size)
	return size, ret
}

func (Device DeviceHandle) GetTotalMemory(Type RSMI_memory_type) (uint32, RSMI_status) {
	return DeviceGetTotalMemory(Device, Type)
}

// rocm_smi.DeviceGetUsedMemory()
func DeviceGetUsedMemory(Device DeviceHandle, Type RSMI_memory_type) (uint32, RSMI_status) {
	var size uint32
	ret := rsmi_dev_memory_usage_get(Device.index, Type, &size)
	return size, ret
}

func (Device DeviceHandle) GetUsedMemory(Type RSMI_memory_type) (uint32, RSMI_status) {
	return DeviceGetUsedMemory(Device, Type)
}

// rocm_smi.DeviceGetMemoryUtilization()
func DeviceGetMemoryUtilization(Device DeviceHandle) (uint32, RSMI_status) {
	var percent uint32
	ret := rsmi_dev_memory_busy_percent_get(Device.index, &percent)
	return percent, ret
}

func (Device DeviceHandle) GetMemoryUtilization() (uint32, RSMI_status) {
	return DeviceGetMemoryUtilization(Device)
}

// rocm_smi.DeviceGetMemoryReservedPages()
func DeviceGetMemoryReservedPages(Device DeviceHandle) ([]RSMI_retired_page_record, RSMI_status) {
	var num_records uint32
	records := make([]RSMI_retired_page_record, 0)
	ret := rsmi_dev_memory_reserved_pages_get(Device.index, &num_records, nil)
	if ret == STATUS_SUCCESS {
		if num_records > 0 {
			records := make([]RSMI_retired_page_record, num_records)
			ret = rsmi_dev_memory_reserved_pages_get(Device.index, &num_records, &records[0])
		}
	}
	return records, ret
}

func (Device DeviceHandle) GetMemoryReservedPages() ([]RSMI_retired_page_record, RSMI_status) {
	return DeviceGetMemoryReservedPages(Device)
}

// rocm_smi.DeviceGetFanRpms()
func DeviceGetFanRpms(Device DeviceHandle, Sensor uint32) (int32, RSMI_status) {
	var speed int32
	ret := rsmi_dev_fan_rpms_get(Device.index, Sensor, &speed)
	return speed, ret
}

func (Device DeviceHandle) GetFanRpms(Sensor uint32) (int32, RSMI_status) {
	return DeviceGetFanRpms(Device, Sensor)
}

// rocm_smi.DeviceGetFanSpeed()
func DeviceGetFanSpeed(Device DeviceHandle, Sensor uint32) (int32, RSMI_status) {
	var speed int32
	ret := rsmi_dev_fan_speed_get(Device.index, Sensor, &speed)
	return speed, ret
}

func (Device DeviceHandle) GetFanSpeed(Sensor uint32) (int32, RSMI_status) {
	return DeviceGetFanSpeed(Device, Sensor)
}

// rocm_smi.DeviceGetMaxFanSpeed()
func DeviceGetMaxFanSpeed(Device DeviceHandle, Sensor uint32) (uint32, RSMI_status) {
	var speed uint32
	ret := rsmi_dev_fan_speed_max_get(Device.index, Sensor, &speed)
	return speed, ret
}

func (Device DeviceHandle) GetMaxFanSpeed(Sensor uint32) (uint32, RSMI_status) {
	return DeviceGetMaxFanSpeed(Device, Sensor)
}

// rocm_smi.DeviceGetTemperatureMetric()
func DeviceGetTemperatureMetric(Device DeviceHandle, Sensor RSMI_temperature_type, Metric RSMI_temperature_metric) (int32, RSMI_status) {
	var temp int32
	ret := rsmi_dev_temp_metric_get(Device.index, uint32(Sensor), Metric, &temp)
	return temp, ret
}

func (Device DeviceHandle) GetTemperatureMetric(Sensor RSMI_temperature_type, Metric RSMI_temperature_metric) (int32, RSMI_status) {
	return DeviceGetTemperatureMetric(Device, Sensor, Metric)
}

// rocm_smi.DeviceGetVoltageMetric()
func DeviceGetVoltageMetric(Device DeviceHandle, Sensor RSMI_voltage_type, Metric RSMI_voltage_metric) (int32, RSMI_status) {
	var voltage int32
	ret := rsmi_dev_volt_metric_get(Device.index, Sensor, Metric, &voltage)
	return voltage, ret
}

func (Device DeviceHandle) GetVoltageMetric(Sensor RSMI_voltage_type, Metric RSMI_voltage_metric) (int32, RSMI_status) {
	return DeviceGetVoltageMetric(Device, Sensor, Metric)
}

// rocm_smi.DeviceResetFan()
func DeviceResetFan(Device DeviceHandle, Sensor uint32) RSMI_status {
	ret := rsmi_dev_fan_reset(Device.index, Sensor)
	return ret
}

func (Device DeviceHandle) ResetFan(Sensor uint32) RSMI_status {
	return DeviceResetFan(Device, Sensor)
}

// rocm_smi.DeviceSetFanSpeed()
func DeviceSetFanSpeed(Device DeviceHandle, Sensor uint32, Speed uint32) RSMI_status {
	ret := rsmi_dev_fan_speed_set(Device.index, Sensor, Speed)
	return ret
}

func (Device DeviceHandle) SetFanSpeed(Sensor uint32, Speed uint32) RSMI_status {
	return DeviceSetFanSpeed(Device, Sensor, Speed)
}

// rocm_smi.DeviceGetBusyPercent()
func DeviceGetBusyPercent(Device DeviceHandle) (uint32, RSMI_status) {
	var util uint32
	ret := rsmi_dev_busy_percent_get(Device.index, &util)
	return util, ret
}

func (Device DeviceHandle) GetBusyPercent() (uint32, RSMI_status) {
	return DeviceGetBusyPercent(Device)
}

// rocm_smi.DeviceGetUtilizationCounters()
func DeviceGetUtilizationCounters(Device DeviceHandle) ([]RSMI_utilization_counter, uint32, RSMI_status) {
	var util []RSMI_utilization_counter
	var timestamp uint32
	var count uint32 = uint32(UTILIZATION_COUNTER_LAST - UTILIZATION_COUNTER_FIRST + 1)
	util = make([]RSMI_utilization_counter, count)
	for i := int(UTILIZATION_COUNTER_FIRST); i <= int(UTILIZATION_COUNTER_LAST); i++ {
		util[i].Type = uint32(i)
		util[i].Value = 0
	}
	ret := rsmi_utilization_count_get(Device.index, &util[0], count, &timestamp)
	return util, timestamp, ret
}

func (Device DeviceHandle) GetUtilizationCounters() ([]RSMI_utilization_counter, uint32, RSMI_status) {
	return DeviceGetUtilizationCounters(Device)
}

// rocm_smi.DeviceGetPerfLevel()
func DeviceGetPerfLevel(Device DeviceHandle) (RSMI_dev_perf_level, RSMI_status) {
	var level RSMI_dev_perf_level
	ret := rsmi_dev_perf_level_get(Device.index, &level)
	return level, ret
}

func (Device DeviceHandle) GetPerfLevel() (RSMI_dev_perf_level, RSMI_status) {
	return DeviceGetPerfLevel(Device)
}

// rocm_smi.DeviceSetDeterminismMode()
func DeviceSetDeterminismMode(Device DeviceHandle, Clock uint32) RSMI_status {
	ret := rsmi_perf_determinism_mode_set(Device.index, Clock)
	return ret
}

func (Device DeviceHandle) SetDeterminismMode(Clock uint32) RSMI_status {
	return DeviceSetDeterminismMode(Device, Clock)
}

// rocm_smi.DeviceGetOverdriveLevel()
func DeviceGetOverdriveLevel(Device DeviceHandle) (uint32, RSMI_status) {
	var level uint32
	ret := rsmi_dev_overdrive_level_get(Device.index, &level)
	return level, ret
}

func (Device DeviceHandle) GetOverdriveLevel() (uint32, RSMI_status) {
	return DeviceGetOverdriveLevel(Device)
}

// rocm_smi.DeviceGetClockFrequency()
func DeviceGetClockFrequency(Device DeviceHandle, Clock RSMI_clk_type) (RSMI_frequencies, RSMI_status) {
	var freqs RSMI_frequencies
	ret := rsmi_dev_gpu_clk_freq_get(Device.index, Clock, &freqs)
	return freqs, ret
}

func (Device DeviceHandle) GetClockFrequency(Clock RSMI_clk_type) (RSMI_frequencies, RSMI_status) {
	return DeviceGetClockFrequency(Device, Clock)
}

// rocm_smi.DeviceReset()
func DeviceReset(Device DeviceHandle) RSMI_status {
	ret := rsmi_dev_gpu_reset(int32(Device.index))
	return ret
}

func (Device DeviceHandle) Reset() RSMI_status {
	return DeviceReset(Device)
}

// rocm_smi.DeviceGetVoltageFrequencyCurve()
func DeviceGetVoltageFrequencyCurve(Device DeviceHandle) (RSMI_od_volt_freq_data, RSMI_status) {
	var data RSMI_od_volt_freq_data
	ret := rsmi_dev_od_volt_info_get(Device.index, &data)
	return data, ret
}

func (Device DeviceHandle) GetVoltageFrequencyCurve() (RSMI_od_volt_freq_data, RSMI_status) {
	return DeviceGetVoltageFrequencyCurve(Device)
}

// rocm_smi.DeviceGetMetrics()
func DeviceGetMetrics(Device DeviceHandle) (RSMI_gpu_metrics, RSMI_status) {
	var data RSMI_gpu_metrics
	ret := rsmi_dev_gpu_metrics_info_get(Device.index, &data)
	return data, ret
}

func (Device DeviceHandle) GetMetrics() (RSMI_gpu_metrics, RSMI_status) {
	return DeviceGetMetrics(Device)
}

// rocm_smi.DeviceSetClockRange()
func DeviceSetClockRange(Device DeviceHandle, MinFreq uint32, MaxFreq uint32, Clock RSMI_clk_type) RSMI_status {
	ret := rsmi_dev_clk_range_set(Device.index, MinFreq, MaxFreq, Clock)
	return ret
}

func (Device DeviceHandle) SetClockRange(MinFreq uint32, MaxFreq uint32, Clock RSMI_clk_type) RSMI_status {
	return DeviceSetClockRange(Device, MinFreq, MaxFreq, Clock)
}

// rocm_smi.DeviceSetClockInfo()
func DeviceSetClockInfo(Device DeviceHandle, Level RSMI_freq_ind, ClockFreq uint32, Clock RSMI_clk_type) RSMI_status {
	ret := rsmi_dev_od_clk_info_set(Device.index, Level, ClockFreq, Clock)
	return ret
}

func (Device DeviceHandle) SetClockInfo(Level RSMI_freq_ind, ClockFreq uint32, Clock RSMI_clk_type) RSMI_status {
	return DeviceSetClockInfo(Device, Level, ClockFreq, Clock)
}

// rocm_smi.DeviceSetVoltageInfo()
func DeviceSetVoltageInfo(Device DeviceHandle, Vpoint uint32, ClockFreq uint32, Voltage uint32) RSMI_status {
	ret := rsmi_dev_od_volt_info_set(Device.index, Vpoint, ClockFreq, Voltage)
	return ret
}

func (Device DeviceHandle) SetVoltageInfo(Vpoint uint32, ClockFreq uint32, Voltage uint32) RSMI_status {
	return DeviceSetVoltageInfo(Device, Vpoint, ClockFreq, Voltage)
}

// rocm_smi.DeviceGetVoltageFrequencyCurveRegions()
func DeviceGetVoltageFrequencyCurveRegions(Device DeviceHandle) ([]RSMI_freq_volt_region, RSMI_status) {
	var num_regions uint32 = 0
	regions := make([]RSMI_freq_volt_region, 0)
	ret := rsmi_dev_od_volt_curve_regions_get(Device.index, &num_regions, nil)
	if ret == STATUS_SUCCESS && num_regions > 0 {
		regions := make([]RSMI_freq_volt_region, num_regions)
		ret = rsmi_dev_od_volt_curve_regions_get(Device.index, &num_regions, &regions[0])
	}
	return regions, ret
}

func (Device DeviceHandle) GetVoltageFrequencyCurveRegions() ([]RSMI_freq_volt_region, RSMI_status) {
	return DeviceGetVoltageFrequencyCurveRegions(Device)
}

// rocm_smi.DeviceGetPowerProfile
func DeviceGetPowerProfile(Device DeviceHandle, Sensor uint32) (RSMI_power_profile_status, RSMI_status) {
	var status RSMI_power_profile_status
	ret := rsmi_dev_power_profile_presets_get(Device.index, Sensor, &status)
	return status, ret
}

func (Device DeviceHandle) GetPowerProfile(Sensor uint32) (RSMI_power_profile_status, RSMI_status) {
	return DeviceGetPowerProfile(Device, Sensor)
}

// rocm_smi.DeviceSetPerfLevel()
var DeviceSetPerfLevel = DeviceSetPerfLevel_v1

func DeviceSetPerfLevel_v2(Device DeviceHandle, Level RSMI_dev_perf_level) RSMI_status {
	ret := rsmi_dev_perf_level_set(int32(Device.index), Level)
	return ret
}

// rocm_smi.DeviceSetPerfLevel_v1()
func DeviceSetPerfLevel_v1(Device DeviceHandle, Level RSMI_dev_perf_level) RSMI_status {
	ret := rsmi_dev_perf_level_set_v1(Device.index, Level)
	return ret
}

func (Device DeviceHandle) SetPerfLevel(Level RSMI_dev_perf_level) RSMI_status {
	return DeviceSetPerfLevel(Device, Level)
}

var DeviceSetOverdriveLevel = DeviceSetOverdriveLevel_v1

// rocm_smi.DeviceSetPerfLevel()
func DeviceSetOverdriveLevel_v2(Device DeviceHandle, Overdrive uint32) RSMI_status {
	ret := rsmi_dev_overdrive_level_set(int32(Device.index), Overdrive)
	return ret
}

// rocm_smi.DeviceSetOverdriveLevel_v1()
func DeviceSetOverdriveLevel_v1(Device DeviceHandle, Overdrive uint32) RSMI_status {
	ret := rsmi_dev_overdrive_level_set_v1(Device.index, Overdrive)
	return ret
}

func (Device DeviceHandle) SetOverdriveLevel(Overdrive uint32) RSMI_status {
	return DeviceSetOverdriveLevel(Device, Overdrive)
}

// rocm_smi.DeviceSetClockFrequency()
func DeviceSetClockFrequency(Device DeviceHandle, Clock RSMI_clk_type, FreqMask uint32) RSMI_status {
	ret := rsmi_dev_gpu_clk_freq_set(Device.index, Clock, FreqMask)
	return ret
}

func (Device DeviceHandle) SetClockFrequency(Clock RSMI_clk_type, FreqMask uint32) RSMI_status {
	return DeviceSetClockFrequency(Device, Clock, FreqMask)
}

// rocm_smi.DeviceGetVbiosVersionString()
func DeviceGetVbiosVersionString(Device DeviceHandle) (string, RSMI_status) {
	var version []byte = make([]byte, 100)
	vptr := &version[0]
	ret := rsmi_dev_vbios_version_get(Device.index, vptr, 100)
	return string(version), ret
}

func (Device DeviceHandle) GetVbiosVersionString() (string, RSMI_status) {
	return DeviceGetVbiosVersionString(Device)
}

// rocm_smi.DeviceGetFirmwareVersion()
func DeviceGetFirmwareVersion(Device DeviceHandle, Block RSMI_fw_block) (uint32, RSMI_status) {
	var version uint32
	ret := rsmi_dev_firmware_version_get(Device.index, Block, &version)
	return version, ret
}

func (Device DeviceHandle) GetFirmwareVersion(Block RSMI_fw_block) (uint32, RSMI_status) {
	return DeviceGetFirmwareVersion(Device, Block)
}

// rocm_smi.DeviceGetEccCount()
func DeviceGetEccCount(Device DeviceHandle, Block RSMI_gpu_block) (RSMI_error_count, RSMI_status) {
	var counts RSMI_error_count
	ret := rsmi_dev_ecc_count_get(Device.index, Block, &counts)
	return counts, ret
}

func (Device DeviceHandle) GetEccCount(Block RSMI_gpu_block) (RSMI_error_count, RSMI_status) {
	return DeviceGetEccCount(Device, Block)
}

// rocm_smi.DeviceGetEccStatus()
func DeviceGetEccStatus(Device DeviceHandle, Block RSMI_gpu_block) (RSMI_ras_err_state, RSMI_status) {
	var state RSMI_ras_err_state
	ret := rsmi_dev_ecc_status_get(Device.index, Block, &state)
	return state, ret
}

func (Device DeviceHandle) GetEccStatus(Block RSMI_gpu_block) (RSMI_ras_err_state, RSMI_status) {
	return DeviceGetEccStatus(Device, Block)
}

// rocm_smi.DeviceGetEccMask()
func DeviceGetEccMask(Device DeviceHandle) (uint32, RSMI_status) {
	var mask uint32
	ret := rsmi_dev_ecc_enabled_get(Device.index, &mask)
	return mask, ret
}

func (Device DeviceHandle) GetEccMask() (uint32, RSMI_status) {
	return DeviceGetEccMask(Device)
}

// SKIP rsmi_dev_counter_group_supported
// SKIP rsmi_dev_counter_create
// SKIP rsmi_dev_counter_destroy
// SKIP rsmi_counter_control
// SKIP rsmi_counter_read
// SKIP rsmi_counter_available_counters_get
