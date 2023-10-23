# Authentication and Authorization

## Challenges

- Central dependency

authentication and authorization logic must be handled separately by each microservice. You could use the same code in all microservices, but this requires that all microservices support a specific language or framework.

- Violating the single responsibility principle

microservices are supposed to fulfill only one function. If you add global authentication and authorization logic to microservices, they now perform an additional function, making them less reliable and more difficult to manage.

- Complexity

authentication and authorization in microservices can lead to very complex scenarios. Consider that there might be users, microservices, and third-party systems accessing every microservice. This complexity can make implementation and maintenance difficult.

## Approaches

- Edge-Level Authorization

You can use an API Gateway to centralize authentication and authorization for all downstream microservices. The gateway enforces authentication and access control for each microservice. In this case, NIST recommends implementing mitigation controls such as mutual authentication between microservices to prevent direct anonymous connections to internal services.

- Service-Level Authorization

