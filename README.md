# obs-srt-bridge

A bridge that allows OBS Studio to display live **SRT** (Secure Reliable Transport) connection statistics and automatically switch scenes based on connection health.  
The project consists of a tiny Go server that embeds a React single-page application which is loaded in OBS as a _Browser Source_.

## Related Tools

This tool is designed to work together with the following tools. Please check them out as well:

- **[srt-live-reporter](https://github.com/e04/srt-live-reporter)** - A proxy that provides SRT statistics via WebSocket
- **[go-srtla](https://github.com/e04/go-srtla)** - A cross-platform SRTLA receiver

Also, be sure to check out [this gist](https://gist.github.com/e04/3914d98d6d0a55c689ab724ac6896081).

## Features

- **Live SRT Statistics Display**: Real-time monitoring of SRT connection quality and statistics.
- **Multiple Display Types**: Choose between simple text display, graph visualization, or no display.
- **Automatic Scene Switching**: Automatically switches OBS scenes based on connection health.
  - Switches to the "online" scene when SRT connection is stable and healthy.
  - Switches to the "offline" scene when connection is lost or experiencing issues.
  - Configurable scene names via URL parameters.
- **Reliable Scene Switching Based on Packet Loss**: Unlike tools that only monitor bitrate, this bridge utilizes detailed SRT statistics such as packet loss. This enables more intelligent and accurate switching decisions, correctly identifying connection health issues that simpler methods might miss.

## Download

You can download prebuilt binaries from the [Releases](https://github.com/e04/obs-srt-bridge/releases/).

## Run

```bash
$ ./obs-srt-bridge -port 9999
Browser Source URL: http://localhost:9999/app
```

Leave the process running while OBS is open.

### Command-line flags

| Flag    | Default | Description                                   |
| ------- | ------- | --------------------------------------------- |
| `-port` | `9999`  | HTTP port that serves the browser source HTML |

## Adding the Browser Source to OBS

1. In OBS Studio choose **+ → Browser Source**.
2. Enter the URL shown when you started the program, for example:
   - `http://localhost:9999/app?type=simple&wsport=8888&onlineSceneName=ONLINE&offlineSceneName=OFFLINE`
3. **For automatic scene switching**: If you are using the `onlineSceneName` and `offlineSceneName` parameters, you must set the **Page permission** to **Advanced access** in the Browser Source properties. This allows the application to control OBS scene switching.
   - Also, please **uncheck** `Restart playback when source becomes active` in your UDP/SRT ingest media source properties to prevent playback issues during scene switching.

All query parameters are optional:

| Parameter          | Default   | Meaning                                                                                                                    |
| ------------------ | --------- | -------------------------------------------------------------------------------------------------------------------------- |
| `type`             | `simple`  | `simple`, `graph` or `none`                                                                                                |
| `wsport`           | `8888`    | Port where your SRT statistics WebSocket is listening (use [srt-live-reporter](https://github.com/e04/srt-live-reporter) ) |
| `onlineSceneName`  | `ONLINE`  | Scene name to switch to when the connection is healthy                                                                     |
| `offlineSceneName` | `OFFLINE` | Scene name to switch to when the connection is lost or poor                                                                |

### `type` Details

| type   |                                                                                                                                                        |
| ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| simple | <img width="345" alt="スクリーンショット 2025-06-29 22 42 24" src="https://github.com/user-attachments/assets/ce8dd8b6-fb3b-44e8-aacc-f74f24d3b2b5" /> |
| graph  | <img width="347" alt="スクリーンショット 2025-06-29 22 41 47" src="https://github.com/user-attachments/assets/bd77524d-f5ae-43ce-84b9-616bca1e6110" /> |
| none   | (none, just for switching scene)                                                                                                                              |

**Metric Explanations (left to right):**

- Bitrate (Mbps)
- RTT (ms)
- Packet Loss (%)

## Build

```bash
# Clone
$ git clone https://github.com/e04/obs-srt-bridge.git
$ cd obs-srt-bridge

# 1. Build the Browser Source bundle
$ cd frontend
$ npm ci           # installs dependencies
$ npm run build    # outputs to frontend/dist
$ cd ..

# 2. Build the Go binary (this embeds the bundle produced above)
$ go build -o obs-srt-bridge
```

## License

This project is released under the terms of the **MIT License**. See the [LICENSE](LICENSE) file for details.
