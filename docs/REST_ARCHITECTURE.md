# StoryGraph API v1

REST API Documentation for communication between the client and the server.

__Note:__ All requests return a valid *JSON* objects.

**********



## Stories

### List Stories
Returns a lsit of stories.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories* | `GET` | - | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
[
    {
        "id": 1,
        "title": "Baron Hedgehog",
    },
    {
        "id": 2,
        "title": "Giggles",
    },
    {
        "id": 3,
        "title": "Wreckage",
    },
    {
        "id": 4,
        "title": "Friends in arms",
    },
    {
        "id": 5,
        "title": "Neverspring",
    },
],
```
__Note:__ If the requesting user does not have any stories, then the returned JSON is:

```javascript
[]
```

### Get Story
Returns story details.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}* | `GET` | *storyID* - ID of the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
{
    "id": 1,
    "title": "Baron Hedgehog",
    "lastEdited": "05-05-2020",
    "currentEvent": "The battle of Noxville",
}
``` 

### Upload Story
Uploads a empty story

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/upload* | `POST` | - | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized <br> __422__ - Unprocessable Entity (probably the story title is not unique) | 

__Sample Request:__
```javascript
{
    "title": "<New Tilte>"
}
``` 

__Sample Response:__
```javascript
{
    "id": "<new_id>"
}
``` 

### Update a story
Uploads a empty story

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}* | `PATCH` | *storyID* - *(required)* ID of the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

__Sample Request:__
```javascript
{
    "title": "New Title",
    "currentEvent": "The battle of Noxville",
}
``` 

__Sample Response:__
```javascript
{
    "id": "<story_id>"
}
``` 

### Delete a story
Delete a story

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}* | `DELETE` | *storyID* - *(required)* ID of the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 



## Events

### List Events
Returns a list of events.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events* | `GET` |  *storyID* - *(required)* ID for the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
[
    {
        "id": 0,
        "name": "Frodo meets Gandalf",
    },
    {
        "id": 1,
        "name": "Sam, Pippin and Merry join the company",
    }, 
    {
        "id": 2,
        "name": "Gandalf departs",
    },
    {
        "id": 3,
        "name": "Hobbits set of to Tom Bombadil's cottage",
    }
]
```
__Note:__ If the requested story does not have any events, then the returned JSON is:

```javascript
[]
```

### Get Event
Returns event details.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events/{eventID}* | `GET` | *storyID* - *(required)* ID for the requested story <br>*eventID* - *(required)* ID for the requested event | __200__ - OK <br>__404__ - Story/Event not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
{
    "id": 0,
    "name": "Frodo meets Gandalf",
    "changelog":
    [
        {
            "subject": {
                "type": "vertex",
                "name": "Gandalf"
            },
            "action": "add",
            "description": "A wise magician."
        },
        {
            "subject": {
                "type": "vertex",
                "name": "Frodo"
            },
            "action": "add",
            "description": "Fat little hobbit. Son of Bilbo."
        },
        {
            "subject": {
                "type": "edge",
                "edge": [0, 1], // from, to
                "name": "Frodo and Gandalf become allies"
            },
            "action": "add",
            "description": "Gandalf and Frodo are allies.",
        }
    ]
}
``` 

### Upload Event
Appends an empty event to the end of the timeline

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events* | `POST` | *storyID* - *(required)* Title for the new story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized <br>__422__ - Unprocessable Entity (probably the event title is not unique)| 

__Sample Request:__
```javascript
{
    "name": "<title>"
}
``` 

__Sample Response:__
```javascript
{
    "id": "<new_event_id>"
}
``` 

### Update a Event
Update a story by appending or replacing it's changelog

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events/{eventID}* | `PATCH` | *storyID* - *(required)* ID for the requested story <br>*eventID* - *(required)* ID for the requested event <br>*title* - New title for the event | __200__ - OK <br>__404__ - Story/Event not found <br> __401__ - Unauthorized | 

__Sample Request:__
```javascript
{
    "name": "Frodo meets Gandalf",
    "changelog":
    [
        {
            "subject": {
                "type": "vertex",
                "name": "Gandalf"
            },
            "action": "add",
            "description": "A wise magician."
        },
        {
            "subject": {
                "type": "vertex",
                "name": "Frodo"
            },
            "action": "add",
            "description": "Fat little hobbit. Son of Bilbo."
        },
        {
            "subject": {
                "type": "edge",
                "edge": [0, 1], // from, to
                "name": "Frodo and Gandalf become allies"
            },
            "action": "add",
            "description": "Gandalf and Frodo are allies.",
        }
    ]
}
``` 
__Sample Response:__ Returns event details after update
```javascript
{
    "id": 0
}
``` 

### Delete an Event
Delete an event

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events/{eventID}* | `DELETE` | *storyID* - *(required)* ID for the requested story <br>*eventID* - *(required)* ID for the requested event | __200__ - OK <br>__404__ - Story/Event not found <br> __401__ - Unauthorized | 



## Weenies

### List Weenies
Returns a list of weenies.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/weenies* | `GET` |  *storyID* - *(required)* ID for the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
[
    {
        "id": 0,
        "name": "Frodo",
    },
    {
        "id": 1,
        "name": "Sam",
    }, 
    {
        "id": 2,
        "name": "Gandalf",
    },
    {
        "id": 3,
        "name": "Pippin",
    }
]
```
__Note:__ If the requested story does not have any weenies, then the returned JSON is:

```javascript
[]
```

### Get Weenie
Returns event details.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/weenies/{weenieID}* | `GET` | *storyID* - *(required)* ID for the requested story <br>*weenieID* - *(required)* ID for the requested event | __200__ - OK <br>__404__ - Story/Weenie not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
{
    "id": 0,
    "name": "Frodo",
    "status": "alive",
    "location": "Middle-earth",
    ....
}
``` 

### Upload Weenie
Appends an empty event to the end of the timeline

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/weenies* | `POST` | *storyID* - *(required)* ID for the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

__Sample Request:__
```javascript
{
    "name": "Frodo",
    "status": "alive",
    "location": "Middle-earth"
}
``` 

__Sample Response:__
```javascript
{
    "id": 0,
    "name": "Frodo",
    "status": "alive",
    "location": "Middle-earth"
}
``` 

### Update a Weenie
Uploads a weenie

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/weenies/{weenieID}* | `PATCH` | *storyID* - *(required)* ID for the requested story <br>*weenieID* - *(required)* ID for the requested weenie | __200__ - OK <br>__404__ - Story/Event not found <br> __401__ - Unauthorized | 

__Note:__ If a property already exists the value will be replaced with one from the request

__[1] Sample Request:__
```javascript
{
    "name": "Frodo",
    "status": "alive",
    "location": "Middle-earth",
    "age": 50
}
``` 

__Sample Response:__
```javascript
{
    "id": 0,
    "name": "Frodo",
    "status": "alive",
    "location": "Middle-earth",
    "age": 50
}
``` 

### Delete a Weenie
Delete a Weenie 

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events/{weenieID}* | `DELETE` | *storyID* - *(required)* ID for the requested story <br>*weenieID* - *(required)* ID for the requested weenie | __200__ - OK <br>__404__ - Story/Weenies not found <br> __401__ - Unauthorized | 
