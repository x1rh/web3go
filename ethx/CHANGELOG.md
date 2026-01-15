## ethx/v0.0.9
Refactor:
- Refactor all managers; each manager now handles a single-chain entity. Multi-chain should run as separate programs in different processes.
- Rename app to protocol.
- Remove unnecessary utils.
- Refactor client.
  - Add: an eth.Client URL must be either HTTP or WSS.
  - Keep Client minimal; prefer its own methods and avoid unnecessary wrappers.

