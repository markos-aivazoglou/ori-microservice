# ORI Microservice Coding Challenge

## Implementation: User venue check-in service

## TDD Awareness
There are enough tests on various parts to demonstrate an effective testing strategy as well as Go test best practices.
Ideally, integration tests would be written, and run separately.

## CI/CD Awareness
Similar to the above point, it would not make sense to demonstrate a specific ci/cd configuration/implementation as it's unlikely to be runnable.
Simple automation for generating the correct version of protobuf is added to the makefile(needed by CI/CD during build to track correct versions).
Also, different testing strategies can be run by CI/CD using the makefile. Of course, there are different strategies to an effective delivery pipeline
but here we demonstrate in simple terms the good practice.

## Operations Supported:
* User Check-in with friends in a venue (gRPC only)
* Getting all unique checked-in venues for a given user(gRPC and http)

## Alignment with the 12 factors
1) Codebase: This code base is checked into version control(along with it's configuration)
2) Dependencies: go modules satisfies what most modern languages do to satisfy the second factor
3) Config: In this example, there are two things satisfying this factor. First, for local development there is an .env file exporting the necessary env vars, and second, for a kubernetes environment, we make use of config-map.yaml to define those variables. In turn, our server and client read those variables in their main functions.
4) Backing Services: To demonstrate this, our client application treats our CheckIn server as an external resource, using an env variable to read the host name. Even though this example doesn't make any use of external databases or message queues, it would follow the same principles to connect to them.
5) Build, release, run: Our makefile will build go executable binaries, and use those to build docker images ready to be deployed. The release can happen using the k8s resource files along with the new built docker images to finaly run our application on a kubernetes cluster. In a real world setting, to keep track of build/release/deployment versions, a tool like helm to use different image versions would be used. In this case, for simplicity reasons the "latest" image is used, which is not the best practice for real world situations
6) Processes: Docker always run in one process, and it's the modern approach to run an application. This way we can make sure we run our applications in isolation from each other. Even though in this synthetic example, the repository is implemented in memory, in real world scenario an external database would be used.
7) Port Binding: Typical container best practice, exposes the ports where our application is bound to. In this example, our gRPC server binds to port 7777, and our http server binds to port 6666.
8) Concurrency: Go provides a comprehensive and complete model of concurrency using goroutines. Whether we want so serve multiple http requests, or batch process we can use the build in mechanisms, or make our own.
9) Disposability: As our application(s) run in Kubernetes, it's easy to dispose and recreate pods at will or automatically. The app will also listen for SIGTERM and terminate gracefully
10) Dev/Prod parity: As you can see by looking at the codebase, and following the instructions this readme provides, this microservice and its client application can run in minikube -a fairly good emulation of a kubernetes cluster. Thus, with probably almost no modification you can run this in production too using the same files. Also, tests are not relying on different backend mechanisms
11) Logs: Kubernetes does well with this factor. kubectl logs shows that logs are indeed streamed to stdout and then consumed by a log aggregator very easily. Public cloud providers have tools that by default will aggregate logs from vms and kubernetes resources(eg Stackdriver)
12) Admin processes: Again, with kubectl we can remotely control our apps and provide access to admins in a fine grained fashion. Go specifically doesn't provide an interpreter as it's a statically typed compiled language, but we can easily write our own utility commands and interact with our app for administrative purposes.

## Alignment with Cloud Native principles
* The app is build and bundled to run in a container environment. In this case Kubernetes
* It's written in a microservices fashion
* It is optimised to run on a PaaS
* It is optimised to run on the Cloud as infrastructure, configuration and application definitions are abstracted in a cloud friendly fashion(kubernetes) but we could use terraform to provision our clusters, etc
* Running on kubernetes we have the benefits of elasticity(eg horizontal scaling) and updatability(quick rollouts and rollbacks)
* Deployment, maintenance etc can be highly automated as it can be briefly demonstrated by the use of makefile

## Accommodating an eventstore
This example implementation follows architecture principles that allow for any kind of datastore to be used. Hexagonal architecture is employed to make it easy
for developers to swap out external services(dbs, event stores, etc). The Domain Model pattern is used(drawn from Domain-Driven Design principles) to allow for future extention into Event Based systems. Our application layer can easily use events instead of saved state, to interact with objects and we can easily plug in an event-bus(implemented using sqs, or kafka or google pub/sub) to allow for message exchanging

## Access from outside the cluster
A NodePort kubernetes service is defined in the server k8s resources, along with an nginx kubernetes ingress controller to allow access to the "outside world".
Please follow the instructions in the "Run The App" section to be able to run a request from your browser and see some venues our user has checked in.

## Run Unit Tests for the app
Run `make unit-test`

There is a makefile rule for integration tests, but they are not implemented. It is there to demonstrate its usefulness for the CI/CD
## Run The App
What are the interactions:
1) The server will import a single user with a predefined user id to make the interactions possible
2) The client will check in in two different venues using that user id and print the result of the check in
3) The client will ask the server for all the check ins that user has done(should see 2 venues)

### Run Locally from terminal
If you want to run the app from the console then:
1) Run `make run-local-server` (on one terminal)
2) Run `make run-local-client` (on another terminal)


Finally, you can also visit `localhost:6666/user/195b5c7f-4bc7-461b-8438-8beb9f9fcd16/venues` on the browser or use curl
to see the results of number 3 over http

### Build images and deploy to minikube cluster
This assumes you have a minikube cluster installed and running successfully on your system

1) Run `eval $(minikube docker-env)`. This will switch to minikube's docker daemon (to switch back to local docker daemon run `eval $(minikube docker-env -u)`)
2) Run `make build-all`. You can use this rule locally too, but to run on minikube the images need to run on minikube's docker to avoid using a docker registry
3) Run `deploy server`. This will deploy all resources related to the server (check logs)
4) Once everything related to server is deployed and running, Run `deploy client` (check logs). you should be able to see the results mentioned in the interactions above
5) An ingress controller is also deployed so that you can access the service from outside the cluster. Optionally, follow the steps below to test the endpoint from outside the cluster. The reason this is not automated is because it requires sudo access to modify /etc/hosts.
    1) Run `minikube ip`
    2) Add the result ip from previous step and `checkin-service.info` in your /etc/hosts
        example:
            `172.17.0.2 checkin-service.info`
    3) Access http:/checkin-service.info/user/195b5c7f-4bc7-461b-8438-8beb9f9fcd16/venues, and confirm that you are seeing results served by the http server