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

func TestFuncVariants(t *testing.T) {
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
		t.Logf("  handle index: %v", devHandle.Index())
		t.Logf("  handle id: %v", devHandle.ID())
		for f := range devHandle.Supported() {
			t.Logf("  handle function: %v", f)
		}
		
	}

	
}
