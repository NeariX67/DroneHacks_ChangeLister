# DroneHacks Change Lister

This tool outputs any changed parameters from the drone hacks parameter export file.
You can also specify regular expression to filter the name and index of the parameter.

## Getting started

Either go to the [releases page](https://github.com/NeariX67/DroneHacks_ChangeLister/releases) or clone the repo and build it yourself with ```go build```

## Usage

```
  -difference float
        Display parameters whose difference (Default and actual value) is greater than this (default 0.001)
  -filterIndex string
        Regular expression to filter for parameter indexes (optional)
  -filterName string
        Regular expression to filter for parameter names (optional)
  -i string
        path to file, usually has the extension DHP or JSON (required)
```