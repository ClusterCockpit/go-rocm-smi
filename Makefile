
CFORGO = ${HOME}/go/bin/c-for-go
CGO = cgo

all: generate

UINT64_TYPE = unsigned long long
INT64_TYPE = long long
#TEST_VERBOSE = -v

generate:
	cd pkg && \
		sed -e "s/uint64_t/$(UINT64_TYPE)/g" -e "s/int64_t/$(INT64_TYPE)/g" rocm_smi/rocm_smi/rocm_smi.h.orig > rocm_smi/rocm_smi/rocm_smi.h && \
		$(CFORGO) -ccincl --ccdefs ../rocm_smi.yml && \
		cd -
	cd pkg/rocm_smi && \
		$(CGO) -godefs types.go > types.go.expand && \
		mv types.go.expand types.go && \
		rm -rf _obj && \
		cd -

clean:
	rm -f pkg/rocm_smi/cgo_helpers.go pkg/rocm_smi/cgo_helpers.h pkg/rocm_smi/cgo_helpers.c
	rm -f pkg/rocm_smi/const.go pkg/rocm_smi/doc.go pkg/rocm_smi/types.go
	rm -f pkg/rocm_smi/rocm_smi.go

test:
	cd pkg/rocm_smi && go build && go test $(TEST_VERBOSE)

.PHONY: fmt
fmt:
	cd pkg/rocm_smi && go fmt


# Examine Go source code and reports suspicious constructs
.PHONY: vet
vet:
	cd pkg/rocm_smi && go vet ./...
