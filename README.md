# RestingGopher

##### REST API written in golang

![gopher image](doc/gopher.png)

## About

This project is all about rest api for a popular news website.It's just for educational purpose , this code is not intended to be ready for production.

## Guidelines

The design of the API is based on small principles that I will note here.

#### HTTP Verbs

The API will support these main **http verbs**

API consumers are capable of sending **GET**, **POST**,
**PUT**, and **DELETE** verbs, which greatly enhance the clarity of a given request.

Generally, the four primary HTTP verbs are used as follows:

* **GET**
Read a specific resource (by an identifier) or a collection of resources.

* **PUT**
Update a specific resource (by an identifier) or a collection of resources. Can also be used to create a specific resource if the resource identifier is know before-hand.

* **DELETE**
Remove/delete a specific resource by an identifier.

* **POST**
Create a new resource. Also a catch-all verb for operations that don't fit into the other categories.

##### Note
> **GET** requests must not change any underlying resource data. Measurements and tracking which update data may still occur, but the resource data identified by the **URI** should not change.

#### Resource Naming

Appropriate **resource names** provide context for a service request, increasing understandability of the **API**. **Resources** are viewed hierarchically via their URI names, offering consumers a friendly, easily-understood hierarchy of **resources** to leverage in their applications.
* Use identifiers in your URLs instead of in the query-string. Using URL query-string parameters is fantastic for filtering, but not for resource names. **Example /users/12345**
* Leverage the hierarchical nature of the URL to imply structure.
* Design for your clients, not for your data.
* Resource names should be nouns. Avoid verbs as resource names, to improve clarity. Use the HTTP methods to specify the verb portion of the request.
* Use plurals in URL segments to keep your API URIs consistent across all HTTP methods, using the collection metaphor. **Example  /customers/33245/orders/8769/lineitems/1**
* Avoid using collection verbiage in URLs. For example 'customer_list' as a resource. Use pluralization to indicate the collection metaphor (e.g. customers vs. customer_list).
* Use lower-case in URL segments, separating words with **underscores**  or **hyphens** Some servers ignore case so it's best to be clear.
* Keep URLs as short as possible, with as few segments as makes sense.

#### Status codes
Response **status codes** are part of the **HTTP** specification. There are quite a number of them to address the most common situations. In the spirit of having our **RESTful** services embrace the **HTTP** specification, our **Web APIs** should return relevant HTTP **status codes**. For example, when a resource is successfully created (e.g. from a **POST** request), the **API** should return HTTP **status code** 201

### License
Apache 2.0