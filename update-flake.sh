#!/usr/bin/env bash
go mod vendor
nix hash path vendor > go.mod.sri
rm -Rf vendor
