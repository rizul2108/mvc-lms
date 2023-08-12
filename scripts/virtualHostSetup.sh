#!/bin/bash

read -p "Enter your domain name: " domain_name
read -p "Enter your email: " server_admin_email

sudo apt update
sudo apt install apache2

sudo a2enmod proxy proxy_http

sudo tee /etc/apache2/sites-available/$domain_name.conf > /dev/null <<EOF
<VirtualHost *:80>
    ServerName $domain_name
    ServerAdmin $server_admin_email
    ProxyPreserveHost On
    ProxyPass / http://127.0.0.1:9000/
    ProxyPassReverse / http://127.0.0.1:9000/
    TransferLog /var/log/apache2/${domain_name}_access.log
    ErrorLog /var/log/apache2/${domain_name}_error.log
</VirtualHost>
EOF

sudo a2ensite $domain_name.conf

sudo a2dissite 000-default.conf

sudo apache2ctl configtest

echo "127.0.0.1  $domain_name" | sudo tee -a /etc/hosts

sudo systemctl restart apache2

open http://$domain_name