#!/usr/bin/env bash
echo "::: starting :::"

cd domain/core

echo "::: starting generate :::"
gen --connstr "root:empalgoreng@tcp(127.0.0.1:3306)/e-wallet?&parseTime=True" --database e-wallet --json

echo "::: finish :::"
