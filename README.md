# KB

KB is a Go  API that shows computer availability in the Kilburn Building at the University of Manchester

## Using the API

To get the availability in the whole school:

`GET /api/status`

```json
{
  "e-c07kilf1601.it.manchester.ac.uk": {
    "status": "free",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "Tootil 1"
  },
  "e-c07kilf901.it.manchester.ac.uk": {
    "status": "used",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "Tootil 0"
  },
  "e-c07kilf3101.it.manchester.ac.uk": {
    "status": "offline",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "LF31"
  }
}
```

To get the availability in one lab:

`GET /api/status/labs/:lab`

```json
{
  "e-c07kilf901.it.manchester.ac.uk": {
    "status": "free",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "Tootil 0"
  },
  "e-c07kilf902.it.manchester.ac.uk": {
    "status": "free",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "Tootil 0"
  },
  "e-c07kilf903.it.manchester.ac.uk": {
    "status": "used",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "Tootil 0"
  }
}
```

To check if a particular machine is available:

`GET /api/status/machines/:machine`

```json
{
  "status": "free",
  "timestamp": "2000-01-01T00:00:00.000000000+01:00",
  "lab": "Tootil 0"
}
```



## Updating the state

The reporter feeds the data to the API via the `/report/` endpoint:

`POST /report/`

```
{
  "e-c07kilf901.it.manchester.ac.uk": {
    "status": "free",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "Tootil 0"
  },
  "e-c07kilf902.it.manchester.ac.uk": {
    "status": "free",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "Tootil 0"
  },
  "e-c07kilf903.it.manchester.ac.uk": {
    "status": "used",
    "timestamp": "2000-01-01T00:00:00.000000000+01:00",
    "lab": "Tootil 0"
  }
}
```

Will return a 200 OK and the state will be updated