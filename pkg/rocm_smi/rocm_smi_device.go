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
//import "strings"

const defaultRsmiStringLength uint32 = 1000
var ctoGoFuncMapping = map[string]string {
    "rsmi_dev_fan_reset": "DeviceResetFan",
    "rsmi_dev_fan_speed_set": "DeviceSetFanSpeed",
    "rsmi_dev_volt_metric_get": "DeviceGetVoltageMetric",
    "rsmi_dev_temp_metric_get": "DeviceGetTemperatureMetric",
    "rsmi_dev_busy_percent_get": "DeviceGetBusyPercent",
    "rsmi_utilization_count_get": "DeviceGetUtilizationCounters",
    "rsmi_dev_fan_speed_max_get": "DeviceGetMaxFanSpeed",
    "rsmi_dev_fan_speed_get": "DeviceGetFanSpeed",
    "rsmi_dev_fan_rpms_get": "DeviceGetFanRpms",
    "rsmi_dev_memory_reserved_pages_get": "DeviceGetMemoryReservedPages",
    "rsmi_dev_memory_busy_percent_get": "DeviceGetMemoryUtilization",
    "rsmi_dev_memory_usage_get": "DeviceGetUsedMemory",
    "rsmi_dev_memory_total_get": "DeviceGetTotalMemory",
    "rsmi_dev_power_profile_set": "DeviceSetPowerProfile",
    "rsmi_dev_power_cap_set": "DeviceSetPowerCap",
    "rsmi_dev_energy_count_get": "DeviceGetEnergyCount",
    "rsmi_dev_power_cap_range_get": "DeviceGetPowerCapRange",
    "rsmi_dev_power_cap_default_get": "DeviceGetDefaultPowerCap",
    "rsmi_dev_power_cap_get": "DeviceGetPowerCap",
    "rsmi_dev_power_ave_get": "DeviceGetPowerAverage",
    "rsmi_topo_numa_affinity_get": "DeviceGetNumaAffinity",
    "rsmi_dev_pci_replay_counter_get": "DeviceGetPciReplayCounter",
    "rsmi_dev_pci_throughput_get": "DeviceGetPciThroughput",
    "rsmi_dev_pci_bandwidth_set": "DeviceSetPciBandwidth",
    "rsmi_dev_pci_bandwidth_get": "DeviceGetPciBandwidth",
    "rsmi_dev_pci_id_get": "DeviceGetPciId",
    "rsmi_dev_unique_id_get": "DeviceGetUniqueId",
    "rsmi_dev_drm_render_minor_get": "DeviceGetDrmRenderMinor",
    "rsmi_dev_subsystem_id_get": "DeviceGetSubsystemId",
    "rsmi_dev_subsystem_name_get": "DeviceGetSubsystemName",
    "rsmi_dev_serial_number_get": "DeviceGetSerial",
    "rsmi_dev_vram_vendor_get": "DeviceGetVramVendor",
    "rsmi_dev_vendor_id_get": "DeviceGetVendorId",
    "rsmi_dev_vendor_name_get": "DeviceGetVendorName",
    "rsmi_dev_sku_get": "DeviceGetSku",
    "rsmi_dev_name_get": "DeviceGetName",
    "rsmi_dev_brand_get": "DeviceGetBrand",
    "rsmi_dev_perf_level_get": "DeviceGetPerfLevel",
    "rsmi_perf_determinism_mode_set": "DeviceSetDeterminismMode",
    "rsmi_dev_overdrive_level_get": "DeviceGetOverdriveLevel",
    "rsmi_dev_gpu_clk_freq_get": "DeviceGetClockFrequency",
    "rsmi_dev_od_volt_info_get": "DeviceGetVoltageFrequencyCurve",
    "rsmi_dev_gpu_metrics_info_get": "DeviceGetMetrics",
    "rsmi_dev_clk_range_set": "DeviceSetClockRange",
    "rsmi_dev_od_clk_info_set": "DeviceSetClockInfo",
    "rsmi_dev_od_volt_info_set": "DeviceSetVoltageInfo",
    "rsmi_dev_od_volt_curve_regions_get": "DeviceGetVoltageFrequencyCurveRegions",
    "rsmi_dev_power_profile_presets_get": "DeviceGetPowerProfile",
    "rsmi_dev_perf_level_set": "DeviceSetPerfLevel_v2",
    "rsmi_dev_perf_level_set_v1": "DeviceSetPerfLevel_v1",
    "rsmi_dev_vbios_version_get": "DeviceGetVbiosVersionString",
    "rsmi_dev_firmware_version_get": "DeviceGetFirmwareVersion",
    "rsmi_dev_overdrive_level_set_v1": "DeviceSetOverdriveLevel_v1",
    "rsmi_dev_overdrive_level_set": "DeviceSetOverdriveLevel_v2",
    "rsmi_dev_gpu_clk_freq_set": "DeviceSetClockFrequency",
    "rsmi_dev_ecc_count_get": "DeviceGetEccCount",
    "rsmi_dev_ecc_status_get": "DeviceGetEccStatus",
    "rsmi_dev_ecc_enabled_get": "DeviceGetEccMask",
    "rsmi_dev_xgmi_error_status": "DeviceXgmiErrorStatus",
    "rsmi_dev_xgmi_error_reset": "DeviceXgmiErrorReset",
    "rsmi_dev_xgmi_hive_id_get": "DeviceXgmiHiveId",
    "rsmi_topo_get_numa_node_number": "DeviceGetNumaNode",
    "rsmi_topo_get_link_weight": "DeviceGetLinkWeight",
    "rsmi_minmax_bandwidth_get": "DeviceGetMinMaxBandwidth",
    "rsmi_topo_get_link_type": "DeviceGetLinkType",
    "rsmi_is_P2P_accessible": "DeviceIsP2PAccessible",
    "rsmi_event_notification_init": "DeviceInitEventNotification",
    "rsmi_event_notification_mask_set": "DeviceSetEventNotificationMask",
    "rsmi_event_notification_get": "GetEventNotification",
    "rsmi_event_notification_stop": "DeviceStopEventNotification",
    "rsmi_dev_counter_group_supported": "DeviceCounterGroupSupported",
    "rsmi_counter_available_counters_get": "DeviceCounterGetAvailable",
    "rsmi_dev_counter_create": "DeviceCounterCreate",
    "rsmi_counter_control": "CounterControl",
    "rsmi_counter_read": "CounterRead",
    "rsmi_dev_counter_destroy": "CounterDestroy",
    "rsmi_compute_process_info_get": "ComputeProcesses",
    "rsmi_compute_process_info_by_pid_get": "ComputeProcessByPid",
    "rsmi_compute_process_gpus_get": "ComputeProcessGpus",
}

func bytes2String(bytes []byte) string {
	s := make([]byte, 0)
	for _, v := range bytes {
		s = append(s, v)
	}
	return string(s)
}

// NumMonitorDevices gets the number of devices that have monitor information.
// The number of devices which have monitors is returned. Monitors are
// referenced by the index which can be between 0 and the returned num_devices - 1.
// 
// Returns STATUS_SUCCESS upon successful call.
func NumMonitorDevices() (int, RSMI_status) {
	var DeviceCount uint32
	ret := rsmi_num_monitor_devices(&DeviceCount)
	return int(DeviceCount), ret
}

// DeviceGetHandleByIndex gets the device handle associated with the device with provided device index.
// It also reads the supported functions and their supported arguments. It retrieves also the **not** unique ID.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
func DeviceGetHandleByIndex(Index int) (DeviceHandle, RSMI_status) {
	var index uint32 = uint32(Index)
	handle := DeviceHandle{
		handle: 0,
		index:  index,
		supported: nil,
	}
	ret := rsmi_dev_id_get(handle.index, &handle.handle)
	if ret == STATUS_SUCCESS {
		B2S := func(bs [200]int8) string {
			b := make([]byte, 0)
			for _, v := range bs {
				if v == 0 {
					break
				}
				b = append(b, byte(v))
			}
			return string(b)
		}
		var funcs RSMI_helper_function
		ret = rsmi_helper_func_variants_get(handle.index, &funcs)
		if ret == STATUS_SUCCESS {
			handle.supported = make(map[string]map[uint64][]uint64)
			for f := 0; f < int(funcs.Functions); f++ {
				name := B2S(funcs.Functionlist[f].Name)
				handle.supported[name] = make(map[uint64][]uint64)
				for s := 0; s < int(funcs.Functionlist[f].Sensors); s++ {
					sdata := funcs.Functionlist[f].Sensorlist[s]
					sensor := uint64(sdata.Variant_id)
					l := make([]uint64, 0)
					for v := 0; v < int(sdata.Num_variants); v++ {
						l = append(l, uint64(sdata.Variantlist[v]))
					}
					handle.supported[name][sensor] = l
				}
				if goname, ok := ctoGoFuncMapping[name]; ok {
				    handle.supported[goname] = handle.supported[name]
				}
			}
		}
	}
	return handle, ret
}

