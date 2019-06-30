[![Go Report Card](https://goreportcard.com/badge/github.com/durandj/git-stats)](https://goreportcard.com/report/github.com/durandj/git-stats)
[![GitHub](https://img.shields.io/github/license/durandj/git-stats.svg)](https://github.com/durandj/git-stats/blob/master/License.md)

# git-stats

A CLI tool for generating stats on Git repositories.

Have you ever wanted to know statistics about your Git repository?
Things like who's committed to it, how long PR's take to approve, that
kind of thing? That's what this tool hopes to provide.

Stats are as easy to generate as running:

`git stats`

## Install

nstallation should be pretty straight forward. Just head on over to
the [releases page](https://github.com/durandj/git-stats/releases)
and download the latest version for your desired platform, mark it
as executable (if on a Unix type system) and put it somewhere on your
`PATH`.

For example, if I wanted to install version `latest` (this isn't a
real version) for 64 bit Linux I would do:

```
wget https://github.com/durandj/git-stats/releases/download/vLatest/git-ignore_vLatest_linux_amd64
chmod +x git-ignore_vLatest_linux_amd64
sudo mv git-ignore_vLatest_linux_amd64 /usr/local/bin
```

## Developing

Make sure first install the following dependencies:

### Dependencies

 * Golang 1.12+ (I prefer using [goenv](https://github.com/syndbg/goenv) for this)
 * [Ginkgo](http://onsi.github.io/ginkgo)
 * [GolangCI-Lint](https://github.com/golangci/golangci-lint)
 * [Taskfile](https://taskfile.org)

### Tasks

Taskfile provides a way of running different scripts easily (similar
to how Makefiles work but better).

To build the code you can just do `task build`.
To run tests run `task test`.
To lint the code run `task lint`.
