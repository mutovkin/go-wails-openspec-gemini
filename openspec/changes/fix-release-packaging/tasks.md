## 1. Investigation
- [x] 1.1 Verify Wails documentation on `.dmg` and NSIS installer generation commands and requirements.
- [x] 1.2 Check if "GitHub Packages" requirement refers to Docker containers or just better visibility of Release Assets. (Decision: Focus on Release Assets as zip/dmg/exe first, as that is standard for Desktop Apps).

## 2. Windows & MacOS Packaging
- [x] 2.1 Modify `.github/actions/build-wails/action.yml` to support platform-specific packaging flags or steps.
- [x] 2.2 Update `.github/workflows/release.yml` to upload the renamed artifacts.

## 3. Linux Packaging
- [x] 3.1 Explicitly add Linux to the build matrix in `release.yml`.
- [x] 3.2 Ensure `.AppImage` is generated (via `wails build -package` if supported or `nfpm`). If not, fallback to `.zip` or `.tar.gz`. Rename to `LotteryPicker-vX.Y.Z-linux-amd64.AppImage` (or extension).

## 4. Verification
- [x] 4.1 Validate workflows with `actionlint`.