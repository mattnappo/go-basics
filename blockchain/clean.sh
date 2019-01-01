#!/bin/bash
find ./ -type d -name data -exec rm -rf {} \;
rmdir common/test/test
rmdir common/test/