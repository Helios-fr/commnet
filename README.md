# Commnet: Lan Communication Network
This tool is designed to run an encrypted lan messaging server/client.

## Communication standard
### General Overview
All data transfered using this tool is provided in the below format:
```
username, encrypted_message
```
Where `username` is the username of the sender, and `encrypted_message` is the message encrypted using the AES encryption algorithm with the key provided by the receiver.

### Ports
The server recieves data on port `6875`, and the client provides their username and public key on port `6876`. When the program is run, the client will attempt to reconnect to all peers known to it. If no peers are known, the client will begin to scan the network for other clients.

### The message
Once decrypted by the Encryption module, the message is catagoized into two types:
 - `message`: A message to be displayed to the user.
 - `data`: Data coresponding to a clients request, to be stored in the database. Examples include:
    - `user_info`: Information about a user containing their username, public key, and local ip address as stored by another peer. In this case the username and public key are stored in the users database, and the local ip address is stored in the peers list.
    - `peers`: A list of all peers known to another peer. This is used to update the peers list.
    - `users`: A list of all users known to another peer. This is used to update the users list in the database. 


