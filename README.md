# t1m3l1n3
command line timeline server and client, decentralized micro-blogging
follow/unfollow network, in golang.

![logo](https://repository-images.githubusercontent.com/352111156/bb7aef80-938d-11eb-8f1d-64f98d44309e)

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
can start a new node with the same username you just picked can
be given to someone else.

Universes are chained together. And just like the blockchain, you 
want to attach your node to a good chain that's going to be accepted
in lots of other places. The world can only vet and maintain so many
different chains of universes. If your chain can't govern itself
you might break apart and join other lead nodes.

It's all about a decentralized twitter/timeline system that one
company does not control. All it takes is $5 a month to rent a
digitalocean server and place your one node in the system.

As more and more people start to adopt TL, it will be harder to get
first name usernames and still have your chain taken seriously by
the N major other chains. And for free speach and who gets to set
the rules? It's the person with the $5 credit card bill's decision.
There will be certain community standards some chains allow.

# So anyone can take my username for themselves in a new node?

Yes, but they would be all alone, claiming to be you in a node that no
one else agrees is valid.

# What if a bunch of people agree it's valid?

Then yes, there is competition for your username. The same way the username
"goodname" on twitter is in competition for "goodname" on another social
media site.

# Which one wins?

They both co-exist but you vote with your attention and your wallet. You
can choose to accept the TL chain where most of your friends are. You can
chose to spend your $5 a month node on the chain you like. You can even get
two nodes and spend $10 a month or as much as you want.

# So how many chains are there?

A lot. Most of them small. Over time only a few will dominate the scene as
"the big chains" and you can probably guess there will be one that leans left
politically and one that leans right. Is there a center one? Sure. Is there more
than one center one? Sure. When you see something that offends you, define your
own guidelines and break off into a new chain.

# What about illegal content or hate speach?

Read your contract with digitalocean or whoever you are renting your $5 a month
server from. I woudln't want illegal content on my server! When you are in a major
chain, certain messages are flagged for violations and removed. Can everyone agree
on which messages MUST be removed? No. Can a bunch of people agree on some basic
guidelines? Yes.

# Can you detach your name/public key from one node and move to another?

Yes, when someone installs the TL client for the first time, they are given 
the default chain. But you can configure your client to speak to another chain.
Think of some very popular blogger like Matt Taibbi. You know where the REAL
Matt Taibbi is blogging from right? At the time of writing this (March 2021) 
it's substack. But if he were to suddenly abandon that account and appear 
somewhere else, you'd follow that new Matt Taibbi account on some other chain.

# Where does my $5/mo go, to the universe or node creator?

It goes to digitalocean or aws or any cloud provider you want to use to rent a
server that has CPU, RAM, Internet Access, and hard drive space. With one node
you are the sysop for 64 people. You are paying the bill for these 64 people to
talk and follow each other you are part of this decentralized social media
system. The whole point is if many individuals pay the hosting costs, there is 
no one big company paying ALL the hosting and therefore making ALL the rules.

But 64 people isn't very many people. So you need to band together with another 64
and another 64 and on and on.

# About
anyone is free to take the server binary and run it on a machine with
a nice amount of RAM and publish the IP or domain name of your
public TL server to the main list.

# Download the client
When you run the client it will connect to the main server to get a copy
of the main list. Then it will pick a random server from that list to query
for your timeline ls command.

# Sample Use

```
[~/TL/client] $ ./client 

  client ls        # List recent timelines
  client profile   # List recent timelines
  client post      # Post new timeline with --text=hi
  client auth      # Set your username --name=
  client toggle    # Toggle follow --name=
  client idplease  # tell me what server/node i'm connected to
  client taken     # taken usernames

[~/TL/client] $ ./client ls
Inbox
[~/TL/client] $ ./client post --text=hi
[~/TL/client] $ ./client ls
Inbox
01.               andrew   less than a minute hi
[~/TL/client] $ ./client auth --name=bob
Ok you are now: bob
[~/TL/client] $ ./client post --text=hi
[~/TL/client] $ ./client ls
Inbox
01.                  bob   less than a minute hi
02.               andrew   less than a minute hi
[~/TL/client] $ ./client profile
Profile
01.                  bob   less than a minute hi
[~/TL/client] $ ./client profile --name=andrew
Profile
01.               andrew             1 minute hi
[~/TL/client] $ ./client auth --name=andrew
This username already taken!
[~/TL/client] $ ./client auth --name=andrew2
Ok you are now: andrew2
[~/TL/client] $ ./client idplease
localhost:8080

[~/TL/client] $ ./client taken
{"users":[{"ts":1617232944,"username":"andrew"},{"ts":1617241169,"username":"andrew2"},{"ts":1617238813,"username":"bob"}]}
[~/TL/client] $ ./client auth --name=sue
Ok you are now: sue
[~/TL/client] $ ./client taken
{"users":[{"ts":1617242196,"username":"sue"},{"ts":1617232944,"username":"andrew"},{"ts":1617241169,"username":"andrew2"},{"ts":1617238813,"username":"bob"}]}

```
