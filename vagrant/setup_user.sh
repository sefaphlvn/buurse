#!/bin/bash

# Kullanıcı adı
user_name="sefa"

# Kullanıcı oluştur (eğer mevcut değilse)
id -u $user_name &>/dev/null || useradd -m -s /bin/bash $user_name

#useradd -m -s /bin/bash $user_name

# Kullanıcıya şifre ata
echo "$user_name:12345" | chpasswd

# Kullanıcıya sudo yetkileri ver
usermod -aG root $user_name

# .ssh dizinini oluştur (eğer yoksa)
mkdir -p /home/$user_name/.ssh

# Environment variable'dan alınan SSH anahtarını authorized_keys dosyasına ekle
echo "$PUB_KEY" >> /home/$user_name/.ssh/authorized_keys

# Dosya izinlerini ayarla
chown -R $user_name:$user_name /home/$user_name/.ssh
chown -R $user_name:$user_name /vagrant/suubar
chmod 700 /home/$user_name/.ssh
chmod 600 /home/$user_name/.ssh/authorized_keys

echo "$user_name ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$user_name
chmod 0440 /etc/sudoers.d/$user_name
