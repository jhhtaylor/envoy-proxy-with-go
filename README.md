# envoy-proxy-with-go
Go to `envoy-proxy-with-go/`

**Step 1: Install and set up Envoy**

2.1 Install Envoy. For macOS, you can use Homebrew:

```bash
brew install envoy
```

**Step 2: Running the Demo**

3.1 Start two instances of the Go server:

```bash
PORT=8080 go run server.go
PORT=8081 go run server.go
```

3.2 Start Envoy with the configuration:

```bash
envoy -c envoy.yaml
```

3.3 Now, if you make requests to **`http://localhost:10000`**, Envoy will load balance between the two Go servers:

```bash
curl http://localhost:10000
```

You'll see responses from different server instances based on the ROUND_ROBIN load-balancing policy.