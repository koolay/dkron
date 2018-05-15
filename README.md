# Dkron - Distributed, fault tolerant job scheduling system [![GoDoc](https://godoc.org/github.com/victorcoder/dkron?status.svg)](https://godoc.org/github.com/victorcoder/dkron) [![Build Status](https://travis-ci.org/victorcoder/dkron.svg?branch=master)](https://travis-ci.org/victorcoder/dkron)

Website: http://dkron.io/

Dkron is a distributed cron service, easy to setup and fault tolerant with focus in:

- Easy: Easy to use with a great UI
- Reliable: Completely fault tolerant
- High scalable: Able to handle high volumes of scheduled jobs and thousands of nodes

Dkron is written in Go and leverage the power of distributed key-value stores and serf for providing fault tolerance, reliability and scalability while keeping simple and easily instalable.

Dkron is inspired by the google whitepaper [Reliable Cron across the Planet](https://queue.acm.org/detail.cfm?id=2745840) and by Airbnb Chronos borrowing the same features from it.

Dkron runs on Linux, OSX and Windows. It can be used to run scheduled commands on a server cluster using any combination of servers for each job. It has no single points of failure due to the use of the Gossip protocol and fault tolerant distributed databases.

You can use Dkron to run the most important part of your company, scheduled jobs.

## Status

Currently Dkron is under heavy development but approaching production quality, it has been tested in production for long time and a moderate amount of jobs.

Said that, I encourage you to try it, it's very easy to use, see how it works for you and report any bugs [creating an issue](https://github.com/victorcoder/dkron/issues) in the github project.

## Documentation

Full, comprehensive documentation is viewable on the [Dkron website](http://dkron.io)

## Development Quick start

The best way to test and develop dkron is using docker, you will need [Docker](https://www.docker.com/) installed before proceding.

Clone the repository.

Install dep, follow instructions here https://github.com/golang/dep#setup

Download dependencies:

`dep ensure`

Next, run the included Docker Compose config:

`docker-compose up`

This will start, etc, consul and Dkron instances. To add more Dkron instances to the clusters:

`docker-compose scale dkron=4`

Check the port mapping using `docker-compose ps` and use the browser to navigate to the Dkron dashboard using one of the ports mapped by compose.

To add jobs to the system read the API docs.

## Frontend development

Dkron dashboard is built using a combinations of golang templates and AngularJS code.

To start developing the dashboard enter the `static` directory and run `npm install` to get the frontend dependencies.

Change code in JS files or in templates, then run `make gen` to generate `bindata.go` file. This is a method of embedding resources in Go applications.

### Resources

Chef cookbook
https://supermarket.chef.io/cookbooks/dkron

Python Client Library
https://github.com/oldmantaiter/pydkron

Ruby client
https://github.com/jobandtalent/dkron-rb

## Get in touch

- Twitter: [@dkronio](https://twitter.com/dkronio) or [@victorcoder](https://twitter.com/victorcoder)
- Email: victor at victorcastell.com


## BUG to fix

1. 当一个job失败后变成运行中状态，不能再恢复重新运行

# Sponsor

This project is possible thanks to the Support of Jobandtalent

![](https://upload.wikimedia.org/wikipedia/en/d/db/Jobandtalent_logo.jpg)

