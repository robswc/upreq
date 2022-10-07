![upreq-logo(1)](https://user-images.githubusercontent.com/38849824/193359986-a5d2c43b-8e45-4456-8305-b9041388578f.png)

### Why? Because...
```python
len(pip freeze > requirements.txt) > len(upreq)
```

<!-- ![GitHub file size in bytes](https://img.shields.io/github/size/robswc/upreq?style=for-the-badge) -->
[![twitter](https://img.shields.io/badge/follow-@robswc-blue?style=for-the-badge&logo=twitter)](https://twitter.com/robswc)
<!-- ![downloads](https://img.shields.io/github/downloads/robswc/upreq/total?style=for-the-badge) -->


# Upreq

Upreq is a simple CLI tool that provides shortcuts and feedback for updating your Python project's `requirements.txt` file.
It's written in Go and uses [Cobra](https://github.com/spf13/cobra) as the CLI framework.
It started as a bash script, after I found myself typing `pip freeze > etc` one too many times and well, here we are.

## Features

- Comparing your `requirements.txt` file to your current environment
- Reduces carpel tunnel by turning `pip freeze > requirements.txt` into `upreq`
- Provides feedback on what packages were added/removed
- Flag for automatically adding new requirements to git

## Installation

Since Upreq is a binary executable, you can download the latest release from the [releases page](https://github.com/robswc/upreq/releases) and copy it to your `PATH`.
Below is an example of how to do this on each major OS.

### Linux
```bash
curl -s https://api.github.com/repos/robswc/upreq/releases/latest | grep "browser_download_url.*upreq-linux" | cut -d : -f 2,3 | tr -d \" | wget -qi -
chmod +x upreq-linux
sudo mv upreq-linux /usr/local/bin/upreq
```

### MacOs
```bash
curl -s https://api.github.com/repos/robswc/upreq/releases/latest | grep "browser_download_url.*upreq-darwin" | cut -d : -f 2,3 | tr -d \" | wget -qi -
chmod +x upreq-darwin
sudo mv upreq-darwin /usr/local/bin/upreq
```

### Windows (powershell)
```powershell
Invoke-WebRequest -Uri
"
https://api.github.com/repos/robswc/upreq/releases/latest
" | Select-Object -ExpandProperty content | ConvertFrom-Json | Select-Object -ExpandProperty assets | Select-Object -ExpandProperty browser_download_url | Select-String -Pattern "upreq-windows" | Invoke-WebRequest -OutFile upreq-windows.exe
mkdir "C:\Program Files\upreq"
mv .\upreq-windows.exe "C:\Program Files\upreq\upreq.exe"
$env:Path += ";C:\Program Files\upreq"
```



## Usage

Usage is pretty simple. Just run `upreq` in your project's root directory.

_Note: Be sure to activate your virtual environment before running._

Running just upreq will run `pip freeze` and compare it to your current `requirements.txt` file.
If there are any differences, it will print them out with a `+` or `-` to indicate if it was added or removed.
Finally, it will write any new requirements to your `requirements.txt` file.

### Basic Usage



```bash
upreq
```

### Advanced Usage

```bash
upreq --git  # Automatically add new requirements to git
```

### Help

```bash
upreq --help
```
