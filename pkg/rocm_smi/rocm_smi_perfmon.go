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



// Performance Counter Functions

// DeviceCounterGroupSupported tells if an event group is supported by a given device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
func DeviceCounterGroupSupported(Device DeviceHandle, EventGroup RSMI_event_group) (RSMI_status) {
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_counter_group_supported"]; ok {
		ret = rsmi_dev_counter_group_supported(Device.index, EventGroup)
	}
	return ret
}

// CounterGroupSupported tells if an event group is supported by a given device.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
func (Device DeviceHandle) CounterGroupSupported(EventGroup RSMI_event_group) (RSMI_status) {
	return DeviceCounterGroupSupported(Device, EventGroup)
}

// DeviceCounterCreate creates a performance counter object.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_OUT_OF_RESOURCES unable to allocate memory for counter.
// STATUS_PERMISSION function requires root access.
func DeviceCounterCreate(Device DeviceHandle, Type RSMI_event_type) (RSMI_event_handle, RSMI_status) {
	var event RSMI_event_handle = RSMI_event_handle(0)
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if _, ok := Device.supported["rsmi_dev_counter_create"]; ok {
		ret = rsmi_dev_counter_create(Device.index, Type, &event)
	}
	return event, ret
}

// CounterCreate creates a performance counter object.
//
// STATUS_SUCCESS call was successful.
// STATUS_NOT_SUPPORTED installed software or hardware does not support this function with the given arguments.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_OUT_OF_RESOURCES unable to allocate memory for counter.
// STATUS_PERMISSION function requires root access.
func (Device DeviceHandle) CounterCreate(Type RSMI_event_type) (RSMI_event_handle, RSMI_status) {
	return DeviceCounterCreate(Device, Type)
}

// CounterDestroy deallocates a performance counter object.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_PERMISSION function requires root access.
func CounterDestroy(Handle RSMI_event_handle) (RSMI_status) {
	ret := rsmi_dev_counter_destroy(Handle)
	return ret
}

// CounterDestroy deallocates a performance counter object.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_PERMISSION function requires root access.
func (Handle RSMI_event_handle) CounterDestroy() (RSMI_status) {
	return CounterDestroy(Handle)
}

// CounterControl issues performance counter control commands.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_PERMISSION function requires root access.
func CounterControl(Handle RSMI_event_handle, Cmd RSMI_counter_command) (RSMI_status) {
	ret := rsmi_counter_control(Handle, Cmd, nil)
	return ret
}

// CounterControl issues performance counter control commands.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_PERMISSION function requires root access.
func (Handle RSMI_event_handle) CounterControl(Cmd RSMI_counter_command) (RSMI_status) {
	return CounterControl(Handle, Cmd)
}

// CounterRead reads the current value of a performance counter.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_PERMISSION function requires root access.
func CounterRead(Handle RSMI_event_handle) (RSMI_counter_value, RSMI_status) {
	var count RSMI_counter_value = RSMI_counter_value{
		Value: 0,
		Enabled: 0,
		Running: 0,
	}
	ret := rsmi_counter_read(Handle, &count)
	return count, ret
}

// CounterRead reads the current value of a performance counter.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_PERMISSION function requires root access.
func (Handle RSMI_event_handle) CounterRead() (RSMI_counter_value, RSMI_status) {
	return CounterRead(Handle)
}

// DeviceCounterGetAvailable gets the number of currently available counters.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func DeviceCounterGetAvailable(Device DeviceHandle, EventGroup RSMI_event_group) (uint32, RSMI_status) {
	var count uint32 = 0
	var ret RSMI_status = STATUS_NOT_SUPPORTED
	if groups, ok := Device.supported["rsmi_counter_available_counters_get"]; ok {
		if _, ok := groups[uint64(EventGroup)]; ok {
			ret = rsmi_counter_available_counters_get(Device.index, EventGroup, &count)
		}
	}
	return count, ret
}

// DeviceCounterGetAvailable gets the number of currently available counters.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
func (Device DeviceHandle) CounterGetAvailable(EventGroup RSMI_event_group) (uint32, RSMI_status) {
	return DeviceCounterGetAvailable(Device, EventGroup)
}