// DeviceGetBrand gets the brand string of a gpu device.
// If the sku associated with the device is not found as one of the values
// contained within rsmi_dev_brand_get, then this function will return the
// device marketing name as a string instead of the brand name.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
// Returns STATUS_INSUFFICIENT_SIZE when the vendor name is longer than defaultRsmiStringLength characters.
func DeviceGetBrand(Device DeviceHandle) (string, RSMI_status) {
	if _, ok := Device.supported["rsmi_dev_brand_get"]; !ok {
		return "", STATUS_NOT_SUPPORTED
	}
	var brand []byte = make([]byte, defaultRsmiStringLength)
	bptr := &brand[0]
	ret := rsmi_dev_brand_get(Device.index, bptr, defaultRsmiStringLength)
	return bytes2String(brand), ret
}

// DeviceGetBrand gets the brand string of a gpu device.
// If the sku associated with the device is not found as one of the values
// contained within rsmi_dev_brand_get, then this function will return the
// device marketing name as a string instead of the brand name.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
// Returns STATUS_INSUFFICIENT_SIZE when the vendor name is longer than defaultRsmiStringLength characters.
func (Device DeviceHandle) GetBrand() (string, RSMI_status) {
	return DeviceGetBrand(Device)
}

// DeviceGetName gets the name string of a gpu device.
// If the integer ID associated with the device is not found in one of the system
// files containing device name information (e.g. /usr/share/misc/pci.ids), then this
// function will return the hex device ID as a string.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
// Returns STATUS_INSUFFICIENT_SIZE when the vendor name is longer than defaultRsmiStringLength characters.
func DeviceGetName(Device DeviceHandle) (string, RSMI_status) {
	if _, ok := Device.supported["rsmi_dev_name_get"]; !ok {
		return "", STATUS_NOT_SUPPORTED
	}
	var name []byte = make([]byte, defaultRsmiStringLength)
	nptr := &name[0]
	ret := rsmi_dev_name_get(Device.index, nptr, defaultRsmiStringLength)
	return bytes2String(name), ret
}

// GetName gets the name string of a gpu device.
// If the integer ID associated with the device is not found in one of the system
// files containing device name information (e.g. /usr/share/misc/pci.ids), then this
// function will return the hex device ID as a string.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
// Returns STATUS_INSUFFICIENT_SIZE when the vendor name is longer than defaultRsmiStringLength characters.
func (Device DeviceHandle) GetName() (string, RSMI_status) {
	return DeviceGetName(Device)
}


// DeviceGetSku gets the SKU for a device.
// This function will attempt to obtain the SKU from the Product Information FRU chip, present on server ASICs
//
// Note: There are versions of the rocm_smi library which do not export the function rsmi_dev_sku_get. Therefore the bindings
//       perform a symbol lookup at initialization. If it is not available, the deviceGetSkuFake function is assigned to DeviceGetSku
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
var DeviceGetSku = deviceGetSkuFake

// DeviceGetSkuReal is the actual reading function for the SKU. But it is not supported by some devices.
func deviceGetSkuReal(Device DeviceHandle) (string, RSMI_status) {
	var Sku []byte = make([]byte, defaultRsmiStringLength)
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_sku_get"]; !ok {
		sptr := &Sku[0]
		ret = rsmi_dev_sku_get(Device.index, sptr)
	}
	return bytes2String(Sku), ret
}

// DeviceGetSkuReal is returning 'NA' as SKU because the funtion is not supported by the device
func deviceGetSkuFake(Device DeviceHandle) (string, RSMI_status) {
	return "NA", STATUS_NOT_SUPPORTED
}

// GetSku gets the SKU for a device..
// This function will attempt to obtain the SKU from the Product Information FRU chip, present on server ASICs
//
// Note: There are versions of the rocm_smi library which do not export the function rsmi_dev_sku_get. Therefore the bindings
//       perform a symbol lookup at initialization. If it is not available, the deviceGetSkuFake function is assigned to DeviceGetSku
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
func (Device DeviceHandle) GetSku() (string, RSMI_status) {
	return DeviceGetSku(Device)
}

// DeviceGetVendorName gets the name string for a give vendor ID.
// If the integer ID associated with the vendor is not found in one of
// the system files containing device name information (e.g. /usr/share/misc/pci.ids),
// then this function will return the hex vendor ID as a string.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
// Returns STATUS_INSUFFICIENT_SIZE when the vendor name is longer than defaultRsmiStringLength characters.
func DeviceGetVendorName(Device DeviceHandle) (string, RSMI_status) {
	if _, ok := Device.supported["rsmi_dev_vendor_name_get"]; !ok {
		return "", STATUS_NOT_SUPPORTED
	}
	var Name []byte = make([]byte, defaultRsmiStringLength)
	nptr := &Name[0]
	ret := rsmi_dev_vendor_name_get(Device.index, nptr, defaultRsmiStringLength)
	return bytes2String(Name), ret
}

// GetVendorName gets the name string for a give vendor ID.
// If the integer ID associated with the vendor is not found in one of
// the system files containing device name information (e.g. /usr/share/misc/pci.ids),
// then this function will return the hex vendor ID as a string.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
// Returns STATUS_INSUFFICIENT_SIZE when the vendor name is longer than defaultRsmiStringLength characters.
func (Device DeviceHandle) GetVendorName() (string, RSMI_status) {
	return DeviceGetVendorName(Device)
}

// DeviceGetVendorId gets the device vendor id associated with the device
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
func DeviceGetVendorId(Device DeviceHandle) (uint16, RSMI_status) {
	var id uint16 = 0
	if _, ok := Device.supported["rsmi_dev_vendor_id_get"]; !ok {
		return id, STATUS_NOT_SUPPORTED
	}
	ret := rsmi_dev_vendor_id_get(Device.index, &id)
	return id, ret
}


// GetVendorId gets the device vendor id associated with the device
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_NOT_SUPPORTED when installed software or hardware does not support this function with the given arguments.
// Returns STATUS_INVALID_ARGS when the provided arguments are not valid.
func (Device DeviceHandle) GetVendorId() (uint16, RSMI_status) {
	return DeviceGetVendorId(Device)
}

// DeviceGetVramVendor gets the vram vendor string of a gpu device.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_INSUFFICIENT_SIZE when the vram vendor name is longer than defaultRsmiStringLength characters.
func DeviceGetVramVendor(Device DeviceHandle) (string, RSMI_status) {
	if _, ok := Device.supported["rsmi_dev_vram_vendor_get"]; !ok {
		return "", STATUS_NOT_SUPPORTED
	}
	var Name []byte = make([]byte, defaultRsmiStringLength)
	nptr := &Name[0]
	ret := rsmi_dev_vram_vendor_get(Device.index, nptr, defaultRsmiStringLength)
	return bytes2String(Name), ret
}

// GetVramVendor gets the vram vendor string of a gpu device.
//
// Returns STATUS_SUCCESS when call was successful.
// Returns STATUS_INSUFFICIENT_SIZE when the vram vendor name is longer than defaultRsmiStringLength characters.
func (Device DeviceHandle) GetVramVendor() (string, RSMI_status) {
	return DeviceGetVramVendor(Device)
}

// DeviceGetSerial gets the serial number string for a device
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_INSUFFICIENT_SIZE is returned if version string is larger than defaultRsmiStringLength bytes.
func DeviceGetSerialNumber(Device DeviceHandle) (string, RSMI_status) {
	if _, ok := Device.supported["rsmi_dev_serial_number_get"]; !ok {
		return "", STATUS_NOT_SUPPORTED
	}
	var size uint32 = 100
	var tmp []byte
	var ret RSMI_status = STATUS_SUCCESS
	for size = 100; size < defaultRsmiStringLength; size += 100 {
		tmp = make([]byte, size)
		sptr := &tmp[0]
		ret = rsmi_dev_serial_number_get(Device.index, sptr, size)
	}
	Serial := make([]byte, 0)
	for _, v := range tmp {
		if v != 0 {
			Serial = append(Serial, v)
		}
	}
	return bytes2String(Serial), ret
}

