# Change: Setup GitHub Actions and Git Workflow

## Why
To establish a robust Continuous Integration and Continuous Deployment (CI/CD) pipeline and enforce a structured Git workflow. This automation will ensure code quality, streamline releases, and provide downloadable artifacts for users, while maintaining a clean and linear history for the main branches.

## What Changes
- **CI Pipeline**:
    - Automatic checks on Pull Requests (linting, testing, building).
    - Uses `golangci-lint` GitHub Action.
    - **Feedback**: Explicit PR feedback (comments/status checks) for lint/test/build results.
- **Release Pipeline**:
    - Manual workflow dispatch to trigger semantic version releases (Major, Minor, Patch).
    - Auto-tagging and fast-forward merging from `develop` to `master`.
    - Hotfix workflow: Merge hotfix branches directly to `master` (squashed), release patch version, and auto-rebase `develop` onto `master`.
    - Building and uploading artifacts for macOS and Windows.
    - **Refactoring**: Common build steps extracted into custom composite actions (DRY).
    - **Optimization**: Dependency caching (Go modules, Node dependencies).
    - **Security**: Action pinning to SHA for stability.
- **Git Workflow**:
    - Formal adoption of Git Flow.
    - Enforcement of Linear History on `master` and `develop`.
    - Squash merging for feature/hotfix branches.
    - Conventional Commits requirement.
    - **Documentation**: Feature branch rebase strategy documented.
- **Documentation**:
    - Updates to `project.md` to reflect the new workflow and CI/CD/Release requirements.
    - New `design/release-flow.md` with Mermaid diagrams describing the process.
    - New section on GitHub Repository Configuration (Branch protection rules).

## Impact
- **New Capabilities**: `cicd`, `git-workflow`.
- **Project Configuration**:
    - New `.github/workflows` directory.
    - New `.github/actions` directory for custom actions.
    - Updated `openspec/project.md`.
    - New `design/release-flow.md`.
