# API文档



## 注册接口

URL：/register

Method：POST

Body：

```json
{
    "username":"zcw",
    "password":"123456"
}
```

Response：

```json
{
    "code": 0,
    "data": {
        "id": "ab410050-bd02-492c-aa51-d59d5a3c22b4",
        "user_name": "zcw",
        "pass_word": "123456",
        "image_path": "",
        "register_time": "2023-05-13T05:45:01Z"
    },
    "msg": "success"
}
```

## 登录接口

URL：/login

Method：POST

Body：

```json
// json
{
    "username":"zcw",
    "password":"123456"
}
```

Response：

```json
{
    "code": 0,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYWI0MTAwNTAtYmQwMi00OTJjLWFhNTEtZDU5ZDVhM2MyMmI0IiwidXNlcm5hbWUiOiJ6Y3ciLCJleHAiOjE2ODM5NzI4NjEsImlzcyI6ImJsdWViZWxsIn0.xLlsa-VyEVijvijryGeY43C-qB_JZ-at7XdKdzvztqQ"
    },
    "msg": "success"
}
```

## 查询文件接口

URL：/user/file/{file_id}

Method：GET

Header："Authorization: Bearer Token"

Response：

```json
{
    "code": 0,
    "data": {
        "id": "8ea6e87c-0a81-4ed8-acc3-5aaa468c7217",
        "user_id": "ab410050-bd02-492c-aa51-d59d5a3c22b4",
        "folder_id": "0",
        "file_name": "20230301151918",
        "file_hash": "31334b8b88b8486b5ed9ccc450bb1658b4285aa60dcb9566874efe8dd71b19bf",
        "postfix": ".csv",
        "size": 0,
        "type": 5,
        "download_num": 0,
        "created_at": "2023-05-13T05:45:58+08:00",
        "updated_at": "2023-05-13T05:45:58+08:00"
    },
    "msg": "success"
}
```

## 上传文件接口

URL：/user/file

Method：POST

Header："Authorization: Bearer Token"

Body：

```go
// from
file=binarydata&folder_id
```

Response：

```json
{
    "code": 0,
    "data": {
        "id": "23cedcb5-7f1e-436d-8e92-1bdeb477bdf9",
        "user_id": "476dc087-1b2c-45ba-9838-adecff569da5",
        "folder_id": "0",
        "file_name": "20230509_224230",
        "file_hash": "3a283effdac77cef0461b0bd1d4eb3068e89aad9c4b2f92a946d8ab845a9fbe2",
        "postfix": ".csv",
        "size": 2,
        "type": 5,
        "download_num": 0,
        "created_at": "2023-05-13T16:21:05.84444+08:00",
        "updated_at": "2023-05-13T16:21:05.84444+08:00"
    },
    "msg": "success"
}
```

## 删除文件接口

URL：/user/file/{file_id}

Method：DELETE

Header："Authorization: Bearer Token"

Response：

```json
{
    "code": 0,
    "data": "",
    "msg": "success"
}
```

## 查询目录接口

URL：/folder/{folder_id}

Method：GET

Header："Authorization: Bearer Token"

Response：

```json
{
    "code": 0,
    "data": {
        "current_folder": {
            "id": "",
            "name": "",
            "user_id": "",
            "parent_id": "",
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        "files": [
            {
                "id": "58005313-8096-46d7-bca0-f8d208bfbbb8",
                "user_id": "ab410050-bd02-492c-aa51-d59d5a3c22b4",
                "folder_id": "0",
                "file_name": "goland-2020.3.5.exe",
                "file_hash": "ac7e687b2503d7da07c28b88d50c2d5b3eb259809060632c4e331c6912f997a3",
                "postfix": ".mp4",
                "size": 381770,
                "type": 3,
                "download_num": 3,
                "created_at": "2023-05-13T15:04:07+08:00",
                "updated_at": "2023-05-13T15:09:30+08:00"
            }
        ],
        "folders": []
    },
    "msg": "success"
}
```

## 创建目录接口

URL：/user/folder

Method：POST

Body：

```json
{
    "parent_id":"f305c1c8-7a65-4494-a997-cc1ebd6d9c90",
    "folder_name":"目录22"
}
```

Response：

```json
{
    "code": 0,
    "data": {
        "id": "01068e43-1d91-4e57-b61f-2efe715914c3",
        "name": "122",
        "user_id": "476dc087-1b2c-45ba-9838-adecff569da5",
        "parent_id": "f305c1c8-7a65-4494-a997-cc1ebd6d9c90",
        "created_at": "2023-05-13T16:41:20.31908+08:00",
        "updated_at": "2023-05-13T16:41:20.31908+08:00"
    },
    "msg": "success"
}
```

## 修改目录接口

URL：/user/folder

Method：PUT

Body：

```json
{
    "folder_id":"40ab600a-bb9a-4f4f-9447-e82aafec1188",
    "folder_name":"22"
}
```

Response：

```json
{
    "code": 0,
    "data": {
        "id": "",
        "name": "22",
        "user_id": "",
        "parent_id": "",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "2023-05-13T17:01:50.741008+08:00"
    },
    "msg": "success"
}
```

## 删除目录接口

URL：/user/folder/{folder_id}

Method：DELETE

Response：

## 查询分享接口

URL：/user/share

Method：GET

Response：

```json
{
    "code": 0,
    "data": [
        {
            "id": "406ceaad-f4aa-4c3e-9706-d18b00e579b6",
            "code": "5fc2753dcb",
            "user_id": "476dc087-1b2c-45ba-9838-adecff569da5",
            "file_id": "40ab600a-bb9a-4f4f-9447-e82aafec1188",
            "created_at": "2023-05-13T17:05:17+08:00",
            "updated_at": "2023-05-13T17:05:17+08:00"
        },
        {
            "id": "4a201ba4-3138-43ac-8043-d19ff32b8bef",
            "code": "a72c497bf8",
            "user_id": "476dc087-1b2c-45ba-9838-adecff569da5",
            "file_id": "",
            "created_at": "2023-05-13T17:06:41+08:00",
            "updated_at": "2023-05-13T17:06:41+08:00"
        }
    ],
    "msg": "success"
}
```

## 创建分享接口

URL：/user/share

Method：POST

Body：

```json
{
    "file_id":"40ab600a-bb9a-4f4f-9447-e82aafec1188"
}
```

Response：

```json
{
    "code": 0,
    "data": {
        "id": "406ceaad-f4aa-4c3e-9706-d18b00e579b6",
        "code": "5fc2753dcb",
        "user_id": "476dc087-1b2c-45ba-9838-adecff569da5",
        "file_id": "40ab600a-bb9a-4f4f-9447-e82aafec1188",
        "created_at": "2023-05-13T17:05:16.618745+08:00",
        "updated_at": "2023-05-13T17:05:16.618745+08:00"
    },
    "msg": "success"
}
```

## 删除分享接口

URL：/user/share/{share_id}

Method：DELETE

Response：