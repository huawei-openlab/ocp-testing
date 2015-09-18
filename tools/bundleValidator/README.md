The `scv` verifies:

## spec/bundle
Verify whether a bundle is valid, with all the required files and
all the required attributes.
Since the `bundle` spec is not decide yet, now just check if `config.json`,
`runtime.json` and `rootfs` were accessbile.

The validation work is done by .go files in the `libsv` directory.
These .go files follows the .go files in [specs](https://github.com/opencontainers/specs) closely
in order to make the validation clearly, for example:

```
spec/config.go

// Spec is the base configuration for the container.  It specifies platform
// independent configuration.
type Spec struct {
	// Version is the version of the specification that is supported.
	Version string `json:"version"`
	// Platform is the host information for OS and Arch.
	Platform Platform `json:"platform"`
	// Process is the container's main process.
	Process Process `json:"process"`
	// Root is the root information for the container's filesystem.
	Root Root `json:"root"`
	// Hostname is the container's host name.
	Hostname string `json:"hostname"`
	// Mounts profile configuration for adding mounts to the container's filesystem.
	MountPoints []MountPoint `json:"mounts"`
}
```

```
libsv/config.go

func SpecValid(s specs.Spec, msgs []string) (bool, []string) {
        valid, msgs := checkSemVer(s.Version, msgs)

        ret, msgs := PlatformValid(s.Platform, msgs)
        valid = ret && valid

        ret, msgs = ProcessValid(s.Process, msgs)
        valid = ret && valid

        ret, msgs = RootValid(s.Root, msgs)
        valid = ret && valid

        ret, msgs = StringValid("Spec.Hostname", s.Hostname, msgs)
        valid = ret && valid

        if len(s.MountPoints) > 0 {
                for index := 0; index < len(s.MountPoints); index++ {
                        ret, msgs = MountPointValid(s.MountPoints[index], msgs)
                        valid = ret && valid
                }
        }
        return valid, msgs
}
```

#How To try
It is easy to use this tool, we provide a demo-bundle.


```
make
./scv bundle demo-bundle
./scv config demo-bundle/config.json
./scv runtime demo-bundle/runtime.json linux

```

Also we add a simple config.json/runtime.json generator tool.

```
./scv  -o config.json genconfig
./scv  -o runtime.json genruntime
```

We can use the validation tool to check these json file.
```
./scv c config.json
./scv r runtime.json
```
