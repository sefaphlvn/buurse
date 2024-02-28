#!/bin/bash

NOFILE_CONF="/etc/security/limits.d/limits-nofile.conf"

# Ayarları dosyaya ekle
cat <<EOF > $NOFILE_CONF
* hard nofile 500000
* soft nofile 500000
EOF
