## MODIFIED Requirements

### Requirement: Artifact Generation
The release process MUST generate and publish installable application artifacts with consistent naming.

#### Scenario: Build Artifacts
- **WHEN** a release is created
- **THEN** a macOS artifact is generated as `LotteryPicker-{version}-macos-arm64.zip`
- **AND** a Windows artifact is generated as `LotteryPicker-{version}-windows-amd64.zip`
- **AND** a Linux artifact is generated as `LotteryPicker-{version}-linux-amd64.AppImage` (or `.zip`)
- **AND** these artifacts are attached to the GitHub Release
