# git-workflow Specification

## Purpose
TBD - created by archiving change setup-github-actions. Update Purpose after archive.
## Requirements
### Requirement: Git Flow Branching
The repository MUST follow a modified Git Flow branching strategy.

#### Scenario: Branch Roles
- **WHEN** developing features, **THEN** use `feature/*` branches merging into `develop`.
- **WHEN** preparing a standard release, **THEN** `develop` is merged to `master`.
- **WHEN** hotfixing, **THEN** use `hotfix/*` branches merging directly into `master`.

### Requirement: Commit Convention
All commits MUST adhere to the Conventional Commits v1.0.0 specification.

#### Scenario: Commit Message Format
- **WHEN** a commit is created
- **THEN** the message must start with a type (feat, fix, chore, etc.)
- **AND** include a description of the change
- **AND** include "Why" and "How" details in the body for non-trivial changes

### Requirement: Linear History
The `master` and `develop` branches MUST maintain a linear history.

#### Scenario: Merge Strategy
- **WHEN** merging a feature branch into `develop`
- **THEN** the branch is squashed into a single commit
- **WHEN** merging `develop` into `master` for release
- **THEN** a fast-forward merge is performed

### Requirement: History Synchronization
Deviations between `master` and `develop` MUST be reconciled automatically where possible.

#### Scenario: Hotfix Rebase
- **WHEN** a hotfix is merged and released on `master`
- **THEN** `develop` is automatically rebased on top of the new `master` commit

#### Scenario: Feature Rebase
- **WHEN** a feature branch is ready for merge
- **THEN** it SHOULD be rebased onto the latest `develop` to ensure a clean squash merge

### Requirement: Repository Configuration
The repository configuration MUST enforce the workflow.

#### Scenario: Branch Protection
- **WHEN** configuring the repository
- **THEN** `master` and `develop` MUST require Pull Requests
- **AND** Status Checks (CI) MUST pass before merging
- **AND** Linear history MUST be required

