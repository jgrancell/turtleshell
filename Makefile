TEST?=$$(go list ./... | grep -v vendor)
WORKDIR=$$(pwd)
BINARY=$$(pwd | xargs basename)
VERSION=$$(grep version main.go | head -n1 | cut -d\" -f2)
GOBIN=${GOPATH}/bin

default: build

build:
	go build -o ${BINARY}
	chmod +x ${BINARY}

install: build
	mkdir -p ${GOBIN}
	mv ${BINARY} ${GOPATH}/bin/${BINARY}

binaries: build
	rm -rf packaging/binaries
	mkdir -p packaging/binaries
	bash packaging/generate-binaries.sh ${BINARY} ${WORKDIR}

package: binaries
	bash packaging/generate-containers.sh ${WORKDIR}/packaging

test:
	echo $(TEST) | xargs -t -n4 go test -timeout=30s -parallel=4 -coverprofile=cover.out
	go tool cover -html=cover.out

clean:
	rm -rf packaging/binaries
	rm -rf packaging/containers/rootfs
	rm -rf packaging/containers/workdir
	rm -rf packaging/workdir