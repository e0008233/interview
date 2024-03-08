https://www.youtube.com/watch?v=NHqdG-aZxOk

1. clarifications
1B users in 150 countries
1B videos viewed per day
10B videos uploaded per year
success metrics: time in app 1 hour/day

2. non-functional assumptions
10-second videos on average -> 1MB per video
10B/year * 1MB = 10 PB videos per year -> *10 for replication -> 100PB -> blob storage1
10B/year * 1KB meta data = 10TB of video metadata -> nosql storage
10B/year/365 -> 30M/day -> 1000/second -> 1GB/ second ingress

3. High level design
Tiktok app -> LB -> App service 
                 -> upload service  -> video blob storage (AWS S3) - tiered storage, replication
                 -> for you service
              -> CDN 

4. drill down
user/video meta data
Use of DB
for you feed
5. bottlenecks
Vide playing in a real time
6. enhancement
offline some for you generator services to mobile devices