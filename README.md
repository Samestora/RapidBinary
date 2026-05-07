# Rapid Binary

## Rapid Binary Server (rbserver)

A high-performance, concurrent replacement for `python -m http.server` written in Go. **rbserver** allows you to serve files instantly and receive uploads via a modern web interface.

### Features
- **Concurrent by Default**: Handles multiple transfers simultaneously without blocking.
- **Modern UI**: Prettified directory listings instead of plain text.
- **Upload Support**: Optional `--upload` flag enables a drag-and-drop web UI.
- **Zero Dependencies**: Compiles to a single static binary.
- **Logging**: Optional `--log` flag for real-time traffic monitoring.

### Installation

#### Using the Build Scripts
**Windows:**
```batch
build.bat build
```
**Linux/Mac**
```bash
chmod +x build.sh
./build.sh build
```

### Usage
#### Simple File Serving
Serve the current directory on port 8080:

```bash
./rbserver --mode serve --port 8080
```
#### Serving with Uploads
Enable the upload interface (accessible at /upload):

```bash
./rbserver --mode serve --port 8080 --upload
```
#### Full Configuration
```bash
./rbserver --mode serve --port 9000 --dir ./my-files --upload --log
```

## Project Structure :
cmd/rbserver: Entry point and CLI flag handling.  
internal/server: Core HTTP logic and directory rendering.  
ui/: Embedded HTML templates and CSS.  
build/: Destination for compiled binaries.  

## License
[GNU GPL V3](./LICENSE)
