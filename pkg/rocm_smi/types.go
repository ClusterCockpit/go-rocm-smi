// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs /home/hpc/unrz/unrz139/Work/go-rocm-smi/pkg/rocm_smi/types.go

package rocm_smi

import "unsafe"

type RSMI_event_handle uint32

type RSMI_counter_value struct {
	Value	uint64
	Enabled	uint64
	Running	uint64
}

type RSMI_evt_notification_data struct {
	Ind	uint32
	Event	uint32
	Message	[64]int8
}

type RSMI_bit_field uint64

type RSMI_utilization_counter struct {
	Type	uint32
	Value	uint64
}

type RSMI_retired_page_record struct {
	Address		uint64
	Size		uint64
	Status		uint32
	Pad_cgo_0	[4]byte
}

type RSMI_power_profile_status struct {
	Available_profiles	uint64
	Current			int64
	Num_profiles		uint32
	Pad_cgo_0		[4]byte
}

type RSMI_frequencies struct {
	Supported	uint32
	Current		uint32
	Frequency	[32]uint64
}

type RSMI_pcie_bandwidth struct {
	Rate	RSMI_frequencies
	Lanes	[32]uint32
}

type RSMI_version struct {
	Major	uint32
	Minor	uint32
	Patch	uint32
	Build	*int8
}

type RSMI_range struct {
	Lower_bound	uint64
	Upper_bound	uint64
}

type RSMI_od_vddc_point struct {
	Frequency	uint64
	Voltage		uint64
}

type RSMI_freq_volt_region struct {
	Freq_range	RSMI_range
	Volt_range	RSMI_range
}

type RSMI_od_volt_curve struct {
	Points [3]RSMI_od_vddc_point
}

type RSMI_od_volt_freq_data struct {
	Curr_sclk_range		RSMI_range
	Curr_mclk_range		RSMI_range
	Sclk_freq_limits	RSMI_range
	Mclk_freq_limits	RSMI_range
	Curve			RSMI_od_volt_curve
	Num_regions		uint32
	Pad_cgo_0		[4]byte
}

type RSMI_gpu_metrics struct {
	Common_header			_Ctype_struct_metrics_table_header_t
	Temperature_edge		uint16
	Temperature_hotspot		uint16
	Temperature_mem			uint16
	Temperature_vrgfx		uint16
	Temperature_vrsoc		uint16
	Temperature_vrmem		uint16
	Average_gfx_activity		uint16
	Average_umc_activity		uint16
	Average_mm_activity		uint16
	Average_socket_power		uint16
	Energy_accumulator		uint64
	System_clock_counter		uint64
	Average_gfxclk_frequency	uint16
	Average_socclk_frequency	uint16
	Average_uclk_frequency		uint16
	Average_vclk0_frequency		uint16
	Average_dclk0_frequency		uint16
	Average_vclk1_frequency		uint16
	Average_dclk1_frequency		uint16
	Current_gfxclk			uint16
	Current_socclk			uint16
	Current_uclk			uint16
	Current_vclk0			uint16
	Current_dclk0			uint16
	Current_vclk1			uint16
	Current_dclk1			uint16
	Throttle_status			uint32
	Current_fan_speed		uint16
	Pcie_link_width			uint16
	Pcie_link_speed			uint16
	Padding				uint16
	Gfx_activity_acc		uint32
	Mem_actvity_acc			uint32
	Temperature_hbm			[4]uint16
}

type RSMI_error_count struct {
	Correctable_err		uint64
	Uncorrectable_err	uint64
}

type RSMI_process_info struct {
	Process_id	uint32
	Pasid		uint32
	Vram_usage	uint64
	Sdma_usage	uint64
	Cu_occupancy	uint32
	Pad_cgo_0	[4]byte
}

type RSMI_func_id_iter_handle *_Ctype_struct_rsmi_func_id_iter_handle

const sizeofRSMI_func_id_value = unsafe.Sizeof([8]byte{})

type RSMI_func_id_value [sizeofRSMI_func_id_value]byte

type RSMI_helper_variants struct {
	Num_variants	int32
	Variant_id	uint64
	Variantlist	[100]uint64
}

type RSMI_helper_function_variants struct {
	Name		[200]int8
	Sensors		int32
	Sensorlist	[30]RSMI_helper_variants
}

type RSMI_helper_function struct {
	Functions	int32
	Functionlist	[300]RSMI_helper_function_variants
}