// GetSerialNumber gets the serial number string for a device
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_INSUFFICIENT_SIZE is returned if version string is larger than defaultRsmiStringLength bytes.
func (Device DeviceHandle) GetSerialNumber() (string, RSMI_status) {
	return DeviceGetSerialNumber(Device)
}

// DeviceGetSubsystemName gets the name string for the device subsytem
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_INSUFFICIENT_SIZE is returned if version string is larger than defaultRsmiStringLength bytes.
func DeviceGetSubsystemName(Device DeviceHandle) (string, RSMI_status) {
	if _, ok := Device.supported["rsmi_dev_subsystem_name_get"]; !ok {
		return "", STATUS_NOT_SUPPORTED
	}
	var Name []byte = make([]byte, defaultRsmiStringLength)
	nptr := &Name[0]
	ret := rsmi_dev_subsystem_name_get(Device.index, nptr, defaultRsmiStringLength)
	return bytes2String(Name), ret
}

// GetSubsystemName gets the name string for the device subsytem
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_INSUFFICIENT_SIZE is returned if version string is larger than defaultRsmiStringLength bytes.
func (Device DeviceHandle) GetSubsystemName() (string, RSMI_status) {
	return DeviceGetSubsystemName(Device)
}

// DeviceGetSubsystemId gets the subsystem device id associated with the device
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetSubsystemId(Device DeviceHandle) (uint16, RSMI_status) {
	var id uint16 = 0
	if _, ok := Device.supported["rsmi_dev_subsystem_id_get"]; !ok {
		return id, STATUS_NOT_SUPPORTED
	}
	ret := rsmi_dev_subsystem_id_get(Device.index, &id)
	return id, ret
}

// DeviceGetSubsystemId gets the subsystem device id associated with the device
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetSubsystemId() (uint16, RSMI_status) {
	return DeviceGetSubsystemId(Device)
}

// DeviceGetDrmRenderMinor gets the drm minor number associated with the device
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetDrmRenderMinor(Device DeviceHandle) (uint32, RSMI_status) {
	var minor uint32 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_drm_render_minor_get"]; ok {
		ret = rsmi_dev_drm_render_minor_get(Device.index, &minor)
	}
	return minor, ret
}

// DeviceGetDrmRenderMinor gets the drm minor number associated with the device
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetDrmRenverMinor() (uint32, RSMI_status) {
	return DeviceGetDrmRenderMinor(Device)
}


// DeviceGetUniqueId gets Unique ID
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetUniqueId(Device DeviceHandle) (uint64, RSMI_status) {
	var id uint64 = 0
	if _, ok := Device.supported["rsmi_dev_unique_id_get"]; !ok {
		return id, STATUS_NOT_SUPPORTED
	}
	ret := rsmi_dev_unique_id_get(Device.index, &id)
	return id, ret
}

// DeviceGetUniqueId gets Unique ID
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetUniqueId() (uint64, RSMI_status) {
	return DeviceGetUniqueId(Device)
}

// DeviceGetPciId gets the unique PCI device identifier associated with the device
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetPciId(Device DeviceHandle) (uint64, RSMI_status) {
	var id uint64 = 0
	if _, ok := Device.supported["rsmi_dev_pci_id_get"]; !ok {
		return id, STATUS_NOT_SUPPORTED
	}
	ret := rsmi_dev_pci_id_get(Device.index, &id)
	return id, ret
}

// DeviceGetPciId gets the unique PCI device identifier associated with the device
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetPciId() (uint64, RSMI_status) {
	return DeviceGetPciId(Device)
}

type Pci_info struct {
	Domain   uint32 // PCI domain
	Bus      uint8 // PCI bus
	Device   uint8 // PCI device
	Function uint8 // PCI function
}

// DeviceGetPciInfo gets the unique PCI device identifier associated with the device split into its parts: PCI domain, PCI bus, PCI device and function.
//
// Note: This is an own addition to simply usage, it uses DeviceGetPciId to retrieve the actual PCI identifer and splits it into the parts.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetPciInfo(Device DeviceHandle) (Pci_info, RSMI_status) {
	var id uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	info := Pci_info{
		Domain:   0,
		Bus:      0,
		Device:   0,
		Function: 0,
	}
	id, ret = DeviceGetPciId(Device)
	if ret == STATUS_SUCCESS {
		info.Domain = uint32((id >> 32) & 0xffffffff)
		info.Bus = uint8((id >> 8) & 0xff)
		info.Device = uint8((id >> 3) & 0x1f)
		info.Function = uint8(id & 0x7)
	}
	return info, ret
}

// GetPciInfo gets the unique PCI device identifier associated with the device split into its parts: PCI domain, PCI bus, PCI device and function.
//
// Note: This is an own addition to simply usage, it uses DeviceGetPciId to retrieve the actual PCI identifer and splits it into the parts.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetPciInfo() (Pci_info, RSMI_status) {
	return DeviceGetPciInfo(Device)
}

// DeviceGetPciBandwidth gets the list of possible PCIe bandwidths that are available.
//
// STATUS_SUCCESS call was successful.
func DeviceGetPciBandwidth(Device DeviceHandle) (RSMI_pcie_bandwidth, RSMI_status) {
	var info RSMI_pcie_bandwidth = RSMI_pcie_bandwidth {
		Rate : RSMI_frequencies {
			Supported: 0,
			Current: 0,
		},
	}
	if _, ok := Device.supported["rsmi_dev_pci_bandwidth_get"]; !ok {
		return info, STATUS_NOT_SUPPORTED
	}
	ret := rsmi_dev_pci_bandwidth_get(Device.index, &info)
	return info, ret
}

// GetPciBandwidth gets the list of possible PCIe bandwidths that are available.
//
// STATUS_SUCCESS call was successful.
func (Device DeviceHandle) GetPciBandwidth() (RSMI_pcie_bandwidth, RSMI_status) {
	return DeviceGetPciBandwidth(Device)
}

// DeviceSetPciBandwidth controls the set of allowed PCIe bandwidths that can be used.
//
// STATUS_SUCCESS call was successful.
// STATUS_PERMISSION function requires root access.
func DeviceSetPciBandwidth(Device DeviceHandle, Mask uint64) RSMI_status {
	ret := STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_pci_bandwidth_set"]; ok {
		ret = rsmi_dev_pci_bandwidth_set(Device.index, Mask)
	}
	return ret
}

// SetPciBandwidth controls the set of allowed PCIe bandwidths that can be used.
//
// STATUS_SUCCESS call was successful.
// STATUS_PERMISSION function requires root access.
func (Device DeviceHandle) SetPciBandwidth(Mask uint64) RSMI_status {
	return DeviceSetPciBandwidth(Device, Mask)
}

// DeviceGetPciThroughput gets PCIe traffic information for the device.
// This function returns the number of bytes sent and received in 1 second and the maximum possible packet size.
//
// Note: The function blocks execution for 1 second.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
func DeviceGetPciThroughput(Device DeviceHandle) (uint64, uint64, uint64, RSMI_status) {
	var sent uint64 = 0
	var recv uint64 = 0
	var max_pkts_size uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_pci_throughput_get"]; ok {
		ret = rsmi_dev_pci_throughput_get(Device.index, &sent, &recv, &max_pkts_size)
	}
	return sent, recv, max_pkts_size, ret
}

// GetPciThroughput gets PCIe traffic information for the device.
// This function returns the number of bytes sent and received in 1 second and the maximum possible packet size.
//
// Note: The function blocks execution for 1 second.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
func (Device DeviceHandle) GetPciThroughput() (uint64, uint64, uint64, RSMI_status) {
	return DeviceGetPciThroughput(Device)
}

// DeviceGetPciReplayCounter gets the PCIe replay counter.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetPciReplayCounter(Device DeviceHandle) (uint64, RSMI_status) {
	var counter uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_pci_replay_counter_get"]; ok {
		ret = rsmi_dev_pci_replay_counter_get(Device.index, &counter)
	}
	return counter, ret
}

// GetPciReplayCounter gets the PCIe replay counter.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetPciReplayCounter() (uint64, RSMI_status) {
	return DeviceGetPciReplayCounter(Device)
}

