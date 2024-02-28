#!/bin/bash

if [ ! -d "/vagrant/buurse" ]; then
    git clone https://github.com/sefaphlvn/buurse.git /vagrant/buurse
else
    cd /vagrant/buurse
    git pull
    echo "Repo already cloned."
fi

chown sefa:sefa /vagrant/buurse