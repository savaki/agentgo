go-agent
========

Go 1.1.1 Agent for ThoughtWorks Go 13.2

go-agent creates Amazon AMIs with the following packages installed:

* go 1.1.1 (the language from Google) 
* go-agent 13.2
* Oracle Java 6

## Dependencies

You'll need to install the following to run this:

* [Vagrant](http://www.vagrantup.com/) - allows you to test configuration location
* [Packer.io](http://www.packer.io/) - creates Amazon EC2 AMIs

## Environment Variables

You'll need to set the following environment variables to run this:

* AWS_ACCESS_KEY_ID 
* AWS_SECRET_ACCESS_KEY

## Runtime Configuration

When you deploy the AMI, you'll need to specify user data that tells the go-agent where the go-server is.  

Here's an example:

```
{
	"go-server": "127.0.0.1",
	"go-port":   "8080"
}
```

**Note:** all values are strings





