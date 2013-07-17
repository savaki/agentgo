#!/bin/bash

#---------------------------------------------------------------------------
# Install MySQL 5.5
#---------------------------------------------------------------------------

echo mysql-server-5.5 mysql-server/root_password password password | sudo debconf-set-selections
echo mysql-server-5.5 mysql-server/root_password_again password password | sudo debconf-set-selections
echo mysql-server-5.5 mysql-server/root_password seen true | sudo debconf-set-selections
echo mysql-server-5.5 mysql-server/root_password_again seen true | sudo debconf-set-selections

echo installing mysql 5.5
sudo apt-get -y install mysql-server-5.5
