# BGDK

**bearaujus golang development kit** or `bgdk` is the development kit
with the main aim is making the development easier and testable.

# INSTALLATION


```shell
go get "github.com/bearaujus/bgdk@v1.0.1"
```

# PACKAGES

> ### [WORKER](https://github.com/bearaujus/bgdk/tree/master/worker)
> `worker` can execute many jobs asynchronously.
> This package also can do retries and use a custom error listener function
> to listen to the job when the job occurs an error.

&nbsp;

> ### [UTIL](https://github.com/bearaujus/bgdk/tree/master/util)
> `util` is an `bgdk` utilities packages.

- [UTIL/JSON](https://github.com/bearaujus/bgdk/tree/master/util/json)

`json` is utilities for the JSON files. This package implementing core functions from [encoding/json](https://cs.opensource.google/go/go/+/master:/src/encoding/json/).
This package also has some additional I/O features such as JSON writer and JSON reader.

&nbsp;

- [UTIL/YAML](https://github.com/bearaujus/bgdk/tree/master/util/yaml)

`yaml` is utilities for the YAML files. This package implementing core functions from [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3).
This package also has some additional I/O features such as YAML writer and YAML reader.

&nbsp;

- [UTIL/PTRCONV](https://github.com/bearaujus/bgdk/tree/master/util/ptrconv)

`ptrconv` is utilities for converting an object to object ptr (*object) and vice versa. 
This package also can handle nil ptr object (*object) conversion.

&nbsp;

[Back to top](#bgdk) 
