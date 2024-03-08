- twitter search https://www.youtube.com/watch?v=eUvwRAxmEgY
- twitter: https://www.youtube.com/watch?v=o5n85GRKuzk

1. Background
   - social media platform
   - one user can follow many other users, and can be followed by others
2. Functional requirement
    - follow others
    - create tweets
    - view feed
    - tweet includes images and videos
3. Numbers 
    - total: 500M users, active users: 200M
    - 100 tweets/user/day -> 20B tweet read/day -> 1KB text, 1MB each tweets, 20PB read/day
    - 50M tweet create/day -> 50GB data/day
    - pop star 100M
4. High level design
   - Client -> LB -> server -> cache -> user data DB (relationship), rds/graphDB
                            -> CDN -> object storage   

5. design details
   - DB, read/write replica, sharding  
   - For non-popular creators -> pub/sub -> update feed cache