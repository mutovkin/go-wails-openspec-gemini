## 1. CI Workflow
- [x] 1.1 Create `.github/actions/build-wails/action.yml` composite action to encapsulate Wails setup and build steps (DRY).
    - [x] 1.1.1 Implement dependency caching (Go/Node).
- [x] 1.2 Create `.github/workflows/ci.yml` to run on PRs targeting `develop` and `master` (hotfixes).
    - [x] 1.2.1 Configure explicit PR comments/annotations for failures (formatting, linting, tests).
- [x] 1.3 Implement `golangci-lint` step using the official action (pinned).
- [x] 1.4 Implement `go test` step.
- [x] 1.5 Implement `wails build` step using the custom composite action.

## 2. Release Workflow
- [x] 2.1 Create `.github/workflows/release.yml` triggered by `workflow_dispatch` with version input.
- [x] 2.2 Implement semantic version bumping and tagging logic.
- [x] 2.3 Implement fast-forward merge logic from `develop` to `master` for standard releases.
- [x] 2.4 Implement hotfix logic: Tag `master` directly and auto-rebase `develop` onto `master` after release.
- [x] 2.5 Configure Wails build for macOS and Windows artifacts using custom action.
- [x] 2.6 Configure GitHub Release creation and artifact upload.
- [x] 2.7 Verify action pinning for all used external actions.
- [x] 2.8 Create Makefile rules to validate GitHub Actions (`make lint-actions`).

## 3. Documentation & Governance
- [x] 3.1 Update `openspec/project.md` with the new Git Flow, branching, rebase strategy, and commit conventions.
- [x] 3.2 Create `design/release-flow.md` with Mermaid diagrams illustrating standard and hotfix release flows.
- [x] 3.3 Add "Repository Configuration" section to documentation covering Branch Protection rules (Require PR, Status Checks, Linear History).
- [x] 3.4 Update `README.md` with workflow details.