// DeviceGetNumaAffinity gets the NUMA node associated with a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetNumaAffinity(Device DeviceHandle) (uint32, RSMI_status) {
	var id uint32 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_topo_numa_affinity_get"]; ok {
		ret = rsmi_topo_numa_affinity_get(Device.index, &id)
	}
	return id, ret
}

// GetNumaAffinity gets the NUMA node associated with a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetNumaAffinity() (uint32, RSMI_status) {
	return DeviceGetNumaAffinity(Device)
}

// DeviceGetPowerAverage gets the average power consumption of the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetPowerAverage(Device DeviceHandle, Sensor uint32) (uint64, RSMI_status) {
	var power uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_power_ave_get"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_power_ave_get(Device.index, Sensor, &power)
		} else if defaults, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaults) > 0 {
				ret = rsmi_dev_power_ave_get(Device.index, uint32(defaults[0]), &power)
			}
		}
	}
	return power, ret
}

// GetPowerAverage gets the average power consumption of the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetPowerAverage(Sensor uint32) (uint64, RSMI_status) {
	return DeviceGetPowerAverage(Device, Sensor)
}

// DeviceGetPowerCap gets the cap on power which, when reached, causes the system to take action to reduce power.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetPowerCap(Device DeviceHandle, Sensor uint32) (uint64, RSMI_status) {
	var power uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_power_cap_get"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_power_cap_get(Device.index, Sensor, &power)
		} else if defaults, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaults) > 0 {
				ret = rsmi_dev_power_cap_get(Device.index, uint32(defaults[0]), &power)
			}
		}
	}
	return power, ret
}

// GetPowerCap gets the cap on power which, when reached, causes the system to take action to reduce power.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetPowerCap(Sensor uint32) (uint64, RSMI_status) {
	return DeviceGetPowerCap(Device, Sensor)
}

// DeviceGetDefaultPowerCap gets the default power cap for the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetDefaultPowerCap(Device DeviceHandle) (uint64, RSMI_status) {
	var power uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	// if sensors, ok := Device.supported["rsmi_dev_power_cap_default_get"]; ok {
	// 	if _, ok := sensors[uint64(Sensor)]; ok {
	// 		ret = rsmi_dev_power_cap_default_get(Device.index, &power)
	// 	} else if defaults, ok := sensors[DEFAULT_VARIANT]; ok {
	// 		if len(defaults) > 0 {
	// 			ret = rsmi_dev_power_cap_default_get(Device.index, &power)
	// 		}
	// 	}
	// }
	if _, ok := Device.supported["rsmi_dev_power_cap_default_get"]; ok {
		ret = rsmi_dev_power_cap_default_get(Device.index, &power)
	}
	return power, ret
}

// GetDefaultPowerCap gets the default power cap for the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetDefaultPowerCap() (uint64, RSMI_status) {
	return DeviceGetDefaultPowerCap(Device)
}

// DeviceGetPowerCapRange gets the range of valid values (maximum and minimum) for the power cap.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetPowerCapRange(Device DeviceHandle, Sensor uint32) (uint64, uint64, RSMI_status) {
	var mini uint64 = 0
	var maxi uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_power_cap_range_get"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_power_cap_range_get(Device.index, Sensor, &maxi, &mini)
		} else if defaults, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaults) > 0 {
				ret = rsmi_dev_power_cap_range_get(Device.index, uint32(defaults[0]), &maxi, &mini)
			}
		}
	}
	return maxi, mini, ret
}

// GetPowerCapRange gets the range of valid values (maximum and minimum) for the power cap.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetPowerCapRange(Sensor uint32) (uint64, uint64, RSMI_status) {
	return DeviceGetPowerCapRange(Device, Sensor)
}

// DeviceGetEnergyCount gets the energy accumulator counter of the device.
// It returns the power, the resolution and the timestamp.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetEnergyCount(Device DeviceHandle) (uint64, float32, uint64, RSMI_status) {
	var power uint64 = 0
	var resolution float32 = 0
	var timestamp uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	ret = rsmi_dev_energy_count_get(Device.index, &power, &resolution, &timestamp)
	// Seems to be not listed in the function-variant mapping
	// if _, ok := Device.supported["rsmi_dev_energy_count_get"]; ok {
	// 	ret = rsmi_dev_energy_count_get(Device.index, &power, &resolution, &timestamp)
	// } else {
	// 	fmt.Println("rsmi_dev_energy_count_get not available")
	// }
	return power, resolution, timestamp, ret
}

// GetEnergyCount gets the energy accumulator counter of the device.
// It returns the power, the resolution and the timestamp.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetEnergyCount() (uint64, float32, uint64, RSMI_status) {
	return DeviceGetEnergyCount(Device)
}

// DeviceSetPowerCap sets the power cap value to a set of available settings selectable through a mask.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_PERMISSION function requires root access.
func DeviceSetPowerCap(Device DeviceHandle, Sensor uint32, Mask uint64) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_power_cap_set"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_power_cap_set(Device.index, Sensor, Mask)
		} else if defaults, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaults) > 0 {
				ret = rsmi_dev_power_cap_set(Device.index, uint32(defaults[0]), Mask)
			}
		}
	}
	return ret
}

// SetPowerCap sets the power cap value to a set of available settings selectable through a mask.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_PERMISSION function requires root access.
func (Device DeviceHandle) SetPowerCap(Sensor uint32, Mask uint64) RSMI_status {
	return DeviceSetPowerCap(Device, Sensor, Mask)
}

// DeviceSetPowerProfile set the power profile.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid. (not documented but likely).
// STATUS_PERMISSION function requires root access.
func DeviceSetPowerProfile(Device DeviceHandle, Preset RSMI_power_profile_preset_masks) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_power_profile_set"]; ok {
		ret = rsmi_dev_power_profile_set(Device.index, 0, Preset)
	}
	return ret
}

// SetPowerProfile set the power profile.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid. (not documented but likely).
// STATUS_PERMISSION function requires root access.
func (Device DeviceHandle) SetPowerProfile(Preset RSMI_power_profile_preset_masks) RSMI_status {
	return DeviceSetPowerProfile(Device, Preset)
}

// DeviceGetTotalMemory get the total amount of memory that exists.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetTotalMemory(Device DeviceHandle, Type RSMI_memory_type) (uint64, RSMI_status) {
	var size uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if types, ok := Device.supported["rsmi_dev_memory_total_get"]; ok {
		if _, ok := types[uint64(Type)]; ok {
			ret = rsmi_dev_memory_total_get(Device.index, Type, &size)
		}
	}
	return size, ret
}

// GetTotalMemory get the total amount of memory that exists.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetTotalMemory(Type RSMI_memory_type) (uint64, RSMI_status) {
	return DeviceGetTotalMemory(Device, Type)
}

// DeviceGetUsedMemory gets the current memory usage.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetUsedMemory(Device DeviceHandle, Type RSMI_memory_type) (uint64, RSMI_status) {
	var size uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if types, ok := Device.supported["rsmi_dev_memory_usage_get"]; ok {
		if _, ok := types[uint64(Type)]; ok {
			ret = rsmi_dev_memory_usage_get(Device.index, Type, &size)
		}
	}
	return size, ret
}

// GetUsedMemory gets the current memory usage.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetUsedMemory(Type RSMI_memory_type) (uint64, RSMI_status) {
	return DeviceGetUsedMemory(Device, Type)
}

// DeviceGetMemoryUtilization gets the percentage of time any device memory is being used.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetMemoryUtilization(Device DeviceHandle) (uint32, RSMI_status) {
	var percent uint32 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_memory_busy_percent_get"]; ok {
		ret = rsmi_dev_memory_busy_percent_get(Device.index, &percent)
	}
	return percent, ret
}

// GetMemoryUtilization gets the percentage of time any device memory is being used.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetMemoryUtilization() (uint32, RSMI_status) {
	return DeviceGetMemoryUtilization(Device)
}

