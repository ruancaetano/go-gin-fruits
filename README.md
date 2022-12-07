# go-gin-fruits

Just a fruit crud, yeah useless, this project was created just to test technologies like `gin`, `godog` and `swag`. That way, don't take it too seriously, I recommend relaxing and enjoying a good gin with fruits of your choice 😄.


## How to play with it?

### To run api
In your terminal, after clone this repository, run:

```sh
docker compose up -d
```

After that the swagger will be accessible at `http://localhost:8080/swagger/index.html`

### To run tests

```sh
docker compose up -d
```

```sh
docker exec go-gin-fruits go test ./... -v
```
