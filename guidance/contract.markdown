# Overview of this doc

"people are fairly good at short-term design, and usually awful at long-term design" - Dr. Roy Fielding

In this document we will discuss the principles behind our API and what that means for you the end user of this API.

Even if you won't use the API this is worth a read because it defines who we are as a company as well as the integrity that we strive to keep.

## The contract

This API is a contract from us to you. It's our word, it's everything. It also happens to be what we utilize to build the rest of our tools so we will maintain it as best as we can. 

That means we have thought through the following:

* Who is the API for
* What kind of API we have built
* How do we maintain and monitor each endpoint of the API
* How we will document the API
* How we will let you interact with the API in an easy manner
* How we will manage authentication, provisioning, throttling, and security.
* How will we protect servers against attacks and developer misuse
* How we manage support

## Overall design principles

### Use nouns: think of things as objects we're manipulating

/users over /createUser, /getUser, /deleteUser

Many times API's have this tendency of exposing confusing endpoints like 'ask'. While 'ask' is used sometimes to denote a question (a noun), many times it really is just a verb. When someone hits the ask endpoint what are they asking?

Instead a better approach would be to call the endpoint 'question'. When you POST to /question you would be creating a question. `GET /questions` would return the questions you posed. Etc.

### Use CRUD (Create-Read-Update-Delete) that maps to GET, POST, PUT, PATCH, DELETE
When manipulating objects the important pieces that users need access to are Create, Read, Update, Delete. Anything else can be built using that

Create (POST)
Read (GET)
Update (PUT / PATCH)
Delete (DELETE)

#### PUT / PATCH Updates

Technically PUT is meant to make a full update whereas PATCH only updates the relevant information. This can be a bit confusing so we will always treat PUT / PATCH the same way.

###  Use Hypermedia (HATEOAS)

TODO

### Use Accept/content-type
To round out the API to make it easy to use we will return (as we can) based on Accept/content-type

Internally we will always use Protocol Buffers but we will expose:

* Protocol Buffers ("application/x-protobuf")
* JSON ("application/json")
* XML ("application/xml")
* # RSS ("application/rss+xml")
* # ATOM ("application/atom+xml")
* GZIP ("application/gzip")
* CSV ("text/csv")
* Text ("text/plain")
* HTML ("text/html")

## Who is this API for?

### Current Customers
TODO

### Business Partners
TODO

### Third Parties
TODO

## What kind of API we have built

### Restful but also hypermedia

http://cl.ly/image/3J123A3y2B2N

#### Why are we using this format?
#### What does that mean for development?
#### What does that mean for usability?
#### What does that mean for longevity?

### Fully API driven development

Many times API's have a tendency to be an afterthought. Something to expose to users so they can customize and extend as needed. In our case it is exactly the opposite. We have built the API first, we use the API that you use so we keep it maintained cause we need it to be maintained to run our business. 

### Integrate with other services

An email provider doesn't exist in a vacuum. If you want to send information to other third parties as a part of your flow then you are allowed to do so. This is primarily done through an event api which sends off messages as things progress in your campaign.

## How we maintain and monitor each endpoint of the API

### Maximum flow
### Latency / Throughput / Amdahl's Law

* Lua
* ZeroMQ
* Microservices

## How we will document the API

* RAML

## How we will let you interact with the API in an easy manner

API notebook

## How we will manage authentication, provisioning, throttling, and security.
https://starttls.info/about
Datensparsamkeit
Forward Security

### Authentication

### Provisioning
### Throttling
### Security
#### Datensparsamkeit
#### Forward Security through Ecliptic Curve Cryptography
## How we protect servers against attacks and developer misuse
### Forward Security
## How we manage support