// DeviceGetMemoryReservedPages gets information about reserved ("retired") memory pages.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_INSUFFICIENT_SIZE should not happen, the size is retrieved and a big enough slice allocated.
func DeviceGetMemoryReservedPages(Device DeviceHandle) ([]RSMI_retired_page_record, RSMI_status) {
	var num_records uint32 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	records := make([]RSMI_retired_page_record, 0)
	ret = rsmi_dev_memory_reserved_pages_get(Device.index, &num_records, nil)
	if ret == STATUS_SUCCESS {
		if num_records > 0 {
			records := make([]RSMI_retired_page_record, num_records)
			ret = rsmi_dev_memory_reserved_pages_get(Device.index, &num_records, &records[0])
		}
	}
	// if _, ok := Device.supported["rsmi_dev_memory_reserved_pages_get"]; ok {
	// 	ret = rsmi_dev_memory_reserved_pages_get(Device.index, &num_records, nil)
	// 	if ret == STATUS_SUCCESS {
	// 		if num_records > 0 {
	// 			records := make([]RSMI_retired_page_record, num_records)
	// 			ret = rsmi_dev_memory_reserved_pages_get(Device.index, &num_records, &records[0])
	// 		}
	// 	}
	// }
	return records, ret
}

// GetMemoryReservedPages gets information about reserved ("retired") memory pages.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_INSUFFICIENT_SIZE should not happen, the size is retrieved and a big enough slice allocated.
func (Device DeviceHandle) GetMemoryReservedPages() ([]RSMI_retired_page_record, RSMI_status) {
	return DeviceGetMemoryReservedPages(Device)
}

// DeviceGetFanRpms gets the fan speed in RPMs of the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetFanRpms(Device DeviceHandle, Sensor uint32) (int64, RSMI_status) {
	var speed int64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_fan_rpms_get"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_fan_rpms_get(Device.index, Sensor, &speed)
		} else if defaultsensor, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaultsensor) > 0 {
				ret = rsmi_dev_fan_rpms_get(Device.index, uint32(defaultsensor[0]), &speed)
			}
		}
	}
	return speed, ret
}

// GetFanRpms gets the fan speed in RPMs of the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetFanRpms(Sensor uint32) (int64, RSMI_status) {
	return DeviceGetFanRpms(Device, Sensor)
}

// DeviceGetFanSpeed gets the fan speed for the specified device as a value relative to MAX_FAN_SPEED.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetFanSpeed(Device DeviceHandle, Sensor uint32) (int64, RSMI_status) {
	var speed int64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_fan_speed_get"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_fan_speed_get(Device.index, Sensor, &speed)
		} else if defaultsensor, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaultsensor) > 0 {
				ret = rsmi_dev_fan_speed_get(Device.index, uint32(defaultsensor[0]), &speed)
			}
		}
	}
	return speed, ret
}

// GetFanSpeed gets the fan speed for the specified device as a value relative to MAX_FAN_SPEED.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetFanSpeed(Sensor uint32) (int64, RSMI_status) {
	return DeviceGetFanSpeed(Device, Sensor)
}

// DeviceGetMaxFanSpeed gets the max. fan speed of the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetMaxFanSpeed(Device DeviceHandle, Sensor uint32) (uint64, RSMI_status) {
	var speed uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_fan_speed_max_get"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_fan_speed_max_get(Device.index, Sensor, &speed)
		} else if defaultsensor, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaultsensor) > 0 {
				ret = rsmi_dev_fan_speed_max_get(Device.index, uint32(defaultsensor[0]), &speed)
			}
		}
	}
	return speed, ret
}

// DeviceGetMaxFanSpeed gets the max. fan speed of the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetMaxFanSpeed(Sensor uint32) (uint64, RSMI_status) {
	return DeviceGetMaxFanSpeed(Device, Sensor)
}

// DeviceGetTemperatureMetric gets the temperature metric value for the specified metric, from the specified temperature sensor on the specified device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetTemperatureMetric(Device DeviceHandle, Sensor RSMI_temperature_type, Metric RSMI_temperature_metric) (int64, RSMI_status) {
	var temp int64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_temp_metric_get"]; ok {
		if metrics, ok := sensors[uint64(Sensor)]; ok {
			avail := false
			for _, m := range metrics {
				if m == uint64(Metric) {
					avail = true
					break
				}
			}
			if avail {
				ret = rsmi_dev_temp_metric_get(Device.index, uint32(Sensor), Metric, &temp)
			}
		} else if defaultsensor, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaultsensor) > 0 {
				ret = rsmi_dev_temp_metric_get(Device.index, uint32(defaultsensor[0]), Metric, &temp)
			}
		}
	}
	return temp, ret
}

// GetTemperatureMetric gets the temperature metric value for the specified metric, from the specified temperature sensor on the specified device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetTemperatureMetric(Sensor RSMI_temperature_type, Metric RSMI_temperature_metric) (int64, RSMI_status) {
	return DeviceGetTemperatureMetric(Device, Sensor, Metric)
}

// DeviceGetVoltageMetric gets the voltage metric value for the specified metric, from the
// specified voltage sensor on the specified device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetVoltageMetric(Device DeviceHandle, Sensor RSMI_voltage_type, Metric RSMI_voltage_metric) (int64, RSMI_status) {
	var voltage int64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_volt_metric_get"]; ok {
		if metrics, ok := sensors[uint64(Sensor)]; ok {
			avail := false
			for _, m := range metrics {
				if m == uint64(Metric) {
					avail = true
					break
				}
			}
			if avail {
				ret = rsmi_dev_volt_metric_get(Device.index, Sensor, Metric, &voltage)
			}
		} else if defaultsensor, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaultsensor) > 0 {
				ret = rsmi_dev_volt_metric_get(Device.index, RSMI_voltage_type(defaultsensor[0]), Metric, &voltage)
			}
		}
	}
	return voltage, ret
}

// GetVoltageMetric gets the voltage metric value for the specified metric, from the
// specified voltage sensor on the specified device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetVoltageMetric(Sensor RSMI_voltage_type, Metric RSMI_voltage_metric) (int64, RSMI_status) {
	return DeviceGetVoltageMetric(Device, Sensor, Metric)
}

// DeviceResetFan resets the fan to automatic driver control.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
func DeviceResetFan(Device DeviceHandle, Sensor uint32) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_fan_reset"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_fan_reset(Device.index, Sensor)
		}
	}
	return ret
}

// ResetFan resets the fan to automatic driver control.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
func (Device DeviceHandle) ResetFan(Sensor uint32) RSMI_status {
	return DeviceResetFan(Device, Sensor)
}

// DeviceSetFanSpeed sets the fan speed for the specified device with the provided speed, in RPMs.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_PERMISSION function requires root access.
func DeviceSetFanSpeed(Device DeviceHandle, Sensor uint32, Speed uint64) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_fan_speed_set"]; ok {
		if _, ok := sensors[uint64(Sensor)]; ok {
			ret = rsmi_dev_fan_speed_set(Device.index, Sensor, Speed)
		}
	}
	return ret
}

// SetFanSpeed sets the fan speed for the specified device with the provided speed, in RPMs.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_PERMISSION function requires root access.
func (Device DeviceHandle) SetFanSpeed(Sensor uint32, Speed uint64) RSMI_status {
	return DeviceSetFanSpeed(Device, Sensor, Speed)
}

// DeviceGetBusyPercent gets the percentage of time device is busy doing any processing.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetBusyPercent(Device DeviceHandle) (uint32, RSMI_status) {
	var util uint32 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_busy_percent_get"]; ok {
		ret = rsmi_dev_busy_percent_get(Device.index, &util)
	}
	return util, ret
}

// GetBusyPercent gets the percentage of time device is busy doing any processing.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetBusyPercent() (uint32, RSMI_status) {
	return DeviceGetBusyPercent(Device)
}

// DeviceGetUtilizationCounters gets coarse grain utilization counter of the specified device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetUtilizationCounters(Device DeviceHandle) ([]RSMI_utilization_counter, uint64, RSMI_status) {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	var util []RSMI_utilization_counter
	var timestamp uint64
	var count uint32 = uint32(UTILIZATION_COUNTER_LAST - UTILIZATION_COUNTER_FIRST + 1)
	util = make([]RSMI_utilization_counter, 0 )
	if _, ok := Device.supported["rsmi_utilization_count_get"]; ok {
		for i := int(UTILIZATION_COUNTER_FIRST); i <= int(UTILIZATION_COUNTER_LAST); i++ {
			util = append(util, RSMI_utilization_counter{
				Type: uint32(i),
				Value: 0,
			})
		}
		ret = rsmi_utilization_count_get(Device.index, &util[0], count, &timestamp)
	}
	return util, timestamp, ret
}

// GetUtilizationCounters gets coarse grain utilization counter of the specified device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetUtilizationCounters() ([]RSMI_utilization_counter, uint64, RSMI_status) {
	return DeviceGetUtilizationCounters(Device)
}

