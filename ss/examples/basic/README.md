# `ss` Basic Example

The basic example illustrates how to instantiate a Self-Service API client.
It uses the `LoginAuthenticator` to perform basic authentication using email
and password.

Once authenticated the example uses the API client to instantiate a `ExecutionLocator` and calls its
`index` action. The example then iterates over the returned Self-Service executions
(running Cloud Applications) and prints their status.

