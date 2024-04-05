# ovpnd - Simple .ovpn Files Webserver for OpenVPN Connect

The official OpenVPN client [OpenVPN Connect](https://openvpn.net/client/) also
can fetch client configuration files (.ovpn files) by HTTPS, usually from an
[OpenVPN Access Server](https://openvpn.net/access-server/).
`ovpnd` serves those .ovpn files files as well by implementing the official
[REST API](https://openvpn.net/images/pdf/REST_API.pdf).

## Requirements

You need the following:

- directory with .ovpn files a.k.a. connection profiles in
  [unified format](https://openvpn.net/faq/i-am-having-trouble-importing-my-ovpn-file/)
- for each .ovpn file a corresponding .txt file in the same directory that
  includes an unecrypted password (required for user authentication)
- TLS certificate and key

## Usage

`ovpnd` is distributed as
[docker image](https://hub.docker.com/r/bjoernalbers/ovpnd) for easy deployment.

Getting help:

    $ docker run --rm bjoernalbers/ovpnd -h

Running `ovpnd`:

    $ ls tls
    cert.crt        cert.key
    $ ls profiles
    johndoe.ovpn    johndoe.txt
    $ cat profiles/johndoe.txt
    secret
    $ docker run --rm -p 443:443 -v $(pwd)/tls:/tls -v $(pwd)/profiles:/profiles bjoernalbers/ovpnd -cert /tls/cert.crt -key /tls/cert.key /profiles

Testing:

    $ curl https://openvpn.example.com/rest/GetUserlogin
    <?xml version="1.0" encoding="UTF-8"?>
    <Error>
    <Type>Authorization Required</Type>
    <Synopsis>REST method failed</Synopsis>
    <Message>Invalid username or password</Message>
    </Error>

    $ curl -u johndoe:secret https://openvpn.example.com/rest/GetUserlogin
    content of profile

Running `ovpnd` without TLS if a reverse-proxy already takes care of TLS:

    $ docker run --rm -p 80:80  -v $(pwd)/profiles:/profiles bjoernalbers/ovpnd -no-tls /profiles
