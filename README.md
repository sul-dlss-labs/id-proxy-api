# DLSS Identifiers Proxy API & Microservice for AWS Applications (Hydrox)

Microservice components:
- Swagger API specification for loading into AWS API Gateway
- Lambdas for running each operation tied to DLSS Identifiers Proxy API operation (tied to each route)
- Testing framework and build for API and Lambdas
- Eventually, deploy tools for loading this API into defined environments in AWS via Terraform or other framework

## Guiding Questions around original SURI Proxy API idea

1. Why SURI Proxy?
  - Better permissions management (currently, SURI uses shared passwords across local codebases).
  - Separate the specification from the implementation / make connections more modular as SURI may (or may not?) change.
  - Set up a pattern for building APIs as AWS work continues
  - Capture better analytics for AWS querying SURI or future SDR components via these proxy APIs.
2. SURI Proxy API should be in AWS?
  - Yes, if we can support in production. Given Hydrox is in AWS, SURI Proxy in AWS seems a good step for this.
3. Authentication for the SURI Proxy
  - TBD: method and how we can improve this (but probably having authentication attached to a to-be-created AWS IAM
    Service Role)
  - Proposal of using API keys (a la GitHub) also proposed
  - Protection of SURI: Using SSL?

## Requirements

### For Hydrox:
- Create a new DRUID on a single call for single identifier action
- Support later additions (but no current implementation) for:
  - Request a list of DRUIDs
  - Use other namespaces for identifiers (not DRUIDs)
  - Do not support ability to create a new identifier namespace or identifier spec via this API (should occur in the SURI implementation)

### For APIs generally:
- Version the API
- Elegant solution for retries / connection issues when calling SURI
