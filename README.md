To set up the varlink run `podman varlink --timeout=0 unix:/run/user/<uid>/podman/io.podman` for rootless and `sudo podman varlink --timeout=0 unix:/run/podman/io.podman` for root.
