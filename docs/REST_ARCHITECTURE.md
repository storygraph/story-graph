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

