#!/bin/bash

if [ ! -d "/vagrant/suubar" ]; then
    git clone https://github.com/sefaphlvn/suubar.git /vagrant/suubar
else
    echo "Repo already cloned."
fi

cd /vagrant/suubar
git config --local user.email "sefa.pehlivan@gmail.com"
git config --local user.name "Sefa Pehlivan"
git pull

