# Vidéothèque

## Usage

### Manual Movie Entry

```shell
vt <movie entry>
```

#### Support

- Magnet URI Link (e.g. magnet:?xt=urn:btih:\<hash\>)
- Local Video File (e.g. ~/Movies/\<movie\>.mp4)

### Manual Subtitle Entry

```shell
vt <movie entry> --subtitle <subtitle entry>
```

#### Support

- Local `.srt` or `.vtt` file (e.g. ~/Movies/\<movie\>.srt)

### Additional CLI Options

```shell
vt help
```

#### --config \<path\> or -c \<path\>

Set custom path to a `Config file` described above.

#### --language \<lang-code\> or -l \<lang-code\>

Set a language code, e.g. en-us, fr, pt-br and so on. The default is `en`.

#### --qrcode or -q

Show a QR Code that redirects to the settings page.

Useful for configuring the app on the phone while media plays on television.

## Automatic Subtitle Download

*Requires a `Config file` with Open Subtitles credentials*

Vidéothèque verifies if the movie to be watched is a Foreign language movie and Downloads a subtitle if needed.

## Config File

Set once properties can be grouped on a `.toml` file. Vidéothèque has sensible defaults and the Config file is completly optional.

The program follows the [XDG Base Directory Specification](https://xdgbasedirectoryspecification.com/) and tries to find the file: `$XDG_CONFIG_HOME/videoteque/config.toml`.

### Sample

```toml
language = "fr"

[open-subtitles]
username = "my-username"
password = "my-password"

[server]
port = 1201
```

## Install

### Homebrew

*MacOS or Linux only*

```shell
brew tap yassernasc/videoteque https://github.com/yassernasc/videoteque.git
brew install videoteque
```

### Build From Source

*Requires Golang and Node.js to be installed*

```shell
make install
```
