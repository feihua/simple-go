#!/bin/bash

chmod +x simple-go
pkill -f simple-go
nohup ./simple-go > /dev/null 2>&1 &