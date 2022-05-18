
CFORGO = ${HOME}/go/bin/c-for-go
CGO = ${HOME}/.modules/go-1.16.4/pkg/tool/linux_amd64/cgo

all: generate

generate:
	$(CFORGO) -ccincl --ccdefs rocm_smi.yml
	sed -i -e s/'cStatus, cStatus_string'/'cStatus, \&cStatus_string'/g rocm_smi/rocm_smi.go
	cd rocm_smi && \
		$(CGO) -godefs types.go > types.go.expand && \
		mv types.go.expand types.go && \
		rm -rf _obj && \
		cd -

clean:
	rm -f rocm_smi/cgo_helpers.go rocm_smi/cgo_helpers.h rocm_smi/cgo_helpers.c
	rm -f rocm_smi/const.go rocm_smi/doc.go rocm_smi/types.go
	rm -f rocm_smi/rocm_smi.go

test:
	cd rocm_smi && go build && go test

.PHONY: fmt
fmt:
	cd rocm_smi && go fmt


# Examine Go source code and reports suspicious constructs
.PHONY: vet
vet:
	cd rocm_smi && go vet ./...
