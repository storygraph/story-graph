# StoryGraph API v1

REST API Documentation for communication between the client and the server.

__Note:__ All requests return a valid *JSON* objects.



## Stories
**********

### List Stories
Returns a lsit of stories.
Needs *Colaborator*/*Author* rights for the project and *Read* rights stories.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/project/{project_id}/stories/{story_id}* | `GET` | *project_id* - ID of the project <br> *story_id* - ID of the requested story | __200__ - Everything is Okay <br>__404__ - Story not found <br> __401__ - Not Authenticated <br> __500__ - Server error | 

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
__Note:__ If the requesting user has *Colaborator*/*Author* rights for the project, but has no *Read* right for any of the stories, then the returned JSON is::

```javascript
[]
```



### Get Story
Returns story details.
Needs *Colaborator*/*Author* rights for the project and *Read* rights for the requested story.

| Endpoint | Type | Parameters | Return Codes |
| -------- | ---- | ---------- | ------------ |
| */api/v1/project/{project_id}/stories/{story_id}* | `GET` | *project_id* - ID of the project <br> *story_id* - ID of the requested story | __200__ - Everything is Okay <br>__404__ - Story not found <br> __401__ - Not Authenticated <br> __500__ - Server error | 

__Sample Response:__
```javascript
{
    "id": 1,
    "title": "Baron Hedgehog",
    "lastEdited": "05-05-2020",
    "currentEvent": "The battle of Noxville",
}
``` 


