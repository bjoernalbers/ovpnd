# ovpnd - Distribute .ovpn files to OpenVPN Connect clients by HTTPS

## Motivation

[OpenVPN Connect](https://openvpn.net/client/) has the ability to import
connection profiles (.ovpn files) from an 
[OpenVPN Access Server](https://openvpn.net/access-server/) by HTTPS.
Unfortunately the free OpenVPN server does not have this feature.
But this project provides such a
[webservice](https://openvpn.net/images/pdf/REST_API.pdf), so that your users
can download their connection profiles by themself.

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
