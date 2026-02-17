# Agents Guide

Goal: ship an interview-ready CRUD REST API quickly, with clean, idiomatic Go.

## Defaults

- Prefer the standard library for HTTP and JSON unless asked otherwise.
- Keep handlers small and focused. Validate input early.
- Return consistent error payloads and status codes.
- Keep validation, storage, and HTTP concerns separate.

## Clean Code Rules

- Small functions, clear names, no hidden side effects.
- Avoid unnecessary abstractions. Add interfaces only when they simplify tests or swapping dependencies.
- Favor composition over inheritance.

## Go Proverbs (Selected)

- Don't communicate by sharing memory; share memory by communicating.
- Errors are values.
- The bigger the interface, the weaker the abstraction.
- Make the zero value useful.

## 100 Go Mistakes Reminders

- Always check errors.
- Avoid silent shadowing.
- Don't ignore context cancellation in long operations.
- Use timeouts on IO.
- Prefer `time.Time` in UTC for persistence and serialization.
- Guard against data races.
- Validate inputs, especially query params and JSON payloads.
