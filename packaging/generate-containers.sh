#!/bin/sh

WORKDIR=$1
cd $WORKDIR

## Cleaning up any previous packaging attempts
rm -rf workdir
rm -rf rootfs

mkdir -p workdir
cp binaries/turtleshell-linux-amd64 workdir/tsh
chmod +x workdir/tsh

## Setting up our scratch image root filesystem
mkdir -p rootfs/home
mkdir -p rootfs/root
mkdir -p rootfs/var
mkdir -p rootfs/etc
mkdir -p rootfs/bin
tar -cjvf rootfs.tar.xz rootfs
mv rootfs.tar.xz workdir/

## Build our containers using Docker
cd ../

podman build -t jgrancell/tsh:alpine -f packaging/containerfiles/Containerfile-alpine .
podman build -t jgrancell/tsh:scratch -f packaging/containerfiles/Containerfile-scratch .