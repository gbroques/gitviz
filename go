#!/bin/bash

./node_modules/supervisor/lib/cli-wrapper.js -p 100 -e "js|py" -n error -- app.js
