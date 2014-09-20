
# Marlin Architecture

## Positioning

Marlin is an orchestrator, not an inference engine. It sits in front of a
target LLM endpoint and a set of draft endpoints, decides *what* to propose,
*when* to admit a proposal, and *how long* the proposals should be - then
lets the target verify.
