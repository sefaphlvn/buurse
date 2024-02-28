#!/bin/bash

NF_CONF="/etc/modprobe.d/nf_conntrack.conf"

# Ayarları dosyaya ekle
cat <<EOF > $NF_CONF
options nf_conntrack expect_hashsize=262144 hashsize=262144
EOF

modprobe -r nf_conntrack
modprobe nf_conntrack