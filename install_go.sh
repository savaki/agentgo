#!/bin/bash

set -u
set -e

#---------------------------------------------------------------------------
# Install Google Go
#---------------------------------------------------------------------------
if [ ! -d /usr/local/go ] ; then
  GO=go_1.1.1_amd64.deb

  echo installing go
  wget --quiet http://d3a9nbnkw85yq1.cloudfront.net/ubuntu/precise/${GO}
  sudo dpkg -i ${GO}
  rm -f ${GO}

# update the path and environment variables with go settings
  echo updating /etc/profile
  cp /etc/profile profile.$$
  cat >> profile.$$ <<EOF

# Support for Golang 
export GOPATH=\${HOME}/gocode
export PATH=\${PATH}:/usr/local/go/bin:\${GOPATH}/bin
EOF
  sudo mv profile.$$ /etc/profile
  mkdir ${HOME}/gocode
fi

exit 0
