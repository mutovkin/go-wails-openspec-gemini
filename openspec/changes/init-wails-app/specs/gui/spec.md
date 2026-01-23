## ADDED Requirements

### Requirement: Ticket Generation UI
The User Interface MUST provide a mechanism to generate and display lottery tickets.

#### Scenario: Generate Button
- **WHEN** the application starts
- **THEN** a "Generate Ticket" button is visible
- **AND** the button is enabled

#### Scenario: Display Ticket
- **WHEN** the "Generate Ticket" button is clicked
- **THEN** 6 numbers are displayed in ascending order
- **AND** the numbers are clearly legible
