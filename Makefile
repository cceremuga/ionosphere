name = ionosphere
config = config/config.yml

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

define build_arch
	$(1) go build -o bin/$(2)/ionosphere main.go
	mkdir bin/$(2)/config
	mkdir bin/$(2)/logs
	touch bin/$(2)/logs/PLACEHOLDER
	cp $(config) bin/$(2)/$(config)
	cp README.md bin/$(2)/README.md
	cp LICENSE bin/$(2)/LICENSE.md
endef

define tarball_distro
	tar -zcvf $(1).tar.gz $(1)/
endef

define zip_distro
	zip -r $(1).zip $(1)/
endef

test:
	go test ./... -v

build:
	GOOS=linux GOARCH=amd64 go build -o bin/$(linuxamd64)/ionosphere main.go

run:
	go run main.go

clean:
	rm -rf ./bin

macosarm64 = $(name)-macos-arm64
macosamd64 = $(name)-macos-amd64
linuxamd64 = $(name)-linux-amd64
freebsdamd64 = $(name)-freebsd-amd64
raspberrypi = $(name)-raspberry-pi
windowsamd64 = $(name)-windows-amd64

.ONESHELL:

compile:
	rm -rf ./bin
	$(call build_arch,GOOS=darwin GOARCH=amd64,$(macosamd64))
	$(call build_arch,GOOS=darwin GOARCH=arm64,$(macosarm64))
	$(call build_arch,GOOS=linux GOARCH=amd64,$(linuxamd64))
	$(call build_arch,GOOS=linux GOARCH=arm GOARM=5,$(raspberrypi))
	$(call build_arch,GOOS=freebsd GOARCH=amd64,$(freebsdamd64))
	$(call build_arch,GOOS=windows GOARCH=amd64,$(windowsamd64))
	cd bin/
	$(call tarball_distro,$(macosamd64))
	$(call tarball_distro,$(macosarm64))
	$(call tarball_distro,$(linuxamd64))
	$(call tarball_distro,$(raspberrypi))
	$(call tarball_distro,$(freebsdamd64))
	$(call zip_distro,$(windowsamd64))
