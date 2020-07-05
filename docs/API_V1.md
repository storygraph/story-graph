# StoryGraph API v1

_**:bangbang: Note:** All requests return valid *JSON* objects._

## Stories

<table>
<tr>
<td>Endpoint</td> <td>Method</td> <td>Sample request</td> <td>Sample response</td> <td>Status</td> <td> Description </td>
</tr>
<tr>
<td>

 **/api/v1/stories**
</td>
<td>

 `GET`
</td>
<td>-</td>
<td>

```json
[
    {
        "id": 1,
        "title": "Hobbit",
    },
    {
        "id": 2,
        "title": "Giggles",
    },
]
```

</td>
<td>

__200__

</td>

<td> Returns a list of all the stories. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}**
</td>
<td>

 `GET`
</td>
<td>-</td>
<td>

```json
{
    "id": 1,
    "title": "Hobbit",
    "lastEdited": "25-05-1935",
    "currentEvent": "Roast Mutton",
}
```

</td>
<td>

__200__ <br>__404__

</td>

<td> Returns story details. </td>
</tr>
<tr>
<td>

 **/api/v1/stories**
</td>
<td>

 `POST`
</td>
<td>

```json
{
    "title": "The witcher"
}
```

</td>
<td>

```json
{
    "id": "1"
}
```

</td>
<td>

__200__ <br> __400__

</td>

<td> Uploads an empty story. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}**
</td>
<td>

 `PATCH`
</td>
<td>

```json
{
    "title": "LOTR",
}
```

</td>
<td>-</td>
<td>

__200__ <br> __400__ <br> __404__

</td>

<td> Updates a story. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}**
</td>
<td>

 `DELETE`
</td>
<td>-</td>
<td>-</td>
<td>

__200__ <br> __404__

</td>

<td> Deletes a story. </td>
</tr>
</table>

## Events

<table>
<tr>
<td>Endpoint</td> <td>Method</td> <td>Sample request</td> <td>Sample response</td> <td>Status</td> <td> Description </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/events**
</td>
<td>

 `GET`
</td>
<td>-</td>
<td>

```json
[
    {
        "id": 0,
        "name": "Frodo meets Gandalf",
    }, {
        "id": 1,
        "name": "Sam, Pippin and Merry join the company",
    }
]
```

</td>
<td>

__200__<br>__404__

</td>

<td> Returns a list of all the events of a story. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/events/{eventID}**
</td>
<td>

 `GET`
</td>
<td>-</td>
<td>

```json
{
    "id": 0,
    "name": "Frodo escapes Shelob's lair",
    "actions": [
        {
            "name": "Frodo stabs shelob",
            "delta": {
                "type": "weenie",
                "id": 1234,
                "name": "Shelob",
                "status": "wounded",
            },
        },
        {
            "name": "Frodo leaves the cave",
            "delta": {
                "type": "weenie",
                "id": 1,
                "name": "Frodo",
                "location": [125.0, 135.0],
            },
        },
    ]
}
```

</td>
<td>

__200__<br>__404__

</td>

<td> Returns the details of an event. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/events**
</td>
<td>

 `POST`
</td>
<td>

```json
{
    "name": "New event"
}
```

</td>
<td>

```json
{
    "id": 1234,
}
```

</td>
<td>

__200__

</td>

<td> Appends an empty event to the end of the timeline. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/events/{eventID}**
</td>
<td>

 `PATCH`
</td>
<td>

```json
{
    "id": 0,
    "name": "Frodo escapes Shelob's lair",
    "actions": [
        {
            "name": "Frodo stabs shelob",
            "delta": {
                "type": "weenie",
                "id": 1234,
                "name": "Shelob",
                "status": "wounded",
            },
        },
        {
            "name": "Shelob poisons Frodo",
            "delta": {
                "type": "weenie",
                "id": 1,
                "name": "Frodo",
                "status": "poisoned",
            },
        },
    ]
}
```

</td>
<td>

```json
{
    "id": 0,
}
```

</td>
<td>

__200__<br>__400__<br>__404__

</td>

<td> Updates an event. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/events/{eventID}**
</td>
<td>

 `DELETE`
</td>
<td>-</td>
<td>-</td>
<td>

__200__<br>__404__

</td>

<td> Deletes an event. </td>
</tr>
</table>







## Weenies

<table>
<tr>
<td>Endpoint</td> <td>Method</td> <td>Sample request</td> <td>Sample response</td> <td>Status</td> <td> Description </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/weenies**
</td>
<td>

 `GET`
</td>
<td>-</td>
<td>

```json
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
]
```

</td>
<td>

__200__<br>__404__

</td>

<td> Returns a list of all the weenies of a story. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/weenies/{weenieID}**
</td>
<td>

 `GET`
</td>
<td>-</td>
<td>

```json
{

    "id": 0,
    "name": "Frodo",
    "status": "poisoned",
    "tags": {
        "race": "hobbit", 
        "alliance": "The fellowship of the ring",
    },
    "image": "<image-url>",
    "possesions": [
        {
            "name": "Sting",
            "description": "A mystical sword forged by elfs.",
        },
        {
            "name": "Light of Earendil",
            "description": "A light to guide you in the darkest of ways.",
        }
    ]
}
```

</td>
<td>

__200__<br>__404__

</td>

<td> Returns the details of a weenie. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/weenies**
</td>
<td>

 `POST`
</td>
<td>

```json
{
    "name": "Frodo",
    "status": "poisoned",
    "tags": {
        "race": "hobbit", 
        "alliance": "The fellowship of the ring",
    },
    "image": "<image-url>",
    "possesions": [
        {
            "name": "Sting",
            "description": "A mystical sword forged by elfs.",
        },
        {
            "name": "Light of Earendil",
            "description": "A light to guide you in the darkest of ways.",
        }
    ]
}
```

</td>
<td>

```json
{
    "id": 1234,
}
```

</td>
<td>

__200__<br>__404__

</td>

<td> Adds a new weenie to the story. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/weenies/{weenieID}**
</td>
<td>

 `PATCH`
</td>
<td>

```json
{
    "id": 1234,
    "name": "Frodo",
    "status": "healthy",
}
```

</td>
<td>

```json
{
    "id": 1234,
}
```

</td>
<td>

__200__<br>__404__

</td>

<td> Updates the data of a weenie. </td>
</tr>
<tr>
<td>

 **/api/v1/stories/{storyID}/weenies/{weenieID}**
</td>
<td>

 `DELETE`
</td>
<td>-</td>
<td>-</td>
<td>

__200__<br>__404__

</td>

<td> Deletes a weenie. </td>
</tr>
</table>
