#!/bin/bash

wget https://go.dev/dl/go1.22.0.linux-arm64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.0.linux-arm64.tar.gz
echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile.d/go.sh


