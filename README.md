# Juggler is an experimental network services operating system.
#### Be impolite and strong. Juggler is only for those who have balls.

## How to define a ball (juggler managed service)?
The example bellow is not working (yet). It just shows what we want to achieve.

```
# spec/service/signup.ball <-- spec file name defines a service name
ui-widget: <path-to-ui-widget-sources>
logging:
  <some-setting, like: log-level>: <some-setting-value, like: INFO>
monitoring: true
inputs:
  <subscription-name, like: signup-data-sent>: true
  <subscription-name>:
    <some-setting>:<some-setting-value>
outputs:
  <subscription-name, like: signup-data-sent>: true
  <subscription-name>:
    <some-setting>:<some-setting-value>
```

## How to define a subscription?
The example bellow is not working (yet). It just shows what we want to achieve.

```
# spec/subscription/signup-data-sent.ball <-- spec file name defines a service name
transport: <any-of: http, websocket, ...>
data-spec:
  >- <string of specification (spec lib)>
```