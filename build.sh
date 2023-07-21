#!/bin/bash -ex

go build -o app

set +e
rm -r build
set -e

mkdir -p build/GoTunnel.app/Contents/{MacOS,Resources}
cp app build/GoTunnel.app/Contents/MacOS/
cp Info.plist build/GoTunnel.app/Contents/
cp Icon.icns build/GoTunnel.app/Contents/Resources/
