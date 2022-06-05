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

// ComputeProcesses gets process information about processes currently using GPU.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_INSUFFICIENT_SIZE should not happen as the size is retrieved and a big enough slice allocated.
func ComputeProcesses() ([]RSMI_process_info, RSMI_status) {
	var ret RSMI_status = STATUS_SUCCESS
	var num_procs uint32 = 0
	var procs []RSMI_process_info = make([]RSMI_process_info, 0)

	ret = rsmi_compute_process_info_get(nil, &num_procs)
	if ret == STATUS_SUCCESS {
		if num_procs > 0 {
			procs = make([]RSMI_process_info, num_procs)
			ret = rsmi_compute_process_info_get(&procs[0], &num_procs)
		}
	}

	return procs, ret
}
// ComputeProcessByPid gets process information about a specific process.
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_NOT_FOUND is returned if there was no process information found for the provided Pid.
func ComputeProcessByPid(Pid uint32) (RSMI_process_info, RSMI_status) {
	var proc RSMI_process_info
	ret := rsmi_compute_process_info_by_pid_get(Pid, &proc)
	return proc, ret
}

// ComputeProcessGpus gets the device indices currently being used by a process
//
// STATUS_SUCCESS call was successful.
// STATUS_INVALID_ARGS the provided arguments are not valid.
// STATUS_INSUFFICIENT_SIZE should not happen as the size is retrieved and a big enough slice allocated.
func ComputeProcessGpus(Pid uint32) ([]uint32, RSMI_status) {
	var num_devs uint32 = 0
	var devs []uint32
	ret := rsmi_compute_process_gpus_get(Pid, nil, &num_devs)
	if ret == STATUS_SUCCESS {
		devs = make([]uint32, num_devs)
		ret := rsmi_compute_process_gpus_get(Pid, &devs[0], &num_devs)
		return devs, ret
	}
	return make([]uint32, 0), ret
}
