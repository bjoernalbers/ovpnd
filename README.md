# ovpnd - Distribute .ovpn files to OpenVPN Connect clients by HTTPS

## Motivation

[OpenVPN Connect](https://openvpn.net/client/) has the ability to import
connection profiles (.ovpn files) from an 
[OpenVPN Access Server](https://openvpn.net/access-server/) by HTTPS.
Unfortunately the free OpenVPN server does not have this feature.
But this project provides such a
[webservice](https://openvpn.net/images/pdf/REST_API.pdf), so that your users
can download their connection profiles by themself.
