## Context
The first release attempt exposed gaps in how Wails build artifacts are handled by our CI. We need to transition from "building binaries" to "packaging applications".

## Decisions
- **MacOS**: We will target `.zip` containing the `.app` bundle.
    - *Arch*: ARM64 (as requested).
- **Windows**: We will target `.zip` containing the `.exe` (or installer).
    - *Arch*: AMD64.
- **Linux**: We will target `.tar.gz` archives.
    - *Arch*: AMD64.

## Packaging Strategy & Naming
We will use the following naming pattern: `AppName-Version-OS-Arch.ext`.

- **Mac**: `LotteryPicker-X.Y.Z-macos-arm64.zip`
- **Win**: `LotteryPicker-X.Y.Z-windows-amd64.zip`
- **Linux**: `LotteryPicker-X.Y.Z-linux-amd64.tar.gz`

## GitHub Packages vs Release Assets
- We will stick to **Release Assets**.
