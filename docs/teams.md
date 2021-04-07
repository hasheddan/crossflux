# Teams

The Crossflux organization is divided into teams. The following sections
describe the roles and responsibilities of each team, as well as how new teams
can be added.

## Existing Teams

### Infra

The `infra` team is responsible for building the platform, managing org-level
infrastructure, and onboarding all other teams (see [Adding a New
Team](#adding-a-new-team)).

### Team-1

Team-1 develops and runs the Crossflux Todo app.

## Adding a New Team

When a new team needs to be onboarded to the Crossflux platform, the `infra`
team will execute the following steps in a PR to this repo:

1. Create a new directory under `teams/`.
2. Add directory to top-level `teams/kustomization.yaml`.
3. Add a team-level `kustomization.yaml` that enforces any required overlays,
   such as default `namespace`.
4. Add an entry to this doc describing the purpose of the team.
