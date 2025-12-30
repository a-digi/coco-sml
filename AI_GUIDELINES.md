# AI Guidelines

These guidelines apply to all AI-assisted and automated code generation and review processes in this project.

## Core Principles

1. **SOLID Principles**
    - All generated or reviewed code must adhere to the SOLID principles:
        - **S**ingle Responsibility Principle (Each class/function has exactly one responsibility)
        - **O**pen/Closed Principle (Open for extension, closed for modification)
        - **L**iskov Substitution Principle (Subtypes must be substitutable)
        - **I**nterface Segregation Principle (Interfaces are specific and small)
        - **D**ependency Inversion Principle (Depend on abstractions, not concrete implementations)

2. **DRY Principle (Don't Repeat Yourself)**
    - Redundancies in code must be avoided.
    - Reusable logic should be encapsulated in functions, methods, or modules.
    - Copy-paste code must be avoided.

3. **Separation of Concerns**
    - Functional responsibilities must be clearly separated.
    - Each component, class, or function is responsible for a clearly defined aspect.
    - No mixing of UI, business logic, data access, or infrastructure code.

4. **Documentation and Traceability**
    - All AI-generated or modified code must be sufficiently documented (comments, docstrings, change logs).
    - The motivation for complex or unusual solutions must be explained.
    - AI-generated changes must be clearly marked in commit messages or pull requests.

5. **Testability and Tests**
    - AI-generated code must be testable (e.g., via unit or integration tests).
    - Where possible, automated tests for new or changed functions should be provided.
    - Existing tests must not break due to AI changes.

6. **Error and Exception Handling**
    - Error cases and exceptions must be handled explicitly.
    - No silent errors or unclear error messages should occur.

7. **Security and Data Protection**
    - AI-generated code must not process, store, or transmit sensitive data unprotected.
    - Security-relevant aspects (e.g., input validation, authentication) must be given special attention.

8. **Consistency and Style**
    - Generated code must adhere to the project's code style guides (formatting, naming, structure).
    - In case of uncertainty, use existing code as a reference.

9. **Transparency in Limitations**
    - If the AI cannot directly implement a requested change (e.g., due to lack of write permissions), it must always provide the complete relevant file content – never just code snippets.
    - **Important:** In such cases, it is expressly forbidden to provide only code snippets or excerpts. The complete relevant file content must always be provided.

## Application
- These principles apply to all AI-generated suggestions, automated refactorings, and reviews.
- In case of conflicting goals, always choose the best possible balance between readability, maintainability, and extensibility.
- Violations of these principles must be documented and justified.

---

*Last updated: 25.11.2025*
