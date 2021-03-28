



You are in rooms with 64 people.

By default you follow all, then you slowly unfollow people.

Your entire room is in a room with either 192 or 384 other people.

And that room is in a room with 576 people.

All the N number of 576 universies gossip with each other about
events going on in their worlds.


  576              576
384 192          384 192
 64  64           64  64


1111111111111111111111111111111111111111111111111111111111111111
1111111111111111111111111111111111111111111111111111111111111111
1111111111011111111111111111111011111111111111111111111111111111
1111111111111111111111111111111111111111111111111111111111111111
1111111111111111111111111111111111111111111111110111111111111111

andrew: hello

andrew is user 0, find his row and loop thru inbox delivery for
all users but skip a user if the row says so

bob is user 63, he says hi
find his row, loop, skip if row says so

inboxes:

from, 0-63
to, 0-63
msg, pointer

 














http timelines.org -> DNS round robbin to 1 node:

1. 











s1 -> s2 -> s3

s1


s1: s2
s2: s1

s1: s2
s2: s3
s3: s1


s1 -> s2 
s2 -> s1        // s2 told s1, s2 got reply with s1


s1 -> s2        
s2 -> s3       // s3 told s1, s3 got reply with s1
s3 -> s1       // s1 tells s2 your new out is s3

s1 -> s2        
s2 -> s3       
s3 -> s4       // s1 tells s3 your new out is s4
s4 -> s1      

