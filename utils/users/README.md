## Users Module
This modile handles all elements of user data, including the creation and replication of users from peers, the storage and retrival of user data, and the management of user data.

Functions included in this file:
- Create user (CreateUser: username*, publicKey*, privateKey) --> bool
- Get user (GetUser: username*) --> publicKey, privateKey
- Validate user (ValidateUser: username*, publicKey) --> bool
- Remove user (RemoveUser: username*) --> bool
- Reset DB (resetDB: ) --> bool

### CreateUser --> bool
This function appends a new user to the user data csv file, storing the username, public key, and private key if it is known.

### GetUser --> publicKey, privateKey
This function retrieves the public key for the given username, and the private key if it is known.

### ValidateUser --> bool
This function validates the user by checking the public key against the stored public key for the given username.

### RemoveUser --> bool
This function removes the user from the user data csv file.

### ResetDB --> bool
This function resets the user data csv file to the default state, removing all users.