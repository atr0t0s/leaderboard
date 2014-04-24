Leaderboard
-----------

Leaderboard is aiming to be a multiplayer game services REST API, written in Go. 
At first stage, the purpose is to provide handling of user registrations and authentication,
player data persistence over different games, game and player stat gathering
and presentation, as well as allowing to create custom achievements for each 
game.

-----------

When the first stage is completed we can move to adding support for matchmaking,
and other multiplayer game services such as in-game chat, web access to player 
accounts, forums, and in-game purchases. At an even later stage, the possibility
of cryptocurrency payments can come under consideration.

The project runs on the Revel framework. First make sure you have Go installed
and that your GOPATH is configured. Then install the revel command line:
- go get github.com/revel/cmd/revel

Then
- git clone, make some changes, and run with "revel run leaderboard"
Or
- go get github.com/violarisgeorge/leaderboard/, and run with "revel run github.com/violarisgeorge/leaderboard". 

The second method is preferred as it will automatically insert the source code for
leaderboard in your GOPATH/src folder. If you git clone make sure to copy to that
folder, or somewhere which is relevant based on your revel cmd configuration.

// TODO: Introduced models for db object creation to be clearer, however still
// need to add a db controller for unified mgo requests so the db object will not
// have to be reinstantiated in every function that needs to use the db
// (this project is currently set up to work only with MongoDB. SQL may be 
// introduced later on if needed - this makes the need for a unified db controller
// more obvious, as the DB driver can be changed based on project needs rather
// than forcing everyone to work with mgo)
