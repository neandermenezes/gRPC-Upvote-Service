# gRPC-Upvote-Service
Upvote service using gRPC and golang for Klever

## protobuf messsages:
```
- Post: id, author_name, title, content, like_count
- PostId: id,
```

## rpc methods:
```
- rpc CreatePost (Post) returns (PostId);
- rpc ReadPost (PostId) returns (Post);
- rpc UpdatePost (Post) returns (google.protobuf.Empty);
- rpc DeletePost (PostId) returns (google.protobuf.Empty);
- rpc ListPosts (google.protobuf.Empty) returns (stream Post);
- rpc UpvotePost (PostId) returns (google.protobuf.Empty); 
```

## How to use this repo
```
clone the repository
$ make build
$ ./bin/server
$ ./bin/client
```
