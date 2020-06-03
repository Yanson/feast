# API

The following API docs are available

* [Feast Core gRPC API](https://api.docs.feast.dev/grpc/feast.core.pb.html): This is the gRPC API used by Feast Core. Feast Core has a dual function of schema registry and job manager. This API contains RPCs for creating and managing feature sets, stores, projects, and jobs.
* [Feast Serving gRPC API](https://api.docs.feast.dev/grpc/feast.serving.pb.html): This is the gRPC API used by Feast Serving. It contains RPCs used for the retrieval of online feature data or historical feature data.
* [Feast gRPC Types](https://api.docs.feast.dev/grpc/feast.types.pb.html): These are the gRPC types used by both Feast Core, Feast Serving, and the Go, Java, and Python clients.
* [Go Client SDK](https://godoc.org/github.com/feast-dev/feast/sdk/go): The Go library used for the retrieval of online features from Feast.
* [Java Client SDK](https://javadoc.io/doc/dev.feast/feast-sdk): The Java library used for the retrieval of online features from Feast.
* [Python SDK](https://api.docs.feast.dev/python/): This is the complete reference to the Feast Python SDK. The SDK is used to manage feature sets, features, jobs, projects, and entities. It can also be used to retrieve training datasets or online features from Feast Serving.

