# Agent Working Notes

This file summarizes how AI coding agents should work in this repository.

## Engineering Priorities

- Think like a backend architect: reliability, security, scalability, and performance first.
- Keep architecture clean: `transport -> service -> repository`.
- Prefer safe, auditable defaults (no hardcoded secrets, explicit migrations, conservative rollout).
- Preserve backward compatibility unless a breaking change is explicitly requested.

## Required Delivery Flow (Per Task)

1. Create a new feature branch from `main`.
2. Implement the task with tests/docs updates when relevant.
3. Run:
   - `make preprod`
   - `make docker-compose-up` (verify startup and logs)
   - `make docker-compose-down`
4. Commit with Conventional Commits.
5. Push the branch to origin.
6. Create or provide a PR link.
7. Delete the local feature branch.
8. Sync local `main`:
   - `git fetch --prune origin`
   - `git pull --ff-only`
9. Create and push a new annotated release tag.

## Git Safety

- Never commit or push directly to `main`.
- Delete only local branches unless explicitly instructed otherwise.
- Avoid destructive git operations unless explicitly requested.
