## Create New Song
**POST /songs**
```json
// Request
{
  "title": <string>,
  "artist": <string>,
  "album": <string>
}

// Response
{
  "songs": [
    {
      "ID": <int>,
      "title": <string>,
      "artist": <string>,
      "album": <string>
    }, {...}, ...
  ]
}
```
