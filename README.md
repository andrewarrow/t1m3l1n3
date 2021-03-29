# CLT
command line timelines

# Vision
Picture a 64x64 chessboard and place a pawn in every square.
By default in these universes, you follow 63 other people and they
each follow you.

Now find your row and remove a pawn from column 7 and 14.
You are not following user 7 and user 14 anymore.

When they go to post something, a pointer to their latest message
will be placed in each user's inbox; if the user follows them.

This is what Following []unit64 is for and it means:

```
1111111111111111111111111111111111111111111111111111111111111111
1111111011111101111111111111111111111111111111111111111111111111
1111111111111111111111111111111111111111111111111111111111111111
```

A giant, efficient grid to run a small twitter like message
board for just 64 people.

When this universe first boots up, all usernames are available
and it's first come first served to grab a two to twenty-two length
username. You send in your desired name, and a public key, and
forever more this node will give you and your private key the right
to post as you.

But this is just one node, and this code is open source. Anyone
can start a new node where the same username you just picked can
be given to someone else.

Universes are chained together. And just like the blockchain, you 
want to attach your node to a good chain that's going to be accepted
in lots of other places. The world can only vet and maintain so many
different chains of universes. If your chain can't govern itself
you might break apart and join other lead nodes.

It's all about a decentralized twitter/timeline system that one
company does not control. All it takes is $5 a month to rent a
digitalocean server and place your one node in the system.

As more and more people start to adopt CLT, it will be harder to get
first name usernames and still have your chain taken seriously by
the N major other chains. And for free speach and who gets to set
the rules? It's the person with the $5 credit card bill's decision.
There will be certain community standards some chains allow.




# About
anyone is free to take the server binary and run it on a machine with
a nice amount of RAM and publish the IP or domain name of your
public CLT server to the main list.

# Download the clt client
When you run the client it will connect to the main server to get a copy
of the main list. Then it will pick a random server from that list to query
for your timeline ls command.

# Sample Use

