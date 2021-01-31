# Speedrun

Speedrun is a cloud first command execution utility

## Installation

Download the precompiled binary from here

```bash
curl https://github.com/dawidpogorzelski/speedrun/releases/0.1.0/release
mv speedrun /usr/local/bin
```

## Usage

```bash
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/serviceaccount.json
speedrun init
speedrun key new
speedrun key set
speedrun run whoami
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)