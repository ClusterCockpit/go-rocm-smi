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
	"time"
)

func TestPerfmon(t *testing.T) {
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
		t.Logf("  handle: %v %v", devHandle.Index(), devHandle.ID())
	}

	groups := []RSMI_event_group{EVNT_GRP_XGMI, EVNT_GRP_XGMI_DATA_OUT}
	for _, g := range groups {
		ret = DeviceCounterGroupSupported(devHandle, g)
		if ret == STATUS_NOT_SUPPORTED {
			t.Logf("DeviceCounterGroupSupported(%v): %v (STATUS_NOT_SUPPORTED)", g, ret)
		} else if ret != STATUS_SUCCESS {
			t.Errorf("DeviceCounterGroupSupported(%v): %v", g, ret)
		} else {
			t.Logf("DeviceCounterGroupSupported(%v): %v", g, ret)
		}

		if ret == STATUS_SUCCESS {
			counters, ret := DeviceCounterGetAvailable(devHandle, g)
			if ret == STATUS_NOT_SUPPORTED {
				t.Logf("DeviceCounterGetAvailable(%v): %v (STATUS_NOT_SUPPORTED)", g, ret)
			} else if ret != STATUS_SUCCESS {
				t.Errorf("DeviceCounterGetAvailable(%v): %v", g, ret)
			} else {
				t.Logf("DeviceCounterGetAvailable(%v): %v", g, ret)
				t.Logf("  counters: %v", counters)
			}

			if (counters == 0) {
				continue
			}

			ctr, ret := DeviceCounterCreate(devHandle, RSMI_event_type(g))
			if ret == STATUS_NOT_SUPPORTED {
				t.Logf("DeviceCounterCreate(%v): %v (STATUS_NOT_SUPPORTED)", g, ret)
			} else if ret == STATUS_PERMISSION {
				t.Logf("DeviceCounterCreate(%v): %v (STATUS_PERMISSION)", g, ret)
			} else if ret != STATUS_SUCCESS {
				t.Errorf("DeviceCounterCreate(%v): %v", g, ret)
			} else {
				t.Logf("DeviceCounterCreate(%v): %v", g, ret)
				t.Logf("  ctr: %v", ctr)
			}

			if ret == STATUS_SUCCESS {
				ret = CounterControl(ctr, CNTR_CMD_START)
				if ret == STATUS_NOT_SUPPORTED {
					t.Logf("CounterControl(%v, CNTR_CMD_START): %v (STATUS_NOT_SUPPORTED)", ctr, ret)
				} else if ret == STATUS_PERMISSION {
						t.Logf("CounterControl(%v, CNTR_CMD_START): %v (STATUS_PERMISSION)", ctr, ret)
				} else if ret != STATUS_SUCCESS {
					t.Errorf("CounterControl(%v, CNTR_CMD_START): %v", ctr, ret)
				} else {
					t.Logf("CounterControl(%v, CNTR_CMD_START): %v", ctr, ret)
				}

				if ret == STATUS_SUCCESS {
					time.Sleep(time.Duration(1) * time.Second)

					ret = CounterControl(ctr, CNTR_CMD_STOP)
					if ret == STATUS_NOT_SUPPORTED {
						t.Logf("CounterControl(%v, CNTR_CMD_STOP): %v (STATUS_NOT_SUPPORTED)", ctr, ret)
					} else if ret == STATUS_PERMISSION {
							t.Logf("CounterControl(%v, CNTR_CMD_STOP): %v (STATUS_PERMISSION)", ctr, ret)
					} else if ret != STATUS_SUCCESS {
						t.Errorf("CounterControl(%v, CNTR_CMD_STOP): %v", ctr, ret)
					} else {
						t.Logf("CounterControl(%v, CNTR_CMD_STOP): %v", ctr, ret)
					}

					if ret == STATUS_SUCCESS {
						count, ret := CounterRead(ctr)
						if ret == STATUS_NOT_SUPPORTED {
							t.Logf("CounterRead(%v): %v (STATUS_NOT_SUPPORTED)", ctr, ret)
						} else if ret == STATUS_PERMISSION {
								t.Logf("CounterRead(%v): %v (STATUS_PERMISSION)", ctr, ret)
						} else if ret != STATUS_SUCCESS {
							t.Errorf("CounterRead(%v): %v", ctr, ret)
						} else {
							t.Logf("CounterRead(%v): %v", ctr, ret)
							t.Logf("  timeEnabled: %v", count.Enabled)
							t.Logf("  timeRunning: %v", count.Running)
							t.Logf("  value: %v", count.Value)
						}
					}
				}

				ret = CounterDestroy(ctr)
				if ret == STATUS_NOT_SUPPORTED {
					t.Logf("CounterDestroy(%v): %v (STATUS_NOT_SUPPORTED)", ctr, ret)
				} else if ret == STATUS_PERMISSION {
						t.Logf("CounterDestroy(%v): %v (STATUS_PERMISSION)", g, ret)
				} else if ret != STATUS_SUCCESS {
					t.Errorf("CounterDestroy(%v): %v", ctr, ret)
				} else {
					t.Logf("CounterDestroy(%v): %v", ctr, ret)
				}
			}
		}
	}
}