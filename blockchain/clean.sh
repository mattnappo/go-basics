#!/bin/bash
find ./ -type d -name data -exec rm -rf {} \;
rmdir -p common/test/test