# Juggler is an experimental network services operating system.
#### Be impolite and strong. Juggler is only for those who have balls.

## How to define a ball (juggler managed service)?
The example bellow is not working (yet). It just shows what we want to achieve.

```
# spec/service/signup.ball <-- spec file name defines a service name
ui-widget: <path-to-ui-widget-sources>
build-path: ./build/$service-name

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
repositories:
  <repository-name>:
    <query-name, like: GetUsers>: <any-of: mongodb, postgresql, redis, cassandra, ...>

#
# In perspective:
#
<any-other-interface>:
  <any-other-interface-implementation>:<any-other-settings and drivers>
```

## How to define a subscription?
The example bellow is not working (yet). It just shows what we want to achieve.

```
# spec/subscription/signup-data-sent.ball <-- spec file name defines a service name
transport: <any-of: http, websocket, ...>
format:    <any-of: json, xml, bson, protobuf, ...>
data-spec:
  >- <string of specification (spec lib)>
```