// DeviceGetPerfLevel gets the performance level of the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetPerfLevel(Device DeviceHandle) (RSMI_dev_perf_level, RSMI_status) {
	var level RSMI_dev_perf_level = DEV_PERF_LEVEL_UNKNOWN
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_perf_level_get"]; ok {
		ret = rsmi_dev_perf_level_get(Device.index, &level)
	}
	return level, ret
}

// GetPerfLevel gets the performance level of the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetPerfLevel() (RSMI_dev_perf_level, RSMI_status) {
	return DeviceGetPerfLevel(Device)
}

// DeviceSetDeterminismMode enters performance determinism mode with provided device.
// The performance determinism mode, which enforces a GFXCLK frequency
// SoftMax limit per GPU set by the user. This prevents the GFXCLK PLL from
// stretching when running the same workload on different GPUS, making
// performance variation minimal.
// The performance level is set to DEV_PERF_LEVEL_DETERMINISM.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceSetDeterminismMode(Device DeviceHandle, Clock uint64) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_perf_determinism_mode_set"]; ok {
		ret = rsmi_perf_determinism_mode_set(Device.index, Clock)
	}
	return ret
}

// SetDeterminismMode enters performance determinism mode with provided device.
// The performance determinism mode, which enforces a GFXCLK frequency
// SoftMax limit per GPU set by the user. This prevents the GFXCLK PLL from
// stretching when running the same workload on different GPUS, making
// performance variation minimal.
// The performance level is set to DEV_PERF_LEVEL_DETERMINISM.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) SetDeterminismMode(Clock uint64) RSMI_status {
	return DeviceSetDeterminismMode(Device, Clock)
}

// DeviceGetOverdriveLevel gets the overdrive percent associated with the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetOverdriveLevel(Device DeviceHandle) (uint32, RSMI_status) {
	var level uint32 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_overdrive_level_get"]; ok {
		ret = rsmi_dev_overdrive_level_get(Device.index, &level)
	}
	return level, ret
}

// GetOverdriveLevel gets the overdrive percent associated with the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetOverdriveLevel() (uint32, RSMI_status) {
	return DeviceGetOverdriveLevel(Device)
}

// DeviceGetClockFrequency gets the list of possible system clock speeds of device for a
// specified clock type.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetClockFrequency(Device DeviceHandle, Clock RSMI_clk_type) (RSMI_frequencies, RSMI_status) {
	var freqs RSMI_frequencies = RSMI_frequencies{
		Supported: 0,
		Current: 0,
	}
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if sensors, ok := Device.supported["rsmi_dev_gpu_clk_freq_get"]; ok {
		if _, ok := sensors[uint64(Clock)]; ok {
			ret = rsmi_dev_gpu_clk_freq_get(Device.index, Clock, &freqs)
		} else if defaults, ok := sensors[DEFAULT_VARIANT]; ok {
			if len(defaults) > 0 {
				ret = rsmi_dev_gpu_clk_freq_get(Device.index, RSMI_clk_type(defaults[0]), &freqs)
			}
		}
	}
	return freqs, ret
}

// GetClockFrequency gets the list of possible system clock speeds of device for a
// specified clock type.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetClockFrequency(Clock RSMI_clk_type) (RSMI_frequencies, RSMI_status) {
	return DeviceGetClockFrequency(Device, Clock)
}

// DeviceReset resets the GPU associated with the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceReset(Device DeviceHandle) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_gpu_reset"]; ok {
		ret = rsmi_dev_gpu_reset(int32(Device.index))
	}
	return ret
}

// Reset resets the GPU associated with the device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) Reset() RSMI_status {
	return DeviceReset(Device)
}

// DeviceGetVoltageFrequencyCurve retrieves the voltage/frequency curve information.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetVoltageFrequencyCurve(Device DeviceHandle) (RSMI_od_volt_freq_data, RSMI_status) {
	var data RSMI_od_volt_freq_data = RSMI_od_volt_freq_data{
		Num_regions: 0,
	}
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_od_volt_info_get"]; ok {
		ret = rsmi_dev_od_volt_info_get(Device.index, &data)
	}
	return data, ret
}

// DeviceGetVoltageFrequencyCurve retrieves the voltage/frequency curve information.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetVoltageFrequencyCurve() (RSMI_od_volt_freq_data, RSMI_status) {
	return DeviceGetVoltageFrequencyCurve(Device)
}

// DeviceGetMetrics retrieves the GPU metrics information.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetMetrics(Device DeviceHandle) (RSMI_gpu_metrics, RSMI_status) {
	var data RSMI_gpu_metrics
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_gpu_metrics_info_get"]; ok {
		ret = rsmi_dev_gpu_metrics_info_get(Device.index, &data)
	}
	return data, ret
}

// GetMetrics retrieves the GPU metrics information.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetMetrics() (RSMI_gpu_metrics, RSMI_status) {
	return DeviceGetMetrics(Device)
}

// DeviceSetClockRange sets the clock range information.
// Only usable with clock types CLK_TYPE_SYS and CLK_TYPE_MEM.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceSetClockRange(Device DeviceHandle, MinFreq uint64, MaxFreq uint64, Clock RSMI_clk_type) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if clocks, ok := Device.supported["rsmi_dev_clk_range_set"]; ok {
		if _, ok := clocks[uint64(Clock)]; ok {
			ret = rsmi_dev_clk_range_set(Device.index, MinFreq, MaxFreq, Clock)
		} else if defaults, ok := clocks[DEFAULT_VARIANT]; ok {
			if len(defaults) > 0 {
				ret = rsmi_dev_clk_range_set(Device.index, MinFreq, MaxFreq, RSMI_clk_type(defaults[0]))
			}
		}
	}
	
	return ret
}

// SetClockRange sets the clock range information.
// Only usable with clock types CLK_TYPE_SYS and CLK_TYPE_MEM.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) SetClockRange(MinFreq uint64, MaxFreq uint64, Clock RSMI_clk_type) RSMI_status {
	return DeviceSetClockRange(Device, MinFreq, MaxFreq, Clock)
}

// DeviceSetClockInfo sets the clock frequency information.
// Only usable with clock types CLK_TYPE_SYS and CLK_TYPE_MEM.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceSetClockInfo(Device DeviceHandle, Level RSMI_freq_ind, ClockFreq uint64, Clock RSMI_clk_type) RSMI_status {
	ret := rsmi_dev_od_clk_info_set(Device.index, Level, ClockFreq, Clock)
	return ret
}

// SetClockInfo sets the clock frequency information.
// Only usable with clock types CLK_TYPE_SYS and CLK_TYPE_MEM.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) SetClockInfo(Level RSMI_freq_ind, ClockFreq uint64, Clock RSMI_clk_type) RSMI_status {
	return DeviceSetClockInfo(Device, Level, ClockFreq, Clock)
}

// DeviceSetVoltageInfo sets the voltage curve points.
// the Vpoint argument can be 1,2 or 3 or STATUS_NOT_SUPPORTED is returned.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceSetVoltageInfo(Device DeviceHandle, Vpoint uint32, ClockFreq uint64, Voltage uint64) RSMI_status {
    if Vpoint < 1 || Vpoint > 3 {
        return STATUS_NOT_SUPPORTED
    }
	ret := rsmi_dev_od_volt_info_set(Device.index, Vpoint, ClockFreq, Voltage)
	return ret
}

// SetVoltageInfo sets the voltage curve points.
// the Vpoint argument can be 1,2 or 3 or STATUS_NOT_SUPPORTED is returned.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) SetVoltageInfo(Vpoint uint32, ClockFreq uint64, Voltage uint64) RSMI_status {
	return DeviceSetVoltageInfo(Device, Vpoint, ClockFreq, Voltage)
}

// DeviceGetVoltageFrequencyCurveRegions retrieves the current valid regions in the frequency/voltage space.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
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

// GetVoltageFrequencyCurveRegions retrieves the current valid regions in the frequency/voltage space.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetVoltageFrequencyCurveRegions() ([]RSMI_freq_volt_region, RSMI_status) {
	return DeviceGetVoltageFrequencyCurveRegions(Device)
}

