# StoryGraph API v1

REST API Documentation for communication between the client and the server.

__Note:__ All requests return a valid *JSON* objects.

**********



## Stories

### List Stories
Returns a lsit of stories.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1//stories/list | `GET` | - | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

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
| */api/v1/stories/{story_id}* | `GET` | *story_id* - ID of the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

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
| */api/v1/stories/upload* | `POST` | *title* - Title for the new story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized <br> __422__ - Unprocessable Entity (probably the story title is not unique) | 

__Sample Response:__
```javascript
{
    "id": "<new_id>",
    "title": "<Story Title>",
    "lastEdited": "dd-mm-yyyy",
}
``` 

### Update a story
Uploads a empty story

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{story_id}* | `PATCH` | *story_id* - *(required)* ID of the requested story <br> *title* - (string) New title for the story <br>*current_event_id* - change the id of current event, can be `null` | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
{
    "id": "<story_id>",
    "title": "<Story Title>",
    "lastEdited": "dd-mm-yyyy",
    "currentEvent": "Title of the <current_event_id>",
}
``` 

### Delete a story
Delete a story

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{story_id}* | `DELETE` | *story_id* - *(required)* ID of the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
{
    "message": "Story with id '<story_id>' was successfully deleted",
}
//QUESTION: We can return the deleted story details here?
``` 



## Events

### List Events
Returns a list of events.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events/list | `GET` |  *storyID* - *(required)* ID for the requested story | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized | 

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
| */api/v1/stories/{storyID}/events/upload* | `POST` | *storyID* - *(required)* Title for the new story <br>*title* - *(required)* Title for the new event | __200__ - OK <br>__404__ - Story not found <br> __401__ - Unauthorized <br>__422__ - Unprocessable Entity (probably the event title is not unique)| 

__Sample Response:__
```javascript
{
    "id": <new_event_id>,
    "name": "<title>",
    "changelog": []
}
``` 

### Update a story
Uploads a empty story

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events{story_id}* | `PATCH` | *storyID* - *(required)* ID for the requested story <br>*eventID* - *(required)* ID for the requested event <br>*title* - New title for the event <br><br>*changelog_action* - *(required if changelog is used)* <br>`append` - append items to the end of changelog <br>`replace` - repalce the current changelog with the \<changelog\> parameter <br><br>*changelog* - a valid changelog array in JSON [1] | __200__ - OK <br>__404__ - Story/Event not found <br> __401__ - Unauthorized | 

__[1]__ - Changelog JSON format:
```javascript
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
    }
]
``` 
__Sample Response:__ Returns event details after update
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

### Delete an Event
Delete an event

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/stories/{storyID}/events{story_id}* | `DELETE` | *storyID* - *(required)* ID for the requested story <br>*eventID* - *(required)* ID for the requested event | __200__ - OK <br>__404__ - Story/Event not found <br> __401__ - Unauthorized | 

__Sample Response:__
```javascript
{
    "message": "Event with id '<story_id>' was successfully deleted",
}
//QUESTION: We can return the deleted event details here?
``` 