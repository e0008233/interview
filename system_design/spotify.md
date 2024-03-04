https://www.youtube.com/watch?v=_K-eupuDVEc

1. what is spotify
- songs / music
- playlists
- users
- artists
- podcasts

2. use cases
- finding and playing the music

3. numbers
- users: a billion
- songs: 100 billion
- mp3 audio: 5 mb
- total audio 500 TB = 0.5PB, X3 for replication -> 1.5PB
- 100b per song meta data -> 100GB
- 1kb per user -> 1 TB
- how many song played

4. high level design
app -> LB -> server 1/2/3/...  -> DB

5. drill down (DB)
5.1 DB use case
a. Meta data (users, songs) -> structured, content may be updated frequently, some complex queries -> rds
b. Song audio data -> BLOB data (Binary Large Object data) -> immutable -> Amazon Simple Storage Service (Amazon S3) -> easily scalable

5.2 song db example:
- song_id
- song_url
- artist
- genre
- link to album cover
- mp3 link

6. use case
- finding a song
app -> lb -> server -> rds
elastic search to facilit ate the search
- playing a song
web socket connection for downloading the auto data chunk by chunk
app -> CDN -> aws s3 (server -> cdn)
multi layer of caching

7. refinement
- Geo location of DB
- rds -> elastic search, search request -> elastic search


