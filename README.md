## How to run this project

---

1. Create an .env file whose contents are based on the .env.example file
2. Create a database named db based on DB_NAME in the .env file
3. run the project

```
go run main.go
```

# Testing api with postman

```
import the file in Postman located at ./export/Blogs API.postman_collection.json
```

# Api Spec

### GROUP: Blog

- [1] - Create blog
- [POST] : {root.api}/api/v1/blogs

```json

Request:
{
    "author": "Adam nasrudin",
    "title": "Belajar dengan bahasa Golang",
    "description": "Golang merupakan bahasa yang saat ini kian populer dikalangan backend"
}

Response:
{
    "meta": {
        "message": "Success to create blog",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 6,
        "author": "Adam nasrudin",
        "title": "Belajar dengan bahasa Golang",
        "description": "Golang merupakan bahasa yang saat ini kian populer dikalangan backend",
        "comments": null
    }
}

```

- [2] - List all blogs
- [GET] : {root.api}/api/v1/blogs

```json

Response:
{
    "meta": {
        "message": "List of blogs",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 4,
            "author": "Nasrudin Adam ",
            "title": "Belajar dengan bahasa Golang",
            "description": "Java merupakan bahasa yang saat ini kian populer dikalangan backend",
            "comments": [
                {
                    "id": 6,
                    "author": "radit dika ",
                    "comments": "ouh Golang itu seperti itu yah"
                }
            ]
        },
        {
            "id": 6,
            "author": "Adam nasrudin",
            "title": "Belajar dengan bahasa Golang",
            "description": "Golang merupakan bahasa yang saat ini kian populer dikalangan backend",
            "comments": []
        }
    ]
}
```

- [3] - Details blog
- [GET] : {root.api}/api/v1/blogs/:id

```json

Response:
{
    "meta": {
        "message": "List of Detail blog",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 4,
        "author": "Nasrudin Adam ",
        "title": "Belajar dengan bahasa Golang",
        "description": "Java merupakan bahasa yang saat ini kian populer dikalangan backend",
        "comments": [
            {
                "id": 6,
                "author": "radit dika ",
                "comments": "ouh Golang itu seperti itu yah"
            }
        ]
    }
}
```

- [4] - Update blog
- [PUT] : {root.api}/api/v1/blogs/:id/update

```json
Request:
{
    "author": "Adam ",
    "title": "Belajar Golang",
    "description": "description Belajar Golang merupakan bahasa yang saat ini kian populer dikalangan backend"
}
Response:
{
    "meta": {
        "message": "Success to updated blog",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 4,
        "author": "Adam ",
        "title": "Belajar Golang",
        "description": "description Belajar Golang merupakan bahasa yang saat ini kian populer dikalangan backend",
        "comments": [
            {
                "id": 6,
                "author": "radit dika ",
                "comments": "ouh Golang itu seperti itu yah"
            }
        ]
    }
}
```

- [4] - Delete blog by id
- [DELETE] : {root.api}/api/v1/blogs/:id/delete

```json

Response:
{
    "meta": {
        "message": "Deleted blog",
        "code": 200,
        "status": "success"
    },
    "data": null
}
```

### GROUP: Comment

- [1] - Create comment
- [POST] : {root.api}/api/v1/blogs/:id/comment

```json

Request:
{
    "author": "Hello Radit",
    "comments": "Ini komentar pertamamu"
}

Response:
{
    "meta": {
        "message": "Success to create comment",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 7,
        "author": "Hello Radit",
        "comments": "Ini komentar pertamamu"
    }
}

```

- [2] - Update comment by idcomment
- [PUT] : {root.api}/api/v1/blogs/:id/comment/:idComment/update

```json

Request:
{
    "author": "Hello Radit",
    "comments": "ini Update pertamamu"
}

Response:
{
    "meta": {
        "message": "Success to updated comment",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 7,
        "author": "Hello Radit",
        "comments": "ini Update pertamamu"
    }
}
```

- [3] - Delete comment by idcomment
- [DELETE] : {root.api}/api/v1/blogs/:id/comment/:idComment/delete

```json

Response:
{
    "meta": {
        "message": "Deleted comment",
        "code": 200,
        "status": "success"
    },
    "data": null
}
```
