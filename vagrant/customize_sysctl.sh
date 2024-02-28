#!/bin/bash

# Özel sysctl ayarlarını içeren dosya
SYSCTL_CONF="/etc/sysctl.d/99-custom.conf"

# Ayarları dosyaya ekle
cat <<EOF > $SYSCTL_CONF
net.netfilter.nf_conntrack_max = 1048576
net.ipv4.tcp_window_scaling = 1
net.ipv4.tcp_syncookies = 1
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216
net.ipv4.tcp_rmem = 4096 87380 16777216
net.ipv4.tcp_wmem = 4096 65536 16777216
net.ipv4.ip_local_port_range = 1024 65535
net.ipv4.tcp_keepalive_time = 300
net.ipv4.tcp_keepalive_intvl = 60
net.ipv4.tcp_keepalive_probes = 9
net.ipv4.tcp_max_syn_backlog = 16384
net.core.somaxconn = 16384
net.core.netdev_max_backlog = 50000
net.ipv4.tcp_tw_reuse = 0
net.ipv4.tcp_fin_timeout = 60
fs.file-max = 1048576000
fs.nr_open = 1048576000
fs.inotify.max_queued_events = 16384
fs.inotify.max_user_instances = 8192
fs.inotify.max_user_watches = 262144
user.max_inotify_instances = 8192
user.max_inotify_watches = 262144
net.ipv4.ip_forward = 1
EOF

# Sistemi yeni ayarlarla yapılandır
sysctl -p $SYSCTL_CONF
