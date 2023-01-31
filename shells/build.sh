#!/bin/zsh

function build_forviolaterbinary () {
    go build -o main/binaries_for_violater/$extention ViolaterWarningStatement/main.go
}

function build_piracycheckbinary () {
    echo "not implement"
}

filename="linux-amd.out"

if [[ $2 = "amd" ]]; then
    GOARCH=amd64
    extention="amd"
elif [[ $2 = "arm" ]]; then
    GOARCH=arm64
    extention="arm"
else
    echo "GOARCH SET default(amd64)"
    GOARCH=amd64
    extention="amd"
fi


if [[ $1 = "linux" ]]; then
    GOOS=linux
    extention="linux-$extention.out"
elif [[ $1 = "mac" ]]; then
    GOOS=darwin
    extention="mac-$extention.out"
elif [[ $1 = "windows" ]]; then
    GOOS=windows
    extention="win-$extention.exe"
else
    echo "GOOS SET default(linux)"
    GOOS=linux
    extention="linux-$extention.out"
fi

build_forviolaterbinary
