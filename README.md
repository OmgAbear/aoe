**This project has been developed on Ubuntu 20.04.4 LTS**

*It is not intended to run on any platform/chip*

## Install docker (with compose for your platform)

https://docs.docker.com/get-docker/

``cd /proj/dir/dev``<br/>
``docker compose up --build -d``

The docker container port to host mapping is 8080:8080.
Change it based on your own needs/port availability.
The dockerfile copies the entire project root and builds/runs the server there as is.
I did not use a multistage build.
---

**Dev folder contains a postman collection for the requests**

**Project structure and file naming should be pretty self-explanatory.**
    
    -root
        - cmd holds any available commands - in this case, a basic http server
        - dev holds development relevant docs and tools
        - internal package holds the actual domain and different layers
            - application holds the translation layer between external and internal data and acts as a facade
            - superheroes has the actual domain logic and so on
            - infrastructure holds the repository
            - presentation is the http handler layer and routing
        - static holds the asset (superheroes.json)

**There are no reusable packages made available externally, thus, the use of internal**

## Superhero creation is not persisted to file. ##
The file contents are loaded in memory at startup and any creation is stored in memory only

## Tests ##
You can find a couple of tests in the superheroes package for the purpose of displaying how I tend to tackle testing.
I generally prefer to limit scope of such testing to core/domain functionality while ensuring that other layers are tested in different stages.

## Validation ##
I only validated allowed superpowers. 
While I would normally add input validation and would most likely use the playground validator
I am exhausted from all the home assignments and interviews I've been doing lately. 
Highly opposed to this approach for prod though.

## Http Calls ##
There's a postman collection in dev folder.
Query string params for list are "encrypted" (bool) and "superpowers" (array) `(e.g. ?encrypted=true&superpowers=procrastination&superpowers=somethingElse) `