# randomness Specification

## Purpose
TBD - created by archiving change init-wails-app. Update Purpose after archive.
## Requirements
### Requirement: 6/49 Rules Compliance
The generated tickets MUST adhere to standard 6/49 lottery rules.

#### Scenario: Valid Range
- **WHEN** a ticket is generated
- **THEN** all numbers are between 1 and 49 (inclusive)

#### Scenario: Unique Numbers
- **WHEN** a ticket is generated
- **THEN** all 6 numbers are distinct (no duplicates)

### Requirement: Cryptographic Randomness
The application MUST use a cryptographically secure random number generator.

#### Scenario: Secure Source
- **WHEN** numbers are picked
- **THEN** `crypto/rand` (or equivalent secure source) is used
- **AND** `math/rand` is NOT used for number generation

