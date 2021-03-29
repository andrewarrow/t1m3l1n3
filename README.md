# CLT
command line timelines

# Vision
Picture a 64x64 chessboard and place a pawn in every square.
By default in a universe, you follow 63 other people and they
each follow you.

Now find your row and remove a pawn from column 7 and 14.
You are not following user 7 and user 14 anymore.

When they go to post something a pointer to their latest message
will be placed in each user's inbox, if they user follows them.

This is what Following []unit64 is for and it means:

1111111111111111111111111111111111111111111111111111111111111111
1111111011111101111111111111111111111111111111111111111111111111
1111111111111111111111111111111111111111111111111111111111111111

A giant, efficient grid to run a small twitter like message
board for just 64 people.

When this universe first boots up, all usernames are available
and it's first come first served to grab a two to twenty-two length
username. You send in your desired name, and a public key, and
forever more this node will give you and your private key the right
to post as you.



# About
anyone is free to take the server binary and run it on a machine with
a nice amount of RAM and publish the IP or domain name of your
public CLT server to the main list.

# Download the clt client
When you run the client it will connect to the main server to get a copy
of the main list. Then it will pick a random server from that list to query
for your timeline ls command.

# Incoming sync messages
The server listens for POST NEW TIMELINE messages from individual users as
well as large blocks of timelines from other servers from that main list.

It doesn't matter which server you connect you, eventually that server should
have a copy of 100% of the messages posted to any server in the main list.

# Outgoing sync messages
Every message the server gets from an individual plus the messages it gets
from it's incoming peer it forwards on to its outgoing peer.

So if the main list only has three items like:

1. clt1.com
2. clt2.com
3. clt3.com

When clt2 boots up it gets incoming peer set to clt1 and outgoing peer set to clt3.

Then if two new servers join the main list:

1. clt1.com
2. clt2.com
3. clt3.com
4. clt4.com
5. clt5.com

clt3 gets message it's new outgoing server is clt4. And the loop continues.
clt5 outgoing is clt1.
