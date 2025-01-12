# Go SSE

#### Phase 1

> **Server**

- [/] Handle `POST /messsage`
  - [x]  Publish message to `room:id` message broker channel
  - [ ]  Save messages to message store
- [x] Handle SSE through `/room/:id` endpoint
  - [x]  Handler subs to `room:id`  message broker channel and pushes new messages to subscribed clients
- [x] No Redis, in memory broker
- [ ] In memory DB, no persistence

### OpenAPI

<https://github.com/OpenAPITools/openapi-generator>
<https://github.com/oapi-codegen/oapi-codegen>
<https://github.com/ogen-go/ogen>
<https://medium.com/@printSAN0/automating-swaggerui-for-go-chi-router-without-swaggo-part-1-032698239f2a>
