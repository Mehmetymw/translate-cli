# Translate CLI

Translate CLI is a simple command-line application that uses the Google Translate API to translate text. The application supports setting default languages, translating text, and performing reverse translations.

## Features

- Text translation
- Setting default languages
- Reverse translation

## Usage

### Download Precompiled Binaries

You can download the precompiled binaries and use them directly.

- [Linux](https://github.com/Mehmetymw/translate-cli/releases/download/v0.1/translate-cli-linux)
- [Windows](https://github.com/Mehmetymw/translate-cli/releases/download/v0.1/translate-cli-windows.exe)
- [macOS](https://github.com/Mehmetymw/translate-cli/releases/download/v0.1/translate-cli-macos)

#### Linux

1. Download the file and make it executable:
    ```sh
    chmod +x translate-cli-linux
    ```

2. Run the application:
    ```sh
    ./translate-cli-linux <commands and arguments>
    ```

#### Windows

1. Download the file and run it directly:
    ```sh
    translate-cli-windows.exe <commands and arguments>
    ```

#### macOS

1. Download the file and make it executable:
    ```sh
    chmod +x translate-cli-macos
    ```

2. Run the application:
    ```sh
    ./translate-cli-macos <commands and arguments>
    ```

### Commands

#### Setting Default Languages

Set the default source and target languages:

```sh
translate-cli -set -s <source> -t <target>
ex: translate-cli -set -s en -t tr (English to Turkish)
translate-cli <text>
ex: translate-cli Hello (It will remember setted config and translate English to Turkish directly)
translate-cli -reverse <text>
ex: translate-cli Merhaba (It will remember your setted config and will reversed translate Turkish to English directly)

