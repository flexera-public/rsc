# `cm15` auditail Example

The auditail example polls the RightScale CM 1.5 APIs every second for the last 100
audit entries created by the given user (authenticated user by default).
It uses the `LoginAuthenticator` to perform basic authentication using email
and password.

auditail prints any new or updated audit entry every second.
