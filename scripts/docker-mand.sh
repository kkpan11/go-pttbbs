#!/bin/bash
# Dockerfile: run mand.

nohup sudo -iu bbs /home/bbs/bin/mand --address 0.0.0.0:5151 &>/dev/null &
