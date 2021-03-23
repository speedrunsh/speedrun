<p align="center">
  <a rel="nofollow">
    <img src="docs/logo.png?raw=true" width="200" style="max-width:100%;">
  </a>
</p>


# Speedrun
[![license](https://img.shields.io/badge/license-MPL2-blue.svg)](https://github.com/dpogorzelski/speedrun/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/dpogorzelski/speedrun)](https://goreportcard.com/report/github.com/dpogorzelski/speedrun)
[![Go](https://github.com/dpogorzelski/speedrun/actions/workflows/go.yml/badge.svg)](https://github.com/dpogorzelski/speedrun/actions/workflows/go.yml)

Speedrun is a command execution framework that works at scale.
The point of Speedrun is so that you can run any command across any number of servers, projects and cloud vendors with ease and fast (currently GCP only but AWS and Azure will be supported as well), example:
`speedrun run systemctl stop nginx` to stop nginx across 3k machines.

No hassels with setting up and maintaining a server with agents as speedrun has none. No runtime to install/maintain as it's a single self-contained binary. Speedrun leverages SSH as transport with tight cloud vendor integration to take the burden of mundane things like key generation and propagation/revocation away from the user.

Server selection is made as intuitive as possible, currently based on [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) but in the future will be replaced by a generic selection mechanism to provide a seamless experience across different cloud vendors.


Features:

* native cloud integration (currently Google Cloud only, AWS and Azure coming up!)
* stateless and agentless
* no complex configuration required
* single self-contained binary
* can run against any number of servers and projects seamlessly
* a plugin system is in the plan to allow anyone to integrate execution modules that will wrap complex functionality instead of running raw shell commands


## Installation

Download the precompiled binary from here: [Releases](https://github.com/dpogorzelski/speedrun/releases)

## Usage

```bash
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/serviceaccount.json
speedrun init
speedrun key new
speedrun key set
speedrun run whoami
```

## Examples

Use 1000 concurrent SSH workers

```bash
speedrun run "uname -r" --concurrency 1000
```

Stop Nginx on VMs matching the filter

```bash
speedrun run sudo systemctl stop nginx --filter "labels.env=staging AND labels.app=foobar"
```

Ignore SSH fingerprint mismatch and connect via private IP addresses

```bash
speedrun run "ls -la" --filter "labels.env != prod" --ignore-fingerprint --concurrency 1000 --use-private-ip
```

Use a different config file

```bash
speedrun run whoami -c /path/to/config.toml
```

Set public key on specific instances instead of project metadata (useful if instances are blocking project wide keys):

```bash
speedrun key set --filter "labels.env = dev"
```



## Configuration

Using certain flags repeteadly can be annoying, it's possible to persist their behavior via config file. Default config file is located at `~/.speedrun/config.toml` and can be re-initialized to it's default form via `speedrun init`.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Project Status

This project is in a very early stage so expect a lot of breaking changes in the nearest future

## License

[MPL-2.0](LICENSE)