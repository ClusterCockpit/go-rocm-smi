package main

import (
	"fmt"
	"log"

	rocm_smi "github.com/ClusterCockpit/go-rocm-smi"
)

func main() {
	ret := rocm_smi.Init()
	if ret != rocm_smi.STATUS_SUCCESS {
		log.Fatalf("Unable to initialize ROCM SMI: %v", rocm_smi.StatusStringNoError(ret))
	}
	defer func() {
		ret := rocm_smi.Shutdown()
		if ret != rocm_smi.STATUS_SUCCESS {
			log.Fatalf("Unable to shutdown ROCM SMI: %v", rocm_smi.StatusStringNoError(ret))
		}
	}()

	count, ret := rocm_smi.NumMonitorDevices()
	if ret != rocm_smi.STATUS_SUCCESS {
		log.Fatalf("Unable to get device count: %v", rocm_smi.StatusStringNoError(ret))
	}

	for i := 0; i < count; i++ {
		device, ret := rocm_smi.DeviceGetHandleByIndex(i)
		if ret != rocm_smi.STATUS_SUCCESS {
			log.Fatalf("Unable to get device at index %d: %v", i, rocm_smi.StatusStringNoError(ret))
		}

		uuid, ret := device.GetUniqueId()
		if ret != rocm_smi.STATUS_SUCCESS {
			log.Fatalf("Unable to get uuid of device at index %d: %v", i, rocm_smi.StatusStringNoError(ret))
		}

		fmt.Printf("%v\n", uuid)
	}
}
