# WörterLab
Ein Deutsch-Übersetzungsassistent für Sprachenlernende

## Table of Contents

1. [How to install](#install)
2. [How to use](#start)
3. [Stacks and libraries](#stacks)

## <a name="install">How to install</a>
    
You will need [Python3](https://www.python.org/) and [LibreTranslate](https://github.com/LibreTranslate/LibreTranslate) installed.

<details>
<summary> LibreTranslate instalation </summary>

```sh
sudo apt-get install python3-pip python3-venv python3-dev
```

Go to the place where you want to install LibreTranslate.

```sh
python3 -m venv venv-translate
source venv-translate/bin/activate
pip install libretranslate
libretranslate [args]
```

This will start libretranslate at http://localhost:5000.
You can check the available args [here](https://github.com/LibreTranslate/LibreTranslate?tab=readme-ov-file#configuration-parameters).
</details>

## <a name="start">How to use</a>

TBD

## <a name="stacks">Stacks and libraries</a>

- [LibreTranslate](https://github.com/LibreTranslate/LibreTranslate)
- [Golang](https://go.dev/)
- [NPM](https://nodejs.org/en/download/)
- [Python3](https://www.python.org/)
- [Redis](https://redis.io/)
- [React](https://reactjs.org/)