#!/bin/bash

set -u 
set -e

cat > agent.json <<EOF
{
  "builders":[{
    "type": "amazon-ebs",
    "access_key": "${AWS_ACCESS_KEY_ID}",
    "secret_key": "${AWS_SECRET_ACCESS_KEY}",
    "region": "us-west-2",
    "source_ami": "ami-70f96e40",
    "instance_type": "m1.medium",
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
