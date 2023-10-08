#!/bin/bash

build_version='github.com/Orendev/gokeeper/internal/app/client/configs.BuildVersion'
build_date='github.com/Orendev/gokeeper/internal/app/client/config.BuildDate'
build_version_value='0.0.2'
build_date_value='2023-10-08'
os_all='linux windows darwin'
arch_all='amd64 arm64'
app_name='/cmd/client/dist/keeper'
app='/cmd/client/main.go'

for os in $os_all; do
    for arch in $arch_all; do
      set GOOS=$os
      set GOARCH=$arch
      go build -o `pwd`$app_name"_"$os"_"$arch -ldflags "-X "$build_version"="$build_version_value" -X "$build_date"="$build_version_value `pwd`$app && echo "Success build for arch "$arch" and os "$os || echo "failed build for arch "$arch" and os "$os
    done
done

