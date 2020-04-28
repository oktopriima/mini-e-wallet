#!/usr/bin/env bash
echo "::: starting :::"

cd domain/core

echo "::: starting generate :::"
gen --connstr "root:empalgoreng@tcp(127.0.0.1:3306)/mkproject?&parseTime=True" --database mkproject --json

echo "::: finish :::"
