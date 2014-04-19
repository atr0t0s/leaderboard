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
