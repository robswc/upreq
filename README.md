# upreq
Tiny requirements.txt updater utility

### About
upreq simply calls `pip freeze > requirements.txt` and adds the new requirements.txt file to git!
It also prints the difference between the old requirements.txt and the new requirements.txt.  Future updates will probably include generating commit messages/other helpful features.

### Installation

The following command will:
- download the file from github
- move it to bin
- change permission to allow for execution

```
sudo wget https://raw.githubusercontent.com/robswc/upreq/main/upreq;sudo mv upreq /bin;sudo chmod +x /bin/upreq
```

### Usage
Call `upreq` in your project's root directory.
