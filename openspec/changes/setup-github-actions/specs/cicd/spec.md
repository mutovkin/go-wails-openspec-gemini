## ADDED Requirements

### Requirement: Pull Request Verification
All changes destined for the `develop` branch MUST pass automated verification checks before merging.

#### Scenario: Verify Code Quality
- **WHEN** a Pull Request is opened or updated targeting `develop`
- **THEN** `golangci-lint` is executed and must pass
- **AND** unit tests (`go test`) are executed and must pass
- **AND** the application builds (`wails build`) successfully
- **AND** detailed feedback (success/fail icons, comments) is provided on the PR

#### Scenario: Optimized Builds
- **WHEN** a workflow runs
- **THEN** project dependencies (Go modules, npm packages) are cached
- **AND** cached dependencies are restored if lockfiles match

### Requirement: Release Automation
The system MUST provide an automated mechanism to release new versions of the application.

#### Scenario: Trigger Release
- **WHEN** a release workflow is manually triggered with a semantic version increment (Major, Minor, or Patch)
- **THEN** the next semantic version is calculated
- **AND** a git tag is created
- **AND** the code is merged from `develop` to `master` using fast-forward only
- **AND** a GitHub Release is created

### Requirement: Artifact Generation
The release process MUST generate and publish installable application artifacts.

#### Scenario: Build Artifacts
- **WHEN** a release is created
- **THEN** a macOS application bundle (.app or .dmg) is built
- **AND** a Windows executable (.exe or installer) is built
- **AND** these artifacts are attached to the GitHub Release

### Requirement: Workflow Security
All workflows MUST adhere to security best practices.

#### Scenario: Pin Actions
- **WHEN** using third-party GitHub Actions
- **THEN** the action MUST be referenced by a specific commit SHA, not a tag or branch