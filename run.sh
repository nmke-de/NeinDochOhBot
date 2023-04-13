#!/bin/sh
cd $(dirname $0)
cd $(pwd -P)
NEIN_DOCH_OH_TOKEN=$(cat token) ./NeinDochOhBot
