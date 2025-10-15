# Plan: Implement Bullet Damage to Dudes

## Goal
Implement a damage system where Dudes have hit points that are reduced when bullets hit them.

## Current State Analysis
- Bullets move in a straight line at BulletSpeedPerTick (20 units/tick)
- Dudes can move and shoot bullets
- No collision detection exists between bullets and dudes
- No health/damage system exists
- Game has an event system (EventEmitter) already in place

## Incremental Implementation Steps
- Add Distance Calculation Helper to Point2D
  - Add `DistanceTo(other Point2D) float64` method to Point2D in `internal/game/point2d.go`
  - This will calculate the Euclidean distance between two points
- Add Hit Points to Dude
  - Add `HitPoints int` field to Dude struct in `internal/game/dude.go`
  - Add `MaxHitPoints int` field (for reference/display)
  - Initialize these values in `NewGame()` in `internal/game/game.go` (e.g., 100 HP each)
- Add Collision Detection Between Bullets and Dudes
  - Add `CheckCollisions()` method to Game in `internal/game/game.go`
  - Define a collision radius (e.g., 20 pixels for dudes)
  - Loop through bullets and dudes, check if distance < collision radius
  - For now, just emit a collision event (don't remove entities yet)
  - Call `CheckCollisions()` in `Game.Update()` after entities update
- Apply Damage on Collision
  - Add `Damage int` field to Bullet struct in `internal/game/bullet.go` (e.g., 10 damage)
  - In collision detection, apply bullet damage to dude's HitPoints
  - Emit a damage event with relevant data (which dude, how much damage)
  - play sound when bullets hit
- Remove Dead Dudes and Used Bullets
  - Mark bullets for removal after collision in `CheckCollisions()`
  - Mark dudes for removal when HitPoints <= 0
  - Add entity removal logic to `Game.Update()` - filter Entities slice
  - Emit death event when dude dies
- Add Bullet Owner Tracking (Optional but Recommended)
  - Prevent bullets from damaging the dude that shot them.
  - Add `Owner *Dude` or `OwnerID string` field to Bullet
  - Set owner when creating bullet in `Dude.Update()`
  - In collision detection, skip collisions between bullet and its owner
- Polish? - Add Visual/Audio Feedback (Optional)
  - Emit events for hit/damage/death that GUI can use
  - Consider adding invulnerability frames after being hit
  - Consider bullet lifetime/max range
