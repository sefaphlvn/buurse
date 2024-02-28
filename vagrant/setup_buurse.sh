#!/bin/bash

if [ ! -d "/vagrant/buurse" ]; then
    git clone https://github.com/sefaphlvn/buurse.git /vagrant/buurse
else
    echo "Repo already cloned."
fi

cd /vagrant/buurse
git config --local user.email "sefa.pehlivan@gmail.com"
git config --local user.name "Sefa Pehlivan"
git pull

