# Endpoints (GUI -> FTP Client)

## POST url/connect

Endpoint para conectar con el servidor

### Request

```json
{
  "ipAddress": "127.0.0.1",
  "userName": "johnDoe",
  "password": "pass"
}
```

### Response

```json
{
  "status": "Connection successfully created",
  "successful": true
}
```

```json
{
  "status": "Wrong password",
  "successful": false
}
```

```json
{
  "status": "Unknown user name",
  "successful": false
}
```

etc...

## GET url/close

Endpoint para cerrar la conexión con el servidor

### Response

```json
{
  "status": "Connection successfully closed",
  "successful": true
}
```

## GET url/status

Endpoint para manejar el estado de las operaciones realizadas sobre el servidor. Este request se producirá de forma periódica (cada 300 msec, por ejemplo)

### Response

```json
{
  "status": "File music.mp3 successfully downloaded"
}
```

```json
{
  "status": "File pic.jpeg successfully uploaded"
}
```

etc...

## GET url/list/local

Endpoint para listar el directorio local

### Request

```json
{
  "path": "/"
}
```

### Response 

```json
  {
    "directories": "string of all the directories and files",
    "successful": true
  }
```

## POST url/list/server

Endpoint para listar un directorio del servidor

### Request

```json
  {
    "path": "./Music"
  }
```


### Response 

```json
  {
    "directories": "string of all the directories and files",
    "successful": true
  }
```

## POST url/files/upload

### Request

```json
{
  "source": "/Music/music.mp3",
  "destination": "/Data/"
}
```

### Response

```json
{
  "status": "Start uploading file",
  "successful": true
}
```

```json
{
  "status": "Error while uploading file",
  "successful": false
}
```

## POST url/files/download

### Request

```json
{
  "source": "/Music/music.mp3",
  "destination": "/Data/"
}
```

### Response

```json
{
  "status": "Start downloading file",
  "successful": true
}
```

```json
{
  "status": "Error while downloading file",
  "successful": false
}
```

## POST url/files/remove

### Request

```json
{
  "path": "./Pictures/pic.jpeg"
}
```

### Response

```json
{
  "status": "File successfully removed",
  "successful": true
}
```

```json
{
  "status": "Error while removing file",
  "successful": false
}
```

## POST url/directories/upload

### Request

```json
{
  "source": "/Music/",
  "destination": "/Data/"
}
```

### Response

```json
{
  "status": "Start uploading directory",
  "successful": true
}
```

```json
{
  "status": "Error while uploading directory",
  "successful": false
}
```

## POST url/directories/download

### Request

```json
{
  "source": "/Music/",
  "destination": "/Data/"
}
```

### Response

```json
{
  "status": "Start downloading directory",
  "successful": true
}
```

```json
{
  "status": "Error while downloading directory",
  "successful": false
}
```

## POST url/directories/remove

### Request

```json
{
  "path": "./Pictures"
}
```

### Response

```json
{
  "status": "Directory successfully removed",
  "successful": true
}
```

```json
{
  "status": "Error while removing directory",
  "successful": false
}
```

## POST url/directories/create

### Request

```json
{
  "path": "./Pictures/Nature"
}
```

### Response

```json
{
  "status": "Directory successfully created",
  "successful": true
}
```

```json
{
  "status": "Error while creating directory",
  "successful": false
}
```