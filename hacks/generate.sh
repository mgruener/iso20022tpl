#!/bin/sh

echo "Running go generate..."
PKGS=$(go list ./...)
go generate ${PKGS}