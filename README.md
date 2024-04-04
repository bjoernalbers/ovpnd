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

Build or download the binary and run it on your server:

    ./ovpnd -cert yourcert.crt -key yourkey.key openvpn-profiles/

This will start the webservice on port `443/tcp`.
Then point your users to the server URL.

For authentication the username is the basename of the .ovpn file (without the
suffix) and the password is the content of the corresponding .txt file.
