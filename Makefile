CFORGO = c-for-go
CGO = $(shell go env GOTOOLDIR)/cgo
PKG_DIR = $(shell readlink -f pkg)

.PHONY: all
all: generate

UINT64_TYPE = uint64_t
INT64_TYPE = int64_t
TEST_VERBOSE = -v

.PHONY: generate
generate:
	cd pkg/rocm_smi && \
		sed -e "s/uint64_t/$(UINT64_TYPE)/g" \
		    -e "s/int64_t/$(INT64_TYPE)/g" \
		    -e "s/    bool/    _Bool/g" \
		    -e "s/union id/union id_rename/g" \
		    rocm_smi/rocm_smi.h.orig > rocm_smi/rocm_smi.h
	cd pkg/rocm_smi && $(CFORGO) -ccincl -ccdefs -out $(PKG_DIR) ../../rocm_smi.yml
	cd pkg/rocm_smi && \
		$(CGO) -godefs types.go > types.go.expand && \
		mv types.go.expand types.go && \
		rm -f _cgo_2.o
	cd pkg/rocm_smi && \
		for T in RSMI_status RSMI_event_group RSMI_clk_type RSMI_temperature_type RSMI_voltage_type RSMI_ras_err_state RSMI_freq_ind RSMI_IO_LINK_TYPE RSMI_POWER_TYPE; do \
			sed -i -e "s/type $${T} int32/type $${T} uint32/g" const.go; \
		done && \
		for T in RSMI_init_flags RSMI_power_profile_preset_masks RSMI_gpu_block; do \
			sed -i -e "s/type $${T} int32/type $${T} uint64/g" const.go; \
		done

.PHONY: clean
clean:
	rm -f pkg/rocm_smi/cgo_helpers.go pkg/rocm_smi/cgo_helpers.h pkg/rocm_smi/cgo_helpers.c
	rm -f pkg/rocm_smi/const.go pkg/rocm_smi/doc.go pkg/rocm_smi/types.go
	rm -f pkg/rocm_smi/rocm_smi.go

.PHONY: build
build:
	cd pkg/rocm_smi && go build

.PHONY: test
test: build
	cd pkg/rocm_smi && go test $(TEST_VERBOSE)

.PHONY: fmt
fmt:
	cd pkg/rocm_smi && go fmt


# Examine Go source code and reports suspicious constructs
.PHONY: vet
vet:
	cd pkg/rocm_smi && go vet ./...
