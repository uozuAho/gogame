See the Makefile for how to run stuff

# quick start
- install go
```sh
sudo apt install libc6-dev libgl1-mesa-dev libxcursor-dev \
  libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev \
  libasound2-dev pkg-config

make test
make run
```

# troubleshooting
If you're having GL issues (old hardware?), try:

    LIBGL_ALWAYS_SOFTWARE=1 go run ./cmd/mygame

# todo
- fix tests/build etc.
- keep going with plan ... one of them
