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

