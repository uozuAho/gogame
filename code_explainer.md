Hello humans and robots.

This is a RTS/tower-defense style game, written in go, using ebiten as the game
engine.

The ultimate goals are:
- can run headless: run games (quickly) with no UI
- user can modify game units' behaviour with scripts

Code structure and rules:
- the game starts from the cli in cmd/mygame
- all code using the ebiten game engine goes under internal/gui. Nothing depends
  on this code other than the cli.
