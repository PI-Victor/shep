Go Trello API
================

[![Trello Logo](https://raw.githubusercontent.com/adlio/trello/master/trello-logo.png)](https://www.trello.com)

[![GoDoc](https://godoc.org/github.com/adlio/trello?status.svg)](http://godoc.org/github.com/adlio/trello)
[![Build Status](https://travis-ci.org/adlio/trello.svg)](https://travis-ci.org/adlio/trello)
[![Coverage Status](https://coveralls.io/repos/github/adlio/trello/badge.svg?branch=master)](https://coveralls.io/github/adlio/trello?branch=master)

A #golang package to access the [Trello API](https://www.trello.com/api). Nearly 100% of the
read-only surface area of the API is covered, as is creation and modification of Cards. Pull
requests are welcome for missing features.

My current development focus is documentation, especially enhancing this README.md with more
example use cases.

## Installation

The Go Trello API has been Tested compatible with Go 1.1 on up. Its only dependency is
the `github.com/pkg/errors` package. It otherwise relies only on the Go standard library.

```
go get github.com/adlio/trello
```

## Basic Usage

All interaction starts with a `trello.Client`. Create one with your appKey and token:

```Go
client := trello.NewClient(appKey, token)
```

All API requests accept a trello.Arguments object. This object is a simple
`map[string]string`, converted to query string arguments in the API call.
Trello has sane defaults on API calls. We have a `rello.Defaults` object
which can be used when you desire the default Trello arguments. Internally,
`trello.Defaults()` is an empty map, which translates to an empty query string.

```Go
board, err := client.GetBoard("bOaRdID", trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Trello Boards for a User

Boards can be retrieved directly by their ID (see example above), or by asking
for all boards for a member:

```Go
member, err := client.GetMember("usernameOrId", trello.Defaults())
if err != nil {
  // Handle error
}

boards, err := member.GetBoards(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Trello Lists on a Board

```Go
board, err := client.GetBoard("bOaRdID", trello.Defaults())
if err != nil {
  // Handle error
}

lists, err := board.GetLists(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Trello Cards on a Board

```Go
board, err := client.GetBoard("bOaRdID", trello.Defaults())
if err != nil {
  // Handle error
}

cards, err := board.GetCards(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Trello Cards on a List

```Go
list, err := client.GetList("lIsTID", trello.Defaults())
if err != nil {
  // Handle error
}

cards, err := list.GetCards(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Actions on a Board

```Go
board, err := client.GetBoard("bOaRdID", trello.Defaults())
if err != nil {
  // Handle error
}

actions, err := board.GetActions(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Actions on a Card

```Go
card, err := client.GetCard("cArDID", trello.Defaults())
if err != nil {
  // Handle error
}

actions, err := card.GetActions(trello.Defaults())
if err != nil {
  // Handle error
}
```
