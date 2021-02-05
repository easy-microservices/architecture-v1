# architecture-v1
Dedicated to people who want to learn microservices based on my personal notes and experiences.

# Introduction

Microservices has become a buzzword in today's industry. Everybody would like to follow this architecture and apply it in their business with enough confidence to consider that it fits them without understanding the full requirements and drawbacks they generate. 

My mission is to clear mist that have been generated over the internet by those individuals which write a lot of tutorials, and pages which I found most of them to be very confusing and incomplete. 

This series is dedicated to you, well especially most developers who want to start learning about microservices or maybe to improve their knowledge. You will find here a list of projects, each one containing a separate application which will be in deep dissussed and explained with a lot of concepts, patterns, techniques that are widely used in a microservice architecture. 

In this first project tho, we will start exploring the most important concepts and patterns that you will have to learn in order to begin your journey in building microservices by creating an **e-commerce** application.

## By Definition

Ok, let's start by defining what are microservices. The Microservices Architecture is a sub-architectural style of distributed-systems which involves breaking your application into a collection of individual, self-contained services, which are communicating between them through network calls. ( Mainly just how Andrew Tannenbaum said: "just a bunch of machines which are communication over a network") simple enough huh ?.

The main characteristics of an correct microservice oriented architecture are:

- **SR** (Single Responsabiliy) each microservice has only one job, providing a cohesive set of functionalities built around an particular business domain. I think this term is very familiar to you, is the same one of the five *SOLID* concepts you maybe studied in Object Oriented Programming.

- **Loosely Coupled** Coupling is the deggree of independence between multiple software components (basically is shows how closelly these components are). Loose Coupling states that two components know very little about each other, they are highly detached and can work independenly in a software system. *Pssst remember low coupling is always a best practice to apply when building good software.* 

- **Independent Deployability** This one was a little bit hard to understand and always brought me into confusion. It bassically says that we can independently deploy one service in every desired environment without affectiv other services in our system. Making each service independent deployable will lower the chance of breaking your system by impacting other services. To guarantee this we should ensure first that our services are *loosely coupled*.

Each microservice encapsulates a set of functionalities which were modeled around some business domain ( which is actually just the sphere of activities that software should perform) exposing this logic should only be done through some well defined network endpoints. 

### Benefits

- **Flexibility** 

    When it comes to flexibility, microservices have an open culture, you can easily addopt *Polyglot Technologies*. The fact that our architecture now reliece on multiple distributed services, it give us the advantage of choosing the right/prefered technology stack for each individual service. How can you do this ? because each service needs to be loosely coupled and independently maintaned, this gives you the posibility to choose the stacl that fits for your service, while communication will be handled by some standard format(JSON, XML, Binary, etc).  

    Say your team needs to build two services each one with different requirements: one needs to manage a ton of connection , for this one you choose to go with (Golang or Node.js), other one should have better performance (C++ may be a better option here). At this point the deccision is strictly handled by your team. But at least you can easly do it.


- **Parallel Deployment**
- **Maintability**
- **Highly Scallable**
- **Easy to develop**
- **Increased Productivity**

### Drawbacks
- **Complicated and Expensive**
- **You need dev power**
- **Client impressions**
- **End-To-End Testing**
- **Migrations**


# Get Rid of Monoliths ?

## Comparison with Microservices

# Challanges

## Communication

### Synchronous

### Asynchronous

## Testing

## Reliability

## Concistency

# Domain-Driven Design (DDD)

# Databases


## CAP Theorem

## Eventual vs Strong Concistency

## Do you still want to guarantee ACID ?


# Local Development Environment

# Monitoring and Tracking

# What you should avoid !

## Breaking contracts

# Enough talking, show me some code!
Ok, ok I got you, nobody is satisfied just with some theory , however it is very importat to understand the fundamental concepts first. Now it's time for some hands on practice. In this first part I developed an simple microservice e-commerce system with REST oriented communication.

## Architecture

Below you can see how many microservices I planned to build, each with a short description about their business domain they need should address.

- **product-catalog-service** : stores products and any related data about them.
- **order-service** : manages orders made by users.
- **configuration-service** : distributes configuration to microservices during their bootstrap process.
- **search-service** : search products and their detailes.
- **shipping-service** : deals with shipping details, recipes, and delivery.
- **payment-service** : handles prayment transaction and information.
- **authentication-service** : handles user authentication and authorization.
- **api-gateway** : implements the api gateway pattern ( entrypoint in our system, handles caching,security,compression,aggregationetc).
- **web-ui** : the user interface application.


## Development Stack

- Docker
- Golang
- Ngnix
- React.js
- Redis
- MongoDB
- Postgresql

## Building and Deployment


## References

Below you can find some of the best learning materials which I encountered during my daily searches ( both books and tutorials). I've filtered them so you can benefit from the most usefull and notable papers.

### Books

- [Building Microservices by Sam Newman](https://www.amazon.com/Building-Microservices-Designing-Fine-Grained-Systems/dp/1491950358/ref=sr_1_1?crid=3O8AK4JPH58SQ&dchild=1&keywords=building+microservices&qid=1611162924&sprefix=Building+Micr%2Caps%2C344&sr=8-1)
- [From Monoliths to Microservices by Sam Newman](https://www.amazon.com/Monolith-Microservices-Evolutionary-Patterns-Transform/dp/1492047848/ref=sr_1_1?crid=1EAFELZEVAP4H&dchild=1&keywords=from+monolith+to+microservices&qid=1611163087&sprefix=From+Monolits%2Caps%2C324&sr=8-1)
- [Microservices Patterns by Chris Richardson](https://www.amazon.com/Microservices-Patterns-examples-Chris-Richardson/dp/1617294543/ref=sr_1_1?crid=MNQHS2MZHO5S&dchild=1&keywords=microservices+design+patterns&qid=1611163111&sprefix=Microservices+Design+%2Caps%2C330&sr=8-1)

### Links
- https://microservices.io/ this website belongs to Chris Richardson one of the most experienced architects when It comes to microservices. Here it briefly describes a series of design patterns which are strongly utilized in the Microservice world.

- https://martinfowler.com/ Martin Fowler is one of the veterans of this complex industry, highly experienced in refactoring and building strong architecture. Here you can find some great articles, not only about microservices but about software engineering in general.

