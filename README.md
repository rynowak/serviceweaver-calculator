# serviceweaver-calculator

Experimenting with serviceweaver

## Running the sample

Prerequisites:

- Go 1.20
- NodeJS (any modern version)

0. Install Serviceweaver

```sh
go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
```

1. Build the client application

```sh
cd react-calculator/client
npm install && npm run build
cd ../..
```

2. Run the client application 

```sh
cd react-calculator
npm start
```

Leave it running

3. Run the service (new console)

Open a new console

```sh
weaver multi deploy weaver.toml
```

4. Open browser

Visit `http://localhost:8080` for the calculator