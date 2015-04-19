# `cm15` Basic Example

The basic example illustrates how to instantiate a cm15 client.
It uses the `LoginAuthenticator` to perform basic authentication using email
and password.

Once authenticated the example uses the API client to instantiate a `CloudLocator` and calls the
`index` action using the `extended` view. The example then iterates over the returned clouds and
prints their capabilities as returned by the RightScale API.
