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
  "status": "Connection successfully created"
}
```

```json
{
  "status": "Wrong password"
}
```

```json
{
  "status": "Unknown user name"
}
```

etc...

## GET /close

Endpoint para cerrar la conexión con el servidor

### Response

```json
{
  "status": "Connection successfully closed"
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

```json
{
  "status": "Directory Dir successfully created"
}
```

```json
{
  "status": "Error while removing directory Dir"
}
```

etc...