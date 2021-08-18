# clist
CLI Support Tickets

**Overview**
Clist is a CLI based support ticketing system built in Go.

Clist consists of a client server architecture. 

**Set Up**

The Clist server will require two tables on a PostgreSQL Database. One for storing user information and one to store tickets. The database parameters are stored in the clist-server.json configuration file. The config file also allows you to set the port for the server to run on. Once the server is configured, all you will need to do is start the server executable.

Setting up the Client is also pretty straight forward. The client will look for a clist-client.json config file, which stores the information to connect to the server as well as the login information. Once the client is configured, you can run 'clist register' in your terminal. This will set up your login information on the server side if it does not yet exist in the database. It will also store your username and a hash of your password in the client configuration file in order to autheticate with the server.
