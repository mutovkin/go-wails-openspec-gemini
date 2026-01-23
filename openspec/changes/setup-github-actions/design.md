## Context
We need a professional-grade development workflow for a public Go/Wails project. The goal is to balance rigorous quality checks with automated ease of release.

## Goals / Non-Goals
- **Goals**:
    - Automate semantic versioning.
    - Ensure `master` is always stable and points to the latest release.
    - Enforce code quality via CI with explicit feedback.
    - Keep workflows DRY using composite actions.
    - Secure and optimized workflows (pinned actions, caching).
- **Non-Goals**:
    - Automated deployment to app stores (Mac App Store, Microsoft Store) is out of scope for now; GitHub Releases is the target.
    - Strict deletion of previous release artifacts (Github Actions retention policy is sufficient).

## Decisions
- **Git Flow Adaptation**:
    - `develop` is the default branch for feature PRs.
    - `master` is the stable release branch.
    - **Standard Release**: `develop` is merged to `master` (fast-forward) and tagged.
    - **Hotfix Release**: Hotfix branches merge directly to `master` (squash). A bot then rebases `develop` onto the new `master`.
- **Feature Development**:
    - Feature branches should be rebased on `develop` before merge (enforced via documentation/hooks if possible, but primarily cultural).
- **CI/CD Architecture**:
    - **Composite Actions**: Build logic encapsulated in `.github/actions/build-wails` for reuse.
    - **Dependencies**: Caching enabled for Go modules and Node packages to speed up builds.
    - **Security**: All third-party actions pinned to commit SHAs.
- **Artifacts**:
    - We build for macOS (universal or arm64/amd64) and Windows (amd64).

## Risks / Trade-offs
- **Complexity**: Hotfix workflow with auto-rebase adds complexity.
    - *Mitigation*: The release workflow will handle this logic explicitly.
- **Rebase Conflicts**: Auto-rebasing `develop` might fail.
    - *Mitigation*: Notification for manual intervention.
