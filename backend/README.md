# Backend API
## Content
- [Dictionaries](#dictionaries)
  - [Instagram states](#instagram-states)
  - [Channel states](#channel-states)
  - [Channel actions](#channel-actions)
  - [Channel targets](#channel-targets)
- [Entities](#entities)
  - [Token](#token)
  - [User](#user)
  - [Instagram](#instagram)
  - [Channel](#channel)
- [Errors](#errors)
- [Methods](#methods)
  - [POST /api/v1/login](#post-apiv1login)
  - [POST /api/v1/signup](#post-apiv1signup)
  - [GET /api/v1/user/refresh_token](#get-apiv1userrefresh_token)
  - [GET /api/v1/user](#get-apiv1user)
  - [POST /api/v1/instagrams](#post-apiv1instagrams)
  - [GET /api/v1/instagrams](#get-apiv1instagrams)
  - [GET /api/v1/instagrams/:id](#get-apiv1instagramsid)
  - [GET /api/v1/instagrams/:id/stop](#get-apiv1instagramsidstop)
  - [GET /api/v1/instagrams/:id/start](#get-apiv1instagramsidstart)
  - [PUT /api/v1/instagrams/:id](#put-apiv1instagramsid)
  - [DELETE /api/v1/instagrams/:id](#delete-apiv1instagramsid)
  - [POST /api/v1/intagram/:instagram_id/channels](#post-apiv1instagraminstagram_idchannels)
  - [GET /api/v1/instagram/:instagram_id/channels](#get-apiv1instagraminstagram_idchannels)
  - [GET /api/v1/instagram/:instagram_id/channels/:id](#get-apiv1instagraminstagram_idchannelsid)
  - [GET /api/v1/instagram/:instagram_id/channels/:id/stop](#get-apiv1instagraminstagram_idchannelsidstop)
  - [GET /api/v1/instagram/:instagram_id/channels/:id/start](#get-apiv1instagraminstagram_idchannelsidstart)
  - [PUT /api/v1/instagram/:instagram_id/channels/:id](#put-apiv1instagraminstagram_idchannelsid)
  - [DELETE /api/v1/instagram/:instagram_id/channels/:id](#delete-apiv1instagraminstagram_idchannelsid)
## Dictionaries
#### Instagram states
- start
- stop
- suspected
#### Channel states
- start
- stop
- empty
#### Channel actions
- like
- comment
- follow
- unfollow
#### Channel targets
- followers
- subscriptions
- hashtag
- likes
- comments
## Entities
#### Token
```json
{
    "expire": "2017-03-18T17:34:30+04:00",
    "token": "xxx.xxx.xxx"
}
```
#### User
```json
{
    "id": "1",
    "username": "user@example.com",
    "created_at": "2017-01-21T13:18:32Z",
    "updated_at": "2017-01-21T13:18:32Z"
}
```
#### Instagram
```json
{
    "id": 20,
    "info": {
        "is_business": false,
        "profile_pic_url": "https://profile_pic_url.jpg",
        "hd_profile_pic_url_info": {
            "width": 150,
            "height": 150,
            "url": "https://hd_profile_pic_url_info.jpg"
        },
        "usertags_count": 0,
        "external_lynx_url": "http://l.instagram.com/external_lynx_url",
        "following_count": 64,
        "has_anonymous_profile_picture": true,
        "geo_media_count": 0,
        "external_url": "http://instagram.com/username",
        "username": "username",
        "biography": "biography",
        "has_chaining": true,
        "full_name": "full_name",
        "is_private": false,
        "pk": 1234567890,
        "follower_count": 25,
        "profile_pic_id": "",
        "is_verified": false,
        "hd_profile_pic_versions": [
            {
                "width": 320,
                "height": 320,
                "url": "https://hd_profile_pic_versions.jpg"
            },
            {
                "width": 640,
                "height": 640,
                "url": "https://hd_profile_pic_versions.jpg"
            }
        ],
        "media_count": 5,
        "is_favorite": false
    },
    "state": "start",
    "username": "username",
    "password": "username",
    "trusted": false,
    "hours": {
        "min": 0,
        "max": 0
    },
    "speed": {
        "like": 10,
        "comment": 10,
        "follow": 10,
        "unfollow": 10
    },
    "created_at": "2017-07-21T12:40:45+04:00",
    "updated_at": "2017-07-21T12:40:48.521987752+04:00"
}
```
#### Channel
```json
{
    "id": 10,
    "action": "like",
    "target": "followers",
    "value": "natgeo",
    "comments": [],
    "state": "empty",
    "leads_count": 0,
    "targets_count": 30,
    "followers_count": 0,
    "created_at": "2017-07-16T19:52:08+04:00",
    "updated_at": "2017-07-21T12:49:37+04:00"
}
```
## Errors
#### HTTP 400 BAD REQUEST
```json
{"code": 1010, "message": "Request to instagram failed",          "payload": ""}
{"code": 1020, "message": "Instagram login failed",               "payload": ""}
{"code": 1030, "message": "Error parse instagram response",       "payload": ""}
{"code": 1040, "message": "Error fetch followers list",           "payload": ""}
{"code": 1050, "message": "Instagram account is suspected",       "payload": ""}

{"code": 2010, "message": "Instagram not found",                  "payload": ""}
{"code": 2020, "message": "Instagram collection not found",       "payload": ""}
{"code": 2030, "message": "Instagram not created",                "payload": ""}
{"code": 2040, "message": "Instagram not updated",                "payload": ""}
{"code": 2050, "message": "Instagram not deleted",                "payload": ""}
{"code": 2060, "message": "Undefined action type",                "payload": ""}
{"code": 2070, "message": "Instagram login failed",               "payload": ""}
{"code": 2080, "message": "Instagram info fetching failed",       "payload": ""}
{"code": 2090, "message": "Instagram account is inactive",        "payload": ""}

{"code": 3010, "message": "User not found",                       "payload": ""}
{"code": 3030, "message": "User not created",                     "payload": ""}

{"code": 4010, "message": "Undefined action type",                "payload": ""}
{"code": 4020, "message": "Error during channel action execution","payload": ""}
{"code": 4030, "message": "Undefined target type",                "payload": ""}
{"code": 4040, "message": "Error during channel queue filling",   "payload": ""}
{"code": 4050, "message": "Channel not found",                    "payload": ""}
{"code": 4060, "message": "Channel collection not found",         "payload": ""}
{"code": 4070, "message": "Channel not created",                  "payload": ""}
{"code": 4080, "message": "Channel not updated",                  "payload": ""}
{"code": 4090, "message": "Channel not deleted",                  "payload": ""}
{"code": 4100, "message": "Empty action",                         "payload": ""}

{"code": 9999, "message": "Unknown error",                        "payload": ""}
```
#### HTTP 401 UNAUTHORIZED
```json
{ "code": 401, "message": "Incorrect Username / Password" }
{ "code": 401, "message": "auth header empty" }
{ "code": 401, "message": "signature is invalid" }
```
## Methods
#### POST /api/v1/login
###### Params
```json
Content-Type: application/json
{
    "username": "user@email.com",
    "password": "password"
}
```
###### Response
```json
HTTP 200 OK
{
    "expire": "2017-03-18T17:26:24+04:00",
    "token": "xxx.xxx.xxx"
}
```
#### POST /api/v1/signup
###### Params
```json
Content-Type: application/json
{
    "username": "user@email.com",
    "password": "password"
}
```
###### Response
```json
HTTP 200 OK
{
    "id": "1",
    "username": "user@email.com",
    "created_at": "2017-01-21T13:18:32Z",
    "updated_at": "2017-01-21T13:18:32Z"
}
```
#### GET /api/v1/user/refresh_token
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
{
    "expire": "2017-03-18T17:34:30+04:00",
    "token": "xxx.xxx.xxx"
}
```
#### GET /api/v1/user
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
<User>
```
#### POST /api/v1/instagrams/
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
{
    "username": "instalogin",
    "password": "instapassword",
    "trusted": false,
    "hours": {
        "min": 7,
        "max": 15
    },
    "speed": {
        "like": 10,
        "comment": 10,
        "follow": 10,
        "unfollow": 10
    }
}
```
###### Response
```json
HTTP 200 OK
<Instagram>
```
#### GET /api/v1/instagrams/
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```
HTTP 200 OK
[
    <Instagram_1>,
    <Instagram_2>,
    ...
    <Instagram_n>
]
```
#### GET /api/v1/instagrams/:id
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
<Instagram>
```
#### GET /api/v1/instagrams/:id/stop
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
<Instagram>
```
#### GET /api/v1/instagrams/:id/start
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
<Instagram>
```
#### PUT /api/v1/instagrams/:id
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
{
    "username": "instalogin",
    "password": "instapassword",
    "trusted": false,
    "hours": {
        "min": 7,
        "max": 15
    },
    "speed": {
        "like": 10,
        "comment": 10,
        "follow": 10,
        "unfollow": 10
    }
}
```
###### Response
```json
HTTP 200 OK
<Instagram>
```
#### DELETE /api/v1/instagrams/:id
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
{}
```
#### POST /api/v1/intagram/:instagram_id/channels
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
{
    "action": "like",
    "target": "followers",
    "value": "natgeo",
    "comments": []
}
```
###### Response
```json
HTTP 200 OK
<Channel>
```
#### GET /api/v1/instagram/:instagram_id/channels
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```
HTTP 200 OK
[
    <Channel_1>,
    <Channel_2>,
    ...
    <Channel_n>
]
```
#### GET /api/v1/instagram/:instagram_id/channels/:id
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
<Channel>
```
#### GET /api/v1/instagram/:instagram_id/channels/:id/stop
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
<Channel>
```
#### GET /api/v1/instagram/:instagram_id/channels/:id/start
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
<Channel>
```
#### PUT /api/v1/instagram/:instagram_id/channels/:id
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
{
    "action": "like",
    "target": "followers",
    "value": "natgeo",
    "comments": []
}
```
###### Response
```json
HTTP 200 OK
<Channel>
```
#### DELETE /api/v1/instagram/:instagram_id/channels/:id
###### Params
```json
Content-Type: application/json
Authorization: Bearer xxx.xxx.xxx
```
###### Response
```json
HTTP 200 OK
{}
```