// DeviceGetPowerProfile gets the list of available preset power profiles and an indication of
// which profile is currently active.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetPowerProfile(Device DeviceHandle, Sensor uint32) (RSMI_power_profile_status, RSMI_status) {
	var status RSMI_power_profile_status = RSMI_power_profile_status{
		Available_profiles: 0,
		Current: 0,
		Num_profiles: 0,
	}
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	// Seems to be not listed by rsmi_dev_supported_func_iterator
	// if sensors, ok := Device.supported["rsmi_dev_power_profile_presets_get"]; ok {
	// 	if _, ok := sensors[uint64(Sensor)]; ok {
	// 		ret = rsmi_dev_power_profile_presets_get(Device.index, Sensor, &status)
	// 	} else if defaults, ok := sensors[DEFAULT_VARIANT]; ok {
	// 		if len(defaults) > 0 {
	// 			ret = rsmi_dev_power_profile_presets_get(Device.index, uint32(defaults[0]), &status)
	// 		}
	// 	}
	// }
	ret = rsmi_dev_power_profile_presets_get(Device.index, Sensor, &status)
	return status, ret
}

// GetPowerProfile gets the list of available preset power profiles and an indication of
// which profile is currently active.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetPowerProfile(Sensor uint32) (RSMI_power_profile_status, RSMI_status) {
	return DeviceGetPowerProfile(Device, Sensor)
}

// DeviceSetPerfLevel set the PowerPlay performance level associated with the device.
//
// Note: The RSMI library provides two functions to set the performance level. Which function is called by DeviceSetPerfLevel is
//       determined at initialization.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_PERMISSION function requires root access.
var DeviceSetPerfLevel = deviceSetPerfLevel_v1

func deviceSetPerfLevel_v0(Device DeviceHandle, Level RSMI_dev_perf_level) RSMI_status {
	ret := rsmi_dev_perf_level_set(int32(Device.index), Level)
	return ret
}

// rocm_smi.DeviceSetPerfLevel_v1()
func deviceSetPerfLevel_v1(Device DeviceHandle, Level RSMI_dev_perf_level) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_perf_level_set_v1"]; ok {
		ret = rsmi_dev_perf_level_set_v1(Device.index, Level)
	}
	return ret
}

// SetPerfLevel set the PowerPlay performance level associated with the device.
//
// Note: The RSMI library provides two functions to set the performance level. Which function is called by SetPerfLevel is
//       determined at initialization.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_PERMISSION function requires root access.
func (Device DeviceHandle) SetPerfLevel(Level RSMI_dev_perf_level) RSMI_status {
	return DeviceSetPerfLevel(Device, Level)
}

// DeviceSetOverdriveLevel sets the overdrive percent associated with the device.
//
// Note: The RSMI library provides two functions to set the overdrive level. Which function is called by DeviceSetOverdriveLevel is
//       determined at initialization.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_PERMISSION function requires root access.
var DeviceSetOverdriveLevel = deviceSetOverdriveLevel_v1


func deviceSetOverdriveLevel_v2(Device DeviceHandle, Overdrive uint32) RSMI_status {
	ret := rsmi_dev_overdrive_level_set(int32(Device.index), Overdrive)
	return ret
}

// rocm_smi.DeviceSetOverdriveLevel_v1()
func deviceSetOverdriveLevel_v1(Device DeviceHandle, Overdrive uint32) RSMI_status {
	ret := rsmi_dev_overdrive_level_set_v1(Device.index, Overdrive)
	return ret
}

// SetOverdriveLevel sets the overdrive percent associated with the device.
//
// Note: The RSMI library provides two functions to set the overdrive level. Which function is called by SetOverdriveLevel is
//       determined at initialization.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_PERMISSION function requires root access.
func (Device DeviceHandle) SetOverdriveLevel(Overdrive uint32) RSMI_status {
	return DeviceSetOverdriveLevel(Device, Overdrive)
}

// DeviceSetClockFrequency controls the set of allowed frequencies that can be used for the specified clock.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_PERMISSION function requires root access.
func DeviceSetClockFrequency(Device DeviceHandle, Clock RSMI_clk_type, FreqMask uint64) RSMI_status {
	ret := rsmi_dev_gpu_clk_freq_set(Device.index, Clock, FreqMask)
	return ret
}

// SetClockFrequency controls the set of allowed frequencies that can be used for the specified clock.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_PERMISSION function requires root access.
func (Device DeviceHandle) SetClockFrequency(Clock RSMI_clk_type, FreqMask uint64) RSMI_status {
	return DeviceSetClockFrequency(Device, Clock, FreqMask)
}

// DeviceGetVbiosVersionString gets the VBIOS identifer string.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetVbiosVersionString(Device DeviceHandle) (string, RSMI_status) {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	var version []byte = make([]byte, defaultRsmiStringLength)
	if _, ok := Device.supported["rsmi_dev_vbios_version_get"]; ok {
		vptr := &version[0]
		ret = rsmi_dev_vbios_version_get(Device.index, vptr, defaultRsmiStringLength)
	}
	return bytes2String(version), ret
}

// GetVbiosVersionString gets the VBIOS identifer string.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetVbiosVersionString() (string, RSMI_status) {
	return DeviceGetVbiosVersionString(Device)
}

// DeviceGetFirmwareVersion gets the firmware versions for a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetFirmwareVersion(Device DeviceHandle, Block RSMI_fw_block) (uint64, RSMI_status) {
	var version uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if blocks, ok := Device.supported["rsmi_dev_firmware_version_get"]; ok {
		if _, ok := blocks[uint64(Block)]; ok {
			ret = rsmi_dev_firmware_version_get(Device.index, Block, &version)
		} else if defaults, ok := blocks[DEFAULT_VARIANT]; ok {
			if len(defaults) > 0 {
				ret = rsmi_dev_firmware_version_get(Device.index, RSMI_fw_block(defaults[0]), &version)
			}
		}
	}
	return version, ret
}

// GetFirmwareVersion gets the firmware versions for a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetFirmwareVersion(Block RSMI_fw_block) (uint64, RSMI_status) {
	return DeviceGetFirmwareVersion(Device, Block)
}

// DeviceGetEccCount retrieves the error counts for a GPU block.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetEccCount(Device DeviceHandle, Block RSMI_gpu_block) (RSMI_error_count, RSMI_status) {
	var counts RSMI_error_count = RSMI_error_count{
		Correctable_err: 0,
		Uncorrectable_err: 0,
	}
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if blocks, ok := Device.supported["rsmi_dev_ecc_count_get"]; ok {
		if _, ok := blocks[uint64(Block)]; ok {
			ret = rsmi_dev_ecc_count_get(Device.index, Block, &counts)
		} else if defaults, ok := blocks[DEFAULT_VARIANT]; ok {
			if len(defaults) > 0 {
				ret = rsmi_dev_ecc_count_get(Device.index, RSMI_gpu_block(defaults[0]), &counts)
			}
		}
	}
	return counts, ret
}

// GetEccCount retrieves the error counts for a GPU block.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetEccCount(Block RSMI_gpu_block) (RSMI_error_count, RSMI_status) {
	return DeviceGetEccCount(Device, Block)
}

// DeviceGetEccStatus retrieves the ECC status for a GPU block.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetEccStatus(Device DeviceHandle, Block RSMI_gpu_block) (RSMI_ras_err_state, RSMI_status) {
	var state RSMI_ras_err_state = RAS_ERR_STATE_INVALID
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_ecc_status_get"]; ok {
		ret = rsmi_dev_ecc_status_get(Device.index, Block, &state)
	}
	// // Seems to be not listed by rsmi_dev_supported_func_iterator
	// if blocks, ok := Device.supported["rsmi_dev_ecc_status_get"]; ok {
	// 	fmt.Println(blocks)
	// 	if _, ok := blocks[uint64(Block)]; ok {
	// 		ret = rsmi_dev_ecc_status_get(Device.index, Block, &state)
	// 	} else if defaults, ok := blocks[DEFAULT_VARIANT]; ok {
	// 		if len(defaults) > 0 {
	// 			ret = rsmi_dev_ecc_status_get(Device.index, RSMI_gpu_block(defaults[0]), &state)
	// 		}
	// 	}
	// }
	return state, ret
}

// GetEccStatus retrieves the ECC status for a GPU block.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetEccStatus(Block RSMI_gpu_block) (RSMI_ras_err_state, RSMI_status) {
	return DeviceGetEccStatus(Device, Block)
}

