#!/bin/bash

UPSTREAM_TAG=0.0.20210212

git remote add upstream git@github.com:WireGuard/wireguard-go.git
git fetch upstream
git checkout -b upstream ${UPSTREAM_TAG}
git filter-branch --prune-empty -f --subdirectory-filter tun/wintun
git checkout master
git rebase upstream
git branch -D upstream
git remote remove upstream
