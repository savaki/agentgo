#!/bin/bash

set -u 
set -e

AMI="ami-70f96e40"
REGION="us-west-2"
INSTANCE_TYPE="m1.medium"

cat > packer.json <<EOF
{
  "builders":[{
    "type": "amazon-ebs",
    "access_key": "${AWS_ACCESS_KEY_ID}",
    "secret_key": "${AWS_SECRET_ACCESS_KEY}",
    "region": "${REGION}",
    "source_ami": "${AMI}",
    "instance_type": "${INSTANCE_TYPE}",
    "ssh_username": "ubuntu",
    "ami_name": "go-agent {{.CreateTime}}"
  }],

  "provisioners": [
    {
      "type": "shell",
      "scripts": [
        "install_go.sh",
        "install_go_agent.sh",
        "install_mysql.sh"
      ]
    }
  ]
}

EOF