// DeviceGetEccMask retrieved the enabled ECC bit-mask.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetEccMask(Device DeviceHandle) (uint64, RSMI_status) {
	var mask uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_ecc_enabled_get"]; ok {
		ret = rsmi_dev_ecc_enabled_get(Device.index, &mask)
	}
	return mask, ret
}

// GetEccMask retrieved the enabled ECC bit-mask.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetEccMask() (uint64, RSMI_status) {
	return DeviceGetEccMask(Device)
}



// XGMI Functions

// DeviceXgmiErrorStatus retrieves the XGMI error status for a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceXgmiErrorStatus(Device DeviceHandle) (RSMI_xgmi_status, RSMI_status) {
	var status RSMI_xgmi_status = XGMI_STATUS_NO_ERRORS
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	// if _, ok := Device.supported["rsmi_dev_xgmi_error_status"]; ok {
	// 	ret = rsmi_dev_xgmi_error_status(Device.index, &status)
	// }
	ret = rsmi_dev_xgmi_error_status(Device.index, &status)
	return status, ret
}

// XgmiErrorStatus retrieves the XGMI error status for a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) XgmiErrorStatus() (RSMI_xgmi_status, RSMI_status) {
	return DeviceXgmiErrorStatus(Device)
}

// DeviceXgmiErrorReset resets the XGMI error status for a device.
//
// STATUS_SUCCESS call was successful.
func DeviceXgmiErrorReset(Device DeviceHandle) RSMI_status {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	// if _, ok := Device.supported["rsmi_dev_xgmi_error_reset"]; ok {
	// 	ret = rsmi_dev_xgmi_error_reset(Device.index)
	// }
	ret = rsmi_dev_xgmi_error_reset(Device.index)
	return ret
}

// XgmiErrorReset resets the XGMI error status for a device.
//
// STATUS_SUCCESS call was successful.
func (Device DeviceHandle) XgmiErrorReset() RSMI_status {
	return DeviceXgmiErrorReset(Device)
}

// DeviceXgmiHiveId retrieves the XGMI hive id for a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceXgmiHiveId(Device DeviceHandle) (uint64, RSMI_status) {
	var id uint64 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	// Seems to be not listed by rsmi_dev_supported_func_iterator
	// if _, ok := Device.supported["rsmi_dev_xgmi_hive_id_get"]; ok {
	// 	ret = rsmi_dev_xgmi_hive_id_get(Device.index, &id)
	// }
	ret = rsmi_dev_xgmi_hive_id_get(Device.index, &id)
	return id, ret
}

// XgmiHiveId retrieves the XGMI hive id for a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) XgmiHiveId() (uint64, RSMI_status) {
	return DeviceXgmiHiveId(Device)
}

// Hardware Topology Functions

// DeviceGetNumaNode retrieves the NUMA CPU node number for a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetNumaNode(Device DeviceHandle) (uint32, RSMI_status) {
	var node uint32 = 0
	ret := rsmi_topo_get_numa_node_number(Device.index, &node)
	return node, ret
}

// GetNumaNode retrieves the NUMA CPU node number for a device.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetNumaNode() (uint32, RSMI_status) {
	return DeviceGetNumaNode(Device)
}

// DeviceGetLinkWeight retrieves the weight for a connection between 2 GPUs.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetLinkWeight(SrcDevice DeviceHandle, DstDevice DeviceHandle) (uint64, RSMI_status) {
	var weight uint64 = 0
	ret := rsmi_topo_get_link_weight(SrcDevice.index, DstDevice.index, &weight)
	return weight, ret
}

// GetLinkWeight retrieves the weight for a connection between 2 GPUs.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetLinkWeight(DstDevice DeviceHandle) (uint64, RSMI_status) {
	return DeviceGetLinkWeight(Device, DstDevice)
}

// DeviceGetMinMaxBandwidth retreives minimal and maximal IO link bandwidth between 2 GPUs.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetMinMaxBandwidth(SrcDevice DeviceHandle, DstDevice DeviceHandle) (uint64, uint64, RSMI_status) {
	var mini uint64 = 0
	var maxi uint64 = 0
	ret := rsmi_minmax_bandwidth_get(SrcDevice.index, DstDevice.index, &mini, &maxi)
	return mini, maxi, ret
}

// GetMinMaxBandwidth retreives minimal and maximal IO link bandwidth between 2 GPUs.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetMinMaxBandwidth(DstDevice DeviceHandle) (uint64, uint64, RSMI_status) {
	return DeviceGetMinMaxBandwidth(Device, DstDevice)
}

// DeviceGetLinkType retrieves the hops and the connection type between 2 GPUs.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceGetLinkType(SrcDevice DeviceHandle, DstDevice DeviceHandle) (uint64, RSMI_IO_LINK_TYPE, RSMI_status) {
	var hops uint64 = 0
	var linkType RSMI_IO_LINK_TYPE = IOLINK_TYPE_UNDEFINED
	ret := rsmi_topo_get_link_type(SrcDevice.index, DstDevice.index, &hops, &linkType)
	return hops, linkType, ret
}

// GetLinkType retrieves the hops and the connection type between 2 GPUs.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) GetLinkType(DstDevice DeviceHandle) (uint64, RSMI_IO_LINK_TYPE, RSMI_status) {
	return DeviceGetLinkType(Device, DstDevice)
}

// DeviceIsP2PAccessible returns the P2P availability status between 2 GPUs.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceIsP2PAccessible(SrcDevice DeviceHandle, DstDevice DeviceHandle) (bool, RSMI_status) {
	var access bool = false
	ret := rsmi_is_P2P_accessible(SrcDevice.index, DstDevice.index, &access)
	return access, ret
}

// IsP2PAccessible returns the P2P availability status between 2 GPUs.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) IsP2PAccessible(DstDevice DeviceHandle) (bool, RSMI_status) {
	return DeviceIsP2PAccessible(Device, DstDevice)
}


// Event Notification Functions

// DeviceInitEventNotification prepares to collect event notifications for a GPU.
//
// STATUS_SUCCESS call was successful.
func DeviceInitEventNotification(Device DeviceHandle) RSMI_status {
	ret := rsmi_event_notification_init(Device.index)
	return ret
}

// InitEventNotification prepares to collect event notifications for a GPU.
//
// STATUS_SUCCESS call was successful.
func (Device DeviceHandle) InitEventNotification() RSMI_status {
	return DeviceInitEventNotification(Device)
}

// DeviceSetEventNotificationMask specifies which events to collect for a device. The events
// set in mask are OR'd together.
//
// STATUS_SUCCESS call was successful.
// STATUS_INIT_ERROR is returned if DeviceInitEventNotification() has not been called before.
func DeviceSetEventNotificationMask(Device DeviceHandle, Mask uint64) RSMI_status {
	ret := rsmi_event_notification_mask_set(Device.index, Mask)
	return ret
}

// SetEventNotificationMask specifies which events to collect for a device. The events
// set in mask are OR'd together.
//
// STATUS_SUCCESS call was successful.
// STATUS_INIT_ERROR is returned if DeviceInitEventNotification() has not been called before.
func (Device DeviceHandle) SetEventNotificationMask(Mask uint64) RSMI_status {
	return DeviceSetEventNotificationMask(Device, Mask)
}

// GetEventNotification collects event notifications, waiting a specified amount of time.
//
// STATUS_SUCCESS call was successful.
// STATUS_NO_DATA No events were found to collect.
func GetEventNotification(TimeoutMs int32) ([]RSMI_evt_notification_data, RSMI_status) {
	var num_events uint32 = 0
	data := make([]RSMI_evt_notification_data, 0)
	ret := rsmi_event_notification_get(TimeoutMs, &num_events, nil)
	if ret == STATUS_SUCCESS && num_events > 0 {
		data = make([]RSMI_evt_notification_data, num_events)
		ret = rsmi_event_notification_get(TimeoutMs, &num_events, &data[0])
	}
	return data, ret
}

// DeviceStopEventNotification closes any file handles and free any resources used by event
// notification for a GPU.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS resources for the given device have either already been freed, or were never allocated by DeviceInitEventNotification.
func DeviceStopEventNotification(Device DeviceHandle) RSMI_status {
	ret := rsmi_event_notification_stop(Device.index)
	return ret
}

// StopEventNotification closes any file handles and free any resources used by event
// notification for a GPU.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS resources for the given device have either already been freed, or were never allocated by DeviceInitEventNotification.
func (Device DeviceHandle) StopEventNotification() RSMI_status {
	return DeviceStopEventNotification(Device)
}
