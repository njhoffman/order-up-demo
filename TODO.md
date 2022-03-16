
## Objectives
- [ ] 1: Add a Cancel API endpoint
    You need to implement a cancel API endpoint that:
  - [ ] refunds the customer if they were already charged and
  - [ ] updates an order’s status to cancelled.
    The charge service’s /charge endpoint allows you to pass a negative amountCents value to refund a card for a particular amount.
  - [ ] If the refund failed you should not update the status of the order and return the error.
  - [ ] If an order has already been fulfilled then you must return a 409 Conflict error since we’ve already fulfilled the order and cannot cancel it.
  - [ ] You are expected to write any tests necessary to ensure your endpoint is working property. The other tests can be used as a reference.
- [ ] 2: Add a Fulfill API endpoint
    Employees will be using an internal website to mark orders as fulfilled when they’re ready.
  - [ ] You need to implement a fulfill API endpoint that ensures the order’s already been charged,
  - [ ] makes a PUT /fulfill request to the fulfillment service for each line item
  - [ ] and finally updates the order’s status to fulfilled.
  - [ ] The fulfillment service’s /fulfill endpoint returns a 200 OK if the fulfillment succeeded.
  - [ ] If any of the fulfillment service requests fail you should not update the status of the order and return the error.
  - [ ] The fulfillment service is idempotent as long as an orderID is passed so if an employee tries again you can call /fulfill again on the fulfillment service and it will ignore items that are already be fulfilled.
  - [ ] You are expected to write any tests necessary to ensure your endpoint is working property. The fulfillment service mock has been provided for you. The other tests can be used as a reference.
- [ ] 3: Write database code
  - [ ] Only the function skeletons and tests exist for in the storage package/folder and so you need to decide on a database and add the appropriate initialization code and CRUD code to each function in the storage package
- [ ] 4: Limit concurrent changes
    Our charging service can only support 1 concurrent charge at a time.
  - [ ] The order service needs to limit itself to only having 1 outgoing /charge request at a time and most importantly, any incoming requests should be queued up until any previous ones have completed
  - [ ] You can assume there’s only 1 instance of the order-up service running. A test has already been written to verify that only 1 concurrent request is being sent to the charge service.
- [ ] 5: Document API
    Finally, you need to write documentation for a frontend developer who might be interacting with this service.
  - [ ] You should include all necessary details for them to build the public website and internal site assuming they have no existing knowledge about the service or it’s implementation.
