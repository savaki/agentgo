#!/bin/bash

#---------------------------------------------------------------------------
# Install Java
#---------------------------------------------------------------------------

# automatically accept the oracle license
echo debconf shared/accepted-oracle-license-v1-1 select true | sudo debconf-set-selections
echo debconf shared/accepted-oracle-license-v1-1 seen true | sudo debconf-set-selections

cp /etc/apt/sources.list sources.list.$$
cat >> sources.list.$$ <<EOF
deb http://ppa.launchpad.net/webupd8team/java/ubuntu precise main
deb-src http://ppa.launchpad.net/webupd8team/java/ubuntu precise main
EOF
sudo mv sources.list.$$ /etc/apt/sources.list
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys EEA14886
sudo apt-get update
sudo apt-get install -y oracle-java6-installer


#---------------------------------------------------------------------------
# Install Go Agent
#---------------------------------------------------------------------------

# check to see if the go agent is installed already or not
if [ ! -d /usr/share/go-agent ] ; then 
  GO_AGENT=go-agent-13.2.0-17155.deb
  wget --quiet http://d3a9nbnkw85yq1.cloudfront.net/ubuntu/precise/${GO_AGENT}
  sudo dpkg -i ${GO_AGENT}
  rm -f ${GO_AGENT}

  # add GOPATH to the go-agent profile
  cp /etc/default/profile profile.$$
  cat >> profile.$$ <<EOF

# Support for Golang 
export GOPATH=${HOME}/gocode
export PATH=${PATH}:/usr/local/go/bin:${GOPATH}/bin
EOF

  sudo cp profile.$$ /etc/default/profile
fi  


