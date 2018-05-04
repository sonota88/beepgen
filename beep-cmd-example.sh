#!/bin/bash

set -o pipefail

sr=44100

./beepgen sr=$sr "$@" | aplay --rate=$sr --format=U8 --quiet
