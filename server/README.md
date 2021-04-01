./server start 


In a world with just two complete Universes:

Ua = 64x64 follow matrix
Ub = 64x64 follow matrix

In order for a user in Ub to follow a user in Ua, Ua and Ub have to
link and agree on the same username space.

When linking, either Ua adds Ub to it's UpPeers, and Ub adds Ua to its DownPeers...
or the other way around.

The UpPeers break ties for usernames.

If your universe is new and you want to link to another more establish one, you
are going to have to tell a few of your 64 users that their usernames have to change.





You are in rooms with 64 people.

By default you follow all, then you slowly unfollow people.

Your entire room is in a room with either 192 or 384 other people.

And that room is in a room with 576 people.

All the N number of 576 universies gossip with each other about
events going on in their worlds.


  576              576
384 192          384 192
 64  64           64  64


andrew is user 0, find his row and loop thru inbox delivery for
all users but skip a user if the row says so

bob is user 63, he says hi
find his row, loop, skip if row says so

inboxes:

from, 0-63
to, 0-63
msg, pointer

