# Prerequisites

Since the focus of the session is to enable attendees to code, it would be helpful to have the environment ready before the session starts. This guide will help you to get the local environement setup. 

*It is assumed that you have __admin rights__ to install software and set/modify environment variables in your machine.*

# Installation

## Go
First dependency we have is to install and configure __Go__ itself. 
- Follow the official instructions from [here](https://golang.org/doc/install) to download and install Go.
- Ensure that relevant environment variables mentioned in the instructions are configured properly. 

## Proxy
If you are connected to your organization's network, which uses proxy, make sure to add following environment variables

- `HTTP_PROXY=http://<proxy-url>:<port>`
- or 
- `HTTPS_PROXY=https://<proxy-url>:<port>`
- *(Ensure to replace values for __proxy-url__ and __port__ with correct values)*

## Visual Studio Code
The next dependency we have is to setup a code editor. For that purpose, we'll be using Visual Studio Code.

- Download Visual Studio Code from [here](https://code.visualstudio.com/Download) and install the same.

In addition to the Visual Studio Code, we also need to configure Go Plugin. 

- Follow instructions from [here](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go) to install the plugin.

# Pre-read
Before the session starts, it would help you to have some background about Go. Here is a set of useful URLs you could use to read before the session.

- Go [FAQs](https://golang.org/doc/faq)
- What kind of solutions can be built using Go?
    - [Cloud Solutions](https://go.dev/solutions/cloud/)
    - [Web Development](https://go.dev/solutions/webdev/)
    - [Command Line Interfaces](https://go.dev/solutions/clis/)
    - [DevOps & Site Reliability Engineering](https://go.dev/solutions/devops/)
- [How to write Go code](https://golang.org/doc/code.html)
- [Go Documentation](https://golang.org/doc/)
- [How to effectively write Go](https://golang.org/doc/effective_go.html)

# Coding Practice

If you prefer to play around with Go without installation, visit [here](https://play.golang.org/)
