# GOLESSON/lesson3

We will change dotenv config, connect to mongo db with pointer client, repository pattern created for db crud, bootstrap added for genereted random data with channel (optional: i would to know what is the channel), for practice that two types (channel and classic iteration)

## Note
- Channel defines with make function for allocation
- Pointer reduce to memory footprint, so you should use with struct type as specially function side and return values
- You preffer to use struct functions for manipulation your sturct data, find change sort or another think

While tasks generated before, you dont create new data
```
    [config] init()
    2022/07/01 12:13:52 [config] Initialized .env file
    2022/07/01 12:13:52 [db/db.go] init()
    2022/07/01 12:13:52 [config] Get() successfully
    2022/07/01 12:13:52 [models/task.go] init()
    2022/07/01 12:13:52 Getting collection: task
    2022/07/01 12:13:52 [config] Get() successfully
    2022/07/01 12:13:52 finded to many object in collection len: 2 . Not generate new data
    2022/07/01 12:13:52 silly <nil>
```

Generate new data and insert to mongo
```
    [config] init()
    2022/07/01 12:15:45 [config] Initialized .env file
    2022/07/01 12:15:45 [db/db.go] init()
    2022/07/01 12:15:45 [config] Get() successfully
    2022/07/01 12:15:45 [models/task.go] init()
    2022/07/01 12:15:45 Getting collection: task
    2022/07/01 12:15:45 [config] Get() successfully
    Channel Open  &{ObjectID("000000000000000000000000") 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC Task 0 XVlBzgbaiC MRAjWwhTHc true} true
    Channel Open  &{ObjectID("000000000000000000000000") 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC Task 1 cuAxhxKQFD aFpLSjFbcX false} true
    Channel Close  false
    2022/07/01 12:15:45 silly <nil>
```

## next lesson
- we will try to reduce init func usage
- we will try to add api
- we will try to add cli module for jobs
- we will try to add new app
- we will try to add async and gorotuine


