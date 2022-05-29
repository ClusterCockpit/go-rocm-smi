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