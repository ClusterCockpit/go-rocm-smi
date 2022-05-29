# Introduction

This is an unofficial interface to the AMD ROCM SMI library for Golang applications. It is heavily
inspired by [`go-nvml`](https://github.com/NVIDIA/go-nvml) by also using [`cgo`](https://golang.org/cmd/cgo/), [`c-for-go`](https://c.for-go.com/) and its [`dlopen` wrapper](https://github.com/NVIDIA/go-nvml/tree/main/pkg/dl).

This Golang interface is planned to be used in [cc-metric-collector](https://github.com/ClusterCockpit/cc-metric-collector).

**Disclaimer**: These bindings are created without any collaboration with AMD. Use them as you like but we, the developers of these bindings, are not responsible for any damage or anything that was caused by them. If you want official Golang bindings for the ROCm SMI library, use [this](https://github.com/amd/go_amd_smi) package.

# Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/ClusterCockpit/go-rocm-smi/pkg/rocm_smi"
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
```

The `librocm_smi64.so` is dynamically loaded by the `rocm_smi` package. Make sure that the directory containing this library is in your `LD_LIBRARY_PATH`.

# Documentation
See [pkg.go.dev](https://pkg.go.dev/github.com/ClusterCockpit/go-rocm-smi).

# Generating the bindings

## ROCm SMI Headers

There are three ROCm SMI Headers, all located at `rocm_smi/rocm_smi`
- `rocm_smi.h`
- `rocm_smi64Config.h`
- `kfd_ioctl.h`

The files are copied from ROCm 5.1.0. For the generation, the `rocm_smi.h` header is changed to support [`c-for- go`](https://c.for-go.com/)'s parser.
- All occurences of `uint64_t` are changed to `unsigned long long`, otherwise [`c-for-go`](https://c.for-go.com/) wouldn't use Golang's `uint64` type.
- All occurences of `int64_t` are changed to `long long`, otherwise [`c-for-go`](https://c.for-go.com/) wouldn't use Golang's `int64` type.
- The `union id` is renamed to `union id_rename` to avoid problems with clang. The type is never addressed with the name `id` but a `typedef` name.

## Generation

Calling [`c-for-go`](https://c.for-go.com/) with the `rocm_smi.yml` as input

## Post processing

After the generation, the `types.go` file still contains the C types but it is more suitable to have
Golang types for them. Luckly [`cgo`](https://golang.org/cmd/cgo/) has a bootstrapping option `-godefs` to
generate the Go types.

Before:
```go
type RSMI_pcie_bandwidth C.rsmi_pcie_bandwidth_t
```
After:
```go
type RSMI_pcie_bandwidth struct {
	Rate	RSMI_frequencies
	Lanes	[32]uint32
}
```

## Manual labor

In the end, the generated functions are wrapped to have more Golang style. This is similar to the
wrappers created in [`go-nvml`](https://github.com/NVIDIA/go-nvml). Most of them are straight-forward
with a little bit of casting.

```go
// rocm_smi.DeviceGetSerial()
func DeviceGetSerial(Device DeviceHandle) (string, RSMI_status) {
	var Serial []byte = make([]byte, 100)
	sptr := &Serial[0]
	ret := rsmi_dev_serial_number_get(Device.index, sptr, 100)
	return bytes2String(Serial), ret
}

func (Device DeviceHandle) DeviceGetSerial() (string, RSMI_status) {
	return DeviceGetSerial(Device)
}
```



# The device index and the "device handle"

For most libraries which handle multiple devices ([`go-nvml`](https://github.com/NVIDIA/go-nvml) is an example), the user at first requests a handle for each device, mostly through the logical index in the list of available devices. The official `rocm_smi` library uses the logical index instead but in order to get everything right, you have to do quite some work to know what is supported. The `rocm_smi` provides a feature (`APISupport` in `rocm_smi.h`) to determine which functions are supported for a device and if a function accepts arguments, which ones are valid for this device. An example would be the function to get the firmware version and the list of GPU parts that provide such a version. The `go-rocm-smi` bindings introduce a virtual type `DeviceHandle`, retrivable through the logical index (so similar to [`go-nvml`](https://github.com/ NVIDIA/go-nvml)), which encapsulates the `APISupport` lookup: `DeviceGetHandleByIndex()`. The `DeviceHandle` is used for all device related calls in `go-rocm-smi`. You can get the logical index by `deviceHandle.Index()`, the **not** unique ID of a GPU by `deviceHandle.ID()` and the list of supported functions through `deviceHandle.Supported()`


# Problems
- One big problem is currently, that [`c-for-go`](https://c.for-go.com/) does not generate `uint64` types for the C type `uint64_t`. It is one of the main data type used in the ROCm SMI headers. While I was able to generate underlying code for `uint64_t`, the Golang function still uses `uint32`:
  ```C
  rsmi_status_t rsmi_dev_unique_id_get(uint32_t dv_ind, uint64_t *id);
  ```
  Output:
  ```go
  func rsmi_dev_unique_id_get(Dv_ind uint32, Id *uint32) RSMI_status {
	cDv_ind, cDv_indAllocMap := (C.uint32_t)(Dv_ind), cgoAllocsUnknown
	cId, cIdAllocMap := (*C.uint64_t)(unsafe.Pointer(Id)), cgoAllocsUnknown
	__ret := C.rsmi_dev_unique_id_get(cDv_ind, cId)
	runtime.KeepAlive(cIdAllocMap)
	runtime.KeepAlive(cDv_indAllocMap)
	__v := (RSMI_status)(__ret)
	return __v
  }
  ```
  One can see, that the `cId` is casted to `*C.uint64_t`, but the `Id` variable used by the function is `*uint32`. I was not able to persuade [`c-for-go`](https://c.for-go.com/) to use `uint64`. See also https://github.com/xlab/c-for-go/issues/120. As a workaround, `uint64_t` gets replaced by `unsigned long long` and `int64_t` gets replaced by `long long`, see `Makefile`. Interestingly, the translation of the C types to Golang types with [`cgo`](https://golang.org/cmd/cgo/) generates `uint64` without the type exchange in the header. If we wouldn't use `unsigned long long`, the `uint32` generated by [`c-for-go`](https://c.for-go.com/) would clash with the `uint64` generated by [`cgo`](https://golang.org/cmd/cgo/).

- The symbol `rsmi_dev_sku_get` is defined by the `rocm_smi.h` header but on the test system with ROCm 5.1.0, the symbol lookup fails. There is now an `updateFunctionPointers()` function that is called at `Init()`. This is quite similar the function `updateVersionedSymbols()` in [`go-nvml`](https://github.com/NVIDIA/go-nvml). The `APISupport` feature of the `rocm_smi` library shows, `rsmi_dev_sku_get` is supported by the device.

- The function `rsmi_status_string` cannot use the wrapper generated by [`c-for-go`](https://c.for-go.com/) because it requires a pointer to a `char` array while [`c-for-go`](https://c.for-go.com/) wants to use the `char` array directly. There is a manually created version to get the status string `StatusString()`. One issue is when using it in prints (see example) because `rsmi_status_string` accepts a status and returns a new status and the string. To drop the new status, use `StatusStringNoError()`.

- I havn't found a way to access the `Build` field in `RSMI_version`. It is a `char*` in `rocm_smi` but [`c-for-go`](https://c.for-go.com/) generates an `*int8` entry for it.
