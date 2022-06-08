<!-- This file was autogenerated via cilium-agent --cmdref, do not edit manually-->

## cilium-agent completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	cilium-agent completion fish | source

To load completions for every new session, execute once:

	cilium-agent completion fish > ~/.config/fish/completions/cilium-agent.fish

You will need to start a new shell for this setup to take effect.


```
cilium-agent completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --enable-k8s-api-discovery         Enable discovery of Kubernetes API groups and resources with the discovery API
      --gops-port uint16                 Port for gops server to listen on (default 9890)
      --k8s-api-server string            Kubernetes API server URL
      --k8s-client-burst int             Burst value allowed for the K8s client
      --k8s-client-qps float32           Queries per second limit for the K8s client
      --k8s-heartbeat-timeout duration   Configures the timeout for api-server heartbeat, set to 0 to disable (default 30s)
      --k8s-kubeconfig-path string       Absolute path of the kubernetes kubeconfig file
```

### SEE ALSO

* [cilium-agent completion](cilium-agent_completion.md)	 - Generate the autocompletion script for the specified shell
