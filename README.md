# Go-Tunnel! SSH Tunnel Auto-Connector 🚇💻

This Go program reads SSH tunnel profiles from a `config.json` file and automatically establishes the SSH tunnels for profiles where forwarding is enabled.

## 🌟 Program Features

- 🚀 Automatically establishes SSH tunnels based on profiles in a configuration file.
- 👀 Monitors the SSH tunnels and re-establishes them if they go down.
- 📑 Handles multiple profiles and multiple forwarding options per profile.


## 🛠️ Setup


1. 📦 Clone this repository:
```bash
git clone https://github.com/ryan-shaw/go-tunnel.git
```

2. 📝 Update the config.json file with your SSH tunnel profiles. For example:
```json
{
  "profiles" : [
    {
      "port" : 10022,
      "address" : "your.host.name",
      "forwardings" : [
        {
          "bindPort" : 8080,
          "enabled" : true
        }
      ]
    }
  ]
}
```

3. 🔨 Build the Go program:
```bash
go build -o app
```

4. 🏃 Run
```bash
./app
```
