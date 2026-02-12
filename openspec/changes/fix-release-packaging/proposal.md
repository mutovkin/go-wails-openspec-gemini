# Change: Fix Release Packaging and Artifacts

## Why
The initial release automation successfully built binaries but failed to package them correctly for distribution. MacOS apps were uploaded as raw binaries instead of .app/.dmg, Windows installers were missing, and artifacts were not published to GitHub Packages. Additionally, Linux builds were observed but no artifacts were produced.

## What Changes
- **MacOS Packaging**: 
    - Package the `.app` bundle into a `.zip` archive for distribution.
    - Naming convention: `LotteryPicker-vX.Y.Z-macos-arm64.zip`.
- **Windows Packaging**:
    - Ensure the NSIS installer (if generated) is uploaded.
    - Package into a `.zip` if installer generation fails or is optional.
    - Naming convention: `LotteryPicker-vX.Y.Z-windows-amd64.zip`.
- **Linux Packaging**:
    - Explicitly enable Linux artifact generation.
    - Generate `.AppImage` if possible, or `.tar.gz`.
    - Naming convention: `LotteryPicker-vX.Y.Z-linux-amd64.AppImage` (or `.zip`).
- **GitHub Packages**:
    - Focus on Release Assets with correct naming.
- **Workflow**:
    - Update `.github/workflows/release.yml` and `.github/actions/build-wails` to handle platform-specific packaging and renaming logic.

## Impact
- **Specs**: `cicd` capability updated with specific packaging requirements.
- **Workflows**: `release.yml` updated.
- **Actions**: `build-wails/action.yml` updated.
