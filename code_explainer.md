Hello humans and robots.

This is a RTS/tower-defense style game, written in go, using ebiten as the game
engine.

The ultimate goals are:
- can run headless: run games (quickly) with no UI
- deterministic: the same initial conditions and sequence of inputs will always
  result in the same result
- user can modify game units' behaviour with scripts

Code structure and rules:
- the game starts from the cli in cmd/mygame
- all code using the ebiten game engine goes under internal/gui. Nothing depends
  on this code other than the cli.

Help for writing code:
- ebiten API docs: https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2
- use make to test, build and format code. See Makefile. Do not run custom
  commands unless what you need is not in the Makefile.
